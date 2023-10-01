package rpc

import (
	"context"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/taikoxyz/taiko-client/testutils"
)

func (s *RpcTestSuite) TestWaitReceiptTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err := WaitReceipt(
		ctx, s.cli.L1, types.NewTransaction(0, common.Address{}, common.Big0, 0, common.Big0, []byte{}),
	)
	s.ErrorContains(err, "context deadline exceeded")
}

// TODO: fix this, need to propose/prove/execute tx before this'll work
// func TestWaitReceiptRevert() {
// 	client := s.newTestClient()
// 	testAddrPrivKey, err := crypto.ToECDSA(
// 		common.Hex2Bytes("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"),
// 	)
// 	s.NoError(err)
// 	testAddr := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")

// 	// build transaction
// 	nonce, err := client.L2.PendingNonceAt(context.Background(), testAddr)
// 	s.NoError(err)
// 	data := []byte("invalid")
// 	parent, err := client.L2.BlockByNumber(context.Background(), nil)
// 	s.NoError(err)
// 	baseFee, err := client.TaikoL2.GetBasefee(nil, 1, uint32(parent.GasUsed()))
// 	s.NoError(err)
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
// 	s.NoError(err)
// 	s.Nil(t, client.L2.SendTransaction(context.Background(), signedTx))

// 	_, err2 := WaitReceipt(
// 		context.Background(), client.L2, signedTx,
// 	)
// 	s.ErrorContains(t, err2, "transaction reverted,")
// }

func (s *RpcTestSuite) TestSetHead() {
	s.Nil(SetHead(context.Background(), s.cli.L2RawRPC, common.Big0))
}

func (s *RpcTestSuite) TestStringToBytes32() {
	s.Equal([32]byte{}, StringToBytes32(""))
	s.Equal([32]byte{0x61, 0x62, 0x63}, StringToBytes32("abc"))
}

func (s *RpcTestSuite) TestL1ContentFrom() {
	l2Head, err := s.cli.L2.HeaderByNumber(context.Background(), nil)
	s.NoError(err)

	baseFee, err := s.cli.TaikoL2.GetBasefee(nil, 0, uint32(l2Head.GasUsed))
	s.NoError(err)

	testAddrPrivKey := testutils.ProposerPrivKey

	testAddr := crypto.PubkeyToAddress(testAddrPrivKey.PublicKey)

	nonce, err := s.cli.L2.PendingNonceAt(context.Background(), testAddr)
	s.NoError(err)

	tx := types.NewTransaction(
		nonce,
		testAddr,
		common.Big1,
		100000,
		baseFee,
		[]byte{},
	)
	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(s.cli.L2ChainID), testAddrPrivKey)
	s.NoError(err)
	s.Nil(s.cli.L2.SendTransaction(context.Background(), signedTx))

	content, err := ContentFrom(context.Background(), s.cli.L2RawRPC, testAddr)
	s.NoError(err)

	s.NotZero(len(content["pending"]))
	s.Equal(signedTx.Nonce(), content["pending"][strconv.Itoa(int(signedTx.Nonce()))].Nonce())
}
