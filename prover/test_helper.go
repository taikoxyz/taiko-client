package prover

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/go-resty/resty/v2"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/pkg/jwt"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	capacity "github.com/taikoxyz/taiko-client/prover/capacity_manager"
	"github.com/taikoxyz/taiko-client/prover/server"
	"github.com/taikoxyz/taiko-client/testutils"
)

// FakeProverServer starts a new prover server that has channel listeners to respond and react
// to requests for capacity, which provers can call.
func FakeProverServer(
	s *testutils.ClientSuite,
	proverPrivKey *ecdsa.PrivateKey,
	capacityManager *capacity.CapacityManager,
	url *url.URL,
) *server.ProverServer {
	cli, err := ethclient.Dial(s.L1.WsEndpoint())
	s.NoError(err)
	taikoL1, err := bindings.NewTaikoL1Client(testutils.TaikoL1Address, cli)
	s.NoError(err)
	protocolConfig, err := taikoL1.GetConfig(nil)
	s.Nil(err)
	jwtSecret, err := jwt.ParseSecretFromFile(testutils.JwtSecretFile)
	s.NoError(err)
	rpcClient, err := rpc.NewClient(context.Background(), &rpc.ClientConfig{
		L1Endpoint:        s.L1.WsEndpoint(),
		L2Endpoint:        s.L2.WsEndpoint(),
		TaikoL1Address:    testutils.TaikoL1Address,
		TaikoTokenAddress: testutils.TaikoL1TokenAddress,
		TaikoL2Address:    testutils.TaikoL2Address,
		L2EngineEndpoint:  s.L2.AuthEndpoint(),
		JwtSecret:         string(jwtSecret),
		RetryInterval:     backoff.DefaultMaxInterval,
	})
	s.NoError(err)
	srv, err := server.New(&server.NewProverServerOpts{
		ProverPrivateKey: proverPrivKey,
		MinProofFee:      common.Big1,
		MaxExpiry:        24 * time.Hour,
		CapacityManager:  capacityManager,
		TaikoL1Address:   testutils.TaikoL1Address,
		Rpc:              rpcClient,
		Bond:             protocolConfig.ProofBond,
		IsOracle:         true,
	})
	s.NoError(err)

	go func() {
		if err := srv.Start(fmt.Sprintf(":%v", url.Port())); !errors.Is(err, http.ErrServerClosed) {
			log.Error("Failed to start prover server", "error", err)
		}
	}()

	// Wait till the server fully started.
	s.Nil(backoff.Retry(func() error {
		res, err := resty.New().R().Get(url.String() + "/healthz")
		if err != nil {
			return err
		}
		if !res.IsSuccess() {
			return fmt.Errorf("invalid response status code: %d", res.StatusCode())
		}

		return nil
	}, backoff.NewExponentialBackOff()))

	return srv
}
