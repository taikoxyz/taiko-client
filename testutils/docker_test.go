package testutils

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/suite"
)

func (s *ExampleTestSuite) TestDocker() {
	ctx := context.Background()
	cli, err := ethclient.DialContext(ctx, s.l1ContainerConf.Ports.HttpEndpoint())
	s.NoError(err)
	defer cli.Close()
	genesis, err := cli.BlockByNumber(ctx, big.NewInt(0))
	s.NoError(err)
	s.T().Logf("genesis hash: %s", genesis.Hash().String())
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuite))
}
