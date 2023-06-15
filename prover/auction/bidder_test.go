package auction

import (
	"context"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/suite"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	"github.com/taikoxyz/taiko-client/testutils"
)

type BidderTestSuite struct {
	testutils.ClientTestSuite
	b *Bidder
}

func (s *BidderTestSuite) SetupTest() {
	s.ClientTestSuite.SetupTest()

	l1ProverPrivKey, err := crypto.ToECDSA(common.Hex2Bytes(os.Getenv("L1_PROVER_PRIVATE_KEY")))
	s.Nil(err)

	proverAddress := crypto.PubkeyToAddress(l1ProverPrivKey.PublicKey)

	l1TaikoTokenOwnerPrivKey, err := crypto.ToECDSA(common.Hex2Bytes(os.Getenv("TAIKO_TOKEN_OWNER_PRIVATE_KEY")))
	s.Nil(err)

	testutils.TransferTaikoToken(
		&s.ClientTestSuite,
		common.HexToAddress(os.Getenv("L1_TAIKO_TOKEN_ADDRESS")),
		proverAddress,
		big.NewInt(1000000000000000),
		l1TaikoTokenOwnerPrivKey)

	// Clients
	rpcClient, err := rpc.NewClient(context.Background(), &rpc.ClientConfig{
		L1Endpoint:     os.Getenv("L1_NODE_HTTP_ENDPOINT"),
		L2Endpoint:     os.Getenv("L2_EXECUTION_ENGINE_HTTP_ENDPOINT"),
		TaikoL1Address: common.HexToAddress(os.Getenv("TAIKO_L1_ADDRESS")),
		TaikoL2Address: common.HexToAddress(os.Getenv("TAIKO_L2_ADDRESS")),
		RetryInterval:  10,
	})

	s.Nil(err)

	b := NewBidder(NewAlwaysBidStrategy(), rpcClient, l1ProverPrivKey, proverAddress)
	s.b = b
}

func (s *BidderTestSuite) TestSubmitBid_AlwaysBid() {
	s.Nil(s.b.SubmitBid(context.Background(), big.NewInt(1)))
}
func TestBidderTestSuite(t *testing.T) {
	suite.Run(t, new(BidderTestSuite))
}
