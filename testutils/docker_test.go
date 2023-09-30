package testutils

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/suite"
	"github.com/taikoxyz/taiko-client/bindings"
)

func (s *ClientSuite) TestDocker() {
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
	suite.Run(t, new(ClientSuite))
}

// func TestSome(t *testing.T) {
// 	ctx := context.Background()
// 	base := "http://localhost:34519"
// 	l1 := "http://localhost:34531"
// 	printInfo(t, ctx, base)
// 	printInfo(t, ctx, l1)
// }

// func printInfo(t *testing.T, ctx context.Context, url string) {
// 	r, err := rpc.Dial(url)
// 	assert.NoError(t, err)
// 	e := ethclient.NewClient(r)
// 	latest, err := e.HeaderByNumber(ctx, nil)
// 	assert.NoError(t, err)
// 	// target := (big.NewInt(0)).Sub(latest.Number, common.Big1)
// 	target := latest.Number
// 	t.Logf("latest block: %v", target)
// 	g := gethclient.New(r)
// 	keys := []string{"0x0000000000000000000000000000000000000000000000000000000000000000"}
// 	_, err = g.GetProof(ctx, TaikoL1SignalService, keys, target)
// 	assert.NoError(t, err)
// }
