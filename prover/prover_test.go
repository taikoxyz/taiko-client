package prover

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/suite"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/driver"
	"github.com/taikoxyz/taiko-client/pkg/jwt"
	"github.com/taikoxyz/taiko-client/proposer"
	"github.com/taikoxyz/taiko-client/testutils"
)

type ProverTestSuite struct {
	testutils.ClientTestSuite
	p        *Prover
	d        *driver.Driver
	proposer *proposer.Proposer
}

func (s *ProverTestSuite) SetupTest() {
	s.ClientTestSuite.SetupTest()

	// Init prover
	l1ProverPrivKey, err := crypto.ToECDSA(common.Hex2Bytes(os.Getenv("L1_PROVER_PRIVATE_KEY")))
	s.Nil(err)

	p := new(Prover)
	s.Nil(InitFromConfig(context.Background(), p, (&Config{
		L1Endpoint:      os.Getenv("L1_NODE_ENDPOINT"),
		L2Endpoint:      os.Getenv("L2_NODE_ENDPOINT"),
		TaikoL1Address:  common.HexToAddress(os.Getenv("TAIKO_L1_ADDRESS")),
		TaikoL2Address:  common.HexToAddress(os.Getenv("TAIKO_L2_ADDRESS")),
		L1ProverPrivKey: l1ProverPrivKey,
		Dummy:           true,
	})))
	s.p = p

	// Init driver
	jwtSecret, err := jwt.ParseSecretFromFile(os.Getenv("JWT_SECRET"))
	s.Nil(err)
	s.NotEmpty(jwtSecret)

	throwawayBlocksBuilderPrivKey, err := crypto.ToECDSA(
		common.Hex2Bytes(os.Getenv("THROWAWAY_BLOCKS_BUILDER_PRIV_KEY")),
	)
	s.Nil(err)

	d := new(driver.Driver)
	s.Nil(driver.InitFromConfig(context.Background(), d, &driver.Config{
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
	l1ProposerPrivKey, err := crypto.ToECDSA(common.Hex2Bytes(os.Getenv("L1_PROPOSER_PRIVATE_KEY")))
	s.Nil(err)

	prop := new(proposer.Proposer)

	s.Nil(proposer.InitFromConfig(context.Background(), prop, (&proposer.Config{
		L1Endpoint:              os.Getenv("L1_NODE_ENDPOINT"),
		L2Endpoint:              os.Getenv("L2_NODE_ENDPOINT"),
		TaikoL1Address:          common.HexToAddress(os.Getenv("TAIKO_L1_ADDRESS")),
		TaikoL2Address:          common.HexToAddress(os.Getenv("TAIKO_L2_ADDRESS")),
		L1ProposerPrivKey:       l1ProposerPrivKey,
		L2SuggestedFeeRecipient: common.HexToAddress(os.Getenv("L2_SUGGESTED_FEE_RECIPIENT")),
		ProposeInterval:         1024 * time.Hour, // No need to periodically propose transactions list in unit tests
	})))

	s.proposer = prop
	s.proposer.AfterCommitHook = s.MineL1Confirmations
}

func (s *ProverTestSuite) TestName() {
	s.Equal("prover", s.p.Name())
}

func (s *ProverTestSuite) TestGetProveBlocksTxOpts() {
	opts, err := s.p.getProveBlocksTxOpts(context.Background(), s.RpcClient.L1)
	s.Nil(err)
	s.Equal(proveBlocksGasLimit, opts.GasLimit)
}

func (s *ProverTestSuite) TestOnBlockProposed() {
	// Valid block
	e := testutils.ProposeAndInsertValidBlock(&s.ClientTestSuite, s.proposer, s.d.ChainInserter())
	s.Nil(s.p.onBlockProposed(context.Background(), e, func() {}))
	s.Nil(s.p.submitValidBlockProof(context.Background(), <-s.p.proveValidProofCh))

	// Invalid block
	e = testutils.ProposeAndInsertThrowawayBlock(&s.ClientTestSuite, s.proposer, s.d.ChainInserter())
	s.Nil(s.p.onBlockProposed(context.Background(), e, func() {}))
	s.Nil(s.p.submitInvalidBlockProof(context.Background(), <-s.p.proveInvalidProofCh))
}

func (s *ProverTestSuite) TestOnBlockProposedTxNotFound() {
	s.ErrorContains(
		s.p.onBlockProposed(context.Background(), &bindings.TaikoL1ClientBlockProposed{
			Id:  common.Big2,
			Raw: types.Log{BlockHash: common.Hash{}, TxIndex: 0},
		}, func() {}),
		ethereum.NotFound.Error(),
	)
}

func (s *ProverTestSuite) TestOnBlockVerifiedEmptyBlockHash() {
	s.Nil(s.p.onBlockVerified(context.Background(), &bindings.TaikoL1ClientBlockVerified{BlockHash: common.Hash{}}))
}

func (s *ProverTestSuite) TestIsWhitelisted() {
	isWhitelisted, err := s.p.isWhitelisted(crypto.PubkeyToAddress(s.p.cfg.L1ProverPrivKey.PublicKey))
	s.Nil(err)
	s.True(isWhitelisted)
}

func TestProverTestSuite(t *testing.T) {
	suite.Run(t, new(ProverTestSuite))
}
