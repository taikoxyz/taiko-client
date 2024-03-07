package sender

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/suite"
	"golang.org/x/sync/errgroup"

	"github.com/taikoxyz/taiko-client/internal/utils"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
)

type SenderBlobTestSuite struct {
	suite.Suite
	client       *rpc.EthClient
	sender       *Sender
	beaconClient *rpc.BeaconClient
}

func (s *SenderBlobTestSuite) SetupTest() {
	utils.LoadEnv()
	var err error
	s.client, err = rpc.NewEthClient(context.Background(), os.Getenv("BLOB_GETH_NODE_ENDPOINT"), time.Second*30)
	s.Nil(err)
	priv, err := crypto.ToECDSA(common.FromHex(os.Getenv("L1_PROPOSER_PRIVATE_KEY")))
	s.Nil(err)
	s.sender, err = NewSender(context.Background(), nil, s.client, priv)
	s.Nil(err)
	s.beaconClient, err = rpc.NewBeaconClient(os.Getenv("BLOB_BEACON_NODE_ENDPOINT"), time.Second*30)
	s.Nil(err)
}

func (s *SenderBlobTestSuite) TestSendTransaction() {
	var (
		sender = s.sender
		eg     errgroup.Group
	)
	data, err := os.ReadFile("./sender.go")
	s.Nil(err)
	sidecar, err := rpc.MakeSidecar(data)
	s.Nil(err)
	tx, err := s.client.TransactBlobTx(sender.GetOpts(), nil, nil, sidecar)
	s.Nil(err)

	for i := 0; i < 16; i++ {
		eg.Go(func() error {
			_, err = sender.SendTransaction(tx)
			return err
		})
	}
	s.Nil(eg.Wait())

	for _, confirmCh := range sender.TxToConfirmChannels() {
		confirm := <-confirmCh
		s.Nil(confirm.Err)

		_, err = s.beaconClient.GetBlobs(context.Background(), confirm.Receipt.BlockNumber)
		s.Nil(err)
	}
}

func (s *SenderBlobTestSuite) TearDownTest() {
	s.sender.Close()
	s.client.Close()
}

func TestSenderBlobTestSuite(t *testing.T) {
	suite.Run(t, new(SenderBlobTestSuite))
}
