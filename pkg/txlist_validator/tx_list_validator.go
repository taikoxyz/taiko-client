package txlistvalidator

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/taikoxyz/taiko-client/internal/utils"
)

// TxListValidator is responsible for validating the transactions list in a TaikoL1.proposeBlock transaction.
type TxListValidator struct {
	blockMaxGasLimit  uint64
	maxBytesPerTxList uint64
	chainID           *big.Int
}

// NewTxListValidator creates a new TxListValidator instance based on giving configurations.
func NewTxListValidator(
	blockMaxGasLimit uint64,
	maxBytesPerTxList uint64,
	chainID *big.Int,
) *TxListValidator {
	return &TxListValidator{
		blockMaxGasLimit:  blockMaxGasLimit,
		maxBytesPerTxList: maxBytesPerTxList,
		chainID:           chainID,
	}
}

// ValidateTxList checks whether the transactions list in the TaikoL1.proposeBlock transaction's
// input data is valid, the rules are:
// - If the transaction list is empty, it's valid.
// - If the transaction list is not empty:
//  1. If the transaction list is using calldata, the compressed bytes of the transaction list must be
//     less than or equal to maxBytesPerTxList.
//  2. The transaction list bytes must be able to be RLP decoded into a list of transactions.
func (v *TxListValidator) ValidateTxList(
	blockID *big.Int,
	txListBytes []byte,
	blobUsed bool,
) bool {
	// If the transaction list is empty, it's valid.
	if len(txListBytes) == 0 {
		return true
	}

	if !blobUsed && (len(txListBytes) > int(v.maxBytesPerTxList)) {
		log.Info("Compressed transactions list binary too large", "length", len(txListBytes), "blockID", blockID)
		return false
	}

	var (
		txs types.Transactions
		err error
	)

	if txListBytes, err = utils.Decompress(txListBytes); err != nil {
		log.Info("Failed to decompress tx list bytes", "blockID", blockID, "error", err)
		return false
	}

	if err = rlp.DecodeBytes(txListBytes, &txs); err != nil {
		log.Info("Failed to decode transactions list bytes", "blockID", blockID, "error", err)
		return false
	}

	log.Info("Transaction list is valid", "blockID", blockID)
	return true
}
