package helper

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
	"github.com/taikoxyz/taiko-client/testutils"
)

func ProposeInvalidTxListBytes(s *testutils.ClientTestSuite, proposer testutils.Proposer) {
	invalidTxListBytes := RandomBytes(256)

	s.Nil(proposer.ProposeTxList(context.Background(), &encoding.TaikoL1BlockMetadataInput{
		Proposer:        proposer.L2SuggestedFeeRecipient(),
		TxListHash:      crypto.Keccak256Hash(invalidTxListBytes),
		TxListByteStart: common.Big0,
		TxListByteEnd:   new(big.Int).SetUint64(uint64(len(invalidTxListBytes))),
		CacheTxListInfo: false,
	}, invalidTxListBytes, 1, nil))
}

func ProposeAndInsertEmptyBlocks(
	s *testutils.ClientTestSuite,
	proposer testutils.Proposer,
	calldataSyncer testutils.CalldataSyncer,
) []*bindings.TaikoL1ClientBlockProposed {
	var events []*bindings.TaikoL1ClientBlockProposed
	rpcClient := NewWsRpcClient(s)
	defer rpcClient.Close()
	l1Head, err := rpcClient.L1.HeaderByNumber(context.Background(), nil)
	s.NoError(err)

	sink := make(chan *bindings.TaikoL1ClientBlockProposed)

	sub, err := rpcClient.TaikoL1.WatchBlockProposed(nil, sink, nil, nil)
	s.Nil(err)
	defer func() {
		sub.Unsubscribe()
		close(sink)
	}()

	// RLP encoded empty list
	var emptyTxs []types.Transaction
	encoded, err := rlp.EncodeToBytes(emptyTxs)
	s.Nil(err)
	// 1. Propose empty tx list
	s.Nil(proposer.ProposeTxList(context.Background(), &encoding.TaikoL1BlockMetadataInput{
		Proposer:        proposer.L2SuggestedFeeRecipient(),
		TxListHash:      crypto.Keccak256Hash(encoded),
		TxListByteStart: common.Big0,
		TxListByteEnd:   new(big.Int).SetUint64(uint64(len(encoded))),
		CacheTxListInfo: false,
	}, encoded, 0, nil))
	// 2. Propose invalidate tx list
	ProposeInvalidTxListBytes(s, proposer)

	// 3. Propose empty block (Zero byte txList)
	s.Nil(proposer.ProposeEmptyBlockOp(context.Background()))

	events = append(events, []*bindings.TaikoL1ClientBlockProposed{<-sink, <-sink, <-sink}...)

	_, isPending, err := rpcClient.L1.TransactionByHash(context.Background(), events[len(events)-1].Raw.TxHash)
	s.Nil(err)
	s.False(isPending)

	newL1Head, err := rpcClient.L1.HeaderByNumber(context.Background(), nil)
	s.Nil(err)
	s.Greater(newL1Head.Number.Uint64(), l1Head.Number.Uint64())

	syncProgress, err := rpcClient.L2.SyncProgress(context.Background())
	s.Nil(err)
	s.Nil(syncProgress)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	s.Nil(calldataSyncer.ProcessL1Blocks(ctx, newL1Head))
	return events
}

// ProposeAndInsertValidBlock proposes an valid tx list and then insert it
// into L2 execution engine's local chain.
func ProposeAndInsertValidBlock(
	s *testutils.ClientTestSuite,
	proposer testutils.Proposer,
	calldataSyncer testutils.CalldataSyncer,
) *bindings.TaikoL1ClientBlockProposed {
	rpcClient := NewWsRpcClient(s)
	defer rpcClient.Close()
	l1Head, err := rpcClient.L1.HeaderByNumber(context.Background(), nil)
	s.Nil(err)

	l2Head, err := rpcClient.L2.HeaderByNumber(context.Background(), nil)
	s.Nil(err)

	// Propose txs in L2 execution engine's mempool
	sink := make(chan *bindings.TaikoL1ClientBlockProposed)

	sub, err := rpcClient.TaikoL1.WatchBlockProposed(nil, sink, nil, nil)
	s.Nil(err)
	defer func() {
		sub.Unsubscribe()
		close(sink)
	}()

	baseFee, err := rpcClient.TaikoL2.GetBasefee(nil, 0, uint32(l2Head.GasUsed))
	s.Nil(err)

	nonce, err := rpcClient.L2.PendingNonceAt(context.Background(), testutils.ProposerAddress)
	s.Nil(err)

	tx := types.NewTransaction(
		nonce,
		common.BytesToAddress(RandomBytes(32)),
		common.Big1,
		100000,
		baseFee,
		[]byte{},
	)
	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(rpcClient.L2ChainID), testutils.ProposerPrivKey)
	s.Nil(err)
	s.Nil(rpcClient.L2.SendTransaction(context.Background(), signedTx))

	s.Nil(proposer.ProposeOp(context.Background()))

	event := <-sink

	_, isPending, err := rpcClient.L1.TransactionByHash(context.Background(), event.Raw.TxHash)
	s.Nil(err)
	s.False(isPending)

	receipt, err := rpcClient.L1.TransactionReceipt(context.Background(), event.Raw.TxHash)
	s.Nil(err)
	s.Equal(types.ReceiptStatusSuccessful, receipt.Status)

	newL1Head, err := rpcClient.L1.HeaderByNumber(context.Background(), nil)
	s.Nil(err)
	s.Greater(newL1Head.Number.Uint64(), l1Head.Number.Uint64())

	syncProgress, err := rpcClient.L2.SyncProgress(context.Background())
	s.Nil(err)
	s.Nil(syncProgress)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	s.Nil(calldataSyncer.ProcessL1Blocks(ctx, newL1Head))

	_, err = rpcClient.L2.HeaderByNumber(context.Background(), nil)
	s.Nil(err)
	return event
}

func DepositEtherToL2(s *testutils.ClientTestSuite, depositerPrivKey *ecdsa.PrivateKey, recipient common.Address) {
	rpcClient := NewWsRpcClient(s)
	defer rpcClient.Close()
	config, err := rpcClient.TaikoL1.GetConfig(nil)
	s.Nil(err)

	opts, err := bind.NewKeyedTransactorWithChainID(depositerPrivKey, rpcClient.L1ChainID)
	s.Nil(err)
	opts.Value = config.EthDepositMinAmount

	for i := 0; i < int(config.EthDepositMinCountPerBlock); i++ {
		_, err = rpcClient.TaikoL1.DepositEtherToL2(opts, recipient)
		s.Nil(err)
	}
}
