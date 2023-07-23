package driver

import (
	"context"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
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

	d := new(Driver)
	ctx, cancel := context.WithCancel(context.Background())
	s.cancel = cancel
	s.Nil(InitFromConfig(ctx, d, &Config{
		L1Endpoint:       os.Getenv("L1_NODE_WS_ENDPOINT"),
		L2Endpoint:       os.Getenv("L2_EXECUTION_ENGINE_WS_ENDPOINT"),
		L2EngineEndpoint: os.Getenv("L2_EXECUTION_ENGINE_AUTH_ENDPOINT"),
		TaikoL1Address:   common.HexToAddress(os.Getenv("TAIKO_L1_ADDRESS")),
		TaikoL2Address:   common.HexToAddress(os.Getenv("TAIKO_L2_ADDRESS")),
		JwtSecret:        string(jwtSecret),
	}))
	s.d = d

	// Init proposer
	p := new(proposer.Proposer)

	l1ProposerPrivKey, err := crypto.ToECDSA(common.Hex2Bytes(os.Getenv("L1_PROPOSER_PRIVATE_KEY")))
	s.Nil(err)

	proposeInterval := 1024 * time.Hour // No need to periodically propose transactions list in unit tests
	s.Nil(proposer.InitFromConfig(context.Background(), p, (&proposer.Config{
		L1Endpoint:                 os.Getenv("L1_NODE_WS_ENDPOINT"),
		L2Endpoint:                 os.Getenv("L2_EXECUTION_ENGINE_WS_ENDPOINT"),
		TaikoL1Address:             common.HexToAddress(os.Getenv("TAIKO_L1_ADDRESS")),
		TaikoL2Address:             common.HexToAddress(os.Getenv("TAIKO_L2_ADDRESS")),
		L1ProposerPrivKey:          l1ProposerPrivKey,
		L2SuggestedFeeRecipient:    common.HexToAddress(os.Getenv("L2_SUGGESTED_FEE_RECIPIENT")),
		ProposeInterval:            &proposeInterval, // No need to periodically propose transactions list in unit tests
		MaxProposedTxListsPerEpoch: 1,
	})))
	s.p = p
}

func (s *DriverTestSuite) TestName() {
	s.Equal("driver", s.d.Name())
}

func (s *DriverTestSuite) TestProcessL1Blocks() {
	l1Head1, err := s.d.rpc.L1.HeaderByNumber(context.Background(), nil)
	s.Nil(err)

	l2Head1, err := s.d.rpc.L2.HeaderByNumber(context.Background(), nil)
	s.Nil(err)

	s.Nil(s.d.ChainSyncer().CalldataSyncer().ProcessL1Blocks(context.Background(), l1Head1))

	// Propose a valid L2 block
	testutils.ProposeAndInsertValidBlock(&s.ClientTestSuite, s.p, s.d.ChainSyncer().CalldataSyncer())

	l2Head2, err := s.d.rpc.L2.HeaderByNumber(context.Background(), nil)
	s.Nil(err)

	s.Greater(l2Head2.Number.Uint64(), l2Head1.Number.Uint64())

	// Empty blocks
	testutils.ProposeAndInsertEmptyBlocks(&s.ClientTestSuite, s.p, s.d.ChainSyncer().CalldataSyncer())
	s.Nil(err)

	l2Head3, err := s.d.rpc.L2.HeaderByNumber(context.Background(), nil)
	s.Nil(err)

	s.Greater(l2Head3.Number.Uint64(), l2Head2.Number.Uint64())

	for _, height := range []uint64{l2Head3.Number.Uint64(), l2Head3.Number.Uint64() - 1} {
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

func (s *DriverTestSuite) TestCheckL1Reorg() {
	var testnetL1SnapshotID string
	s.Nil(s.RpcClient.L1RawRPC.CallContext(context.Background(), &testnetL1SnapshotID, "evm_snapshot"))
	s.NotEmpty(testnetL1SnapshotID)

	l1Head1, err := s.d.rpc.L1.HeaderByNumber(context.Background(), nil)
	s.Nil(err)
	l2Head1, err := s.d.rpc.L2.HeaderByNumber(context.Background(), nil)
	s.Nil(err)

	// Propose two L2 blocks
	testutils.ProposeAndInsertValidBlock(&s.ClientTestSuite, s.p, s.d.ChainSyncer().CalldataSyncer())
	testutils.ProposeAndInsertValidBlock(&s.ClientTestSuite, s.p, s.d.ChainSyncer().CalldataSyncer())

	l1Head2, err := s.d.rpc.L1.HeaderByNumber(context.Background(), nil)
	s.Nil(err)
	l2Head2, err := s.d.rpc.L2.HeaderByNumber(context.Background(), nil)
	s.Nil(err)
	s.Greater(l2Head2.Number.Uint64(), l2Head1.Number.Uint64())
	s.Greater(l1Head2.Number.Uint64(), l1Head1.Number.Uint64())

	reorged, _, _, err := s.RpcClient.CheckL1ReorgFromL2EE(context.Background(), l2Head2.Number)
	s.Nil(err)
	s.False(reorged)

	// Reorg back to l2Head1
	var revertRes bool
	s.Nil(s.RpcClient.L1RawRPC.CallContext(context.Background(), &revertRes, "evm_revert", testnetL1SnapshotID))
	s.True(revertRes)

	l1Head3, err := s.d.rpc.L1.HeaderByNumber(context.Background(), nil)
	s.Nil(err)
	s.Equal(l1Head3.Number.Uint64(), l1Head1.Number.Uint64())
	s.Equal(l1Head3.Hash(), l1Head1.Hash())

	// Propose two blocks in another fork
	testutils.ProposeInvalidTxListBytes(&s.ClientTestSuite, s.p)
	testutils.ProposeInvalidTxListBytes(&s.ClientTestSuite, s.p)

	l1Head4, err := s.d.rpc.L1.HeaderByNumber(context.Background(), nil)
	s.Nil(err)

	s.Nil(s.d.ChainSyncer().CalldataSyncer().ProcessL1Blocks(context.Background(), l1Head4))

	l2Head3, err := s.d.rpc.L2.HeaderByNumber(context.Background(), nil)
	s.Nil(err)

	parent, err := s.d.rpc.L2.HeaderByHash(context.Background(), l2Head3.ParentHash)
	s.Nil(err)
	s.Equal(l2Head3.Number.Uint64(), l2Head2.Number.Uint64())
	s.Equal(parent.ParentHash, l2Head1.Hash())
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
