package proposer

import (
	"context"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/ethereum-optimism/optimism/op-service/txmgr"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/stretchr/testify/suite"

	"github.com/taikoxyz/taiko-client/bindings"
	txlistfetcher "github.com/taikoxyz/taiko-client/driver/txlist_fetcher"
	"github.com/taikoxyz/taiko-client/internal/testutils"
	"github.com/taikoxyz/taiko-client/internal/utils"
	"github.com/taikoxyz/taiko-client/pkg/jwt"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
)

type ProposerTestSuite struct {
	testutils.ClientTestSuite
	p      *Proposer
	cancel context.CancelFunc
}

func (s *ProposerTestSuite) SetupTest() {
	s.ClientTestSuite.SetupTest()

	l1ProposerPrivKey, err := crypto.ToECDSA(common.FromHex(os.Getenv("L1_PROPOSER_PRIVATE_KEY")))
	s.Nil(err)

	p := new(Proposer)

	ctx, cancel := context.WithCancel(context.Background())

	jwtSecret, err := jwt.ParseSecretFromFile(os.Getenv("JWT_SECRET"))
	s.Nil(err)
	s.NotEmpty(jwtSecret)

	s.Nil(p.InitFromConfig(ctx, &Config{
		ClientConfig: &rpc.ClientConfig{
			L1Endpoint:        os.Getenv("L1_NODE_WS_ENDPOINT"),
			L2Endpoint:        os.Getenv("L2_EXECUTION_ENGINE_HTTP_ENDPOINT"),
			L2EngineEndpoint:  os.Getenv("L2_EXECUTION_ENGINE_AUTH_ENDPOINT"),
			JwtSecret:         string(jwtSecret),
			TaikoL1Address:    common.HexToAddress(os.Getenv("TAIKO_L1_ADDRESS")),
			TaikoL2Address:    common.HexToAddress(os.Getenv("TAIKO_L2_ADDRESS")),
			TaikoTokenAddress: common.HexToAddress(os.Getenv("TAIKO_TOKEN_ADDRESS")),
		},
		AssignmentHookAddress:      common.HexToAddress(os.Getenv("ASSIGNMENT_HOOK_ADDRESS")),
		L1ProposerPrivKey:          l1ProposerPrivKey,
		L2SuggestedFeeRecipient:    common.HexToAddress(os.Getenv("L2_SUGGESTED_FEE_RECIPIENT")),
		ProposeInterval:            1024 * time.Hour,
		MaxProposedTxListsPerEpoch: 1,
		WaitReceiptTimeout:         12 * time.Second,
		ProverEndpoints:            s.ProverEndpoints,
		OptimisticTierFee:          common.Big256,
		SgxTierFee:                 common.Big256,
		TierFeePriceBump:           common.Big2,
		MaxTierFeePriceBumps:       3,
		ExtraData:                  "test",
		L1BlockBuilderTip:          common.Big0,
		TxmgrConfigs: &txmgr.CLIConfig{
			L1RPCURL:                  os.Getenv("L1_NODE_WS_ENDPOINT"),
			NumConfirmations:          0,
			SafeAbortNonceTooLowCount: txmgr.DefaultBatcherFlagValues.SafeAbortNonceTooLowCount,
			PrivateKey:                common.Bytes2Hex(crypto.FromECDSA(l1ProposerPrivKey)),
			FeeLimitMultiplier:        txmgr.DefaultBatcherFlagValues.FeeLimitMultiplier,
			FeeLimitThresholdGwei:     txmgr.DefaultBatcherFlagValues.FeeLimitThresholdGwei,
			MinBaseFeeGwei:            txmgr.DefaultBatcherFlagValues.MinBaseFeeGwei,
			MinTipCapGwei:             txmgr.DefaultBatcherFlagValues.MinTipCapGwei,
			ResubmissionTimeout:       txmgr.DefaultBatcherFlagValues.ResubmissionTimeout,
			ReceiptQueryInterval:      1 * time.Second,
			NetworkTimeout:            txmgr.DefaultBatcherFlagValues.NetworkTimeout,
			TxSendTimeout:             txmgr.DefaultBatcherFlagValues.TxSendTimeout,
			TxNotInMempoolTimeout:     txmgr.DefaultBatcherFlagValues.TxNotInMempoolTimeout,
		},
	}))

	s.p = p
	s.cancel = cancel
}

func (s *ProposerTestSuite) TestName() {
	s.Equal("proposer", s.p.Name())
}

func parseTxs(client *rpc.Client, event *bindings.TaikoL1ClientBlockProposed) (types.Transactions, error) {
	tx, err := client.L1.TransactionInBlock(context.Background(), event.Raw.BlockHash, event.Raw.TxIndex)
	if err != nil {
		return nil, err
	}

	// Decode transactions list.
	var txListDecoder txlistfetcher.TxListFetcher
	if event.Meta.BlobUsed {
		txListDecoder = txlistfetcher.NewBlobTxListFetcher(client)
	} else {
		txListDecoder = new(txlistfetcher.CalldataFetcher)
	}
	txListBytes, err := txListDecoder.Fetch(context.Background(), tx, &event.Meta)
	if err != nil {
		return nil, err
	}

	txListBytes, err = utils.Decompress(txListBytes)
	if err != nil {
		return nil, err
	}

	var txs types.Transactions
	return txs, rlp.DecodeBytes(txListBytes, &txs)
}

func (s *ProposerTestSuite) getLatestProposedTxs(
	start uint64,
	n int,
	timeout time.Duration,
) ([]types.Transactions, error) {
	sink := make(chan *bindings.TaikoL1ClientBlockProposed)

	sub, err := s.p.rpc.TaikoL1.WatchBlockProposed(&bind.WatchOpts{
		Start: &start,
	}, sink, nil, nil)
	if err != nil {
		return nil, err
	}
	defer sub.Unsubscribe()

	var txLst = make([]types.Transactions, 0, n)
	tick := time.After(timeout)
	for len(txLst) < cap(txLst) {
		select {
		case event := <-sink:
			txs, err := parseTxs(s.RPCClient, event)
			if err != nil {
				return nil, err
			}
			txLst = append(txLst, txs)
		case <-tick:
			break
		}
	}

	return txLst, nil
}

func (s *ProposerTestSuite) TestEmptyBlock() {
	s.p.ForceEmptyBlocksInterval = 3 * time.Second
	s.p.BlockMinGasLimit = math.MaxUint64
	s.p.BlockMinTxListBytes = math.MaxUint64
	s.p.ProposeInterval = time.Second

	number, err := s.RPCClient.L1.BlockNumber(context.Background())
	s.Nil(err)

	// start proposer
	s.Nil(s.p.Start())
	defer s.p.Close()

	txs, err := s.getLatestProposedTxs(number, 1, time.Second*30)
	s.Nil(err)
	s.Equal(true, len(txs) == 1 && len(txs[0]) == 0)
}

func (s *ProposerTestSuite) TestProposeOpNoEmptyBlock() {
	p := s.p

	head, err := s.RPCClient.L1.HeaderByNumber(context.Background(), nil)
	s.Nil(err)

	auth, err := bind.NewKeyedTransactorWithChainID(p.L1ProposerPrivKey, s.RPCClient.L2.ChainID)
	s.Nil(err)

	nonce, err := s.RPCClient.L2.PendingNonceAt(context.Background(), auth.From)
	s.Nil(err)

	gasTipCap, err := s.RPCClient.L2.SuggestGasTipCap(context.Background())
	s.Nil(err)
	gasFeeCap := new(big.Int).Add(
		gasTipCap,
		new(big.Int).Mul(head.BaseFee, big.NewInt(2)),
	)

	batchSize := 100
	for i := 0; i < batchSize; i++ {
		to := common.BytesToAddress(testutils.RandomBytes(32))
		baseTx := types.NewTx(&types.DynamicFeeTx{
			ChainID:   s.RPCClient.L2.ChainID,
			To:        &to,
			Nonce:     nonce + uint64(i),
			GasTipCap: gasTipCap,
			GasFeeCap: gasFeeCap,
			Gas:       2100000,
		})
		tx, err := auth.Signer(auth.From, baseTx)
		s.Nil(err)
		s.Nil(s.RPCClient.L2.SendTransaction(context.Background(), tx))
	}

	preBuiltTxList, err := s.RPCClient.GetPoolContent(
		context.Background(),
		p.proposerAddress,
		p.protocolConfigs.BlockMaxGasLimit,
		rpc.BlockMaxTxListBytes,
		p.LocalAddresses,
		p.MaxProposedTxListsPerEpoch,
	)
	s.Nil(err)
	s.Equal(true, len(preBuiltTxList) > 0)

	var (
		blockMinGasLimit    uint64 = math.MaxUint64
		blockMinTxListBytes uint64 = math.MaxUint64
		txLists                    = make([]types.Transactions, 0, len(preBuiltTxList))
	)
	for _, txs := range preBuiltTxList {
		if txs.EstimatedGasUsed <= blockMinGasLimit {
			blockMinGasLimit = txs.EstimatedGasUsed
		} else {
			break
		}
		if txs.BytesLength <= blockMinTxListBytes {
			blockMinTxListBytes = txs.BytesLength
		} else {
			break
		}
		txLists = append(txLists, txs.TxList)
	}

	// Start proposer
	p.LocalAddressesOnly = false
	p.BlockMinGasLimit = blockMinGasLimit
	p.BlockMinTxListBytes = blockMinTxListBytes
	p.ProposeInterval = time.Second
	s.Nil(p.Start())
	defer p.Close()

	txs, err := s.getLatestProposedTxs(head.Number.Uint64(), len(txLists), time.Minute)
	s.Nil(err)

	for i := 0; i < len(txLists); i++ {
		s.Equal(txLists[i].Len(), txs[i].Len())
	}
}

func (s *ProposerTestSuite) TestProposeOp() {
	// Propose txs in L2 execution engine's mempool
	sink := make(chan *bindings.TaikoL1ClientBlockProposed)

	sub, err := s.p.rpc.TaikoL1.WatchBlockProposed(nil, sink, nil, nil)
	s.Nil(err)
	defer func() {
		sub.Unsubscribe()
		close(sink)
	}()

	nonce, err := s.p.rpc.L2.PendingNonceAt(context.Background(), s.TestAddr)
	s.Nil(err)

	parent, err := s.p.rpc.L2.BlockByNumber(context.Background(), nil)
	s.Nil(err)

	baseFeeInfo, err := s.p.rpc.TaikoL2.GetBasefee(nil, 1, uint32(parent.GasUsed()))
	s.Nil(err)

	to := common.BytesToAddress(testutils.RandomBytes(32))
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   s.RPCClient.L2.ChainID,
		Nonce:     nonce,
		GasTipCap: common.Big0,
		GasFeeCap: new(big.Int).SetUint64(baseFeeInfo.Basefee.Uint64() * 2),
		Gas:       21000,
		To:        &to,
		Value:     common.Big1,
	})

	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(s.p.rpc.L2.ChainID), s.TestAddrPrivKey)
	s.Nil(err)
	s.Nil(s.p.rpc.L2.SendTransaction(context.Background(), signedTx))

	s.Nil(s.p.ProposeOp(context.Background()))

	event := <-sink

	s.Equal(event.Meta.Coinbase, s.p.L2SuggestedFeeRecipient)

	_, isPending, err := s.p.rpc.L1.TransactionByHash(context.Background(), event.Raw.TxHash)
	s.Nil(err)
	s.False(isPending)

	receipt, err := s.p.rpc.L1.TransactionReceipt(context.Background(), event.Raw.TxHash)
	s.Nil(err)
	s.Equal(types.ReceiptStatusSuccessful, receipt.Status)
}

func (s *ProposerTestSuite) TestProposeOpLocalsOnly() {
	s.p.LocalAddresses = []common.Address{common.BytesToAddress(testutils.RandomBytes(20))}
	s.p.LocalAddressesOnly = true

	// Propose txs in L2 execution engine's mempool
	sink := make(chan *bindings.TaikoL1ClientBlockProposed)

	sub, err := s.p.rpc.TaikoL1.WatchBlockProposed(nil, sink, nil, nil)
	s.Nil(err)
	defer func() {
		sub.Unsubscribe()
		close(sink)
	}()

	s.Error(errNoNewTxs, s.p.ProposeOp(context.Background()))
}

func (s *ProposerTestSuite) TestProposeEmptyBlockOp() {
	s.Nil(s.p.ProposeEmptyBlockOp(context.Background()))
}

func (s *ProposerTestSuite) TestCustomProposeOpHook() {
	flag := false

	s.p.CustomProposeOpHook = func() error {
		flag = true
		return nil
	}

	s.Nil(s.p.ProposeOp(context.Background()))
	s.True(flag)
}

func (s *ProposerTestSuite) TestAssignProverSuccessFirstRound() {
	s.SetL1Automine(false)
	defer s.SetL1Automine(true)

	_, _, fee, err := s.p.proverSelector.AssignProver(context.Background(), s.p.tierFees, testutils.RandomHash())

	s.Nil(err)
	s.Equal(fee.Uint64(), s.p.OptimisticTierFee.Uint64())
}

func (s *ProposerTestSuite) TestUpdateProposingTicker() {
	s.p.ProposeInterval = 1 * time.Hour
	s.NotPanics(s.p.updateProposingTicker)

	s.p.ProposeInterval = 0
	s.NotPanics(s.p.updateProposingTicker)
}

func (s *ProposerTestSuite) TestStartClose() {
	s.Nil(s.p.Start())
	s.cancel()
	s.NotPanics(func() { s.p.Close() })
}

func TestProposerTestSuite(t *testing.T) {
	suite.Run(t, new(ProposerTestSuite))
}
