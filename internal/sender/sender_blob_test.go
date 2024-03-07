package sender

import (
	"context"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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

func (s *SenderBlobTestSuite) makeBlobTx(opts *bind.TransactOpts) *types.BlobTx {
	data, err := os.ReadFile("./sender.go")
	s.Nil(err)
	sidecar, err := rpc.MakeSidecar(data)
	s.Nil(err)
	blobTx, err := s.client.CreateBlobTx(opts, nil, nil, sidecar)
	s.Nil(err)
	return blobTx
}

func (s *SenderBlobTestSuite) TestSendTransaction() {
	var (
		sender = s.sender
		eg     errgroup.Group
	)
	blobTx := s.makeBlobTx(sender.GetOpts())

	for i := 0; i < 16; i++ {
		eg.Go(func() error {
			_, err := sender.SendTransaction(types.NewTx(blobTx))
			return err
		})
	}
	s.Nil(eg.Wait())

	for _, confirmCh := range sender.TxToConfirmChannels() {
		confirm := <-confirmCh
		s.Nil(confirm.Err)

		_, err := s.beaconClient.GetBlobs(context.Background(), confirm.Receipt.BlockNumber)
		s.Nil(err)
	}
}

func (s *SenderBlobTestSuite) TestNonce() {
	send := s.sender
	client := s.client
	opts := send.GetOpts()

	// Let max gas price be 2 times of the gas fee cap.
	send.MaxGasFee = opts.GasFeeCap.Uint64() * 2

	nonce, err := client.NonceAt(context.Background(), opts.From, nil)
	s.Nil(err)

	blobTx := s.makeBlobTx(opts)

	_, err = send.SendRawTransaction(nonce+1, &common.Address{}, nil, nil, blobTx.Sidecar)
	s.Equal(true, strings.Contains(err.Error(), "nonce too high"))

	txID, err := send.SendRawTransaction(nonce-1, &common.Address{}, nil, nil, blobTx.Sidecar)
	s.Nil(err)
	confirm := <-send.TxToConfirmChannel(txID)
	s.Nil(confirm.Err)
	s.Equal(nonce, confirm.CurrentTx.Nonce())
}

func (s *SenderBlobTestSuite) TearDownTest() {
	s.sender.Close()
	s.client.Close()
}

func TestSenderBlobTestSuite(t *testing.T) {
	suite.Run(t, new(SenderBlobTestSuite))
}
