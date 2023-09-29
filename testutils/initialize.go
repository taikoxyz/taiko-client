package testutils

import (
	"context"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/taikoxyz/taiko-client/bindings"
	"golang.org/x/sync/errgroup"
)

func init() {
	// Don't change the following initialization order
	var g errgroup.Group
	g.Go(initLog)
	g.Go(initMonoPath)
	g.Go(initJwtSecret)
	g.Go(initTestAccount)
	g.Go(initProverAccount)
	if err := g.Wait(); err != nil {
		panic(err)
	}
	if err := startBaseContainer(context.Background()); err != nil {
		panic(err)
	}
	if err := ensureProverBalance(); err != nil {
		panic(err)
	}
}

func initLog() (err error) {
	level := log.LvlInfo
	if os.Getenv("LOG_LEVEL") != "" {
		level, err = log.LvlFromString(os.Getenv("LOG_LEVEL"))
		if err != nil {
			return fmt.Errorf("invalid log level: %v", os.Getenv("LOG_LEVEL"))
		}
	}
	log.Root().SetHandler(
		log.LvlFilterHandler(level, log.StreamHandler(os.Stdout, log.TerminalFormat(true))),
	)
	return nil
}

func initTestAccount() (err error) {
	TestPrivKey, err = crypto.ToECDSA(common.Hex2Bytes(proposerPrivKey))
	if err != nil {
		panic(err)
	}
	TestAddr = crypto.PubkeyToAddress(TestPrivKey.PublicKey)
	log.Info("Test Account:", "address", TestAddr.Hex())
	return nil
}

func initProverAccount() (err error) {
	ProverPrivKey, err = crypto.ToECDSA(common.Hex2Bytes(proverPrivKey))
	if err != nil {
		return err
	}
	ProverAddr = crypto.PubkeyToAddress(ProverPrivKey.PublicKey)
	log.Info("Prover Account:", "address", ProverAddr.Hex())
	return nil
}

func ensureProverBalance() error {
	cli, err := ethclient.Dial(l1BaseContainer.HttpEndpoint())
	if err != nil {
		return err
	}
	taikoL1, err := bindings.NewTaikoL1Client(TaikoL1Address, cli)
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
	opts, err := bind.NewKeyedTransactorWithChainID(TestPrivKey, chainID)
	if err != nil {
		return err
	}
	premintAmount, _ := new(big.Int).SetString(premintTokenAmount, 10)
	taikoToken, err := bindings.NewTaikoToken(TaikoTokenAddress, cli)
	if err != nil {
		return err
	}
	if _, err = taikoToken.Approve(opts, TaikoL1Address, premintAmount); err != nil {
		return err
	}

	tx, err := taikoL1.DepositTaikoToken(opts, premintAmount)
	if err != nil {
		return err
	}
	log.Debug("DepositTaikoToken for prover ", "tx", tx.Hash().Hex())
	return nil
}
