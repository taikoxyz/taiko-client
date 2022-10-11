package driver

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/stretchr/testify/require"
	"github.com/taikochain/taiko-client/bindings"
)

var (
	testKey, _ = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
)

var testTx1 = types.MustSignNewTx(testKey, types.LatestSigner(params.AllEthashProtocolChanges), &types.LegacyTx{
	Nonce:    0,
	Value:    big.NewInt(12),
	GasPrice: big.NewInt(params.InitialBaseFee),
	Gas:      params.TxGas,
	To:       &common.Address{2},
})

func TestInsertNewHead(t *testing.T) {
	d := newTestDriver(t)

	txList := types.Transactions{testTx1}
	txListBytes, err := rlp.EncodeToBytes(txList)
	require.Nil(t, err)
	require.NotEmpty(t, txListBytes)

	l1Head, err := d.rpc.L1.HeaderByNumber(context.Background(), nil)
	require.Nil(t, err)

	event := &bindings.TaikoL1ClientBlockProposed{
		Id: common.Big1,
		Meta: bindings.LibDataBlockMetadata{
			Id:          common.Big1,
			L1Height:    l1Head.Number,
			L1Hash:      l1Head.Hash(),
			Beneficiary: common.BytesToAddress(randomHash().Bytes()),
			GasLimit:    100000,
			Timestamp:   uint64(time.Now().Unix()),
			TxListHash:  crypto.Keccak256Hash(txListBytes),
			MixHash:     randomHash(),
			ExtraData:   randomHash().Bytes(),
		},
		Raw: types.Log{
			BlockNumber: l1Head.Number.Uint64(),
			BlockHash:   l1Head.Hash(),
		},
	}

	parent, err := d.rpc.L2.HeaderByNumber(context.Background(), nil)
	require.Nil(t, err)

	payload, rpcErr, payloadErr := d.l2ChainInserter.insertNewHead(
		context.Background(),
		event,
		parent,
		new(big.Int).Add(l1Head.Number, common.Big1),
		txListBytes,
		&rawdb.L1Origin{
			BlockID:       event.Id,
			L2BlockHash:   common.Hash{},
			L1BlockHeight: new(big.Int).SetUint64(event.Raw.BlockNumber),
			L1BlockHash:   event.Raw.BlockHash,
		},
	)

	require.Nil(t, rpcErr)
	require.Nil(t, payloadErr)
	require.Equal(t, common.BytesToHash(event.Meta.MixHash[:]), payload.Random)
	require.Less(t, event.Meta.GasLimit, payload.GasLimit)
	require.Equal(t, event.Meta.ExtraData, payload.ExtraData)
	require.Equal(t, event.Meta.Timestamp, payload.Timestamp)
	require.Equal(t, event.Meta.Beneficiary, payload.FeeRecipient)
}
