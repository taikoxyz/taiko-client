package prover

import (
	"context"
	"fmt"

	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
	"github.com/taikoxyz/taiko-client/metrics"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	"github.com/taikoxyz/taiko-client/prover/producer"
)

// proveBlockValid tries to generate a ZK proof to prove the given
// block is valid.
func (p *Prover) proveBlockValid(ctx context.Context, event *bindings.TaikoL1ClientBlockProposed) error {
	l1Origin, err := p.rpc.WaitL1Origin(ctx, event.Id)
	if err != nil {
		return fmt.Errorf("failed to fetch l1Origin, blockID: %d, err: %w", event.Id, err)
	}

	// This should not be reached, only check for safety.
	if l1Origin.Throwaway {
		log.Error("Get a block metadata with invalid transaction list", "l1Origin", l1Origin)
		return nil
	}

	// Get the header of the block to prove from L2 execution engine.
	header, err := p.rpc.L2.HeaderByHash(ctx, l1Origin.L2BlockHash)
	if err != nil {
		return err
	}

	// Request proof.
	opts := &producer.ProofRequestOptions{
		Height:         header.Number,
		L2NodeEndpoint: p.cfg.L2Endpoint,
		Retry:          false,
		Param:          p.cfg.ZkEvmRpcdParamsPath,
	}

	if err := p.proofProducer.RequestProof(opts, event.Id, &event.Meta, header, p.proveValidProofCh); err != nil {
		return err
	}

	metrics.ProverQueuedProofCounter.Inc(1)
	metrics.ProverQueuedValidProofCounter.Inc(1)
	p.l1Current = event.Raw.BlockNumber
	p.lastHandledBlockID = event.Id.Uint64()

	return nil
}

// submitValidBlockProof submits the generated ZK proof to TaikoL1 contract.
func (p *Prover) submitValidBlockProof(ctx context.Context, proofWithHeader *producer.ProofWithHeader) (err error) {
	log.Info(
		"New valid block proof",
		"blockID", proofWithHeader.BlockID,
		"meta", proofWithHeader.Meta,
		"hash", proofWithHeader.Header.Hash(),
		"proof", proofWithHeader.ZkProof,
	)
	var (
		blockID = proofWithHeader.BlockID
		header  = proofWithHeader.Header
		zkProof = proofWithHeader.ZkProof
	)

	metrics.ProverReceivedProofCounter.Inc(1)
	metrics.ProverReceivedValidProofCounter.Inc(1)

	// Get the corresponding L2 block.
	block, err := p.rpc.L2.BlockByHash(ctx, header.Hash())
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

	// Validate TaikoL2.anchor transaction inside the L2 block.
	anchorTx := block.Transactions()[0]
	if err := p.validateAnchorTx(ctx, anchorTx); err != nil {
		return fmt.Errorf("invalid anchor transaction: %w", err)
	}

	// Get and validate this anchor transaction's receipt.
	anchorTxReceipt, err := p.getAndValidateAnchorTxReceipt(ctx, anchorTx)
	if err != nil {
		return fmt.Errorf("failed to fetch anchor transaction receipt: %w", err)
	}

	// Generate the merkel proof (whose root is block's txRoot) of this anchor transaction.
	txRoot, anchorTxProof, err := generateTrieProof(block.Transactions(), 0)
	if err != nil {
		return fmt.Errorf("failed to generate anchor transaction proof: %w", err)
	}

	// Generate the merkel proof (whose root is block's receiptRoot) of this anchor transaction's receipt.
	receipts, err := rpc.GetReceiptsByBlock(ctx, p.rpc.L2RawRPC, block)
	if err != nil {
		return fmt.Errorf("failed to fetch block receipts: %w", err)
	}
	receiptRoot, anchorReceiptProof, err := generateTrieProof(receipts, 0)
	if err != nil {
		return fmt.Errorf("failed to generate anchor receipt proof: %w", err)
	}

	// Double check the calculated roots.
	if txRoot != block.TxHash() || receiptRoot != block.ReceiptHash() {
		return fmt.Errorf(
			"txHash or receiptHash mismatch, txRoot: %s, header.TxHash: %s, receiptRoot: %s, header.ReceiptHash: %s",
			txRoot, header.TxHash, receiptRoot, header.ReceiptHash,
		)
	}

	// Assemble the TaikoL1.proveBlock transaction inputs.
	proofs := [][]byte{}
	for i := 0; i < int(p.protocolConstants.ZKProofsPerBlock.Uint64()); i++ {
		proofs = append(proofs, zkProof)
	}
	proofs = append(proofs, [][]byte{anchorTxProof, anchorReceiptProof}...)

	evidence := &encoding.TaikoL1Evidence{
		Meta:   *proofWithHeader.Meta,
		Header: *encoding.FromGethHeader(header),
		Prover: crypto.PubkeyToAddress(p.cfg.L1ProverPrivKey.PublicKey),
		Proofs: proofs,
	}

	input, err := encoding.EncodeProveBlockInput(evidence, anchorTx, anchorTxReceipt)
	if err != nil {
		return fmt.Errorf("failed to encode TaikoL1.proveBlock inputs: %w", err)
	}

	// Send the TaikoL1.proveBlock transaction.
	txOpts, err := p.getProveBlocksTxOpts(ctx, p.rpc.L1)
	if err != nil {
		return err
	}

	var isUnretryableError bool
	if err := backoff.Retry(func() error {
		tx, err := p.rpc.TaikoL1.ProveBlock(txOpts, blockID, input)
		if err != nil {
			if IsSubmitProofTxErrorRetryable(err) {
				log.Warn("Retry sending TaikoL1.proveBlock transaction", "error", err)
				return err
			}

			isUnretryableError = true
			return nil
		}

		if _, err := rpc.WaitReceipt(ctx, p.rpc.L1, tx); err != nil {
			log.Warn("Failed to wait till transaction executed", "txHash", tx.Hash(), "error", err)
			return err
		}

		return nil
	}, backoff.NewExponentialBackOff()); err != nil {
		return fmt.Errorf("failed to send TaikoL1.proveBlock transaction: %w", err)
	}

	if isUnretryableError {
		return nil
	}

	log.Info(
		"✅ Valid block proved",
		"blockID", proofWithHeader.BlockID,
		"hash", block.Hash(), "height", block.Number(),
		"transactions", block.Transactions().Len(),
	)

	metrics.ProverSentProofCounter.Inc(1)
	metrics.ProverSentValidProofCounter.Inc(1)

	return nil
}
