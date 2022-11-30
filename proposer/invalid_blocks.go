package proposer

import (
	"context"
	"fmt"
	"math/big"
	"math/rand"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/taikoxyz/taiko-client/testutils"
)

var (
	globalEpoch uint64 = 0
)

// ProposeInvalidBlocksOp tries to propose invalid blocks to TaikoL1 contract
// every `interval` normal propose operations.
func (p *Proposer) ProposeInvalidBlocksOp(ctx context.Context, interval uint64) error {
	globalEpoch += 1

	if globalEpoch%interval != 0 {
		return nil
	}

	log.Info("👻 Propose invalid transactions list bytes", "epoch", globalEpoch)

	if err := p.ProposeInvalidTxListBytes(ctx); err != nil {
		return fmt.Errorf("failed to propose invalid transaction list bytes: %w", err)
	}

	log.Info("👻 Propose transactions list including invalid transaction", "epoch", globalEpoch)

	if err := p.proposeTxListIncludingInvalidTx(ctx); err != nil {
		return fmt.Errorf("failed to propose transactions list including invalid transaction: %w", err)
	}

	return nil
}

// ProposeInvalidTxListBytes commits and proposes an invalid transaction list
// bytes to TaikoL1 contract.
func (p *Proposer) ProposeInvalidTxListBytes(ctx context.Context) error {
	invalidTxListBytes := testutils.RandomBytes(256)
	meta, commitTx, err := p.CommitTxList(
		ctx,
		invalidTxListBytes,
		uint64(rand.Int63n(int64(p.poolContentSplitter.maxGasPerBlock))),
		0,
	)
	if err != nil {
		return err
	}

	if p.AfterCommitHook != nil {
		if err := p.AfterCommitHook(); err != nil {
			log.Error("Run AfterCommitHook error", "error", err)
		}
	}

	return p.ProposeTxList(ctx, &commitTxListRes{meta, commitTx, invalidTxListBytes, 1})
}

// proposeTxListIncludingInvalidTx commits and proposes a validly encoded
// transaction list which including an invalid transaction.
func (p *Proposer) proposeTxListIncludingInvalidTx(ctx context.Context) error {
	invalidTx, err := p.generateInvalidTransaction(ctx)
	if err != nil {
		return err
	}

	txListBytes, err := rlp.EncodeToBytes(types.Transactions{invalidTx})
	if err != nil {
		return err
	}

	meta, commitTx, err := p.CommitTxList(ctx, txListBytes, invalidTx.Gas(), 0)
	if err != nil {
		return err
	}

	if p.AfterCommitHook != nil {
		if err := p.AfterCommitHook(); err != nil {
			log.Error("Run AfterCommitHook error", "error", err)
		}
	}

	return p.ProposeTxList(ctx, &commitTxListRes{meta, commitTx, txListBytes, 1})
}

// generateInvalidTransaction creates a transaction with an invalid nonce to
// current L2 world state.
func (p *Proposer) generateInvalidTransaction(ctx context.Context) (*types.Transaction, error) {
	opts, err := bind.NewKeyedTransactorWithChainID(p.l1ProposerPrivKey, p.rpc.L2ChainID)
	if err != nil {
		return nil, err
	}

	nonce, err := p.rpc.L2.PendingNonceAt(ctx, crypto.PubkeyToAddress(p.l1ProposerPrivKey.PublicKey))
	if err != nil {
		return nil, err
	}

	opts.GasLimit = 300000
	opts.NoSend = true
	opts.Nonce = new(big.Int).SetUint64(nonce + 1024)

	return p.rpc.TaikoL2.Anchor(opts, common.Big0, common.BytesToHash(testutils.RandomBytes(32)))
}
