package submitter

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
	"github.com/taikoxyz/taiko-client/metrics"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	proofProducer "github.com/taikoxyz/taiko-client/prover/proof_producer"
)

var _ ProofSubmitter = (*InvalidProofSubmitter)(nil)

// InvalidProofSubmitter is responsible requesting zk proofs for the given invalid L2
// blocks, and submitting the generated proofs to the TaikoL1 smart contract.
type InvalidProofSubmitter struct {
	rpc              *rpc.Client
	proofProducer    proofProducer.ProofProducer
	reusltCh         chan *proofProducer.ProofWithHeader
	proverPrivKey    *ecdsa.PrivateKey
	proverAddress    common.Address
	anchorTxGasLimit uint64
	mutex            *sync.Mutex
}

// NewInvalidProofSubmitter creates a new InvalidProofSubmitter instance.
func NewInvalidProofSubmitter(
	rpc *rpc.Client,
	proofProducer proofProducer.ProofProducer,
	reusltCh chan *proofProducer.ProofWithHeader,
	proverPrivKey *ecdsa.PrivateKey,
	anchorTxGasLimit uint64,
	mutex *sync.Mutex,
) *InvalidProofSubmitter {
	return &InvalidProofSubmitter{
		rpc:              rpc,
		proofProducer:    proofProducer,
		reusltCh:         reusltCh,
		proverPrivKey:    proverPrivKey,
		proverAddress:    crypto.PubkeyToAddress(proverPrivKey.PublicKey),
		anchorTxGasLimit: anchorTxGasLimit,
		mutex:            mutex,
	}
}

// RequestProof implements the ProofSubmitter interface.
func (s *InvalidProofSubmitter) RequestProof(ctx context.Context, event *bindings.TaikoL1ClientBlockProposed) error {
	// Get the throwaway block from L2 execution engine.
	throwAwayBlock, err := s.getThrowAwayBlock(ctx, event)
	if err != nil {
		return err
	}

	log.Debug("Throwaway block", "height", throwAwayBlock.Header().Number, "hash", throwAwayBlock.Header().Hash())

	// Request proof.
	proofOpts := &proofProducer.ProofRequestOptions{
		Height:             throwAwayBlock.Header().Number,
		ProverAddress:      s.proverAddress,
		ProposeBlockTxHash: event.Raw.TxHash,
	}

	if err := s.proofProducer.RequestProof(
		ctx, proofOpts, event.Id, &event.Meta, throwAwayBlock.Header(), s.reusltCh,
	); err != nil {
		return err
	}

	metrics.ProverQueuedProofCounter.Inc(1)
	metrics.ProverQueuedInvalidProofCounter.Inc(1)

	return nil
}

// SubmitProof implements the ProofSubmitter interface.
func (s *InvalidProofSubmitter) SubmitProof(
	ctx context.Context,
	proofWithHeader *proofProducer.ProofWithHeader,
	isOracle bool,
) (err error) {
	log.Info(
		"New invalid block proof",
		"blockID", proofWithHeader.BlockID,
		"beneficiary", proofWithHeader.Meta.Beneficiary,
		"hash", proofWithHeader.Header.Hash(),
		"proof", common.Bytes2Hex(proofWithHeader.ZkProof),
	)
	var (
		blockID = proofWithHeader.BlockID
		header  = proofWithHeader.Header
		zkProof = proofWithHeader.ZkProof
		meta    = proofWithHeader.Meta
	)

	metrics.ProverReceivedProofCounter.Inc(1)
	metrics.ProverReceivedInvalidProofCounter.Inc(1)

	// Get the corresponding L2 throwaway block, which is not in the L2 execution engine's canonical chain.
	block, err := s.rpc.L2.BlockByHash(ctx, header.Hash())
	if err != nil {
		return fmt.Errorf("failed to get the throwaway block (id: %d): %w", blockID, err)
	}

	if block.Transactions().Len() == 0 {
		return fmt.Errorf("invalid throwaway block without any transaction, blockID %s", blockID)
	}

	// Fetch all receipts inside that L2 throwaway block.
	receipts, err := s.rpc.L2.GetThrowawayTransactionReceipts(ctx, header.Hash())
	if err != nil {
		return fmt.Errorf("failed to fetch the throwaway block's transaction receipts (id: %d): %w", blockID, err)
	}

	log.Debug("Throwaway block receipts fetched", "length", receipts.Len())

	// Generate the merkel proof (whose root is block's receiptRoot) of the TaikoL2.invalidateBlock transaction's receipt.
	receiptRoot, receiptProof, err := generateTrieProof(receipts, 0)
	if err != nil {
		return fmt.Errorf("failed to generate anchor receipt proof: %w", err)
	}
	if receiptRoot != header.ReceiptHash { // Double check the calculated root.
		return fmt.Errorf("receipt root mismatch, receiptRoot: %s, block.ReceiptHash: %s", receiptRoot, header.ReceiptHash)
	}

	// Assemble the TaikoL1.proveBlockInvalid transaction inputs.
	txListBytes, err := rlp.EncodeToBytes(block.Transactions())
	if err != nil {
		return fmt.Errorf("failed to encode throwaway block transactions: %w", err)
	}

	circuitsIdx, err := proofProducer.DegreeToCircuitsIdx(proofWithHeader.Degree)
	if err != nil {
		return err
	}

	evidence := &encoding.TaikoL1Evidence{
		Meta: bindings.TaikoDataBlockMetadata{
			Id:          meta.Id,
			L1Height:    meta.L1Height,
			L1Hash:      meta.L1Hash,
			Beneficiary: header.Coinbase,
			GasLimit:    header.GasLimit - s.anchorTxGasLimit,
			Timestamp:   header.Time,
			TxListHash:  crypto.Keccak256Hash(txListBytes),
			MixHash:     header.MixDigest,
			ExtraData:   header.Extra,
		},
		Header:   *encoding.FromGethHeader(header),
		Prover:   s.proverAddress,
		Proofs:   [][]byte{zkProof, receiptProof},
		Circuits: circuitsIdx,
	}

	input, err := encoding.EncodeProveBlockInvalidInput(evidence, meta, receipts[0])
	if err != nil {
		return err
	}

	// Send the TaikoL1.proveBlockInvalid transaction.
	txOpts, err := getProveBlocksTxOpts(ctx, s.rpc.L1, s.rpc.L1ChainID, s.proverPrivKey)
	if err != nil {
		return err
	}

	sendTx := func() (*types.Transaction, error) {
		s.mutex.Lock()
		defer s.mutex.Unlock()

		return s.rpc.TaikoL1.ProveBlockInvalid(txOpts, blockID, input)
	}

	if err := sendTxWithBackoff(ctx, s.rpc, blockID, sendTx, isOracle); err != nil {
		if errors.Is(err, errUnretryable) {
			return nil
		}

		return err
	}

	log.Info(
		"💰 Your invalid block proof was accepted",
		"blockID", proofWithHeader.BlockID,
		"height", block.Number(),
		"hash", header.Hash(),
	)

	metrics.ProverSentProofCounter.Inc(1)
	metrics.ProverSentInvalidProofCounter.Inc(1)
	metrics.ProverLatestProvenBlockIDGauge.Update(proofWithHeader.BlockID.Int64())

	return nil
}

// getThrowAwayBlock keeps waiting till the throwaway block inserted into the L2 chain.
func (s *InvalidProofSubmitter) getThrowAwayBlock(
	ctx context.Context,
	event *bindings.TaikoL1ClientBlockProposed,
) (*types.Block, error) {
	l1Origin, err := s.rpc.WaitL1Origin(ctx, event.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch L1origin, blockID: %d, err: %w", event.Id, err)
	}

	if !l1Origin.Throwaway {
		return nil, fmt.Errorf("invalid throwaway block's L1origin found, blockID: %d: %+v", event.Id, l1Origin)
	}

	return s.rpc.L2.BlockByHash(ctx, l1Origin.L2BlockHash)
}
