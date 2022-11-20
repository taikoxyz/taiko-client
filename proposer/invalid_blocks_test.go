package proposer

import (
	"context"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
)

func (s *ProposerTestSuite) TestProposeInvalidBlocksOp() {
	s.Nil(s.p.ProposeInvalidBlocksOp(context.Background(), 1024))
	s.Nil(s.p.ProposeInvalidBlocksOp(context.Background(), 2))
}

func (s *ProposerTestSuite) TestProposeInvalidTxListBytes() {
	sink := make(chan *bindings.TaikoL1ClientBlockProposed)

	sub, err := s.p.rpc.TaikoL1.WatchBlockProposed(nil, sink, nil)
	s.Nil(err)
	defer func() {
		sub.Unsubscribe()
		close(sink)
	}()

	s.Nil(s.p.ProposeInvalidTxListBytes(context.Background()))

	event := <-sink

	tx, isPending, err := s.p.rpc.L1.TransactionByHash(context.Background(), event.Raw.TxHash)
	s.Nil(err)
	s.False(isPending)

	proposedBytes := s.unpackTxListBytes(tx)
	s.NotNil(rlp.DecodeBytes(proposedBytes, new(types.Transactions)))
}

func (s *ProposerTestSuite) TestProposeTxListIncludingInvalidTx() {
	sink := make(chan *bindings.TaikoL1ClientBlockProposed)

	sub, err := s.p.rpc.TaikoL1.WatchBlockProposed(nil, sink, nil)
	s.Nil(err)
	defer func() {
		sub.Unsubscribe()
		close(sink)
	}()

	s.Nil(s.p.proposeTxListIncludingInvalidTx(context.Background()))

	event := <-sink

	tx, isPending, err := s.p.rpc.L1.TransactionByHash(context.Background(), event.Raw.TxHash)
	s.Nil(err)
	s.False(isPending)

	proposedBytes := s.unpackTxListBytes(tx)

	var txList types.Transactions
	s.Nil(rlp.DecodeBytes(proposedBytes, &txList))

	s.Equal(1, len(txList))

	invalidTx := txList[0]

	invalidTxSender, err := types.Sender(types.LatestSignerForChainID(invalidTx.ChainId()), invalidTx)
	s.Nil(err)

	pendingNonce, err := s.p.rpc.L2.PendingNonceAt(context.Background(), invalidTxSender)
	s.Nil(err)

	s.NotEqual(pendingNonce, invalidTx.Nonce())
}

func (s *ProposerTestSuite) unpackTxListBytes(tx *types.Transaction) []byte {
	method, err := encoding.TaikoL1ABI.MethodById(tx.Data())
	s.Nil(err)
	s.Equal("proposeBlock", method.Name)

	args := map[string]interface{}{}

	s.Nil(method.Inputs.UnpackIntoMap(args, tx.Data()[4:]))

	inputs, ok := args["inputs"].([][]byte)
	s.True(ok)
	s.Equal(2, len(inputs))

	return inputs[1]
}
