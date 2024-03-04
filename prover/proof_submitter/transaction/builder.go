package transaction

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"

	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
)

// TxBuilder will build a transaction with the given nonce.
type TxBuilder func() (*types.Transaction, error)

// ProveBlockTxBuilder is responsible for building ProveBlock transactions.
type ProveBlockTxBuilder struct {
	rpc              *rpc.Client
	proverPrivateKey *ecdsa.PrivateKey
	proverAddress    common.Address
	mutex            *sync.Mutex
}

// NewProveBlockTxBuilder creates a new ProveBlockTxBuilder instance.
func NewProveBlockTxBuilder(
	rpc *rpc.Client,
	proverPrivateKey *ecdsa.PrivateKey,
) *ProveBlockTxBuilder {
	return &ProveBlockTxBuilder{
		rpc:              rpc,
		proverPrivateKey: proverPrivateKey,
		proverAddress:    crypto.PubkeyToAddress(proverPrivateKey.PublicKey),
		mutex:            new(sync.Mutex),
	}
}

// Build creates a new TaikoL1.ProveBlock transaction with the given nonce.
func (a *ProveBlockTxBuilder) Build(
	ctx context.Context,
	blockID *big.Int,
	meta *bindings.TaikoDataBlockMetadata,
	transition *bindings.TaikoDataTransition,
	tierProof *bindings.TaikoDataTierProof,
	txOpts *bind.TransactOpts,
	guardian bool,
) TxBuilder {
	return func() (*types.Transaction, error) {
		a.mutex.Lock()
		defer a.mutex.Unlock()

		var (
			tx  *types.Transaction
			err error
		)

		log.Info(
			"Build proof submission transaction",
			"blockID", blockID,
			"gasLimit", txOpts.GasLimit,
			"nonce", txOpts.Nonce,
			"gasTipCap", txOpts.GasTipCap,
			"gasFeeCap", txOpts.GasFeeCap,
			"guardian", guardian,
		)

		if !guardian {
			input, err := encoding.EncodeProveBlockInput(meta, transition, tierProof)
			if err != nil {
				return nil, err
			}
			if tx, err = a.rpc.TaikoL1.ProveBlock(txOpts, blockID.Uint64(), input); err != nil {
				return nil, err
			}
		} else {
			if tx, err = a.rpc.GuardianProver.Approve(txOpts, *meta, *transition, *tierProof); err != nil {
				return nil, err
			}
		}

		return tx, nil
	}
}
