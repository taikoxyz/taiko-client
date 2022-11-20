package proposer

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/stretchr/testify/require"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
)

func TestProposeInvalidBlocksOp(t *testing.T) {
	p := newTestProposer(t)

	require.Nil(t, p.proposeInvalidBlocksOp(context.Background(), 1024))
	require.Nil(t, p.proposeInvalidBlocksOp(context.Background(), 2))
}

func TestProposeInvalidTxListBytes(t *testing.T) {
	p := newTestProposer(t)
	sink := make(chan *bindings.TaikoL1ClientBlockProposed)

	sub, err := p.rpc.TaikoL1.WatchBlockProposed(nil, sink, nil)
	require.Nil(t, err)
	defer sub.Unsubscribe()

	require.Nil(t, p.proposeInvalidTxListBytes(context.Background()))

	event := <-sink

	tx, isPending, err := p.rpc.L1.TransactionByHash(context.Background(), event.Raw.TxHash)
	require.Nil(t, err)
	require.False(t, isPending)

	proposedBytes := unpackTxListBytes(t, tx)
	require.NotNil(t, rlp.DecodeBytes(proposedBytes, new(types.Transactions)))
}

func TestProposeTxListIncludingInvalidTx(t *testing.T) {
	p := newTestProposer(t)
	sink := make(chan *bindings.TaikoL1ClientBlockProposed)

	sub, err := p.rpc.TaikoL1.WatchBlockProposed(nil, sink, nil)
	require.Nil(t, err)
	defer sub.Unsubscribe()

	require.Nil(t, p.proposeTxListIncludingInvalidTx(context.Background()))

	event := <-sink

	tx, isPending, err := p.rpc.L1.TransactionByHash(context.Background(), event.Raw.TxHash)
	require.Nil(t, err)
	require.False(t, isPending)

	proposedBytes := unpackTxListBytes(t, tx)

	var txList types.Transactions
	require.Nil(t, rlp.DecodeBytes(proposedBytes, &txList))

	require.Equal(t, 1, len(txList))

	invalidTx := txList[0]

	invalidTxSender, err := types.Sender(types.LatestSignerForChainID(invalidTx.ChainId()), invalidTx)
	require.Nil(t, err)

	pendingNonce, err := p.rpc.L2.PendingNonceAt(context.Background(), invalidTxSender)
	require.Nil(t, err)

	require.NotEqual(t, pendingNonce, invalidTx.Nonce())
}

func unpackTxListBytes(t *testing.T, tx *types.Transaction) []byte {
	method, err := encoding.TaikoL1ABI.MethodById(tx.Data())
	require.Nil(t, err)
	require.Equal(t, "proposeBlock", method.Name)

	args := map[string]interface{}{}

	require.Nil(t, method.Inputs.UnpackIntoMap(args, tx.Data()[4:]))

	inputs, ok := args["inputs"].([][]byte)
	require.True(t, ok)
	require.Equal(t, 2, len(inputs))

	return inputs[1]
}
