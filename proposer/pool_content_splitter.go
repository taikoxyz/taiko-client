package proposer

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/les/utils"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/taikoxyz/taiko-client/metrics"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
)

// poolContentSplitter is responsible for splitting the pool content
// which fetched from a `txpool_content` RPC call response into several smaller transactions lists
// and make sure each splitted list satisfies the limits defined in Taiko protocol.
type poolContentSplitter struct {
	chainID                 *big.Int
	shufflePoolContent      bool
	maxTransactionsPerBlock uint64
	blockMaxGasLimit        uint64
	maxBytesPerTxList       uint64
	minTxGasLimit           uint64
}

// split splits the given transaction pool content to make each splitted
// transactions list satisfies the rules defined in Taiko protocol.
func (p *poolContentSplitter) split(poolContent rpc.PoolContent) []types.Transactions {
	var (
		txs                    = poolContent.ToTxsByPriceAndNonce(p.chainID)
		splittedTxLists        = make([]types.Transactions, 0)
		txBuffer               = make([]*types.Transaction, 0, p.maxTransactionsPerBlock)
		gasBuffer       uint64 = 0
	)

	for {
		tx := txs.Peek()
		if tx == nil {
			break
		}

		// If the transaction is invalid, we simply ignore it.
		if err := p.validateTx(tx); err != nil {
			log.Debug("Invalid pending transaction", "hash", tx.Hash(), "error", err)
			metrics.ProposerInvalidTxsCounter.Inc(1)
			txs.Pop() // If this tx is invalid, ignore this sender's other txs in pool.
			continue
		}

		// If the transactions buffer is full, we make all transactions in
		// current buffer a new splitted transaction list, and then reset the
		// buffer.
		if p.isTxBufferFull(tx, txBuffer, gasBuffer) {
			splittedTxLists = append(splittedTxLists, txBuffer)
			txBuffer = make([]*types.Transaction, 0, p.maxTransactionsPerBlock)
			gasBuffer = 0
		}

		txBuffer = append(txBuffer, tx)
		gasBuffer += tx.Gas()

		txs.Shift()
	}

	// Maybe there are some remaining transactions in current buffer,
	// make them a new transactions list too.
	if len(txBuffer) > 0 {
		splittedTxLists = append(splittedTxLists, txBuffer)
	}

	// If the pool content is shuffled, we will only propose the first transactions list.
	if p.shufflePoolContent && len(splittedTxLists) > 0 {
		splittedTxLists = []types.Transactions{p.weightedShuffle(splittedTxLists)[0]}
	}

	return splittedTxLists
}

// validateTx checks whether the given transaction is valid according
// to the rules in Taiko protocol.
func (p *poolContentSplitter) validateTx(tx *types.Transaction) error {
	if tx.Gas() < p.minTxGasLimit || tx.Gas() > p.blockMaxGasLimit {
		return fmt.Errorf(
			"transaction %s gas limit reaches the limits, got=%v, lowerBound=%v, upperBound=%v",
			tx.Hash(), tx.Gas(), p.minTxGasLimit, p.blockMaxGasLimit,
		)
	}

	b, err := rlp.EncodeToBytes(tx)
	if err != nil {
		return fmt.Errorf(
			"failed to rlp encode the pending transaction %s: %w", tx.Hash(), err,
		)
	}

	if len(b) > int(p.maxBytesPerTxList) {
		return fmt.Errorf(
			"size of transaction %s's rlp encoded bytes is bigger than the limit, got=%v, limit=%v",
			tx.Hash(), len(b), p.maxBytesPerTxList,
		)
	}

	return nil
}

// isTxBufferFull checks whether the given transaction can be appended to the
// current transaction list
// NOTE: this function *MUST* be called after using `validateTx` to check every
// inside transaction is valid.
func (p *poolContentSplitter) isTxBufferFull(t *types.Transaction, txs []*types.Transaction, gas uint64) bool {
	if len(txs) >= int(p.maxTransactionsPerBlock) {
		return true
	}

	if gas+t.Gas() > p.blockMaxGasLimit {
		return true
	}

	// Transactions list's RLP encoding error has already been checked in
	// `validateTx`, so no need to check the error here.
	if b, _ := rlp.EncodeToBytes(append([]*types.Transaction{t}, txs...)); len(b) > int(p.maxBytesPerTxList) {
		return true
	}

	return false
}

// weightedShuffle does a weighted shuffle for the given transactions, each transaction's
// gas price will be used as the weight.
func (p *poolContentSplitter) weightedShuffle(txLists []types.Transactions) []types.Transactions {
	shuffled := make([]types.Transactions, 0)

	selector := utils.NewWeightedRandomSelect(func(i interface{}) uint64 {
		var weight uint64 = 1
		for _, tx := range txLists[i.(int)] {
			weight += tx.GasPrice().Uint64()
		}
		return weight
	})

	for i := range txLists {
		selector.Update(i)
	}

	for range txLists {
		idx := selector.Choose().(int)
		shuffled = append(shuffled, txLists[idx])
		selector.Remove(idx)
	}

	return shuffled
}
