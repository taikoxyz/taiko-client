package testutils

import (
	"context"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/stretchr/testify/suite"
)

const (
	premintTokenAmount = "92233720368547758070000000000000"
)

type ClientTestSuite struct {
	suite.Suite
	testnetL1SnapshotID string
	L1                  *gethContainer
	L2                  *gethContainer
	rpcCliL1            *rpc.Client
}

func (s *ClientTestSuite) SetupSuite() {
	s.L1 = s.NewL1(s.getTestName())
	s.L2 = s.NewL2(s.getTestName())
	var err error
	s.rpcCliL1, err = rpc.DialContext(context.Background(), s.L1.HttpEndpoint())
	s.NoError(err)
}

func (s *ClientTestSuite) TearDownSuite() {
	s.StopL1()
	s.StopL2()
}

func (s *ClientTestSuite) SetupTest() {
	s.NoError(s.rpcCliL1.CallContext(context.Background(), &s.testnetL1SnapshotID, "evm_snapshot"))
	s.NotEmpty(s.testnetL1SnapshotID)
}

func (s *ClientTestSuite) TearDownTest() {
	var revertRes bool
	s.NoError(s.rpcCliL1.CallContext(context.Background(), &revertRes, "evm_revert", s.testnetL1SnapshotID))
	s.True(revertRes)
	rpcCli, err := rpc.Dial(s.L2.HttpEndpoint())
	s.NoError(err)
	defer rpcCli.Close()
	s.NoError(gethclient.New(rpcCli).SetHead(context.Background(), common.Big0))
}

func (s *ClientTestSuite) getTestName() string {
	return strings.ReplaceAll(s.T().Name(), "/", "_")
}

func (s *ClientTestSuite) StopL1() {
	s.NoError(s.L1.Stop())
	s.L1 = nil
}

func (s *ClientTestSuite) StopL2() {
	s.NoError(s.L2.Stop())
	s.L2 = nil
}

func (s *ClientTestSuite) NewL1(name string) *gethContainer {
	c, err := newL1Container("L1_" + name)
	s.NoError(err)
	return c
}

func (s *ClientTestSuite) NewL2(name string) *gethContainer {
	c, err := newL2Container("L2_" + name)
	s.NoError(err)
	return c
}

func (s *ClientTestSuite) SetL1Automine(automine bool) {
	cli, err := rpc.DialContext(context.Background(), s.L1.HttpEndpoint())
	s.NoError(err)
	s.NoError(cli.CallContext(context.Background(), nil, "evm_setAutomine", automine))
	cli.Close()
}
