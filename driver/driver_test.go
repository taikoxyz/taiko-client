package driver

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/suite"
	"github.com/taikoxyz/taiko-client/pkg/jwt"
	"github.com/taikoxyz/taiko-client/proposer"
	"github.com/taikoxyz/taiko-client/testutils"
)

type DriverTestSuite struct {
	testutils.ClientTestSuite
	p *proposer.Proposer
	d *Driver
}

func (s *DriverTestSuite) SetupTest() {
	s.ClientTestSuite.SetupTest()

	// Init driver
	jwtSecret, err := jwt.ParseSecretFromFile(os.Getenv("JWT_SECRET"))
	s.Nil(err)
	s.NotEmpty(jwtSecret)

	throwawayBlocksBuilderPrivKey, err := crypto.ToECDSA(
		common.Hex2Bytes(os.Getenv("THROWAWAY_BLOCKS_BUILDER_PRIV_KEY")),
	)
	s.Nil(err)

	d := new(Driver)
	s.Nil(InitFromConfig(context.Background(), d, &Config{
		L1Endpoint:                    os.Getenv("L1_NODE_ENDPOINT"),
		L2Endpoint:                    os.Getenv("L2_NODE_ENDPOINT"),
		L2EngineEndpoint:              os.Getenv("L2_NODE_ENGINE_ENDPOINT"),
		TaikoL1Address:                common.HexToAddress(os.Getenv("TAIKO_L1_ADDRESS")),
		TaikoL2Address:                common.HexToAddress(os.Getenv("TAIKO_L2_ADDRESS")),
		ThrowawayBlocksBuilderPrivKey: throwawayBlocksBuilderPrivKey,
		JwtSecret:                     string(jwtSecret),
	}))
	s.d = d

	// Init proposer
	p := new(proposer.Proposer)

	l1ProposerPrivKey, err := crypto.ToECDSA(common.Hex2Bytes(os.Getenv("L1_PROPOSER_PRIVATE_KEY")))
	s.Nil(err)

	s.Nil(proposer.InitFromConfig(context.Background(), p, (&proposer.Config{
		L1Endpoint:              os.Getenv("L1_NODE_ENDPOINT"),
		L2Endpoint:              os.Getenv("L2_NODE_ENDPOINT"),
		TaikoL1Address:          common.HexToAddress(os.Getenv("TAIKO_L1_ADDRESS")),
		TaikoL2Address:          common.HexToAddress(os.Getenv("TAIKO_L2_ADDRESS")),
		L1ProposerPrivKey:       l1ProposerPrivKey,
		L2SuggestedFeeRecipient: common.HexToAddress(os.Getenv("L2_SUGGESTED_FEE_RECIPIENT")),
		ProposeInterval:         1024 * time.Hour, // No need to periodically propose transactions list in unit tests
	})))
	s.p = p
	s.p.AfterCommitHook = s.MineL1Confirmations
}

func (s *DriverTestSuite) TestName() {
	s.Equal("driver", s.d.Name())
}

func (s *DriverTestSuite) MineL1Confirmations() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return s.d.rpc.L1RawRPC.CallContext(ctx, nil, "hardhat_mine", hexutil.EncodeUint64(4))
}

func (s *DriverTestSuite) TestDoSyncNoNewL2Blocks() {
	s.Nil(s.d.doSync())
}

func TestDriverTestSuite(t *testing.T) {
	suite.Run(t, new(DriverTestSuite))
}
