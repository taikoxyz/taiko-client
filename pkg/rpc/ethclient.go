package rpc

import (
	"context"
	"math/big"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// EthClient is a wrapper for go-ethereum ethclient with a timeout attached.
type EthClient struct {
	*ethclient.Client
	timeout time.Duration
}

// NewEthClientWithTimeout creates a new EthClient instance with the given
// request timeout.
func NewEthClientWithTimeout(
	ethclient *ethclient.Client,
	timeout time.Duration,
) *EthClient {
	if ethclient == nil {
		return nil
	}

	return &EthClient{Client: ethclient, timeout: timeout}
}

// NewEthClientWithDefaultTimeout creates a new EthClient instance with the default
// timeout.
func NewEthClientWithDefaultTimeout(
	ethclient *ethclient.Client,
) *EthClient {
	if ethclient == nil {
		return nil
	}

	return &EthClient{Client: ethclient, timeout: defaultTimeout}
}

// ctxWithTimeoutOrDefault sets a context timeout if the deadline has not passed or is not set,
// and otherwise returns the context as passed in. cancel func is always set to an empty function
// so is safe to defer the cancel.
func (c *EthClient) ctxWithTimeoutOrDefault(ctx context.Context) (context.Context, context.CancelFunc) {
	var (
		ctxWithTimeout                    = ctx
		cancel         context.CancelFunc = func() {}
	)
	if _, ok := ctx.Deadline(); !ok {
		ctxWithTimeout, cancel = context.WithTimeout(ctx, c.timeout)
	}

	return ctxWithTimeout, cancel
}

// ChainID retrieves the current chain ID for transaction replay protection.
func (c *EthClient) ChainID(ctx context.Context) (*big.Int, error) {
	ctxWithTimeout, cancel := c.ctxWithTimeoutOrDefault(ctx)
	defer cancel()

	return c.Client.ChainID(ctxWithTimeout)
}

// BlockByHash returns the given full block.
//
// Note that loading full blocks requires two requests. Use HeaderByHash
// if you don't need all transactions or uncle headers.
func (c *EthClient) BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
	ctxWithTimeout, cancel := c.ctxWithTimeoutOrDefault(ctx)
	defer cancel()

	return c.Client.BlockByHash(ctxWithTimeout, hash)
}

// BlockByNumber returns a block from the current canonical chain. If number is nil, the
// latest known block is returned.
//
// Note that loading full blocks requires two requests. Use HeaderByNumber
// if you don't need all transactions or uncle headers.
func (c *EthClient) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	ctxWithTimeout, cancel := c.ctxWithTimeoutOrDefault(ctx)
	defer cancel()

	return c.Client.BlockByNumber(ctxWithTimeout, number)
}

// BlockNumber returns the most recent block number
func (c *EthClient) BlockNumber(ctx context.Context) (uint64, error) {
	ctxWithTimeout, cancel := c.ctxWithTimeoutOrDefault(ctx)
	defer cancel()

	return c.Client.BlockNumber(ctxWithTimeout)
}

// PeerCount returns the number of p2p peers as reported by the net_peerCount method.
func (c *EthClient) PeerCount(ctx context.Context) (uint64, error) {
	ctxWithTimeout, cancel := c.ctxWithTimeoutOrDefault(ctx)
	defer cancel()

	return c.Client.PeerCount(ctxWithTimeout)
}

// HeaderByHash returns the block header with the given hash.
func (c *EthClient) HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error) {
	ctxWithTimeout, cancel := c.ctxWithTimeoutOrDefault(ctx)
	defer cancel()

	return c.Client.HeaderByHash(ctxWithTimeout, hash)
}

// HeaderByNumber returns a block header from the current canonical chain. If number is
// nil, the latest known header is returned.
func (c *EthClient) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	ctxWithTimeout, cancel := c.ctxWithTimeoutOrDefault(ctx)
	defer cancel()

	return c.Client.HeaderByNumber(ctxWithTimeout, number)
}

// TransactionByHash returns the transaction with the given hash.
func (c *EthClient) TransactionByHash(
	ctx context.Context,
	hash common.Hash,
) (tx *types.Transaction, isPending bool, err error) {
	ctxWithTimeout, cancel := c.ctxWithTimeoutOrDefault(ctx)
	defer cancel()

	return c.Client.TransactionByHash(ctxWithTimeout, hash)
}

// TransactionSender returns the sender address of the given transaction. The transaction
// must be known to the remote node and included in the blockchain at the given block and
// index. The sender is the one derived by the protocol at the time of inclusion.
//
// There is a fast-path for transactions retrieved by TransactionByHash and
// TransactionInBlock. Getting their sender address can be done without an RPC interaction.
func (c *EthClient) TransactionSender(
	ctx context.Context,
	tx *types.Transaction,
	block common.Hash,
	index uint,
) (common.Address, error) {
	ctxWithTimeout, cancel := c.ctxWithTimeoutOrDefault(ctx)
	defer cancel()

	return c.Client.TransactionSender(ctxWithTimeout, tx, block, index)
}

// TransactionCount returns the total number of transactions in the given block.
func (c *EthClient) TransactionCount(ctx context.Context, blockHash common.Hash) (uint, error) {
	ctxWithTimeout, cancel := c.ctxWithTimeoutOrDefault(ctx)
	defer cancel()

	return c.Client.TransactionCount(ctxWithTimeout, blockHash)
}

// TransactionInBlock returns a single transaction at index in the given block.
func (c *EthClient) TransactionInBlock(
	ctx context.Context,
	blockHash common.Hash,
	index uint,
) (*types.Transaction, error) {
	ctxWithTimeout, cancel := c.ctxWithTimeoutOrDefault(ctx)
	defer cancel()

	return c.Client.TransactionInBlock(ctxWithTimeout, blockHash, index)
}

// SyncProgress retrieves the current progress of the sync algorithm. If there's
// no sync currently running, it returns nil.
func (c *EthClient) SyncProgress(ctx context.Context) (*ethereum.SyncProgress, error) {
	ctxWithTimeout, cancel := c.ctxWithTimeoutOrDefault(ctx)
	defer cancel()

	return c.Client.SyncProgress(ctxWithTimeout)
}

// NetworkID returns the network ID for this client.
func (c *EthClient) NetworkID(ctx context.Context) (*big.Int, error) {
	ctxWithTimeout, cancel := c.ctxWithTimeoutOrDefault(ctx)
	defer cancel()

	return c.Client.NetworkID(ctxWithTimeout)
}

// BalanceAt returns the wei balance of the given account.
// The block number can be nil, in which case the balance is taken from the latest known block.
func (c *EthClient) BalanceAt(
	ctx context.Context,
	account common.Address,
	blockNumber *big.Int,
) (*big.Int, error) {
	ctxWithTimeout, cancel := c.ctxWithTimeoutOrDefault(ctx)
	defer cancel()

	return c.Client.BalanceAt(ctxWithTimeout, account, blockNumber)
}

// StorageAt returns the value of key in the contract storage of the given account.
// The block number can be nil, in which case the value is taken from the latest known block.
func (c *EthClient) StorageAt(
	ctx context.Context,
	account common.Address,
	key common.Hash,
	blockNumber *big.Int,
) ([]byte, error) {
	ctxWithTimeout, cancel := c.ctxWithTimeoutOrDefault(ctx)
	defer cancel()

	return c.Client.StorageAt(ctxWithTimeout, account, key, blockNumber)
}

// CodeAt returns the contract code of the given account.
// The block number can be nil, in which case the code is taken from the latest known block.
func (c *EthClient) CodeAt(
	ctx context.Context,
	account common.Address,
	blockNumber *big.Int,
) ([]byte, error) {
	ctxWithTimeout, cancel := c.ctxWithTimeoutOrDefault(ctx)
	defer cancel()

	return c.Client.CodeAt(ctxWithTimeout, account, blockNumber)
}

// NonceAt returns the account nonce of the given account.
// The block number can be nil, in which case the nonce is taken from the latest known block.
func (c *EthClient) NonceAt(
	ctx context.Context,
	account common.Address,
	blockNumber *big.Int,
) (uint64, error) {
	ctxWithTimeout, cancel := c.ctxWithTimeoutOrDefault(ctx)
	defer cancel()

	return c.Client.NonceAt(ctxWithTimeout, account, blockNumber)
}

// PendingBalanceAt returns the wei balance of the given account in the pending state.
func (c *EthClient) PendingBalanceAt(ctx context.Context, account common.Address) (*big.Int, error) {
	ctxWithTimeout, cancel := c.ctxWithTimeoutOrDefault(ctx)
	defer cancel()

	return c.Client.PendingBalanceAt(ctxWithTimeout, account)
}

// PendingStorageAt returns the value of key in the contract storage of the given account in the pending state.
func (c *EthClient) PendingStorageAt(
	ctx context.Context,
	account common.Address,
	key common.Hash,
) ([]byte, error) {
	ctxWithTimeout, cancel := c.ctxWithTimeoutOrDefault(ctx)
	defer cancel()

	return c.Client.PendingStorageAt(ctxWithTimeout, account, key)
}

// PendingCodeAt returns the contract code of the given account in the pending state.
func (c *EthClient) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	ctxWithTimeout, cancel := c.ctxWithTimeoutOrDefault(ctx)
	defer cancel()

	return c.Client.PendingCodeAt(ctxWithTimeout, account)
}

// PendingNonceAt returns the account nonce of the given account in the pending state.
// This is the nonce that should be used for the next transaction.
func (c *EthClient) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	ctxWithTimeout, cancel := c.ctxWithTimeoutOrDefault(ctx)
	defer cancel()

	return c.Client.PendingNonceAt(ctxWithTimeout, account)
}

// PendingTransactionCount returns the total number of transactions in the pending state.
func (c *EthClient) PendingTransactionCount(ctx context.Context) (uint, error) {
	ctxWithTimeout, cancel := c.ctxWithTimeoutOrDefault(ctx)
	defer cancel()

	return c.Client.PendingTransactionCount(ctxWithTimeout)
}

// CallContract executes a message call transaction, which is directly executed in the VM
// of the node, but never mined into the blockchain.
//
// blockNumber selects the block height at which the call runs. It can be nil, in which
// case the code is taken from the latest known block. Note that state from very old
// blocks might not be available.
func (c *EthClient) CallContract(
	ctx context.Context,
	msg ethereum.CallMsg,
	blockNumber *big.Int,
) ([]byte, error) {
	ctxWithTimeout, cancel := c.ctxWithTimeoutOrDefault(ctx)
	defer cancel()

	return c.Client.CallContract(ctxWithTimeout, msg, blockNumber)
}

// CallContractAtHash is almost the same as CallContract except that it selects
// the block by block hash instead of block height.
func (c *EthClient) CallContractAtHash(
	ctx context.Context,
	msg ethereum.CallMsg,
	blockHash common.Hash,
) ([]byte, error) {
	ctxWithTimeout, cancel := c.ctxWithTimeoutOrDefault(ctx)
	defer cancel()

	return c.Client.CallContractAtHash(ctxWithTimeout, msg, blockHash)
}

// PendingCallContract executes a message call transaction using the EVM.
// The state seen by the contract call is the pending state.
func (c *EthClient) PendingCallContract(ctx context.Context, msg ethereum.CallMsg) ([]byte, error) {
	ctxWithTimeout, cancel := c.ctxWithTimeoutOrDefault(ctx)
	defer cancel()

	return c.Client.PendingCallContract(ctxWithTimeout, msg)
}

// SuggestGasPrice retrieves the currently suggested gas price to allow a timely
// execution of a transaction.
func (c *EthClient) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	ctxWithTimeout, cancel := c.ctxWithTimeoutOrDefault(ctx)
	defer cancel()

	return c.Client.SuggestGasPrice(ctxWithTimeout)
}

// SuggestGasTipCap retrieves the currently suggested gas tip cap after 1559 to
// allow a timely execution of a transaction.
func (c *EthClient) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	ctxWithTimeout, cancel := c.ctxWithTimeoutOrDefault(ctx)
	defer cancel()

	return c.Client.SuggestGasTipCap(ctxWithTimeout)
}

// FeeHistory retrieves the fee market history.
func (c *EthClient) FeeHistory(
	ctx context.Context,
	blockCount uint64,
	lastBlock *big.Int,
	rewardPercentiles []float64,
) (*ethereum.FeeHistory, error) {
	ctxWithTimeout, cancel := c.ctxWithTimeoutOrDefault(ctx)
	defer cancel()

	return c.Client.FeeHistory(ctxWithTimeout, blockCount, lastBlock, rewardPercentiles)
}

// EstimateGas tries to estimate the gas needed to execute a specific transaction based on
// the current pending state of the backend blockchain. There is no guarantee that this is
// the true gas limit requirement as other transactions may be added or removed by miners,
// but it should provide a basis for setting a reasonable default.
func (c *EthClient) EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	ctxWithTimeout, cancel := c.ctxWithTimeoutOrDefault(ctx)
	defer cancel()

	return c.Client.EstimateGas(ctxWithTimeout, msg)
}

// SendTransaction injects a signed transaction into the pending pool for execution.
//
// If the transaction was a contract creation use the TransactionReceipt method to get the
// contract address after the transaction has been mined.
func (c *EthClient) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	ctxWithTimeout, cancel := c.ctxWithTimeoutOrDefault(ctx)
	defer cancel()

	return c.Client.SendTransaction(ctxWithTimeout, tx)
}
