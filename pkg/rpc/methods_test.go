package rpc

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/taikoxyz/taiko-client/testutils"
)

func (s *RpcTestSuite) TestL2AccountNonce() {
	client := s.newTestClientWithTimeout()
	defer client.Close()
	nonce, err := client.L2AccountNonce(context.Background(), testutils.ProposerAddress, common.Big0)

	s.NoError(err)
	s.Zero(nonce)
}

func (s *RpcTestSuite) TestGetGenesisL1Header() {
	header, err := s.cli.GetGenesisL1Header(context.Background())

	s.NoError(err)
	s.NotZero(header.Number.Uint64())
}

func (s *RpcTestSuite) TestLatestL2KnownL1Header() {
	header, err := s.cli.LatestL2KnownL1Header(context.Background())

	s.NoError(err)
	s.NotZero(header.Number.Uint64())
}

func (s *RpcTestSuite) TestL2ParentByBlockId() {
	header, err := s.cli.L2ParentByBlockId(context.Background(), common.Big1)
	s.NoError(err)
	s.Zero(header.Number.Uint64())

	_, err = s.cli.L2ParentByBlockId(context.Background(), common.Big2)
	s.Error(err)
}

func (s *RpcTestSuite) TestL2ExecutionEngineSyncProgress() {
	progress, err := s.cli.L2ExecutionEngineSyncProgress(context.Background())
	s.NoError(err)
	s.NotNil(progress)
}

func (s *RpcTestSuite) TestGetProtocolStateVariables() {
	_, err := s.cli.GetProtocolStateVariables(nil)
	s.NoError(err)
}

func (s *RpcTestSuite) TestCheckL1ReorgFromL1Cursor() {
	l1Head, err := s.cli.L1.HeaderByNumber(context.Background(), nil)
	s.NoError(err)

	_, newL1Current, _, err := s.cli.CheckL1ReorgFromL1Cursor(context.Background(), l1Head, l1Head.Number.Uint64())
	s.NoError(err)

	s.Equal(l1Head.Number.Uint64(), newL1Current.Number.Uint64())

	stateVar, err := s.cli.TaikoL1.GetStateVariables(nil)
	s.NoError(err)

	reorged, _, _, err := s.cli.CheckL1ReorgFromL1Cursor(context.Background(), l1Head, stateVar.GenesisHeight)
	s.NoError(err)
	s.False(reorged)

	l1Head.BaseFee = new(big.Int).Add(l1Head.BaseFee, common.Big1)

	reorged, newL1Current, _, err = s.cli.CheckL1ReorgFromL1Cursor(context.Background(), l1Head, stateVar.GenesisHeight)
	s.NoError(err)
	s.True(reorged)
	s.Equal(l1Head.ParentHash, newL1Current.Hash())
}

func (s *RpcTestSuite) TestIsJustSyncedByP2P() {
	_, err := s.cli.IsJustSyncedByP2P(context.Background())
	s.NoError(err)
}

func (s *RpcTestSuite) TestWaitTillL2ExecutionEngineSyncedNewClient() {
	err := s.cli.WaitTillL2ExecutionEngineSynced(context.Background())
	s.NoError(err)
}

func (s *RpcTestSuite) TestWaitTillL2ExecutionEngineSyncedContextErr() {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	err := s.cli.WaitTillL2ExecutionEngineSynced(ctx)
	s.ErrorContains(err, "context canceled")
}

func (s *RpcTestSuite) TestGetPoolContentValid() {
	configs, err := s.cli.TaikoL1.GetConfig(&bind.CallOpts{Context: context.Background()})
	s.NoError(err)
	goldenTouchAddress, err := s.cli.TaikoL2.GOLDENTOUCHADDRESS(nil)
	s.NoError(err)
	parent, err := s.cli.L2.BlockByNumber(context.Background(), nil)
	s.NoError(err)
	baseFee, err := s.cli.TaikoL2.GetBasefee(nil, 1, uint32(parent.GasUsed()))
	s.NoError(err)
	gasLimit := configs.BlockMaxGasLimit
	maxBytes := configs.BlockMaxTxListBytes

	txPools := []common.Address{goldenTouchAddress}

	_, err2 := s.cli.GetPoolContent(
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
	_, err := s.cli.GetStorageRoot(
		context.Background(),
		s.cli.L1GethClient,
		s.L1.TaikoL1SignalService,
		nil)
	s.NoError(err)
}
