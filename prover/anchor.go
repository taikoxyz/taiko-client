package prover

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/taikochain/taiko-client/bindings/encoding"
)

// Address of the golden touch account.
var (
	goldenTouchAddress = common.HexToAddress("0x0000777735367b36bC9B61C50022d9D0700dB4Ec")
)

// validateAnchorTx checks whether the given transaction is a successfully
// executed `TaikoL2.anchor` transaction.
func (p *Prover) validateAnchorTx(ctx context.Context, tx *types.Transaction) error {
	if tx.To() == nil || *tx.To() != p.cfg.TaikoL2Address {
		return fmt.Errorf("invalid TaikoL2.anchor to: %s", tx.To())
	}

	sender, err := types.LatestSignerForChainID(p.rpc.L2ChainID).Sender(tx)
	if err != nil {
		return fmt.Errorf("failed to get TaikoL2.anchor transaction sender: %w", err)
	}

	if sender != goldenTouchAddress {
		return fmt.Errorf("invalid TaikoL2.anchor transaction sender: %s", sender)
	}

	method, err := encoding.TaikoL2ABI.MethodById(tx.Data())
	if err != nil || method.Name != "anchor" {
		return fmt.Errorf("invalid method method, err: %w, methodName: %s", err, method.Name)
	}

	receipt, err := p.rpc.L2.TransactionReceipt(ctx, tx.Hash())
	if err != nil {
		return fmt.Errorf("failed to get TaikoL2.anchor receipt, err: %w", err)
	}

	if receipt.Status != types.ReceiptStatusSuccessful {
		return fmt.Errorf("invalid TaikoL2.anchor receipt status: %d", receipt.Status)
	}

	if len(receipt.Logs) == 0 {
		return fmt.Errorf("HeaderExchanged event not found in TaikoL2.anchor receipt")
	}

	return nil
}
