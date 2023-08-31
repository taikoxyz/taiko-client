package selector

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/suite"
	"github.com/taikoxyz/taiko-client/prover/http"
	"github.com/taikoxyz/taiko-client/testutils"
)

type ProverSelectorTestSuite struct {
	testutils.ClientTestSuite
	s             *ETHFeeSelector
	proverAddress common.Address
	srv           *http.Server
}

func (s *ProverSelectorTestSuite) SetupTest() {
	s.ClientTestSuite.SetupTest()

	proverEndpoint := testutils.LocalRandomProverEndpoint()

	l1ProverPrivKey, err := crypto.ToECDSA(common.Hex2Bytes(os.Getenv("L1_PROVER_PRIVATE_KEY")))
	s.Nil(err)

	srv, err := http.NewServer(http.NewServerOpts{
		ProverPrivateKey:         l1ProverPrivKey,
		MinProofFee:              common.Big1,
		MaxCapacity:              10,
		RequestCurrentCapacityCh: make(chan struct{}),
		ReceiveCurrentCapacityCh: make(chan uint64),
	})
	s.Nil(err)
	s.srv = srv
	s.proverAddress = crypto.PubkeyToAddress(l1ProverPrivKey.PublicKey)

	go func() {
		if err := s.srv.Start(fmt.Sprintf(":%v", proverEndpoint.Port())); err != nil {
			log.Crit("error starting prover http server", "error", err)
		}
	}()

	protocolConfigs, err := s.RpcClient.TaikoL1.GetConfig(nil)
	s.Nil(err)

	s.s, err = NewETHFeeSelector(
		&protocolConfigs,
		s.RpcClient,
		common.HexToAddress(os.Getenv("TAIKO_L1_ADDRESS")),
		common.Big256,
		common.Big2,
		[]*url.URL{proverEndpoint},
		32,
		1*time.Minute,
		1*time.Minute,
	)
	s.Nil(err)
}

func (s *ProverSelectorTestSuite) TestCheckProverBalance() {
	ok, err := s.s.checkProverBalance(context.Background(), s.proverAddress)
	s.Nil(err)
	s.True(ok)
}

func TestProverSelectorTestSuite(t *testing.T) {
	suite.Run(t, new(ProverSelectorTestSuite))
}
