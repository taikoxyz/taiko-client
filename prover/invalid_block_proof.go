package prover

import (
	"context"
	"fmt"

	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
	"github.com/taikoxyz/taiko-client/metrics"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	txListValidator "github.com/taikoxyz/taiko-client/pkg/tx_list_validator"
	"github.com/taikoxyz/taiko-client/prover/producer"
)

// proveBlockInvalid tries to generate a ZK proof to prove the given
// block is invalid.
func (p *Prover) proveBlockInvalid(
	ctx context.Context,
	event *bindings.TaikoL1ClientBlockProposed,
	hint txListValidator.InvalidTxListReason,
	invalidTxIndex int,
) error {
	// Get the throwaway block from L2 execution engine.
	throwAwayBlock, err := p.getThrowAwayBlock(ctx, event)
	if err != nil {
		return err
	}

	log.Debug("Throwaway block", "header", throwAwayBlock.Header())

	// Request proof.
	proofOpts := &producer.ProofRequestOptions{
		Height:     throwAwayBlock.Header().Number,
		L2Endpoint: p.cfg.L2Endpoint,
		Retry:      false,
		Param:      p.cfg.ZkEvmRpcdParamsPath,
	}

	if err := p.proofProducer.RequestProof(
		proofOpts, event.Id, &event.Meta, throwAwayBlock.Header(), p.proveInvalidProofCh,
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
		"beneficiary", proofWithHeader.Meta.Beneficiary,
		"hash", proofWithHeader.Header.Hash(),
		"proof", proofWithHeader.ZkProof,
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
	block, err := p.rpc.L2.BlockByHash(ctx, header.Hash())
	if err != nil {
		return fmt.Errorf("failed to get the throwaway block (id: %d): %w", blockID, err)
	}

	if block.Transactions().Len() == 0 {
		return fmt.Errorf("invalid throwaway block without any transaction, blockID %s", blockID)
	}

	// Fetch all receipts inside that L2 throwaway block.
	receipts, err := p.rpc.L2.GetThrowawayTransactionReceipts(ctx, header.Hash())
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
	proofs := [][]byte{}
	for i := 0; i < int(p.protocolConstants.ZKProofsPerBlock.Uint64()); i++ {
		proofs = append(proofs, zkProof)
	}
	proofs = append(proofs, receiptProof)

	txListBytes, err := rlp.EncodeToBytes(block.Transactions())
	if err != nil {
		return fmt.Errorf("failed to encode throwaway block transactions: %w", err)
	}

	evidence := &encoding.TaikoL1Evidence{
		Meta: bindings.LibDataBlockMetadata{
			Id:          meta.Id,
			L1Height:    meta.L1Height,
			L1Hash:      meta.L1Hash,
			Beneficiary: header.Coinbase,
			GasLimit:    header.GasLimit - p.protocolConstants.AnchorTxGasLimit.Uint64(),
			Timestamp:   header.Time,
			TxListHash:  crypto.Keccak256Hash(txListBytes),
			MixHash:     header.MixDigest,
			ExtraData:   header.Extra,
		},
		Header: *encoding.FromGethHeader(header),
		Prover: crypto.PubkeyToAddress(p.cfg.L1ProverPrivKey.PublicKey),
		Proofs: proofs,
	}

	input, err := encoding.EncodeProveBlockInvalidInput(evidence, meta, receipts[0])
	if err != nil {
		return err
	}

	// Send the TaikoL1.proveBlockInvalid transaction.
	txOpts, err := p.getProveBlocksTxOpts(ctx, p.rpc.L1)
	if err != nil {
		return err
	}

	var isUnretryableError bool
	if err := backoff.Retry(func() error {
		sendTx := func() (*types.Transaction, error) {
			p.submitProofTxMutex.Lock()
			defer p.submitProofTxMutex.Unlock()

			return p.rpc.TaikoL1.ProveBlockInvalid(txOpts, blockID, input)
		}

		tx, err := sendTx()
		if err != nil {
			if isSubmitProofTxErrorRetryable(err, blockID) {
				log.Info("Retry sending TaikoL1.proveBlockInvalid transaction", "reason", err)
				return err
			}

			isUnretryableError = true
			return nil
		}

		if _, err := rpc.WaitReceipt(ctx, p.rpc.L1, tx); err != nil {
			log.Warn("Failed to wait till transaction executed", "blockID", blockID, "txHash", tx.Hash(), "error", err)
			return err
		}

		return nil
	}, backoff.NewExponentialBackOff()); err != nil {
		return fmt.Errorf("failed to send TaikoL1.proveBlockInvalid transaction: %w", err)
	}

	if isUnretryableError {
		return nil
	}

	log.Info(
		"❎ Invalid block proved",
		"blockID", proofWithHeader.BlockID,
		"height", block.Number(),
		"hash", header.Hash(),
	)

	metrics.ProverSentProofCounter.Inc(1)
	metrics.ProverSentInvalidProofCounter.Inc(1)

	return nil
}

// getThrowAwayBlock keeps waiting till the throwaway block inserted into the L2 chain.
func (p *Prover) getThrowAwayBlock(
	ctx context.Context,
	event *bindings.TaikoL1ClientBlockProposed,
) (*types.Block, error) {
	l1Origin, err := p.rpc.WaitL1Origin(ctx, event.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch L1origin, blockID: %d, err: %w", event.Id, err)
	}

	if !l1Origin.Throwaway {
		return nil, fmt.Errorf("invalid throwaway block's L1origin found, blockID: %d: %+v", event.Id, l1Origin)
	}

	return p.rpc.L2.BlockByHash(ctx, l1Origin.L2BlockHash)
}
