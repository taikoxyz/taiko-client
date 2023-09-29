package testutils

import (
	"context"
	"net/url"
	"strings"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/stretchr/testify/suite"
)

const (
	premintTokenAmount = "92233720368547758070000000000000"
)

type ClientSuite struct {
	suite.Suite
	L1              *gethContainer
	L2              *gethContainer
	ProverEndpoints []*url.URL
}

func (s *ClientSuite) SetupSuite() {
}

func (s *ClientSuite) TearDownSuite() {
}

func (s *ClientSuite) SetupTest() {
	s.Reset()
	s.ProverEndpoints = []*url.URL{LocalRandomProverEndpoint()}
}

func (s *ClientSuite) Reset() {
	s.ResetL1()
	s.ResetL2()
}

func (s *ClientSuite) ResetL1() {
	if s.L1 != nil {
		s.NoError(s.L1.Stop())
	}
	s.L1 = s.NewL1()
}

func (s *ClientSuite) ResetL2() {
	if s.L2 != nil {
		s.NoError(s.L2.Stop())
	}
	s.L2 = s.NewL2()
}

func (s *ClientSuite) TearDownTest() {
	s.StopL1()
	s.StopL2()
}

func (s *ClientSuite) StopL1() {
	s.NoError(s.L1.Stop())
	s.L1 = nil
}

func (s *ClientSuite) StopL2() {
	s.NoError(s.L2.Stop())
	s.L2 = nil
}

func (s *ClientSuite) NewL1() *gethContainer {
	name := strings.ReplaceAll(s.T().Name(), "/", "_")
	c, err := newL1Container("L1_" + name)
	s.NoError(err)
	return c
}

func (s *ClientSuite) NewL2() *gethContainer {
	name := strings.ReplaceAll(s.T().Name(), "/", "_")
	c, err := newL2Container("L2_" + name)
	s.NoError(err)
	return c
}

func (s *ClientSuite) SetL1Automine(automine bool) {
	cli, err := rpc.DialContext(context.Background(), s.L1.HttpEndpoint())
	s.NoError(err)
	s.NoError(cli.CallContext(context.Background(), nil, "evm_setAutomine", automine))
	cli.Close()
}
