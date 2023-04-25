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
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
	"github.com/taikoxyz/taiko-client/metrics"
	"github.com/taikoxyz/taiko-client/oracle"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	anchorTxValidator "github.com/taikoxyz/taiko-client/prover/anchor_tx_validator"
	proofProducer "github.com/taikoxyz/taiko-client/prover/proof_producer"
)

var _ ProofSubmitter = (*OracleProofSubmitter)(nil)

// OracleProofSubmitter is responsible requesting zk proofs for the given valid L2
// blocks, and submitting the generated proofs to the TaikoL1 smart contract.
type OracleProofSubmitter struct {
	rpc               *rpc.Client
	resultCh          chan *proofProducer.ProofWithHeader
	anchorTxValidator *anchorTxValidator.AnchorTxValidator
	proverPrivKey     *ecdsa.PrivateKey
	proverAddress     common.Address
	mutex             *sync.Mutex
}

// NewOracleProofSubmitter creates a new OracleProofSubmitter instance.
func NewOracleProofSubmitter(
	rpc *rpc.Client,
	resultCh chan *proofProducer.ProofWithHeader,
	taikoL2Address common.Address,
	proverPrivKey *ecdsa.PrivateKey,
	mutex *sync.Mutex,
) (*OracleProofSubmitter, error) {
	anchorValidator, err := anchorTxValidator.New(taikoL2Address, rpc.L2ChainID, rpc)
	if err != nil {
		return nil, err
	}

	return &OracleProofSubmitter{
		rpc:               rpc,
		resultCh:          resultCh,
		anchorTxValidator: anchorValidator,
		proverPrivKey:     proverPrivKey,
		proverAddress:     crypto.PubkeyToAddress(proverPrivKey.PublicKey),
		mutex:             mutex,
	}, nil
}

// RequestProof implements the ProofSubmitter interface.
func (s *OracleProofSubmitter) RequestProof(ctx context.Context, event *bindings.TaikoL1ClientBlockProposed) error {
	l1Origin, err := s.rpc.WaitL1Origin(ctx, event.Id)
	if err != nil {
		return fmt.Errorf("failed to fetch l1Origin, blockID: %d, err: %w", event.Id, err)
	}

	// Get the header of the block to prove from L2 execution engine.
	header, err := s.rpc.L2.HeaderByHash(ctx, l1Origin.L2BlockHash)
	if err != nil {
		return err
	}

	metrics.ProverQueuedProofCounter.Inc(1)
	metrics.ProverQueuedValidProofCounter.Inc(1)

	// "proof" for Oracle Proof is actually signed in SubmitProof, since
	// we already have the evidence and to not duplicate RPC calls.
	// "Proof" and "Degree" fields are not used here.
	s.resultCh <- &proofProducer.ProofWithHeader{
		Header:  header,
		BlockID: event.Id,
		Meta:    &event.Meta,
	}

	return nil
}

// SubmitProof implements the ProofSubmitter interface.
func (s *OracleProofSubmitter) SubmitProof(
	ctx context.Context,
	proofWithHeader *proofProducer.ProofWithHeader,
) (err error) {
	log.Info(
		"New oracle block proof",
		"blockID", proofWithHeader.BlockID,
		"beneficiary", proofWithHeader.Meta.Beneficiary,
		"hash", proofWithHeader.Header.Hash(),
	)
	var (
		blockID = proofWithHeader.BlockID
		header  = proofWithHeader.Header
	)

	metrics.ProverReceivedProofCounter.Inc(1)
	metrics.ProverReceivedValidProofCounter.Inc(1)

	// Get the corresponding L2 block.
	block, err := s.rpc.L2.BlockByHash(ctx, header.Hash())
	if err != nil {
		return fmt.Errorf("failed to get L2 block with given hash %s: %w", header.Hash(), err)
	}

	log.Debug(
		"Get the L2 block to prove",
		"blockID", blockID,
		"hash", block.Hash(),
		"root", header.Root.String(),
		"transactions", len(block.Transactions()),
	)

	if block.Transactions().Len() == 0 {
		return fmt.Errorf("invalid block without anchor transaction, blockID %s", blockID)
	}

	// Validate TaikoL2.anchor transaction inside the L2 block.
	anchorTx := block.Transactions()[0]
	if err := s.anchorTxValidator.ValidateAnchorTx(ctx, anchorTx); err != nil {
		return fmt.Errorf("invalid anchor transaction: %w", err)
	}

	// Get and validate this anchor transaction's receipt.
	_, err = s.anchorTxValidator.GetAndValidateAnchorTxReceipt(ctx, anchorTx)
	if err != nil {
		return fmt.Errorf("failed to fetch anchor transaction receipt: %w", err)
	}

	signalRoot, err := s.anchorTxValidator.GetAnchoredSignalRoot(ctx, anchorTx)
	if err != nil {
		return err
	}

	parent, err := s.rpc.L2.BlockByHash(ctx, block.ParentHash())
	if err != nil {
		return err
	}

	blockInfo, err := s.rpc.TaikoL1.GetBlock(nil, blockID)
	if err != nil {
		return err
	}

	// signature should be done with proof set to nil, verifierID set to 0,
	// and prover set to 0 address.
	evidence := &encoding.TaikoL1Evidence{
		MetaHash:      blockInfo.MetaHash,
		ParentHash:    block.ParentHash(),
		BlockHash:     block.Hash(),
		SignalRoot:    signalRoot,
		Graffiti:      [32]byte{},
		Prover:        common.HexToAddress("0x0000000000000000000000000000000000000000"),
		ParentGasUsed: uint32(parent.GasUsed()),
		GasUsed:       uint32(block.GasUsed()),
		VerifierId:    0,
		Proof:         nil,
	}

	_, _, err = oracle.HashSignAndSetEvidenceForOracleProof(evidence, s.proverPrivKey)
	if err != nil {
		return fmt.Errorf("failed to sign evidence: %w", err)
	}

	input, err := encoding.EncodeProveBlockInput(evidence)
	if err != nil {
		return fmt.Errorf("failed to encode TaikoL1.proveBlock inputs: %w", err)
	}

	// Send the TaikoL1.proveBlock transaction.
	txOpts, err := getProveBlocksTxOpts(ctx, s.rpc.L1, s.rpc.L1ChainID, s.proverPrivKey)
	if err != nil {
		return err
	}

	sendTx := func() (*types.Transaction, error) {
		s.mutex.Lock()
		defer s.mutex.Unlock()

		return s.rpc.TaikoL1.ProveBlock(txOpts, blockID, input)
	}

	if err := sendTxWithBackoff(ctx, s.rpc, blockID, sendTx); err != nil {
		if errors.Is(err, errUnretryable) {
			return nil
		}

		return err
	}

	log.Info(
		"✅ Oracle block proved",
		"blockID", proofWithHeader.BlockID,
		"hash", block.Hash(), "height", block.Number(),
		"transactions", block.Transactions().Len(),
	)

	metrics.ProverSentProofCounter.Inc(1)
	metrics.ProverSentValidProofCounter.Inc(1)
	metrics.ProverLatestProvenBlockIDGauge.Update(proofWithHeader.BlockID.Int64())

	return nil
}
