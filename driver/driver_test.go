package driver

import (
	"context"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/suite"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
	"github.com/taikoxyz/taiko-client/pkg/jwt"
	"github.com/taikoxyz/taiko-client/proposer"
	"github.com/taikoxyz/taiko-client/testutils"
)

type DriverTestSuite struct {
	testutils.ClientTestSuite
	cancel context.CancelFunc
	p      *proposer.Proposer
	d      *Driver
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
	ctx, cancel := context.WithCancel(context.Background())
	s.cancel = cancel
	s.Nil(InitFromConfig(ctx, d, &Config{
		L1Endpoint:                    os.Getenv("L1_NODE_WS_ENDPOINT"),
		L2Endpoint:                    os.Getenv("L2_EXECUTION_ENGINE_WS_ENDPOINT"),
		L2EngineEndpoint:              os.Getenv("L2_EXECUTION_ENGINE_AUTH_ENDPOINT"),
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

	proposeInterval := 1024 * time.Hour // No need to periodically propose transactions list in unit tests
	s.Nil(proposer.InitFromConfig(context.Background(), p, (&proposer.Config{
		L1Endpoint:              os.Getenv("L1_NODE_WS_ENDPOINT"),
		L2Endpoint:              os.Getenv("L2_EXECUTION_ENGINE_WS_ENDPOINT"),
		TaikoL1Address:          common.HexToAddress(os.Getenv("TAIKO_L1_ADDRESS")),
		TaikoL2Address:          common.HexToAddress(os.Getenv("TAIKO_L2_ADDRESS")),
		L1ProposerPrivKey:       l1ProposerPrivKey,
		L2SuggestedFeeRecipient: common.HexToAddress(os.Getenv("L2_SUGGESTED_FEE_RECIPIENT")),
		ProposeInterval:         &proposeInterval, // No need to periodically propose transactions list in unit tests
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

func (s *DriverTestSuite) TestProcessL1Blocks() {
	l1Head1, err := s.d.rpc.L1.HeaderByNumber(context.Background(), nil)
	s.Nil(err)

	l2Head1, err := s.d.rpc.L2.HeaderByNumber(context.Background(), nil)
	s.Nil(err)

	s.Nil(s.d.ChainSyncer().CalldataSyncer().ProcessL1Blocks(context.Background(), l1Head1))

	// Propose an invalid L2 block
	testutils.ProposeAndInsertThrowawayBlock(&s.ClientTestSuite, s.p, s.d.ChainSyncer().CalldataSyncer())

	l2Head2, err := s.d.rpc.L2.HeaderByNumber(context.Background(), nil)
	s.Nil(err)

	s.Equal(l2Head2.Number.Uint64(), l2Head1.Number.Uint64())

	// Propose a valid L2 block
	testutils.ProposeAndInsertValidBlock(&s.ClientTestSuite, s.p, s.d.ChainSyncer().CalldataSyncer())

	l2Head3, err := s.d.rpc.L2.HeaderByNumber(context.Background(), nil)
	s.Nil(err)

	s.Greater(l2Head3.Number.Uint64(), l2Head2.Number.Uint64())

	// Empty blocks
	testutils.ProposeAndInsertEmptyBlocks(&s.ClientTestSuite, s.p, s.d.ChainSyncer().CalldataSyncer())
	s.Nil(err)

	l2Head4, err := s.d.rpc.L2.HeaderByNumber(context.Background(), nil)
	s.Nil(err)

	s.Equal(l2Head3.Number.Uint64()+2, l2Head4.Number.Uint64())

	for _, height := range []uint64{l2Head4.Number.Uint64(), l2Head4.Number.Uint64() - 1} {
		header, err := s.d.rpc.L2.HeaderByNumber(context.Background(), new(big.Int).SetUint64(height))
		s.Nil(err)

		txCount, err := s.d.rpc.L2.TransactionCount(context.Background(), header.Hash())
		s.Nil(err)
		s.Equal(uint(1), txCount)

		anchorTx, err := s.d.rpc.L2.TransactionInBlock(context.Background(), header.Hash(), 0)
		s.Nil(err)

		method, err := encoding.TaikoL2ABI.MethodById(anchorTx.Data())
		s.Nil(err)
		s.Equal("anchor", method.Name)
	}
}

func (s *DriverTestSuite) TestDoSyncNoNewL2Blocks() {
	s.Nil(s.d.doSync())
}

func (s *DriverTestSuite) TestStartClose() {
	s.Nil(s.d.Start())
	s.cancel()
	s.d.Close()
}

func TestDriverTestSuite(t *testing.T) {
	suite.Run(t, new(DriverTestSuite))
}
