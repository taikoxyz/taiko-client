package prover

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/taikochain/client-mono/bindings"
	"github.com/taikochain/client-mono/bindings/encoding"
	"github.com/taikochain/client-mono/prover/producer"
	"github.com/taikochain/client-mono/util"
)

// proveBlockValid tries to generate a ZK proof to prove the given
// block is valid.
func (p *Prover) proveBlockValid(ctx context.Context, event *bindings.TaikoL1ClientBlockProposed) error {
	l1Origin, err := p.waitForL1Origin(ctx, event.Id)
	if err != nil {
		return fmt.Errorf("failed to fetch l1Origin, blockID: %d, err: %w", event.Id, err)
	}

	if l1Origin.Throwaway {
		log.Crit("Get a block metadata with invalid transaction list", "l1Origin", l1Origin)
	}

	header, err := p.l2RPC.HeaderByHash(ctx, l1Origin.L2BlockHash)
	if err != nil {
		return err
	}

	opts := &producer.ProofRequestOptions{
		Height:         header.Number,
		L2NodeEndpoint: p.cfg.L2Endpoint,
		Retry:          false,
		Param:          p.cfg.ZkEvmRpcdParamsPath,
	}

	return p.proofProducer.RequestProof(opts, event.Id, header, p.proveValidResultCh)
}

// submitValidBlockProof submits the generated ZK proof to TaikoL1 contract.
func (p *Prover) submitValidBlockProof(proofWithHeader *producer.ProofWithHeader) (err error) {
	log.Info(
		"New valid block proof",
		"blockID", proofWithHeader.BlockID,
		"hash", proofWithHeader.Header.Hash(),
		"proof", proofWithHeader.ZkProof,
	)
	var (
		blockID = proofWithHeader.BlockID
		header  = proofWithHeader.Header
		zkProof = proofWithHeader.ZkProof
	)

	meta, err := p.getBlockMetadataByID(blockID)
	if err != nil {
		return fmt.Errorf("failed to fetch L2 block with given block ID %s: %w", blockID, err)
	}

	txOpts, err := p.getProveBlockTxOpts(p.ctx)
	if err != nil {
		return err
	}

	block, err := p.l2RPC.BlockByHash(p.ctx, header.Hash())
	if err != nil {
		return fmt.Errorf("failed to get L2 block with given hash %s: %w", header.Hash(), err)
	}

	log.Info(
		"Get the L2 block to prove",
		"blockID", blockID,
		"hash", block.Hash(),
		"transactions", len(block.Transactions()),
	)

	log.Info("StateRoot", "root", header.Root.String(), "zkProof", zkProof, "block", header)

	anchorTx := block.Transactions()[0]

	if err := p.validateAnchorTx(p.ctx, anchorTx); err != nil {
		return fmt.Errorf("invalid anchor transaction: %w", err)
	}

	anchorTxReceipt, err := p.l2RPC.TransactionReceipt(p.ctx, anchorTx.Hash())
	if err != nil {
		return fmt.Errorf("failed to fetch anchor transaction receipt: %w", err)
	}

	txRoot, anchorTxProof, err := generateTrieProof(block.Transactions(), 0)
	if err != nil {
		return fmt.Errorf("failed to generate anchor transaction proof: %w", err)
	}

	receipts, err := util.GetReceiptsByBlock(p.ctx, p.l2RPC, block)
	if err != nil {
		return fmt.Errorf("failed to fetch block receipts: %w", err)
	}

	receiptRoot, anchorReceiptProof, err := generateTrieProof(receipts, 0)
	if err != nil {
		return fmt.Errorf("failed to generate anchor receipt proof: %w", err)
	}

	if txRoot != block.TxHash() || receiptRoot != block.ReceiptHash() {
		return fmt.Errorf(
			"txHash or receiptHash mismatch, txRoot: %s, header.TxHash: %s, receiptRoot: %s, header.ReceiptHash: %s",
			txRoot, header.TxHash, receiptRoot, header.ReceiptHash,
		)
	}

	evidence := &encoding.TaikoL1Evidence{
		Meta:   *meta,
		Header: *encoding.FromGethHeader(header),
		Prover: crypto.PubkeyToAddress(p.cfg.L1ProverPrivKey.PublicKey),
		Proofs: [][]byte{zkProof, anchorTxProof, anchorReceiptProof},
	}

	input, err := encoding.EncodeProveBlockInput(evidence, anchorTx, anchorTxReceipt)
	if err != nil {
		return fmt.Errorf("failed to encode TaikoL1.proveBlock inputs: %w", err)
	}

	tx, err := p.taikoL1.ProveBlock(txOpts, blockID, input)
	if err != nil {
		return fmt.Errorf("failed to send TaikoL1.proveBlock transaction: %w", err)
	}

	if _, err := util.WaitForTx(p.ctx, p.l1RPC, tx); err != nil {
		return fmt.Errorf("failed to wait till transaction executed: %w", err)
	}

	log.Info(
		"✅ New valid block proved",
		"blockID", proofWithHeader.BlockID,
		"hash", block.Hash(), "height", block.Number(),
		"transactions", block.Transactions().Len(),
	)

	return nil
}
