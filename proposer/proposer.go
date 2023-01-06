package proposer

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"math/rand"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
	"github.com/taikoxyz/taiko-client/metrics"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	"github.com/urfave/cli/v2"
)

// Proposer keep proposing new transactions from L2 execution engine's tx pool at a fixed interval.
type Proposer struct {
	// RPC clients
	rpc *rpc.Client

	// Private keys and account addresses
	l1ProposerPrivKey       *ecdsa.PrivateKey
	l2SuggestedFeeRecipient common.Address

	// Proposing configuration
	proposingInterval *time.Duration
	proposingTimer    *time.Timer
	commitSlot        uint64

	poolContentSplitter *poolContentSplitter

	// Constants in LibConstants
	protocolConstants *bindings.ProtocolConstants

	// Only for testing purposes
	CustomProposeOpHook func() error
	AfterCommitHook     func() error

	ctx context.Context
	wg  sync.WaitGroup
}

// New initializes the given proposer instance based on the command line flags.
func (p *Proposer) InitFromCli(ctx context.Context, c *cli.Context) error {
	cfg, err := NewConfigFromCliContext(c)
	if err != nil {
		return err
	}

	return InitFromConfig(ctx, p, cfg)
}

// InitFromConfig initializes the proposer instance based on the given configurations.
func InitFromConfig(ctx context.Context, p *Proposer, cfg *Config) (err error) {
	p.l1ProposerPrivKey = cfg.L1ProposerPrivKey
	p.l2SuggestedFeeRecipient = cfg.L2SuggestedFeeRecipient
	p.proposingInterval = cfg.ProposeInterval
	p.wg = sync.WaitGroup{}
	p.ctx = ctx

	// RPC clients
	if p.rpc, err = rpc.NewClient(p.ctx, &rpc.ClientConfig{
		L1Endpoint:     cfg.L1Endpoint,
		L2Endpoint:     cfg.L2Endpoint,
		TaikoL1Address: cfg.TaikoL1Address,
		TaikoL2Address: cfg.TaikoL2Address,
	}); err != nil {
		return fmt.Errorf("initialize rpc clients error: %w", err)
	}

	proposerAddress := crypto.PubkeyToAddress(cfg.L1ProposerPrivKey.PublicKey)
	isWhitelisted, err := p.rpc.IsProposerWhitelisted(proposerAddress)
	if err != nil {
		return fmt.Errorf("failed to check whether current proposer %s is whitelisted: %w", proposerAddress, err)
	}

	if !isWhitelisted {
		return fmt.Errorf("proposer %s is not whitelisted", proposerAddress)
	}

	// Protocol constants
	if p.protocolConstants, err = p.rpc.GetProtocolConstants(nil); err != nil {
		return fmt.Errorf("failed to get protocol constants: %w", err)
	}

	log.Info("Protocol constants", "constants", p.protocolConstants)

	p.poolContentSplitter = &poolContentSplitter{
		shufflePoolContent: cfg.ShufflePoolContent,
		blockMaxTxs:        p.protocolConstants.BlockMaxTxs.Uint64(),
		blockMaxGasLimit:   p.protocolConstants.BlockMaxGasLimit.Uint64(),
		txListMaxBytes:     p.protocolConstants.TxListMaxBytes.Uint64(),
		txMinGasLimit:      p.protocolConstants.TxMinGasLimit.Uint64(),
	}
	p.commitSlot = cfg.CommitSlot

	return nil
}

// Start starts the proposer's main loop.
func (p *Proposer) Start() error {
	p.wg.Add(1)
	go p.eventLoop()
	return nil
}

// eventLoop starts the main loop of Taiko proposer.
func (p *Proposer) eventLoop() {
	defer func() {
		p.proposingTimer.Stop()
		p.wg.Done()
	}()

	blockVerifiedCh := make(chan *bindings.TaikoL1ClientBlockVerified, 1000)
	sub, err := p.rpc.TaikoL1.WatchBlockVerified(nil, blockVerifiedCh, nil)
	if err != nil {
		log.Crit("Create TaikoL1.BlockVerified subscription error", "error", err)
	}

	defer sub.Unsubscribe()

	syncNotify := make(chan struct{}, 1)
	reqSync := func() {
		select {
		case syncNotify <- struct{}{}:
		default:
		}
	}

	for {
		p.updateProposingTicker()

		select {
		case <-p.ctx.Done():
			return
		case <-blockVerifiedCh:
			log.Info("verified!!")
			reqSync()
		case <-syncNotify:
			if err := p.ProposeOp(p.ctx); err != nil {
				log.Error("ProposeOp error", "err", err)
				time.Sleep(3 * time.Second)
				continue
			}
		}
	}
}

// Close closes the proposer instance.
func (p *Proposer) Close() {
	p.wg.Wait()
}

type commitTxListRes struct {
	meta        *bindings.LibDataBlockMetadata
	commitTx    *types.Transaction
	txListBytes []byte
	txNum       uint
}

// ProposeOp performs a proposing operation, fetching transactions
// from L2 execution engine's tx pool, splitting them by proposing constraints,
// and then proposing them to TaikoL1 contract.
func (p *Proposer) ProposeOp(ctx context.Context) error {
	if p.CustomProposeOpHook != nil {
		return p.CustomProposeOpHook()
	}
	syncProgress, err := p.rpc.L2.SyncProgress(ctx)
	if err != nil {
		return fmt.Errorf("failed to get L2 execution engine sync progress: %w", err)
	}
	if syncProgress != nil {
		log.Info("L2 execution engine is syncing", "progress", syncProgress)
		return nil
	}

	log.Info("Start fetching L2 execution engine's transaction pool content")

	pendingContent, _, err := p.rpc.L2PoolContent(ctx)
	if err != nil {
		return fmt.Errorf("failed to fetch transaction pool content: %w", err)
	}

	log.Info("Fetching L2 pending transactions finished", "length", pendingContent.ToTxLists().Len())

	var commitTxListResQueue []*commitTxListRes
	for i, txs := range p.poolContentSplitter.split(pendingContent) {
		txListBytes, err := rlp.EncodeToBytes(txs)
		if err != nil {
			return fmt.Errorf("failed to encode transactions: %w", err)
		}

		meta, commitTx, err := p.CommitTxList(ctx, txListBytes, sumTxsGasLimit(txs), i)
		if err != nil {
			return fmt.Errorf("failed to commit transactions: %w", err)
		}

		commitTxListResQueue = append(commitTxListResQueue, &commitTxListRes{
			meta:        meta,
			commitTx:    commitTx,
			txListBytes: txListBytes,
			txNum:       uint(len(txs)),
		})
	}

	if p.AfterCommitHook != nil {
		if err := p.AfterCommitHook(); err != nil {
			log.Error("Run AfterCommitHook error", "error", err)
		}
	}

	if len(commitTxListResQueue) == 0 {
		log.Info("No faucet txs")
		time.Sleep(60 * time.Second)
	}

	var i = 0
	for _, res := range commitTxListResQueue {
		if i > 2 {
			break
		}
		if err := p.ProposeTxList(ctx, res.meta, res.commitTx, res.txListBytes, res.txNum); err != nil {
			return fmt.Errorf("failed to propose transactions: %w", err)
		}
		i += 1
	}

	return nil
}

func (p *Proposer) CommitTxList(ctx context.Context, txListBytes []byte, gasLimit uint64, splittedIdx int) (
	*bindings.LibDataBlockMetadata,
	*types.Transaction,
	error,
) {
	// Assemble the block context and commit the txList
	meta := &bindings.LibDataBlockMetadata{
		Id:          common.Big0,
		L1Height:    common.Big0,
		L1Hash:      common.Hash{},
		Beneficiary: p.l2SuggestedFeeRecipient,
		GasLimit:    gasLimit,
		TxListHash:  crypto.Keccak256Hash(txListBytes),
		CommitSlot:  common.Big0.Uint64(),
	}
	// log.Info("CommitSlot1", "slot", meta.CommitSlot)

	if p.protocolConstants.CommitDelayConfirmations.Cmp(common.Big0) == 0 {
		log.Debug("No commit delay confirmation, skip committing transactions list")
		return meta, nil, nil
	}

	opts, err := getTxOpts(ctx, p.rpc.L1, p.l1ProposerPrivKey, p.rpc.L1ChainID, 1.1)
	if err != nil {
		return nil, nil, err
	}

	commitHash := common.BytesToHash(encoding.EncodeCommitHash(meta.Beneficiary, meta.TxListHash))
	commitTx, err := p.rpc.TaikoL1.CommitBlock(opts, meta.CommitSlot, commitHash)
	if err != nil {
		return nil, nil, err
	}

	return meta, commitTx, nil
}

func (p *Proposer) ProposeTxList(
	ctx context.Context,
	meta *bindings.LibDataBlockMetadata,
	commitTx *types.Transaction,
	txListBytes []byte,
	txNum uint,
) error {
	if p.protocolConstants.CommitDelayConfirmations.Cmp(common.Big0) > 0 {
		receipt, err := rpc.WaitReceipt(ctx, p.rpc.L1, commitTx)
		if err != nil {
			return err
		}

		if receipt.Status != types.ReceiptStatusSuccessful {
			log.Error("Failed to commit transactions list", "txHash", receipt.TxHash)
			return nil
		}

		log.Info(
			"Commit block finished, wait some L1 blocks confirmations before proposing",
			"commitHeight", receipt.BlockNumber,
			"confirmations", p.protocolConstants.CommitDelayConfirmations,
		)

		meta.CommitHeight = receipt.BlockNumber.Uint64()

		if err := rpc.WaitConfirmations(
			ctx, p.rpc.L1, p.protocolConstants.CommitDelayConfirmations.Uint64(), receipt.BlockNumber.Uint64(),
		); err != nil {
			return fmt.Errorf("wait L1 blocks confirmations error, commitHash %s: %w", receipt.BlockNumber, err)
		}
	}

	// Propose the transactions list
	inputs, err := encoding.EncodeProposeBlockInput(meta, txListBytes)
	if err != nil {
		return err
	}

	// log.Info("CommitSlot2", "slot", meta.CommitSlot)

	// key1, err := crypto.HexToECDSA("eed858a6f8b22fea58762091e0237ed59b3555f83e72d5e398a3f436464a4306")
	// if err != nil {
	// 	return err
	// }

	// opts1, err := getTxOpts(ctx, p.rpc.L1, key1, p.rpc.L1ChainID, 10)
	// if err != nil {
	// 	return err
	// }

	// tx1, err := p.rpc.TaikoL1.VerifyBlocks(opts1, new(big.Int).SetUint64(100))
	// if err != nil {
	// 	log.Error("tx error", "err", err)
	// } else {
	// 	log.Info("tx1", "hash", tx1.Hash())
	// }

	opts, err := getTxOpts(ctx, p.rpc.L1, p.l1ProposerPrivKey, p.rpc.L1ChainID, 1.5)
	if err != nil {
		return err
	}

	proposeTx, err := p.rpc.TaikoL1.ProposeBlock(opts, inputs)
	if err != nil {
		return err
	}

	log.Info("tx2", "tx", proposeTx.Hash())
	fmt.Println(proposeTx.Hash())

	if _, err := rpc.WaitReceipt(ctx, p.rpc.L1, proposeTx); err != nil {
		return err
	}

	log.Info("📝 Propose transactions succeeded")
	time.Sleep(1 * time.Second)

	metrics.ProposerProposedTxListsCounter.Inc(1)
	metrics.ProposerProposedTxsCounter.Inc(int64(txNum))

	return nil
}

// updateProposingTicker updates the internal proposing timer.
func (p *Proposer) updateProposingTicker() {
	if p.proposingTimer != nil {
		p.proposingTimer.Stop()
	}

	var duration time.Duration
	if p.proposingInterval != nil {
		duration = *p.proposingInterval
	} else {
		// Random number between 12 - 60
		randomSeconds := rand.Intn((60 - 11)) + 12
		duration = time.Duration(randomSeconds) * time.Second
	}

	p.proposingTimer = time.NewTimer(duration)
}

// Name returns the application name.
func (p *Proposer) Name() string {
	return "proposer"
}

// sumTxsGasLimit calculates the accumulated gas limit of all transactions in the list.
func sumTxsGasLimit(txs []*types.Transaction) uint64 {
	var total uint64
	for i := range txs {
		total += txs[i].Gas()
	}
	return total
}

// getTxOpts creates a bind.TransactOpts instance using the given private key.
func getTxOpts(
	ctx context.Context,
	cli *ethclient.Client,
	privKey *ecdsa.PrivateKey,
	chainID *big.Int,
	gasTip float32,
) (*bind.TransactOpts, error) {
	opts, err := bind.NewKeyedTransactorWithChainID(privKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to generate prepareBlock transaction options: %w", err)
	}

	gasTipCap, err := cli.SuggestGasTipCap(ctx)
	if err != nil {
		if rpc.IsMaxPriorityFeePerGasNotFoundError(err) {
			gasTipCap = rpc.FallbackGasTipCap
		} else {
			return nil, err
		}
	}

	l1Head, err := cli.BlockNumber(ctx)
	if err != nil {
		return nil, err
	}

	nonce, err := cli.NonceAt(ctx, crypto.PubkeyToAddress(privKey.PublicKey), new(big.Int).SetUint64(l1Head))
	// nonce, err := cli.PendingNonceAt(ctx, crypto.PubkeyToAddress(privKey.PublicKey))
	if err != nil {
		return nil, err
	}

	// gasPrice, err := cli.SuggestGasPrice(ctx)
	// if err != nil {
	// 	return nil, err
	// }

	// log.Info("gasPrice", "gasPrice", gasPrice)
	opts.GasTipCap = new(big.Int).SetUint64(gasTipCap.Uint64() * uint64(gasTip))
	// opts.Nonce = new(big.Int).SetUint64(nonce)
	log.Info("Nonce", "nonce", nonce)
	// opts.GasPrice = big.NewInt(1499999992 * 10)

	return opts, nil
}
