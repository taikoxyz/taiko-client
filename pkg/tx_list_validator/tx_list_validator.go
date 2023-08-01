package tx_list_validator

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
)

// InvalidTxListReason represents a reason why a transactions list is invalid.
type InvalidTxListReason uint8

// All invalid transactions list reasons.
const (
	HintNone InvalidTxListReason = iota
	HintOK
)

type TxListValidator struct {
	blockMaxGasLimit        uint64
	maxTransactionsPerBlock uint64
	maxBytesPerTxList       uint64
	chainID                 *big.Int
}

// NewTxListValidator creates a new TxListValidator instance based on giving configurations.
func NewTxListValidator(
	blockMaxGasLimit uint64,
	maxTransactionsPerBlock uint64,
	maxBytesPerTxList uint64,
	chainID *big.Int,
) *TxListValidator {
	return &TxListValidator{
		blockMaxGasLimit:        blockMaxGasLimit,
		maxTransactionsPerBlock: maxTransactionsPerBlock,
		maxBytesPerTxList:       maxBytesPerTxList,
		chainID:                 chainID,
	}
}

// ValidateTxList checks whether the transactions list in the TaikoL1.proposeBlock transaction's
// input data is valid.
func (v *TxListValidator) ValidateTxList(
	blockID *big.Int,
	proposeBlockTxInput []byte,
) (txListBytes []byte, hint InvalidTxListReason, txIdx int, err error) {
	if txListBytes, err = encoding.UnpackTxListBytes(proposeBlockTxInput); err != nil {
		return nil, HintNone, 0, err
	}

	if len(txListBytes) == 0 {
		return txListBytes, HintOK, 0, nil
	}

	hint, txIdx = v.isTxListValid(blockID, txListBytes)

	return txListBytes, hint, txIdx, nil
}

// isTxListValid checks whether the transaction list is valid.
func (v *TxListValidator) isTxListValid(blockID *big.Int, txListBytes []byte) (hint InvalidTxListReason, txIdx int) {
	if len(txListBytes) > int(v.maxBytesPerTxList) {
		log.Info("Transactions list binary too large", "length", len(txListBytes), "blockID", blockID)
		return HintNone, 0
	}

	var txs types.Transactions
	if err := rlp.DecodeBytes(txListBytes, &txs); err != nil {
		log.Info("Failed to decode transactions list bytes", "blockID", blockID, "error", err)
		return HintNone, 0
	}

	log.Debug("Transactions list decoded", "blockID", blockID, "length", len(txs))

	if txs.Len() > int(v.maxTransactionsPerBlock) {
		log.Info("Too many transactions", "blockID", blockID, "count", txs.Len())
		return HintNone, 0
	}

	log.Info("Transaction list is valid", "blockID", blockID)
	return HintOK, 0
}
