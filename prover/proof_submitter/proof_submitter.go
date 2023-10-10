package submitter

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
	"github.com/taikoxyz/taiko-client/metrics"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	anchorTxValidator "github.com/taikoxyz/taiko-client/prover/anchor_tx_validator"
	proofProducer "github.com/taikoxyz/taiko-client/prover/proof_producer"
	"github.com/taikoxyz/taiko-client/prover/proof_submitter/evidence"
)

var _ Submitter = (*ProofSubmitter)(nil)

// ProofSubmitter is responsible requesting proofs for the given L2
// blocks, and submitting the generated proofs to the TaikoL1 smart contract.
type ProofSubmitter struct {
	rpc                        *rpc.Client
	proofProducer              proofProducer.ProofProducer
	resultCh                   chan *proofProducer.ProofWithHeader
	evidenceAssembler          *evidence.Assembler
	txSender                   *TxSender
	proverPrivKey              *ecdsa.PrivateKey
	proverAddress              common.Address
	taikoL2Address             common.Address
	l1SignalService            common.Address
	l2SignalService            common.Address
	graffiti                   [32]byte
	proveBlockTxGasLimit       *uint64
	txReplacementTipMultiplier uint64
	proveBlockMaxTxGasTipCap   *big.Int
	mutex                      *sync.Mutex
}

// New creates a new ProofSubmitter instance.
func New(
	rpcClient *rpc.Client,
	proofProducer proofProducer.ProofProducer,
	resultCh chan *proofProducer.ProofWithHeader,
	taikoL2Address common.Address,
	proverPrivKey *ecdsa.PrivateKey,
	graffiti string,
	submissionMaxRetry uint64,
	retryInterval time.Duration,
	waitReceiptTimeout time.Duration,
	proveBlockTxGasLimit *uint64,
	txReplacementTipMultiplier uint64,
	proveBlockMaxTxGasTipCap *big.Int,
) (*ProofSubmitter, error) {
	anchorValidator, err := anchorTxValidator.New(taikoL2Address, rpcClient.L2ChainID, rpcClient)
	if err != nil {
		return nil, err
	}

	l1SignalService, err := rpcClient.TaikoL1.Resolve0(nil, rpc.StringToBytes32("signal_service"), false)
	if err != nil {
		return nil, err
	}

	l2SignalService, err := rpcClient.TaikoL2.Resolve0(nil, rpc.StringToBytes32("signal_service"), false)
	if err != nil {
		return nil, err
	}

	maxRetry := &submissionMaxRetry
	if proofProducer.Tier() == encoding.TierGuardianID {
		maxRetry = nil
	}

	return &ProofSubmitter{
		rpc:                        rpcClient,
		proofProducer:              proofProducer,
		resultCh:                   resultCh,
		evidenceAssembler:          evidence.NewAssembler(rpcClient, anchorValidator, graffiti),
		txSender:                   NewTxSender(rpcClient, retryInterval, maxRetry, waitReceiptTimeout),
		proverPrivKey:              proverPrivKey,
		proverAddress:              crypto.PubkeyToAddress(proverPrivKey.PublicKey),
		l1SignalService:            l1SignalService,
		l2SignalService:            l2SignalService,
		taikoL2Address:             taikoL2Address,
		graffiti:                   rpc.StringToBytes32(graffiti),
		proveBlockTxGasLimit:       proveBlockTxGasLimit,
		txReplacementTipMultiplier: txReplacementTipMultiplier,
		proveBlockMaxTxGasTipCap:   proveBlockMaxTxGasTipCap,
		mutex:                      new(sync.Mutex),
	}, nil
}

// RequestProof implements the Submitter interface.
func (s *ProofSubmitter) RequestProof(ctx context.Context, event *bindings.TaikoL1ClientBlockProposed) error {
	l1Origin, err := s.rpc.WaitL1Origin(ctx, event.BlockId)
	if err != nil {
		return fmt.Errorf("failed to fetch l1Origin, blockID: %d, err: %w", event.BlockId, err)
	}

	// Get the header of the block to prove from L2 execution engine.
	block, err := s.rpc.L2.BlockByHash(ctx, l1Origin.L2BlockHash)
	if err != nil {
		return fmt.Errorf("failed to get the current L2 block by hash (%s): %w", l1Origin.L2BlockHash, err)
	}

	if block.Transactions().Len() == 0 {
		return errors.New("no transaction in block")
	}

	parent, err := s.rpc.L2.BlockByHash(ctx, block.ParentHash())
	if err != nil {
		return fmt.Errorf("failed to get the L2 parent block by hash (%s): %w", block.ParentHash(), err)
	}

	blockInfo, err := s.rpc.TaikoL1.GetBlock(&bind.CallOpts{Context: ctx}, event.BlockId.Uint64())
	if err != nil {
		return err
	}

	signalRoot, err := s.rpc.GetStorageRoot(ctx, s.rpc.L2GethClient, s.l2SignalService, block.Number())
	if err != nil {
		return fmt.Errorf("failed to get L2 signal service storage root: %w", err)
	}

	// Request proof.
	opts := &proofProducer.ProofRequestOptions{
		BlockID:            block.Number(),
		ProverAddress:      s.proverAddress,
		ProposeBlockTxHash: event.Raw.TxHash,
		L1SignalService:    s.l1SignalService,
		L2SignalService:    s.l2SignalService,
		TaikoL2:            s.taikoL2Address,
		MetaHash:           blockInfo.MetaHash,
		BlockHash:          block.Hash(),
		ParentHash:         block.ParentHash(),
		SignalRoot:         signalRoot,
		EventL1Hash:        event.Raw.BlockHash,
		Graffiti:           common.Bytes2Hex(s.graffiti[:]),
		GasUsed:            block.GasUsed(),
		ParentGasUsed:      parent.GasUsed(),
	}

	if err := s.proofProducer.RequestProof(
		ctx,
		opts,
		event.BlockId,
		&event.Meta,
		block.Header(),
		s.resultCh,
	); err != nil {
		return fmt.Errorf("failed to request proof (id: %d): %w", event.BlockId, err)
	}

	metrics.ProverQueuedProofCounter.Inc(1)

	return nil
}

// SubmitProof implements the Submitter interface.
func (s *ProofSubmitter) SubmitProof(
	ctx context.Context,
	proofWithHeader *proofProducer.ProofWithHeader,
) (err error) {
	log.Info(
		"New block proof",
		"blockID", proofWithHeader.BlockID,
		"proposer", proofWithHeader.Meta.Coinbase,
		"hash", proofWithHeader.Header.Hash(),
		"proof", common.Bytes2Hex(proofWithHeader.Proof),
		"tier", proofWithHeader.Tier,
	)

	metrics.ProverReceivedProofCounter.Inc(1)

	evidence, err := s.evidenceAssembler.AssembleEvidence(ctx, proofWithHeader)
	if err != nil {
		return fmt.Errorf("failed to assemble evidence: %w", err)
	}

	input, err := encoding.EncodeEvidence(evidence)
	if err != nil {
		return fmt.Errorf("failed to encode TaikoL1.proveBlock inputs: %w", err)
	}

	// Send the TaikoL1.proveBlock transaction.
	sendTx := func(nonce *big.Int) (*types.Transaction, error) {
		s.mutex.Lock()
		defer s.mutex.Unlock()

		txOpts, err := getProveBlocksTxOpts(ctx, s.rpc.L1, s.rpc.L1ChainID, s.proverPrivKey)
		if err != nil {
			return nil, err
		}

		if s.proveBlockTxGasLimit != nil {
			txOpts.GasLimit = *s.proveBlockTxGasLimit
		}

		if nonce != nil {
			txOpts.Nonce = nonce

			if txOpts, err = rpc.IncreaseGasTipCap(
				ctx,
				s.rpc,
				txOpts,
				s.proverAddress,
				new(big.Int).SetUint64(s.txReplacementTipMultiplier),
				s.proveBlockMaxTxGasTipCap,
			); err != nil {
				return nil, err
			}
		}

		return s.rpc.TaikoL1.ProveBlock(txOpts, proofWithHeader.BlockID.Uint64(), input)
	}

	if err := s.txSender.Send(ctx, proofWithHeader, sendTx); err != nil {
		if errors.Is(err, errUnretryable) {
			return nil
		}

		return err
	}

	metrics.ProverSentProofCounter.Inc(1)
	metrics.ProverLatestProvenBlockIDGauge.Update(proofWithHeader.BlockID.Int64())

	return nil
}

// Producer returns the inner proof producer.
func (s *ProofSubmitter) Producer() proofProducer.ProofProducer {
	return s.proofProducer
}

// Tier returns the proof tier of the current proof submitter.
func (s *ProofSubmitter) Tier() uint16 {
	return s.proofProducer.Tier()
}
