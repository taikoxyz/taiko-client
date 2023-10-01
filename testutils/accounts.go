package testutils

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/taikoxyz/taiko-client/bindings"
)

const (
	ProposerPrivateKey = "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	ProverPrivateKey   = "59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"
	ownerAddress       = "0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC"
)

// variables need to be initialized
var (
	ProposerPrivKey         *ecdsa.PrivateKey
	ProposerAddress         common.Address
	ProverPrivKey           *ecdsa.PrivateKey
	ProverAddr              common.Address
	L2SuggestedFeeRecipient common.Address
)

func initTestAccount() (err error) {
	ProposerPrivKey, err = crypto.ToECDSA(common.Hex2Bytes(ProposerPrivateKey))
	if err != nil {
		panic(err)
	}
	ProposerAddress = crypto.PubkeyToAddress(ProposerPrivKey.PublicKey)
	log.Info("Test Account:", "address", ProposerAddress.Hex())
	L2SuggestedFeeRecipient = ProposerAddress
	return nil
}

func initProverAccount() (err error) {
	ProverPrivKey, err = crypto.ToECDSA(common.Hex2Bytes(ProverPrivateKey))
	if err != nil {
		return err
	}
	ProverAddr = crypto.PubkeyToAddress(ProverPrivKey.PublicKey)
	log.Info("Prover Account:", "address", ProverAddr.Hex())
	return nil
}

func ensureProverBalance(c *gethContainer) error {
	cli, err := ethclient.Dial(c.HttpEndpoint())
	if err != nil {
		return err
	}
	taikoL1, err := bindings.NewTaikoL1Client(c.TaikoL1Address, cli)
	if err != nil {
		return err
	}
	tokenBalance, err := taikoL1.GetTaikoTokenBalance(nil, ProverAddr)
	if err != nil {
		return err
	}
	if tokenBalance.Cmp(common.Big0) > 0 {
		return nil
	}
	chainID, err := cli.ChainID(context.Background())
	if err != nil {
		return err
	}
	opts, err := bind.NewKeyedTransactorWithChainID(ProposerPrivKey, chainID)
	if err != nil {
		return err
	}
	premintAmount, _ := new(big.Int).SetString(premintTokenAmount, 10)
	taikoToken, err := bindings.NewTaikoToken(c.TaikoL1TokenAddress, cli)
	if err != nil {
		return err
	}
	if _, err = taikoToken.Approve(opts, c.TaikoL1Address, premintAmount); err != nil {
		return err
	}

	tx, err := taikoL1.DepositTaikoToken(opts, premintAmount)
	if err != nil {
		return err
	}
	log.Debug("DepositTaikoToken for prover ", "tx", tx.Hash().Hex())
	return nil
}
