package rpc

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/taikoxyz/taiko-client/testutils"
)

var testAddress = common.HexToAddress("0x98f86166571FE624778203d87A8eD6fd84695B79")

func (s *RpcTestSuite) TestL2AccountNonce() {
	client := s.newTestClientWithTimeout()
	defer client.Close()
	nonce, err := client.L2AccountNonce(context.Background(), testAddress, common.Big0)

	s.NoError(err)
	s.Zero(nonce)
}

func (s *RpcTestSuite) TestGetGenesisL1Header() {
	client := s.newTestClient()
	defer client.Close()
	header, err := client.GetGenesisL1Header(context.Background())

	s.NoError(err)
	s.NotZero(header.Number.Uint64())
}

func (s *RpcTestSuite) TestLatestL2KnownL1Header() {
	client := s.newTestClient()
	defer client.Close()
	header, err := client.LatestL2KnownL1Header(context.Background())

	s.NoError(err)
	s.NotZero(header.Number.Uint64())
}

func (s *RpcTestSuite) TestL2ParentByBlockId() {
	client := s.newTestClient()
	defer client.Close()
	header, err := client.L2ParentByBlockId(context.Background(), common.Big1)
	s.NoError(err)
	s.Zero(header.Number.Uint64())

	_, err = client.L2ParentByBlockId(context.Background(), common.Big2)
	s.Error(err)
}

func (s *RpcTestSuite) TestL2ExecutionEngineSyncProgress() {
	client := s.newTestClient()
	defer client.Close()
	progress, err := client.L2ExecutionEngineSyncProgress(context.Background())
	s.NoError(err)
	s.NotNil(progress)
}

func (s *RpcTestSuite) TestGetProtocolStateVariables() {
	client := s.newTestClient()
	defer client.Close()
	_, err := client.GetProtocolStateVariables(nil)
	s.NoError(err)
}

func (s *RpcTestSuite) TestCheckL1ReorgFromL1Cursor() {
	client := s.newTestClient()
	defer client.Close()
	l1Head, err := client.L1.HeaderByNumber(context.Background(), nil)
	s.NoError(err)

	_, newL1Current, _, err := client.CheckL1ReorgFromL1Cursor(context.Background(), l1Head, l1Head.Number.Uint64())
	s.NoError(err)

	s.Equal(l1Head.Number.Uint64(), newL1Current.Number.Uint64())

	stateVar, err := client.TaikoL1.GetStateVariables(nil)
	s.NoError(err)

	reorged, _, _, err := client.CheckL1ReorgFromL1Cursor(context.Background(), l1Head, stateVar.GenesisHeight)
	s.NoError(err)
	s.False(reorged)

	l1Head.BaseFee = new(big.Int).Add(l1Head.BaseFee, common.Big1)

	reorged, newL1Current, _, err = client.CheckL1ReorgFromL1Cursor(context.Background(), l1Head, stateVar.GenesisHeight)
	s.NoError(err)
	s.True(reorged)
	s.Equal(l1Head.ParentHash, newL1Current.Hash())
}

func (s *RpcTestSuite) TestIsJustSyncedByP2P() {
	client := s.newTestClient()
	defer client.Close()
	_, err := client.IsJustSyncedByP2P(context.Background())
	s.NoError(err)
}

func (s *RpcTestSuite) TestWaitTillL2ExecutionEngineSyncedNewClient() {
	client := s.newTestClient()
	defer client.Close()
	err := client.WaitTillL2ExecutionEngineSynced(context.Background())
	s.NoError(err)
}

func (s *RpcTestSuite) TestWaitTillL2ExecutionEngineSyncedContextErr() {
	client := s.newTestClient()
	defer client.Close()
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	err := client.WaitTillL2ExecutionEngineSynced(ctx)
	s.ErrorContains(err, "context canceled")
}

func (s *RpcTestSuite) TestGetPoolContentValid() {
	client := s.newTestClient()
	defer client.Close()
	configs, err := client.TaikoL1.GetConfig(&bind.CallOpts{Context: context.Background()})
	s.NoError(err)
	goldenTouchAddress, err := client.TaikoL2.GOLDENTOUCHADDRESS(nil)
	s.NoError(err)
	parent, err := client.L2.BlockByNumber(context.Background(), nil)
	s.NoError(err)
	baseFee, err := client.TaikoL2.GetBasefee(nil, 1, uint32(parent.GasUsed()))
	s.NoError(err)
	gasLimit := configs.BlockMaxGasLimit
	maxBytes := configs.BlockMaxTxListBytes

	txPools := []common.Address{goldenTouchAddress}

	_, err2 := client.GetPoolContent(
		context.Background(),
		goldenTouchAddress,
		baseFee,
		gasLimit,
		maxBytes.Uint64(),
		txPools,
		defaultMaxTransactionsPerBlock,
	)
	s.NoError(err2)
}

func (s *RpcTestSuite) TestGetStorageRootNewestBlock() {
	client := s.newTestClient()
	defer client.Close()
	_, err := client.GetStorageRoot(
		context.Background(),
		client.L1GethClient,
		testutils.TaikoL1SignalService,
		nil)
	s.NoError(err)
}
