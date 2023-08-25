package rpc

import (
	"context"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

func TestWaitReceiptTimeout(t *testing.T) {
	client := newTestClient(t)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := WaitReceipt(
		ctx, client.L1, types.NewTransaction(0, common.Address{}, common.Big0, 0, common.Big0, []byte{}),
	)

	require.ErrorContains(t, err, "context deadline exceeded")
}

// TODO: fix this, need to propose/prove/execute tx before this'll work
// func TestWaitReceiptRevert(t *testing.T) {
// 	client := newTestClient(t)
// 	testAddrPrivKey, err := crypto.ToECDSA(
// 		common.Hex2Bytes("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"),
// 	)
// 	require.Nil(t, err)
// 	testAddr := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")

// 	// build transaction
// 	nonce, err := client.L2.PendingNonceAt(context.Background(), testAddr)
// 	require.Nil(t, err)
// 	data := []byte("invalid")
// 	parent, err := client.L2.BlockByNumber(context.Background(), nil)
// 	require.Nil(t, err)
// 	baseFee, err := client.TaikoL2.GetBasefee(nil, 1, uint32(parent.GasUsed()))
// 	require.Nil(t, err)
// 	tx := types.NewTx(&types.DynamicFeeTx{
// 		ChainID:   client.L2ChainID,
// 		Nonce:     nonce,
// 		GasTipCap: common.Big0,
// 		GasFeeCap: new(big.Int).SetUint64(baseFee.Uint64() * 2),
// 		Gas:       uint64(22000),
// 		To:        &testAddr,
// 		Value:     common.Big0,
// 		Data:      data,
// 	})

// 	// sign transaction and send
// 	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(client.L2ChainID), testAddrPrivKey)
// 	require.Nil(t, err)
// 	require.Nil(t, client.L2.SendTransaction(context.Background(), signedTx))

// 	_, err2 := WaitReceipt(
// 		context.Background(), client.L2, signedTx,
// 	)
// 	require.ErrorContains(t, err2, "transaction reverted,")
// }

func TestSetHead(t *testing.T) {
	require.Nil(t, SetHead(context.Background(), newTestClient(t).L2RawRPC, common.Big0))
}

func TestStringToBytes32(t *testing.T) {
	require.Equal(t, [32]byte{}, StringToBytes32(""))
	require.Equal(t, [32]byte{0x61, 0x62, 0x63}, StringToBytes32("abc"))
}

func TestL1ContentFrom(t *testing.T) {
	client := newTestClient(t)
	l2Head, err := client.L2.HeaderByNumber(context.Background(), nil)
	require.Nil(t, err)

	baseFee, err := client.TaikoL2.GetBasefee(nil, 0, uint32(l2Head.GasUsed))
	require.Nil(t, err)

	testAddrPrivKey, err := crypto.ToECDSA(common.Hex2Bytes(os.Getenv("L1_PROPOSER_PRIVATE_KEY")))
	require.Nil(t, err)

	testAddr := crypto.PubkeyToAddress(testAddrPrivKey.PublicKey)

	nonce, err := client.L2.PendingNonceAt(context.Background(), testAddr)
	require.Nil(t, err)

	tx := types.NewTransaction(
		nonce,
		testAddr,
		common.Big1,
		100000,
		baseFee,
		[]byte{},
	)
	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(client.L2ChainID), testAddrPrivKey)
	require.Nil(t, err)
	require.Nil(t, client.L2.SendTransaction(context.Background(), signedTx))

	content, err := ContentFrom(context.Background(), client.L2RawRPC, testAddr)
	require.Nil(t, err)

	require.NotZero(t, len(content["pending"]))
	require.Equal(t, signedTx.Nonce(), content["pending"][strconv.Itoa(int(signedTx.Nonce()))].Nonce())
}
