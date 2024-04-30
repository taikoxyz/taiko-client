package proposer

import (
	"bytes"
	"context"
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/ethereum-optimism/optimism/op-challenger/sender"
	"github.com/ethereum-optimism/optimism/op-service/txmgr"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/urfave/cli/v2"

	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
	"github.com/taikoxyz/taiko-client/internal/metrics"
	"github.com/taikoxyz/taiko-client/internal/utils"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	selector "github.com/taikoxyz/taiko-client/proposer/prover_selector"
	builder "github.com/taikoxyz/taiko-client/proposer/transaction_builder"
)

var (
	proverAssignmentTimeout    = 30 * time.Minute
	requestProverServerTimeout = 12 * time.Second
)

// Proposer keep proposing new transactions from L2 execution engine's tx pool at a fixed interval.
type Proposer struct {
	// configurations
	*Config

	// RPC clients
	rpc *rpc.Client

	// Private keys and account addresses
	proposerAddress common.Address

	proposingTimer *time.Timer

	tiers    []*rpc.TierProviderTierWithID
	tierFees []encoding.TierFee

	// Prover selector
	proverSelector selector.ProverSelector

	// Transaction builder
	txBuilder builder.ProposeBlockTransactionBuilder

	// Protocol configurations
	protocolConfigs *bindings.TaikoDataConfig

	lastProposedAt time.Time

	txSender *sender.TxSender

	ctx context.Context
	wg  sync.WaitGroup
}

// InitFromCli New initializes the given proposer instance based on the command line flags.
func (p *Proposer) InitFromCli(ctx context.Context, c *cli.Context) error {
	cfg, err := NewConfigFromCliContext(c)
	if err != nil {
		return err
	}

	return p.InitFromConfig(ctx, cfg)
}

// InitFromConfig initializes the proposer instance based on the given configurations.
func (p *Proposer) InitFromConfig(ctx context.Context, cfg *Config) (err error) {
	p.proposerAddress = crypto.PubkeyToAddress(cfg.L1ProposerPrivKey.PublicKey)
	p.ctx = ctx
	p.Config = cfg
	p.lastProposedAt = time.Now()

	// RPC clients
	if p.rpc, err = rpc.NewClient(p.ctx, cfg.ClientConfig); err != nil {
		return fmt.Errorf("initialize rpc clients error: %w", err)
	}

	// Protocol configs
	protocolConfigs, err := p.rpc.TaikoL1.GetConfig(&bind.CallOpts{Context: ctx})
	if err != nil {
		return fmt.Errorf("failed to get protocol configs: %w", err)
	}
	p.protocolConfigs = &protocolConfigs

	log.Info("Protocol configs", "configs", p.protocolConfigs)

	if p.tiers, err = p.rpc.GetTiers(ctx); err != nil {
		return err
	}
	if err := p.initTierFees(); err != nil {
		return err
	}

	txmgr, err := txmgr.NewSimpleTxManager(
		"proposer",
		log.Root(),
		&metrics.TxMgrMetrics,
		*cfg.TxmgrConfigs,
	)
	if err != nil {
		return err
	}
	p.txSender = sender.NewTxSender(p.ctx, log.Root(), txmgr, p.Config.MaxProposedTxListsPerEpoch)

	if p.proverSelector, err = selector.NewETHFeeEOASelector(
		&protocolConfigs,
		p.rpc,
		p.proposerAddress,
		cfg.TaikoL1Address,
		cfg.AssignmentHookAddress,
		p.tierFees,
		cfg.TierFeePriceBump,
		cfg.ProverEndpoints,
		cfg.MaxTierFeePriceBumps,
		proverAssignmentTimeout,
		requestProverServerTimeout,
	); err != nil {
		return err
	}

	if cfg.BlobAllowed {
		p.txBuilder = builder.NewBlobTransactionBuilder(
			p.rpc,
			p.L1ProposerPrivKey,
			p.proverSelector,
			p.Config.L1BlockBuilderTip,
			cfg.TaikoL1Address,
			cfg.L2SuggestedFeeRecipient,
			cfg.AssignmentHookAddress,
			cfg.ProposeBlockTxGasLimit,
			cfg.ExtraData,
		)
	} else {
		p.txBuilder = builder.NewCalldataTransactionBuilder(
			p.rpc,
			p.L1ProposerPrivKey,
			p.proverSelector,
			p.Config.L1BlockBuilderTip,
			cfg.L2SuggestedFeeRecipient,
			cfg.TaikoL1Address,
			cfg.AssignmentHookAddress,
			cfg.ProposeBlockTxGasLimit,
			cfg.ExtraData,
		)
	}

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

	for {
		p.updateProposingTicker()

		select {
		case <-p.ctx.Done():
			return
		// proposing interval timer has been reached
		case <-p.proposingTimer.C:
			metrics.ProposerProposeEpochCounter.Add(1)

			// Attempt a proposing operation
			if err := p.ProposeOp(p.ctx); err != nil {
				log.Error("Proposing operation error", "error", err)
				continue
			}
		}
	}
}

// Close closes the proposer instance.
func (p *Proposer) Close(_ context.Context) {
	p.wg.Wait()
}

// fetchPoolContent fetches the transaction pool content from L2 execution engine.
func (p *Proposer) fetchPoolContent(filterPoolContent bool) ([]types.Transactions, error) {
	// Fetch the pool content.
	preBuiltTxList, err := p.rpc.GetPoolContent(
		p.ctx,
		p.proposerAddress,
		p.protocolConfigs.BlockMaxGasLimit,
		rpc.BlockMaxTxListBytes,
		p.LocalAddresses,
		p.MaxProposedTxListsPerEpoch,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch transaction pool content: %w", err)
	}

	txLists := []types.Transactions{}
	for i, txs := range preBuiltTxList {
		// Filter the pool content if the filterPoolContent flag is set.
		if txs.EstimatedGasUsed < p.MinGasUsed && txs.BytesLength < p.MinTxListBytes && filterPoolContent {
			log.Info(
				"Pool content skipped",
				"index", i,
				"estimatedGasUsed", txs.EstimatedGasUsed,
				"minGasUsed", p.MinGasUsed,
				"bytesLength", txs.BytesLength,
				"minBytesLength", p.MinTxListBytes,
			)
			break
		}
		txLists = append(txLists, txs.TxList)
	}
	// If the pool content is empty and the checkPoolContent flag is not set, return an empty list.
	if !filterPoolContent && len(txLists) == 0 {
		log.Info(
			"Pool content is empty, proposing an empty block",
			"lastProposedAt", p.lastProposedAt,
			"minProposingInternal", p.MinProposingInternal,
		)
		txLists = append(txLists, types.Transactions{})
	}

	// If LocalAddressesOnly is set, filter the transactions by the local addresses.
	if p.LocalAddressesOnly {
		var (
			localTxsLists []types.Transactions
			signer        = types.LatestSignerForChainID(p.rpc.L2.ChainID)
		)
		for _, txs := range txLists {
			var filtered types.Transactions
			for _, tx := range txs {
				sender, err := types.Sender(signer, tx)
				if err != nil {
					return nil, err
				}

				for _, localAddress := range p.LocalAddresses {
					if sender == localAddress {
						filtered = append(filtered, tx)
					}
				}
			}

			if filtered.Len() != 0 {
				localTxsLists = append(localTxsLists, filtered)
			}
		}
		txLists = localTxsLists
	}

	log.Info("Transactions lists count", "count", len(txLists))

	return txLists, nil
}

// ProposeOp performs a proposing operation, fetching transactions
// from L2 execution engine's tx pool, splitting them by proposing constraints,
// and then proposing them to TaikoL1 contract.
func (p *Proposer) ProposeOp(ctx context.Context) error {
	// Check if it's time to propose unfiltered pool content.
	var (
		filterPoolContent = time.Now().Before(p.lastProposedAt.Add(p.MinProposingInternal))
		txListsBytes      = make([][]byte, 0)
	)

	// Wait until L2 execution engine is synced at first.
	if err := p.rpc.WaitTillL2ExecutionEngineSynced(ctx); err != nil {
		return fmt.Errorf("failed to wait until L2 execution engine synced: %w", err)
	}

	log.Info(
		"Start fetching L2 execution engine's transaction pool content",
		"filterPoolContent", filterPoolContent,
		"lastProposedAt", p.lastProposedAt,
	)

	// Fetch the pool content with the given limits.
	txLists, err := p.fetchPoolContent(filterPoolContent)
	if err != nil {
		return err
	}

	// If the pool content is empty, return.
	if len(txLists) == 0 {
		return nil
	}

	// Propose all L2 transactions lists.
	for i, txs := range txLists {
		if i >= int(p.MaxProposedTxListsPerEpoch) {
			return nil
		}

		txListBytes, err := rlp.EncodeToBytes(txs)
		if err != nil {
			return fmt.Errorf("failed to encode transactions: %w", err)
		}

		txListsBytes = append(txListsBytes, txListBytes)
	}

	for i, err := range p.ProposeTxLists(ctx, txListsBytes) {
		if err != nil {
			log.Error(
				"Failed to send TaikoL1.proposeBlock transaction",
				"index", i,
				"error", err,
			)
			continue
		}

		metrics.ProposerProposedTxListsCounter.Add(1)
		metrics.ProposerProposedTxsCounter.Add(float64(len(txLists[i])))

		log.Info("📝 Propose transactions succeeded", "txs", len(txLists[i]))
		p.lastProposedAt = time.Now()
	}

	return nil
}

// ProposeTxLists proposes the given transaction lists to TaikoL1 contract.
func (p *Proposer) ProposeTxLists(ctx context.Context, txListsBytes [][]byte) []error {
	txCandidates := make([]txmgr.TxCandidate, 0)

	for i, txListBytes := range txListsBytes {
		compressedTxListBytes, err := utils.Compress(txListBytes)
		if err != nil {
			log.Error("Failed to compress transactions list", "index", i, "error", err)
			break
		}

		candidate, err := p.txBuilder.Build(
			ctx,
			p.tierFees,
			p.IncludeParentMetaHash,
			compressedTxListBytes,
		)
		if err != nil {
			log.Error("Failed to build TaikoL1.proposeBlock transaction", "error", err)
			break
		}

		txCandidates = append(txCandidates, *candidate)
	}

	if len(txCandidates) == 0 {
		return []error{}
	}

	// Send the transactions to the TaikoL1 contract, and if any of them fails, try
	// to parse the custom error.
	errors := p.txSender.SendAndWaitDetailed("proposeBlock", txCandidates...)
	for i, err := range errors {
		if err == nil {
			continue
		}

		// If a transaction is reverted on chain, the error string returned by txSender will like this:
		// fmt.Errorf("%w purpose: %v hash: %v", ErrTransactionReverted, txPurpose, rcpt.Receipt.TxHash)
		// Then we try parsing the custom error for more details in log.
		if strings.Contains(err.Error(), "purpose: ") && strings.Contains(err.Error(), "hash: ") {
			txHash := strings.Split(err.Error(), "hash: ")[1]
			receipt, err := p.rpc.L1.TransactionReceipt(ctx, common.HexToHash(txHash))
			if err != nil {
				log.Error("Failed to fetch receipt", "txHash", txHash, "error", err)
				continue
			}
			errors[i] = encoding.TryParsingCustomErrorFromReceipt(ctx, p.rpc.L1, p.proposerAddress, receipt)
		}
	}

	return errors
}

// updateProposingTicker updates the internal proposing timer.
func (p *Proposer) updateProposingTicker() {
	if p.proposingTimer != nil {
		p.proposingTimer.Stop()
	}

	var duration time.Duration
	if p.ProposeInterval != 0 {
		duration = p.ProposeInterval
	} else {
		// Random number between 12 - 120
		randomSeconds := rand.Intn(120-11) + 12 // nolint: gosec
		duration = time.Duration(randomSeconds) * time.Second
	}

	p.proposingTimer = time.NewTimer(duration)
}

// Name returns the application name.
func (p *Proposer) Name() string {
	return "proposer"
}

// initTierFees initializes the proving fees for every proof tier configured in the protocol for the proposer.
func (p *Proposer) initTierFees() error {
	for _, tier := range p.tiers {
		log.Info(
			"Protocol tier",
			"id", tier.ID,
			"name", string(bytes.TrimRight(tier.VerifierName[:], "\x00")),
			"validityBond", utils.WeiToEther(tier.ValidityBond),
			"contestBond", utils.WeiToEther(tier.ContestBond),
			"provingWindow", tier.ProvingWindow,
			"cooldownWindow", tier.CooldownWindow,
		)

		switch tier.ID {
		case encoding.TierOptimisticID:
			p.tierFees = append(p.tierFees, encoding.TierFee{Tier: tier.ID, Fee: p.OptimisticTierFee})
		case encoding.TierSgxID:
			p.tierFees = append(p.tierFees, encoding.TierFee{Tier: tier.ID, Fee: p.SgxTierFee})
		case encoding.TierGuardianMinorityID:
			p.tierFees = append(p.tierFees, encoding.TierFee{Tier: tier.ID, Fee: common.Big0})
		case encoding.TierGuardianMajorityID:
			// Guardian prover should not charge any fee.
			p.tierFees = append(p.tierFees, encoding.TierFee{Tier: tier.ID, Fee: common.Big0})
		default:
			return fmt.Errorf("unknown tier: %d", tier.ID)
		}
	}

	return nil
}
