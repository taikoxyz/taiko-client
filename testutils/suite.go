package testutils

import (
	"context"
	"crypto/ecdsa"
	"math"
	"math/big"
	"os"

	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/suite"
	"github.com/taikoxyz/taiko-client/pkg/jwt"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
)

type ClientTestSuite struct {
	suite.Suite
	testnetL1SnapshotID string
	RpcClient           *rpc.Client
	TestAddrPrivKey     *ecdsa.PrivateKey
	TestAddr            common.Address
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
		L1Endpoint:       os.Getenv("L1_NODE_WS_ENDPOINT"),
		L2Endpoint:       os.Getenv("L2_EXECUTION_ENGINE_WS_ENDPOINT"),
		TaikoL1Address:   common.HexToAddress(os.Getenv("TAIKO_L1_ADDRESS")),
		TaikoL2Address:   common.HexToAddress(os.Getenv("TAIKO_L2_ADDRESS")),
		L2EngineEndpoint: os.Getenv("L2_EXECUTION_ENGINE_AUTH_ENDPOINT"),
		JwtSecret:        string(jwtSecret),
		RetryInterval:    backoff.DefaultMaxInterval,
	})
	s.Nil(err)

	s.RpcClient = rpcCli

	l1ProposerPrivKey, err := crypto.ToECDSA(common.Hex2Bytes(os.Getenv("L1_PROPOSER_PRIVATE_KEY")))
	s.Nil(err)

	opts, err := bind.NewKeyedTransactorWithChainID(l1ProposerPrivKey, rpcCli.L1ChainID)
	s.Nil(err)

	balance, err := rpcCli.TaikoL1.GetTaikoTokenBalance(nil, crypto.PubkeyToAddress(l1ProposerPrivKey.PublicKey))
	s.Nil(err)

	if balance.Cmp(common.Big0) == 0 {
		tx, err := rpcCli.TaikoL1.DepositTaikoToken(opts, new(big.Int).SetUint64(uint64(math.Pow(2, 32))))
		s.Nil(err)

		receipt, err := rpc.WaitReceipt(context.Background(), rpcCli.L1, tx)
		s.Nil(err)
		s.Equal(types.ReceiptStatusSuccessful, receipt.Status)
	}

	s.Nil(rpcCli.L1RawRPC.CallContext(context.Background(), &s.testnetL1SnapshotID, "evm_snapshot"))
	s.NotEmpty(s.testnetL1SnapshotID)
}

func (s *ClientTestSuite) TearDownTest() {
	var revertRes bool
	s.Nil(s.RpcClient.L1RawRPC.CallContext(context.Background(), &revertRes, "evm_revert", s.testnetL1SnapshotID))
	s.True(revertRes)

	s.Nil(rpc.SetHead(context.Background(), s.RpcClient.L2RawRPC, common.Big0))
}

func (s *ClientTestSuite) SetL1Automine(automine bool) {
	s.Nil(s.RpcClient.L1RawRPC.CallContext(context.Background(), nil, "evm_setAutomine", automine))
}
