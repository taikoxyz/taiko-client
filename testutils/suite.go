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
	glogger := log.NewGlogHandler(log.NewTerminalHandlerWithLevel(os.Stdout, log.LevelInfo, true))
	log.SetDefault(log.NewLogger(glogger))

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

	l1ProverPrivKey, err := crypto.ToECDSA(common.FromHex(os.Getenv("L1_PROVER_PRIVATE_KEY")))
	s.Nil(err)

	s.ProverEndpoints = []*url.URL{LocalRandomProverEndpoint()}
	s.proverServer = NewTestProverServer(s, l1ProverPrivKey, s.ProverEndpoints[0])

	balance, err := rpcCli.TaikoToken.BalanceOf(nil, crypto.PubkeyToAddress(l1ProverPrivKey.PublicKey))
	s.Nil(err)

	if balance.Cmp(common.Big0) == 0 {
		// Do not verify zk && sgx proofs in tests.
		securityConcilPrivKey, err := crypto.ToECDSA(common.FromHex(os.Getenv("L1_SECURITY_COUNCIL_PRIVATE_KEY")))
		s.Nil(err)
		s.setAddress(securityConcilPrivKey, rpc.StringToBytes32("tier_sgx_and_pse_zkevm"), common.Address{})
		s.setAddress(securityConcilPrivKey, rpc.StringToBytes32("tier_sgx"), common.Address{})

		ownerPrivKey, err := crypto.ToECDSA(common.FromHex(os.Getenv("L1_CONTRACT_OWNER_PRIVATE_KEY")))
		s.Nil(err)

		// Transfer some tokens to provers.
		balance, err := rpcCli.TaikoToken.BalanceOf(nil, crypto.PubkeyToAddress(ownerPrivKey.PublicKey))
		s.Nil(err)
		s.Greater(balance.Cmp(common.Big0), 0)

		opts, err := bind.NewKeyedTransactorWithChainID(ownerPrivKey, rpcCli.L1ChainID)
		s.Nil(err)
		proverBalance := new(big.Int).Div(balance, common.Big2)
		s.Greater(proverBalance.Cmp(common.Big0), 0)

		tx, err := rpcCli.TaikoToken.Transfer(opts, crypto.PubkeyToAddress(l1ProverPrivKey.PublicKey), proverBalance)
		s.Nil(err)
		_, err = rpc.WaitReceipt(context.Background(), rpcCli.L1, tx)
		s.Nil(err)

		decimal, err := rpcCli.TaikoToken.Decimals(nil)
		s.Nil(err)

		// Increase allowance for AssignmentHook and TaikoL1
		opts, err = bind.NewKeyedTransactorWithChainID(l1ProverPrivKey, rpcCli.L1ChainID)
		s.Nil(err)

		bigInt := new(big.Int).Exp(big.NewInt(1_000_000_000), new(big.Int).SetUint64(uint64(decimal)), nil)
		_, err = rpcCli.TaikoToken.Approve(
			opts,
			common.HexToAddress(os.Getenv("ASSIGNMENT_HOOK_ADDRESS")),
			bigInt,
		)
		s.Nil(err)

		_, err = rpcCli.TaikoToken.Approve(
			opts,
			common.HexToAddress(os.Getenv("TAIKO_L1_ADDRESS")),
			bigInt,
		)
		s.Nil(err)

		_, err = rpc.WaitReceipt(context.Background(), rpcCli.L1, tx)
		s.Nil(err)
	}
	s.Nil(rpcCli.L1RawRPC.CallContext(context.Background(), &s.testnetL1SnapshotID, "evm_snapshot"))
	s.NotEmpty(s.testnetL1SnapshotID)
}

func (s *ClientTestSuite) setAddress(ownerPrivKey *ecdsa.PrivateKey, name [32]byte, address common.Address) {
	var (
		salt = RandomHash()
	)

	controller, err := bindings.NewTaikoTimelockController(
		common.HexToAddress(os.Getenv("TIMELOCK_CONTROLLER")),
		s.RpcClient.L1,
	)
	s.Nil(err)

	opts, err := bind.NewKeyedTransactorWithChainID(ownerPrivKey, s.RpcClient.L1ChainID)
	s.Nil(err)

	addressManagerABI, err := bindings.AddressManagerMetaData.GetAbi()
	s.Nil(err)

	data, err := addressManagerABI.Pack("setAddress", s.RpcClient.L1ChainID.Uint64(), name, address)
	s.Nil(err)

	tx, err := controller.Schedule(
		opts,
		common.HexToAddress(os.Getenv("ROLLUP_ADDRESS_MANAGER_CONTRACT_ADDRESS")),
		common.Big0,
		data,
		[32]byte{},
		salt,
		common.Big0,
	)
	s.Nil(err)

	_, err = rpc.WaitReceipt(context.Background(), s.RpcClient.L1, tx)
	s.Nil(err)

	tx, err = controller.Execute(
		opts,
		common.HexToAddress(os.Getenv("ROLLUP_ADDRESS_MANAGER_CONTRACT_ADDRESS")),
		common.Big0,
		data,
		[32]byte{},
		salt,
	)
	s.Nil(err)

	_, err = rpc.WaitReceipt(context.Background(), s.RpcClient.L1, tx)
	s.Nil(err)
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

func (s *ClientTestSuite) IncreaseTime(time uint64) {
	var result uint64
	s.Nil(s.RpcClient.L1RawRPC.CallContext(context.Background(), &result, "evm_increaseTime", time))
	s.NotNil(result)
}
