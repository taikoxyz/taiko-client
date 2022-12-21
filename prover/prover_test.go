package prover

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/suite"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/driver"
	"github.com/taikoxyz/taiko-client/pkg/jwt"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
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

	// Whitelist current prover
	whitelisted, err := s.RpcClient.IsProverWhitelisted(crypto.PubkeyToAddress(l1ProverPrivKey.PublicKey))
	s.Nil(err)

	if !whitelisted {
		l1ContractOwnerPrivateKey, err := crypto.ToECDSA(common.Hex2Bytes(os.Getenv("L1_CONTRACT_OWNER_PRIVATE_KEY")))
		s.Nil(err)

		opts, err := bind.NewKeyedTransactorWithChainID(l1ContractOwnerPrivateKey, s.RpcClient.L1ChainID)
		s.Nil(err)
		opts.GasTipCap = rpc.FallbackGasTipCap

		tx, err := s.RpcClient.TaikoL1.WhitelistProver(opts, crypto.PubkeyToAddress(l1ProverPrivKey.PublicKey), true)
		s.Nil(err)

		receipt, err := rpc.WaitReceipt(context.Background(), s.RpcClient.L1, tx)
		s.Nil(err)
		s.Equal(types.ReceiptStatusSuccessful, receipt.Status)
	}

	p := new(Prover)
	s.Nil(InitFromConfig(context.Background(), p, (&Config{
		L1Endpoint:               os.Getenv("L1_NODE_ENDPOINT"),
		L2Endpoint:               os.Getenv("L2_EXECUTION_ENGINE_ENDPOINT"),
		TaikoL1Address:           common.HexToAddress(os.Getenv("TAIKO_L1_ADDRESS")),
		TaikoL2Address:           common.HexToAddress(os.Getenv("TAIKO_L2_ADDRESS")),
		L1ProverPrivKey:          l1ProverPrivKey,
		Dummy:                    true,
		MaxConcurrentProvingJobs: 1,
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
		L2Endpoint:                    os.Getenv("L2_EXECUTION_ENGINE_ENDPOINT"),
		L2EngineEndpoint:              os.Getenv("L2_EXECUTION_ENGINE_AUTH_ENDPOINT"),
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

	proposeInterval := 1024 * time.Hour // No need to periodically propose transactions list in unit tests
	s.Nil(proposer.InitFromConfig(context.Background(), prop, (&proposer.Config{
		L1Endpoint:              os.Getenv("L1_NODE_ENDPOINT"),
		L2Endpoint:              os.Getenv("L2_EXECUTION_ENGINE_ENDPOINT"),
		TaikoL1Address:          common.HexToAddress(os.Getenv("TAIKO_L1_ADDRESS")),
		TaikoL2Address:          common.HexToAddress(os.Getenv("TAIKO_L2_ADDRESS")),
		L1ProposerPrivKey:       l1ProposerPrivKey,
		L2SuggestedFeeRecipient: common.HexToAddress(os.Getenv("L2_SUGGESTED_FEE_RECIPIENT")),
		ProposeInterval:         &proposeInterval, // No need to periodically propose transactions list in unit tests
	})))

	s.proposer = prop
	s.proposer.AfterCommitHook = s.MineL1Confirmations
}

func (s *ProverTestSuite) TestName() {
	s.Equal("prover", s.p.Name())
}

func (s *ProverTestSuite) TestGetProveBlocksTxOpts() {
	_, err := s.p.getProveBlocksTxOpts(context.Background(), s.RpcClient.L1)
	s.Nil(err)
}

func (s *ProverTestSuite) TestOnBlockProposed() {
	// Valid block
	e := testutils.ProposeAndInsertValidBlock(&s.ClientTestSuite, s.proposer, s.d.ChainSyncer())
	s.Nil(s.p.onBlockProposed(context.Background(), e, func() {}))
	s.Nil(s.p.submitValidBlockProof(context.Background(), <-s.p.proveValidProofCh))

	// Empty blocks
	for _, e = range testutils.ProposeAndInsertEmptyBlocks(&s.ClientTestSuite, s.proposer, s.d.ChainSyncer()) {
		s.Nil(s.p.onBlockProposed(context.Background(), e, func() {}))
		s.Nil(s.p.submitValidBlockProof(context.Background(), <-s.p.proveValidProofCh))
	}

	// Invalid block
	e = testutils.ProposeAndInsertThrowawayBlock(&s.ClientTestSuite, s.proposer, s.d.ChainSyncer())
	s.Nil(s.p.onBlockProposed(context.Background(), e, func() {}))
	s.Nil(s.p.submitInvalidBlockProof(context.Background(), <-s.p.proveInvalidProofCh))
}

func (s *ProverTestSuite) TestOnBlockVerifiedEmptyBlockHash() {
	s.Nil(s.p.onBlockVerified(context.Background(), &bindings.TaikoL1ClientBlockVerified{
		Id:        common.Big1,
		BlockHash: common.Hash{}},
	))
}

func (s *ProverTestSuite) TestIsSubmitProofTxErrorRetryable() {
	s.True(isSubmitProofTxErrorRetryable(errors.New(testAddr.String())))
	s.False(isSubmitProofTxErrorRetryable(errors.New("L1:proof:tooMany")))
	s.False(isSubmitProofTxErrorRetryable(errors.New("L1:tooLate")))
	s.False(isSubmitProofTxErrorRetryable(errors.New("L1:prover:dup")))
	s.False(isSubmitProofTxErrorRetryable(errors.New("L1:" + testAddr.String())))
}

func TestProverTestSuite(t *testing.T) {
	suite.Run(t, new(ProverTestSuite))
}
