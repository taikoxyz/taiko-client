package testutils

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/suite"
	"github.com/taikoxyz/taiko-client/bindings"
)

func (s *ClientSuite) TestDocker() {
	ctx := context.Background()
	endpoint := s.l1Container.HttpEndpoint()
	cli, err := ethclient.DialContext(ctx, endpoint)
	s.NoError(err)
	defer cli.Close()
	taikoL1, err := bindings.NewTaikoL1Client(common.HexToAddress("0x0DCd1Bf9A1b36cE34237eEaFef220932846BCD82"), cli)
	s.NoError(err)
	stateVars, err := taikoL1.GetStateVariables(nil)
	s.NoError(err)
	s.T().Logf("state vars: %v", stateVars.GenesisHeight)
}

func TestClientSuite(t *testing.T) {
	suite.Run(t, new(ClientSuite))
}
