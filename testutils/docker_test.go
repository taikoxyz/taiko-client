package testutils

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/suite"
	"github.com/taikoxyz/taiko-client/bindings"
)

func (s *ClientTestSuite) TestDocker() {
	ctx := context.Background()
	endpoint := s.L1.HttpEndpoint()
	cli, err := ethclient.DialContext(ctx, endpoint)
	s.NoError(err)
	defer cli.Close()
	taikoL1, err := bindings.NewTaikoL1Client(s.L1.TaikoL1Address, cli)
	s.NoError(err)
	stateVars, err := taikoL1.GetStateVariables(nil)
	s.NoError(err)
	s.T().Logf("state vars: %v", stateVars.GenesisHeight)
	latest, err := cli.HeaderByNumber(ctx, nil)
	s.NoError(err)
	s.T().Logf("latest block: %v", latest.Number)
}

func TestClientSuite(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}
