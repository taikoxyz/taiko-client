package testutils

import (
	"context"
	"crypto/ecdsa"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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
		common.Hex2Bytes("2bdd21761a483f71054e14f5b827213567971c676928d9a1808cbfa4b7501200"),
	)
	s.Nil(err)

	s.TestAddrPrivKey = testAddrPrivKey
	s.TestAddr = common.HexToAddress("0xDf08F82De32B8d460adbE8D72043E3a7e25A3B39")

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
	})
	s.Nil(err)

	s.Nil(rpcCli.L1RawRPC.CallContext(context.Background(), &s.testnetL1SnapshotID, "evm_snapshot"))
	s.NotEmpty(s.testnetL1SnapshotID)

	s.RpcClient = rpcCli
}

func (s *ClientTestSuite) TearDownTest() {
	var revertRes bool
	s.Nil(s.RpcClient.L1RawRPC.CallContext(context.Background(), &revertRes, "evm_revert", s.testnetL1SnapshotID))
	s.True(revertRes)

	s.Nil(rpc.SetHead(context.Background(), s.RpcClient.L2RawRPC, common.Big0))
}

func (s *ClientTestSuite) MineL1Confirmations() error {
	return s.RpcClient.L1RawRPC.CallContext(context.Background(), nil, "hardhat_mine", hexutil.EncodeUint64(4))
}
