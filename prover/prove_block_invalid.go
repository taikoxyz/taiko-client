package prover

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
	"github.com/taikoxyz/taiko-client/metrics"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	"github.com/taikoxyz/taiko-client/prover/producer"
)

// proveBlockInvalid tries to generate a ZK proof to prove the given
// block is invalid.
func (p *Prover) proveBlockInvalid(
	ctx context.Context,
	event *bindings.TaikoL1ClientBlockProposed,
	hint InvalidTxListReason,
	invalidTxIndex int,
) error {
	// Get the throwaway block from L2 node.
	throwAwayBlock, err := p.getThrowAwayBlock(ctx, event)
	if err != nil {
		return err
	}

	log.Debug("Throwaway block", "header", throwAwayBlock.Header())

	proofOpts := &producer.ProofRequestOptions{
		Height:         throwAwayBlock.Header().Number,
		L2NodeEndpoint: p.cfg.L2Endpoint,
		Retry:          false,
		Param:          p.cfg.ZkEvmRpcdParamsPath,
	}

	if err := p.proofProducer.RequestProof(
		proofOpts, event.Id, throwAwayBlock.Header(), p.proveInvalidProofCh,
	); err != nil {
		return err
	}

	metrics.ProverQueuedProofCounter.Inc(1)
	metrics.ProverQueuedInvalidProofCounter.Inc(1)

	return nil
}

// submitInvalidBlockProof submits the generated ZK proof to TaikoL1 contract.
func (p *Prover) submitInvalidBlockProof(
	ctx context.Context,
	proofWithHeader *producer.ProofWithHeader,
) (err error) {
	log.Info(
		"New invalid block proof",
		"blockID", proofWithHeader.BlockID,
		"hash", proofWithHeader.Header.Hash(),
		"proof", proofWithHeader.ZkProof,
	)
	var (
		blockID = proofWithHeader.BlockID
		header  = proofWithHeader.Header
		zkProof = proofWithHeader.ZkProof
	)

	metrics.ProverReceivedProofCounter.Inc(1)
	metrics.ProverReceivedInvalidProofCounter.Inc(1)

	block, err := p.rpc.L2.BlockByHash(ctx, header.Hash())
	if err != nil {
		return fmt.Errorf("failed to fetch throwaway block: %w", err)
	}

	// Fetch the invalid block metadata
	targetMeta, err := p.rpc.GetBlockMetadataByID(blockID)
	if err != nil {
		return fmt.Errorf("failed to fetch L2 block with given block ID %s: %w", blockID, err)
	}

	// Fetch the transaction receipts in that throwaway block.
	receipts, err := p.rpc.L2.GetThrowawayTransactionReceipts(ctx, header.Hash())
	if err != nil {
		return fmt.Errorf("failed to fetch invalidateBlock transaction receipt: %w", err)
	}

	log.Debug("Throwaway block receipts", "length", receipts.Len())

	receiptRoot, receiptProof, err := generateTrieProof(receipts, 0)
	if err != nil {
		return fmt.Errorf("failed to generate anchor receipt proof: %w", err)
	}

	if receiptRoot != header.ReceiptHash {
		return fmt.Errorf("receipt root mismatch, receiptRoot: %s, block.ReceiptHash: %s", receiptRoot, header.ReceiptHash)
	}

	txListBytes, err := rlp.EncodeToBytes(block.Transactions())
	if err != nil {
		return fmt.Errorf("failed to encode throwaway block transactions: %w", err)
	}

	evidence := &encoding.TaikoL1Evidence{
		Meta: bindings.LibDataBlockMetadata{
			Id:          targetMeta.Id,
			L1Height:    targetMeta.L1Height,
			L1Hash:      targetMeta.L1Hash,
			Beneficiary: header.Coinbase,
			GasLimit:    header.GasLimit - p.anchorGasLimit,
			Timestamp:   header.Time,
			TxListHash:  crypto.Keccak256Hash(txListBytes),
			MixHash:     header.MixDigest,
			ExtraData:   header.Extra,
		},
		Header: *encoding.FromGethHeader(header),
		Prover: crypto.PubkeyToAddress(p.cfg.L1ProverPrivKey.PublicKey),
		Proofs: [][]byte{zkProof, receiptProof},
	}

	input, err := encoding.EncodeProveBlockInvalidInput(evidence, targetMeta, receipts[0])
	if err != nil {
		return err
	}

	txOpts, err := p.getProveBlocksTxOpts(ctx)
	if err != nil {
		return err
	}

	tx, err := p.rpc.TaikoL1.ProveBlockInvalid(txOpts, blockID, input)
	if err != nil {
		return fmt.Errorf("failed to send TaikoL1.proveBlockInvalid transaction: %w", err)
	}

	if _, err := rpc.WaitReceipt(ctx, p.rpc.L1, tx); err != nil {
		return fmt.Errorf("failed to wait till transaction executed: %w", err)
	}

	log.Info(
		"❎ New invalid block proved",
		"blockID", proofWithHeader.BlockID,
		"height", block.Number(),
		"hash", header.Hash(),
	)

	metrics.ProverSentProofCounter.Inc(1)
	metrics.ProverSentInvalidProofCounter.Inc(1)

	return nil
}

// getThrowAwayBlock keeps waitting till the throwaway block inserted into the L2 chain.
func (p *Prover) getThrowAwayBlock(
	ctx context.Context,
	event *bindings.TaikoL1ClientBlockProposed,
) (*types.Block, error) {
	l1Origin, err := p.rpc.WaitL1Origin(ctx, event.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch l1Origin, blockID: %d, err: %w", event.Id, err)
	}

	if !l1Origin.Throwaway {
		return nil, fmt.Errorf("invalid L1origin found: %+v", l1Origin)
	}

	return p.rpc.L2.BlockByHash(ctx, l1Origin.L2BlockHash)
}
