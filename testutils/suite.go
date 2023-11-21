package testutils

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"net/url"
	"os"

	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/suite"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/pkg/jwt"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	capacity "github.com/taikoxyz/taiko-client/prover/capacity_manager"
	"github.com/taikoxyz/taiko-client/prover/server"
)

type ClientTestSuite struct {
	suite.Suite
	testnetL1SnapshotID string
	RpcClient           *rpc.Client
	TestAddrPrivKey     *ecdsa.PrivateKey
	TestAddr            common.Address
	ProverEndpoints     []*url.URL
	AddressManager      *bindings.AddressManager
	proverServer        *server.ProverServer
}

func (s *ClientTestSuite) SetupTest() {
	// Default logger
	log.Root().SetHandler(
		log.LvlFilterHandler(log.LvlInfo, log.StreamHandler(os.Stdout, log.TerminalFormat(true))),
	)

	if os.Getenv("LOG_LEVEL") != "" {
		level, err := log.LvlFromString(os.Getenv("LOG_LEVEL"))
		if err != nil {
			log.Crit("Invalid log level", "level", os.Getenv("LOG_LEVEL"))
		}
		log.Root().SetHandler(
			log.LvlFilterHandler(level, log.StreamHandler(os.Stdout, log.TerminalFormat(true))),
		)
	}

	testAddrPrivKey, err := crypto.ToECDSA(
		common.Hex2Bytes("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"),
	)
	s.Nil(err)

	s.TestAddrPrivKey = testAddrPrivKey
	s.TestAddr = common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")

	jwtSecret, err := jwt.ParseSecretFromFile(os.Getenv("JWT_SECRET"))
	s.Nil(err)
	s.NotEmpty(jwtSecret)

	rpcCli, err := rpc.NewClient(context.Background(), &rpc.ClientConfig{
		L1Endpoint:            os.Getenv("L1_NODE_WS_ENDPOINT"),
		L2Endpoint:            os.Getenv("L2_EXECUTION_ENGINE_WS_ENDPOINT"),
		TaikoL1Address:        common.HexToAddress(os.Getenv("TAIKO_L1_ADDRESS")),
		TaikoL2Address:        common.HexToAddress(os.Getenv("TAIKO_L2_ADDRESS")),
		TaikoTokenAddress:     common.HexToAddress(os.Getenv("TAIKO_TOKEN_ADDRESS")),
		GuardianProverAddress: common.HexToAddress(os.Getenv("GUARDIAN_PROVER_CONTRACT_ADDRESS")),
		L2EngineEndpoint:      os.Getenv("L2_EXECUTION_ENGINE_AUTH_ENDPOINT"),
		JwtSecret:             string(jwtSecret),
		RetryInterval:         backoff.DefaultMaxInterval,
	})
	s.Nil(err)
	s.RpcClient = rpcCli

	l1ProverPrivKey, err := crypto.ToECDSA(common.Hex2Bytes(os.Getenv("L1_PROVER_PRIVATE_KEY")))
	s.Nil(err)

	s.ProverEndpoints = []*url.URL{LocalRandomProverEndpoint()}
	s.proverServer = NewTestProverServer(s, l1ProverPrivKey, capacity.New(1024), s.ProverEndpoints[0])

	allowance, err := rpcCli.TaikoToken.Allowance(
		nil,
		crypto.PubkeyToAddress(l1ProverPrivKey.PublicKey),
		common.HexToAddress("TAIKO_L1_ADDRESS"),
	)
	s.Nil(err)

	if allowance.Cmp(common.Big0) == 0 {
		// Do not verify zk && sgx proofs in tests.
		addressManager, err := bindings.NewAddressManager(
			common.HexToAddress(os.Getenv("ADDRESS_MANAGER_CONTRACT_ADDRESS")),
			rpcCli.L1,
		)
		s.Nil(err)

		chainID, err := rpcCli.L1.ChainID(context.Background())
		s.Nil(err)

		ownerPrivKey, err := crypto.ToECDSA(common.Hex2Bytes(os.Getenv("L1_CONTRACT_OWNER_PRIVATE_KEY")))
		s.Nil(err)

		opts, err := bind.NewKeyedTransactorWithChainID(ownerPrivKey, rpcCli.L1ChainID)
		s.Nil(err)

		tx, err := addressManager.SetAddress(
			opts,
			chainID.Uint64(),
			rpc.StringToBytes32("tier_sgx_and_pse_zkevm"),
			common.Address{},
		)
		s.Nil(err)
		_, err = rpc.WaitReceipt(context.Background(), rpcCli.L1, tx)
		s.Nil(err)

		tx, err = addressManager.SetAddress(opts, chainID.Uint64(), rpc.StringToBytes32("tier_sgx"), common.Address{})
		s.Nil(err)
		_, err = rpc.WaitReceipt(context.Background(), rpcCli.L1, tx)
		s.Nil(err)

		// Transfer some tokens to provers.
		balance, err := rpcCli.TaikoToken.BalanceOf(nil, crypto.PubkeyToAddress(ownerPrivKey.PublicKey))
		s.Nil(err)
		s.Greater(balance.Cmp(common.Big0), 0)

		opts, err = bind.NewKeyedTransactorWithChainID(ownerPrivKey, rpcCli.L1ChainID)
		s.Nil(err)
		proverBalance := new(big.Int).Div(balance, common.Big2)
		s.Greater(proverBalance.Cmp(common.Big0), 0)

		tx, err = rpcCli.TaikoToken.Transfer(
			opts,
			crypto.PubkeyToAddress(l1ProverPrivKey.PublicKey), proverBalance,
		)
		s.Nil(err)
		_, err = rpc.WaitReceipt(context.Background(), rpcCli.L1, tx)
		s.Nil(err)

		// Deposit taiko tokens for provers.
		opts, err = bind.NewKeyedTransactorWithChainID(l1ProverPrivKey, rpcCli.L1ChainID)
		s.Nil(err)

		_, err = rpcCli.TaikoToken.Approve(opts, common.HexToAddress(os.Getenv("ASSIGNMENT_HOOK_ADDRESS")), proverBalance)
		s.Nil(err)

		_, err = rpc.WaitReceipt(context.Background(), rpcCli.L1, tx)
		s.Nil(err)
	}

	s.Nil(rpcCli.L1RawRPC.CallContext(context.Background(), &s.testnetL1SnapshotID, "evm_snapshot"))
	s.NotEmpty(s.testnetL1SnapshotID)
}

func (s *ClientTestSuite) TearDownTest() {
	var revertRes bool
	s.Nil(s.RpcClient.L1RawRPC.CallContext(context.Background(), &revertRes, "evm_revert", s.testnetL1SnapshotID))
	s.True(revertRes)

	s.Nil(rpc.SetHead(context.Background(), s.RpcClient.L2RawRPC, common.Big0))
	s.Nil(s.proverServer.Shutdown(context.Background()))
}

func (s *ClientTestSuite) SetL1Automine(automine bool) {
	s.Nil(s.RpcClient.L1RawRPC.CallContext(context.Background(), nil, "evm_setAutomine", automine))
}
