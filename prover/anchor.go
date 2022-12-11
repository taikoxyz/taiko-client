package prover

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
)

// validateAnchorTx checks whether the given transaction is a valid `TaikoL2.anchor` transaction.
func (p *Prover) validateAnchorTx(ctx context.Context, tx *types.Transaction) error {
	if tx.To() == nil || *tx.To() != p.cfg.TaikoL2Address {
		return fmt.Errorf("invalid TaikoL2.anchor transaction to: %s, want: %s", tx.To(), p.cfg.TaikoL2Address)
	}

	sender, err := types.LatestSignerForChainID(p.rpc.L2ChainID).Sender(tx)
	if err != nil {
		return fmt.Errorf("failed to get TaikoL2.anchor transaction sender: %w", err)
	}

	if sender != bindings.GoldenTouchAddress {
		return fmt.Errorf("invalid TaikoL2.anchor transaction sender: %s", sender)
	}

	method, err := encoding.TaikoL2ABI.MethodById(tx.Data())
	if err != nil || method.Name != "anchor" {
		return fmt.Errorf("invalid TaikoL2.anchor transaction selector, err: %w", err)
	}

	return nil
}

// getAndValidateAnchorTxReceipt gets and validates the `TaikoL2.anchor` transaction's receipt.
func (p *Prover) getAndValidateAnchorTxReceipt(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	receipt, err := p.rpc.L2.TransactionReceipt(ctx, tx.Hash())
	if err != nil {
		return nil, fmt.Errorf("failed to get TaikoL2.anchor transaction receipt, err: %w", err)
	}

	if receipt.Status != types.ReceiptStatusSuccessful {
		return nil, fmt.Errorf("invalid TaikoL2.anchor transaction receipt status: %d", receipt.Status)
	}

	if len(receipt.Logs) == 0 {
		return nil, fmt.Errorf("no event found in TaikoL2.anchor transaction receipt")
	}

	return receipt, nil
}
