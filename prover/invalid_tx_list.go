package prover

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
)

// InvalidTxListReason represents a reason why a transactions list is invalid,
// must match the definitions in LibInvalidTxList.sol:
//
//	enum Reason {
//		OK,
//		BINARY_TOO_LARGE,
//		BINARY_NOT_DECODABLE,
//		BLOCK_TOO_MANY_TXS,
//		BLOCK_GAS_LIMIT_TOO_LARGE,
//		TX_INVALID_SIG,
//		TX_GAS_LIMIT_TOO_SMALL
//	}
type InvalidTxListReason uint8

// All invalid transactions list reasons.
const (
	HintOK InvalidTxListReason = iota
	HintBinaryTooLarge
	HintBinaryNotDecodable
	HintBlockTooManyTxs
	HintBlockGasLimitTooLarge
	HintTxInvalidSig
	HintTxGasLimitTooSmall
)

// isTxListValid checks whether the transaction list is valid, must match
// the validation rule defined in LibInvalidTxList.sol.
// ref: https://github.com/taikochain/taiko-mono/blob/main/packages/bindings/contracts/libs/LibInvalidTxList.sol
func (p *Prover) isTxListValid(blockID *big.Int, txListBytes []byte) (hint InvalidTxListReason, txIdx int) {
	if len(txListBytes) > int(p.maxTxlistBytes) {
		log.Warn("Transactions list binary too large, length: %s", len(txListBytes), "blockID", blockID)
		return HintBinaryTooLarge, 0
	}

	var txs types.Transactions
	if err := rlp.DecodeBytes(txListBytes, &txs); err != nil {
		log.Warn("Failed to decode transactions list bytes", "blockID", blockID, "error", err)
		return HintBinaryNotDecodable, 0
	}

	log.Info("Transactions list decoded", "blockID", blockID, "length", len(txs))

	if txs.Len() > int(p.maxBlockNumTxs) {
		log.Warn("Too many transactions", "blockID", blockID, "count", txs.Len())
		return HintBlockTooManyTxs, 0
	}

	sumGasLimit := uint64(0)
	for _, tx := range txs {
		sumGasLimit += tx.Gas()
	}

	if sumGasLimit > p.maxBlocksGasLimit {
		log.Warn("Accumulate gas limit too large", "blockID", blockID, "sumGasLimit", sumGasLimit)
		return HintBlockGasLimitTooLarge, 0
	}

	signer := types.LatestSignerForChainID(p.chainID)

	for i, tx := range txs {
		sender, err := types.Sender(signer, tx)
		if err != nil || sender == (common.Address{}) {
			log.Warn("Invalid transaction signature", "error", err)
			return HintTxInvalidSig, i
		}

		if tx.Gas() < p.minTxGasLimit {
			log.Warn("Transaction gas limit too small", "gasLimit", tx.Gas())
			return HintTxGasLimitTooSmall, i
		}
	}

	log.Info("Transaction list is valid", "blockID", blockID)
	return HintOK, 0
}

// unpackTxListBytes unpacks the L2 transaction list from a L1 block's calldata.
func (p *Prover) unpackTxListBytes(tx *types.Transaction) ([]byte, error) {
	method, err := p.taikoL1Abi.MethodById(tx.Data())
	if err != nil {
		return nil, err
	}

	// Only check for safety.
	if method.Name != "proposeBlock" {
		return nil, errInvalidProposeBlockTx
	}

	args := map[string]interface{}{}

	if err := method.Inputs.UnpackIntoMap(args, tx.Data()[4:]); err != nil {
		return nil, errInvalidProposeBlockTx
	}

	inputs, ok := args["inputs"].([][]byte)

	if !ok || len(inputs) < 2 {
		return nil, errInvalidProposeBlockTx
	}

	return inputs[1], nil
}
