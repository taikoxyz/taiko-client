package testutils

import (
	"crypto/ecdsa"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
)

var (
	proposerPrivKey = "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	proverPrivKey   = "59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"
	TestPrivKey     *ecdsa.PrivateKey
	TestAddr        common.Address
	ProverPrivKey   *ecdsa.PrivateKey
	ProverAddr      common.Address
)

func init() {
	initLog()
	initTestAccount()
	initProverAccount()
}

func initLog() {
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
}

func initTestAccount() {
	var err error
	TestPrivKey, err = crypto.ToECDSA(common.Hex2Bytes(proposerPrivKey))
	if err != nil {
		panic(err)
	}
	TestAddr = crypto.PubkeyToAddress(TestPrivKey.PublicKey)
}

func initProverAccount() {
	var err error
	ProverPrivKey, err = crypto.ToECDSA(common.Hex2Bytes(proposerPrivKey))
	if err != nil {
		panic(err)
	}
	ProverAddr = crypto.PubkeyToAddress(ProverPrivKey.PublicKey)
}
