package sender_test

import (
	"context"
	"math/big"
	"os"
	"runtime"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"golang.org/x/sync/errgroup"

	"github.com/taikoxyz/taiko-client/internal/sender"
	"github.com/taikoxyz/taiko-client/internal/testutils"
	"github.com/taikoxyz/taiko-client/internal/utils"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
)

type SenderTestSuite struct {
	testutils.ClientTestSuite
	sender *sender.Sender
}

func (s *SenderTestSuite) TestNormalSender() {
	var eg errgroup.Group
	eg.SetLimit(runtime.NumCPU())
	for i := 0; i < 5; i++ {
		i := i
		eg.Go(func() error {
			addr := common.BigToAddress(big.NewInt(int64(i)))
			_, err := s.sender.SendRawTransaction(s.sender.Opts.Nonce.Uint64(), &addr, big.NewInt(1), nil)
			return err
		})
	}
	s.Nil(eg.Wait())

	for _, confirmCh := range s.sender.TxToConfirmChannels() {
		confirm := <-confirmCh
		s.Nil(confirm.Err)
	}
}

// Test touch max gas price and replacement.
func (s *SenderTestSuite) TestReplacement() {
	send := s.sender
	client := s.RPCClient.L1

	// Let max gas price be 2 times of the gas fee cap.
	send.MaxGasFee = send.Opts.GasFeeCap.Uint64() * 2

	nonce, err := client.NonceAt(context.Background(), send.Opts.From, nil)
	s.Nil(err)

	pendingNonce, err := client.PendingNonceAt(context.Background(), send.Opts.From)
	s.Nil(err)
	// Run test only if mempool has no pending transactions.
	if pendingNonce > nonce {
		return
	}

	nonce++
	baseTx := &types.DynamicFeeTx{
		ChainID:   client.ChainID,
		To:        &common.Address{},
		GasFeeCap: big.NewInt(int64(send.MaxGasFee - 1)),
		GasTipCap: big.NewInt(int64(send.MaxGasFee - 1)),
		Nonce:     nonce,
		Gas:       21000,
		Value:     big.NewInt(1),
		Data:      nil,
	}
	rawTx, err := send.Opts.Signer(send.Opts.From, types.NewTx(baseTx))
	s.Nil(err)
	err = client.SendTransaction(context.Background(), rawTx)
	s.Nil(err)

	// Replace the transaction with a higher nonce.
	_, err = send.SendRawTransaction(nonce, &common.Address{}, big.NewInt(1), nil)
	s.Nil(err)

	time.Sleep(time.Second * 6)
	// Send a transaction with a next nonce and let all the transactions be confirmed.
	_, err = send.SendRawTransaction(nonce-1, &common.Address{}, big.NewInt(1), nil)
	s.Nil(err)

	for _, confirmCh := range send.TxToConfirmChannels() {
		confirm := <-confirmCh
		// Check the replaced transaction's gasFeeTap touch the max gas price.
		if confirm.CurrentTx.Nonce() == nonce {
			s.Equal(send.MaxGasFee, confirm.CurrentTx.GasFeeCap().Uint64())
		}
		s.Nil(confirm.Err)
	}

	_, err = client.TransactionReceipt(context.Background(), rawTx.Hash())
	s.Equal("not found", err.Error())
}

// Test nonce too low.
func (s *SenderTestSuite) TestNonceTooLow() {
	client := s.RPCClient.L1
	send := s.sender

	nonce, err := client.NonceAt(context.Background(), send.Opts.From, nil)
	s.Nil(err)
	pendingNonce, err := client.PendingNonceAt(context.Background(), send.Opts.From)
	s.Nil(err)
	// Run test only if mempool has no pending transactions.
	if pendingNonce > nonce {
		return
	}

	txID, err := send.SendRawTransaction(nonce-3, &common.Address{}, big.NewInt(1), nil)
	s.Nil(err)
	confirm := <-send.TxToConfirmChannel(txID)
	s.Nil(confirm.Err)
	s.Equal(nonce, confirm.CurrentTx.Nonce())
}

func (s *SenderTestSuite) SetupTest() {
	s.ClientTestSuite.SetupTest()

	ctx := context.Background()
	priv, err := crypto.ToECDSA(common.FromHex(os.Getenv("L1_PROPOSER_PRIVATE_KEY")))
	s.Nil(err)

	s.sender, err = sender.NewSender(ctx, &sender.Config{
		MaxGasFee:      20000000000,
		GasGrowthRate:  50,
		MaxRetrys:      0,
		GasLimit:       2000000,
		MaxWaitingTime: time.Second * 10,
	}, s.RPCClient.L1, priv)
	s.Nil(err)
}

func (s *SenderTestSuite) TearDownTest() {
	s.sender.Close()
	s.ClientTestSuite.TearDownTest()
}

func TestSenderTestSuite(t *testing.T) {
	suite.Run(t, new(SenderTestSuite))
}

func TestBlockTx(t *testing.T) {
	//t.SkipNow()
	// Load environment variables.
	utils.LoadEnv()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, err := rpc.NewClient(ctx, &rpc.ClientConfig{
		L1Endpoint:        os.Getenv("L1_NODE_WS_ENDPOINT"),
		L2Endpoint:        os.Getenv("L2_EXECUTION_ENGINE_HTTP_ENDPOINT"),
		TaikoL1Address:    common.HexToAddress(os.Getenv("TAIKO_L1_ADDRESS")),
		TaikoL2Address:    common.HexToAddress(os.Getenv("TAIKO_L2_ADDRESS")),
		TaikoTokenAddress: common.HexToAddress(os.Getenv("TAIKO_TOKEN_ADDRESS")),
		L1BeaconEndpoint:  "http://localhost:3500",
	})
	assert.NoError(t, err)
	l1Client := client.L1

	priv := os.Getenv("L1_PROPOSER_PRIVATE_KEY")
	sk, err := crypto.ToECDSA(common.FromHex(priv))
	assert.NoError(t, err)

	send, err := sender.NewSender(ctx, nil, l1Client, sk)
	assert.NoError(t, err)
	opts := send.Opts

	balance, err := l1Client.BalanceAt(ctx, opts.From, nil)
	assert.NoError(t, err)
	t.Logf("address: %s, balance: %s", opts.From.String(), balance.String())

	data, dErr := os.ReadFile("./sender.go")
	assert.NoError(t, dErr)
	//data := []byte{'s'}
	sideCar, sErr := rpc.MakeSidecar(data)
	assert.NoError(t, sErr)

	nonce, err := l1Client.NonceAt(ctx, opts.From, nil)
	assert.NoError(t, err)

	pendingNonce, err := l1Client.PendingNonceAt(ctx, opts.From)
	assert.NoError(t, err)
	if pendingNonce > nonce {
		return
	}

	tx, err := l1Client.TransactBlobTx(opts, nil, nil, sideCar)
	assert.NoError(t, err)
	txID, err := send.SendTransaction(tx)
	assert.NoError(t, err)

	confirm := <-send.TxToConfirmChannel(txID)
	assert.NoError(t, confirm.Err)

	receipt := confirm.Receipt
	t.Log("blob hash: ", tx.BlobHashes()[0].String())
	t.Log("block number: ", receipt.BlockNumber.Uint64())
	t.Log("tx hash: ", receipt.TxHash.String())

	sidecars, err := client.GetBlobs(ctx, receipt.BlockNumber)
	assert.NoError(t, err)

	t.Log(len(sidecars))
}
