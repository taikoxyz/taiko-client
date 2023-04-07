package testutils

import (
	"context"
	"math/big"
	"math/rand"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
)

func ProposeInvalidTxListBytes(s *ClientTestSuite, proposer Proposer) {
	configs, err := s.RpcClient.TaikoL1.GetConfig(nil)
	s.Nil(err)

	invalidTxListBytes := RandomBytes(256)

	s.Nil(proposer.ProposeTxList(context.Background(), &encoding.TaikoL1BlockMetadataInput{
		Beneficiary:     proposer.L2SuggestedFeeRecipient(),
		GasLimit:        uint32(rand.Int63n(configs.BlockMaxGasLimit.Int64())),
		TxListHash:      crypto.Keccak256Hash(invalidTxListBytes),
		TxListByteStart: common.Big0,
		TxListByteEnd:   new(big.Int).SetUint64(uint64(len(invalidTxListBytes))),
		CacheTxListInfo: 0,
	}, invalidTxListBytes, 1))
}

func ProposeAndInsertEmptyBlocks(
	s *ClientTestSuite,
	proposer Proposer,
	calldataSyncer CalldataSyncer,
) []*bindings.TaikoL1ClientBlockProposed {
	var events []*bindings.TaikoL1ClientBlockProposed

	l1Head, err := s.RpcClient.L1.HeaderByNumber(context.Background(), nil)
	s.Nil(err)

	sink := make(chan *bindings.TaikoL1ClientBlockProposed)

	sub, err := s.RpcClient.TaikoL1.WatchBlockProposed(nil, sink, nil)
	s.Nil(err)
	defer func() {
		sub.Unsubscribe()
		close(sink)
	}()

	// Zero byte txList
	s.Nil(proposer.ProposeEmptyBlockOp(context.Background()))

	// RLP encoded empty list
	var emptyTxs []types.Transaction
	encoded, err := rlp.EncodeToBytes(emptyTxs)
	s.Nil(err)

	s.Nil(proposer.ProposeTxList(context.Background(), &encoding.TaikoL1BlockMetadataInput{
		Beneficiary:     proposer.L2SuggestedFeeRecipient(),
		GasLimit:        21000,
		TxListHash:      crypto.Keccak256Hash(encoded),
		TxListByteStart: common.Big0,
		TxListByteEnd:   new(big.Int).SetUint64(uint64(len(encoded))),
		CacheTxListInfo: 0,
	}, encoded, 0))

	ProposeInvalidTxListBytes(s, proposer)

	events = append(events, []*bindings.TaikoL1ClientBlockProposed{<-sink, <-sink, <-sink}...)

	_, isPending, err := s.RpcClient.L1.TransactionByHash(context.Background(), events[len(events)-1].Raw.TxHash)
	s.Nil(err)
	s.False(isPending)

	newL1Head, err := s.RpcClient.L1.HeaderByNumber(context.Background(), nil)
	s.Nil(err)
	s.Greater(newL1Head.Number.Uint64(), l1Head.Number.Uint64())

	syncProgress, err := s.RpcClient.L2.SyncProgress(context.Background())
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
	s *ClientTestSuite,
	proposer Proposer,
	calldataSyncer CalldataSyncer,
) *bindings.TaikoL1ClientBlockProposed {
	l1Head, err := s.RpcClient.L1.HeaderByNumber(context.Background(), nil)
	s.Nil(err)

	l2Head, err := s.RpcClient.L2.HeaderByNumber(context.Background(), nil)
	s.Nil(err)

	// Propose txs in L2 execution engine's mempool
	sink := make(chan *bindings.TaikoL1ClientBlockProposed)

	sub, err := s.RpcClient.TaikoL1.WatchBlockProposed(nil, sink, nil)
	s.Nil(err)
	defer func() {
		sub.Unsubscribe()
		close(sink)
	}()

	nonce, err := s.RpcClient.L2.PendingNonceAt(context.Background(), s.TestAddr)
	s.Nil(err)

	tx := types.NewTransaction(
		nonce,
		common.BytesToAddress(RandomBytes(32)),
		common.Big1,
		100000,
		common.Big1,
		[]byte{},
	)
	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(s.RpcClient.L2ChainID), s.TestAddrPrivKey)
	s.Nil(err)
	s.Nil(s.RpcClient.L2.SendTransaction(context.Background(), signedTx))

	s.Nil(proposer.ProposeOp(context.Background()))

	event := <-sink

	_, isPending, err := s.RpcClient.L1.TransactionByHash(context.Background(), event.Raw.TxHash)
	s.Nil(err)
	s.False(isPending)

	receipt, err := s.RpcClient.L1.TransactionReceipt(context.Background(), event.Raw.TxHash)
	s.Nil(err)
	s.Equal(types.ReceiptStatusSuccessful, receipt.Status)

	newL1Head, err := s.RpcClient.L1.HeaderByNumber(context.Background(), nil)
	s.Nil(err)
	s.Greater(newL1Head.Number.Uint64(), l1Head.Number.Uint64())

	syncProgress, err := s.RpcClient.L2.SyncProgress(context.Background())
	s.Nil(err)
	s.Nil(syncProgress)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	s.Nil(calldataSyncer.ProcessL1Blocks(ctx, newL1Head))

	newL2Head, err := s.RpcClient.L2.HeaderByNumber(context.Background(), nil)
	s.Nil(err)

	s.Greater(newL2Head.Number.Uint64(), l2Head.Number.Uint64())

	return event
}

// RandomHash generates a random blob of data and returns it as a hash.
func RandomHash() common.Hash {
	var hash common.Hash
	if n, err := rand.Read(hash[:]); n != common.HashLength || err != nil {
		panic(err)
	}
	return hash
}

// RandomBytes generates a random bytes.
func RandomBytes(size int) (b []byte) {
	b = make([]byte, size)
	if _, err := rand.Read(b); err != nil {
		log.Crit("Generate random bytes error", "error", err)
	}
	return
}

// SignatureFromRSV creates the signature bytes from r,s,v.
func SignatureFromRSV(r, s string, v byte) []byte {
	return append(append(hexutil.MustDecode(r), hexutil.MustDecode(s)...), v)
}
