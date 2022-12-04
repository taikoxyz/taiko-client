package prover

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/metrics"
	eventIterator "github.com/taikoxyz/taiko-client/pkg/chain_iterator/event_iterator"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	txListValidator "github.com/taikoxyz/taiko-client/pkg/tx_list_validator"
	"github.com/taikoxyz/taiko-client/prover/producer"
	"github.com/urfave/cli/v2"
)

var (
	// Gas limit of TaikoL1.proveBlock and TaikoL1.proveBlockInvalid transactions.
	// TODO: tune this value based when the on-chain solidity verifier is available.
	proveBlocksGasLimit uint64 = 1000000
	maxPendingProofs           = 10
)

// Prover keep trying to prove new proposed blocks valid/invalid.
type Prover struct {
	// Configurations
	cfg *Config

	// Clients
	rpc *rpc.Client

	// Contract configurations
	txListValidator   *txListValidator.TxListValidator
	protocolConstants *bindings.ProtocolConstants

	// States
	lastVerifiedHeader   *types.Header
	lastVerifiedL1Height uint64
	l1Current            uint64

	// Subscriptions
	blockProposedCh  chan *bindings.TaikoL1ClientBlockProposed
	blockProposedSub event.Subscription
	blockVerifiedCh  chan *bindings.TaikoL1ClientBlockVerified
	blockVerifiedSub event.Subscription
	proveNotify      chan struct{}

	// Proof related
	proveValidProofCh   chan *producer.ProofWithHeader
	proveInvalidProofCh chan *producer.ProofWithHeader
	proofProducer       producer.ProofProducer

	ctx context.Context
	wg  sync.WaitGroup
}

// New initializes the given prover instance based on the command line flags.
func (p *Prover) InitFromCli(ctx context.Context, c *cli.Context) error {
	cfg, err := NewConfigFromCliContext(c)
	if err != nil {
		return err
	}

	return InitFromConfig(ctx, p, cfg)
}

// InitFromConfig initializes the prover instance based on the given configurations.
func InitFromConfig(ctx context.Context, p *Prover, cfg *Config) (err error) {
	log.Debug("Prover configurations", "config", cfg)

	p.cfg = cfg
	p.ctx = ctx

	// Clients
	if p.rpc, err = rpc.NewClient(p.ctx, &rpc.ClientConfig{
		L1Endpoint:     cfg.L1Endpoint,
		L2Endpoint:     cfg.L2Endpoint,
		TaikoL1Address: cfg.TaikoL1Address,
		TaikoL2Address: cfg.TaikoL2Address,
	}); err != nil {
		return err
	}

	proverAddress := crypto.PubkeyToAddress(p.cfg.L1ProverPrivKey.PublicKey)
	isWhitelisted, err := p.isWhitelisted(proverAddress)
	if err != nil {
		return fmt.Errorf("failed to check whether current prover %s is whitelisted: %w", proverAddress, err)
	}

	if !isWhitelisted {
		return fmt.Errorf("prover %s is not whitelisted", proverAddress)
	}

	// Constants
	if p.protocolConstants, err = p.rpc.GetProtocolConstants(nil); err != nil {
		return fmt.Errorf("failed to get protocol constants: %w", err)
	}

	log.Info("Protocol constants", "constants", p.protocolConstants)

	p.txListValidator = txListValidator.NewTxListValidator(
		p.protocolConstants.BlockMaxGasLimit.Uint64(),
		p.protocolConstants.BlockMaxTxs.Uint64(),
		p.protocolConstants.TxListMaxBytes.Uint64(),
		p.protocolConstants.TxMinGasLimit.Uint64(),
		p.rpc.L2ChainID,
	)
	p.blockProposedCh = make(chan *bindings.TaikoL1ClientBlockProposed, p.protocolConstants.MaxProposedBlocks.Uint64())
	p.blockVerifiedCh = make(chan *bindings.TaikoL1ClientBlockVerified, p.protocolConstants.MaxProposedBlocks.Uint64())
	p.proveValidProofCh = make(chan *producer.ProofWithHeader, p.protocolConstants.MaxProposedBlocks.Uint64())
	p.proveInvalidProofCh = make(chan *producer.ProofWithHeader, p.protocolConstants.MaxProposedBlocks.Uint64())
	p.proveNotify = make(chan struct{}, 1)
	if err := p.initL1Current(); err != nil {
		return fmt.Errorf("initialize L1 current cursor error: %w", err)
	}

	if cfg.Dummy {
		p.proofProducer = new(producer.DummyProofProducer)
	} else {
		if p.proofProducer, err = producer.NewZkevmRpcdProducer(cfg.ZKEvmRpcdEndpoint); err != nil {
			return err
		}
	}

	return nil
}

// Start starts the main loop of the L2 block prover.
func (p *Prover) Start() error {
	p.wg.Add(1)
	p.startSubscription()
	go p.eventLoop()

	return nil
}

// eventLoop starts the main loop of Taiko prover.
func (p *Prover) eventLoop() {
	defer func() {
		p.wg.Done()
	}()

	// reqProving requests performing a proving operation, won't block
	// if we are already proving.
	reqProving := func() {
		select {
		case p.proveNotify <- struct{}{}:
		default:
		}
	}

	// Call reqProving() right away to catch up with the latest state.
	reqProving()

	for {
		select {
		case <-p.ctx.Done():
			return
		case proofWithHeader := <-p.proveValidProofCh:
			if err := p.submitValidBlockProof(p.ctx, proofWithHeader); err != nil {
				log.Error("Prove valid block error", "error", err)
			}
		case proofWithHeader := <-p.proveInvalidProofCh:
			if err := p.submitInvalidBlockProof(p.ctx, proofWithHeader); err != nil {
				log.Error("Prove invalid block error", "error", err)
			}
		case <-p.proveNotify:
			if err := p.proveOp(); err != nil {
				log.Error("Prove new blocks error", "error", err)
			}
		case <-p.blockProposedCh:
			reqProving()
		case e := <-p.blockVerifiedCh:
			if err := p.onBlockVerified(p.ctx, e); err != nil {
				log.Error("Handle BlockVerified event error", "error", err)
			}
		}
	}
}

// Close closes the prover instance.
func (p *Prover) Close() {
	p.closeSubscription()
	p.wg.Wait()
}

// proveOp perfors a proving operation, find current unproven blocks, then
// request generating proofs for them.
func (p *Prover) proveOp() error {
	isHalted, err := p.rpc.TaikoL1.IsHalted(nil)
	if err != nil {
		return err
	}

	if isHalted {
		log.Warn("L2 chain halted")
		return nil
	}

	iter, err := eventIterator.NewBlockProposedIterator(p.ctx, &eventIterator.BlockProposedIteratorConfig{
		Client:               p.rpc.L1,
		TaikoL1:              p.rpc.TaikoL1,
		StartHeight:          new(big.Int).SetUint64(p.l1Current),
		OnBlockProposedEvent: p.onBlockProposed,
	})
	if err != nil {
		return err
	}

	return iter.Iter()
}

// onBlockProposed tries to prove that the newly proposed block is valid/invalid.
func (p *Prover) onBlockProposed(
	ctx context.Context,
	event *bindings.TaikoL1ClientBlockProposed,
	end eventIterator.EndBlockProposeEventIterFunc,
) error {
	if len(p.proveValidProofCh) > maxPendingProofs || len(p.proveInvalidProofCh) > maxPendingProofs {
		end()
		return nil
	}
	log.Info("Proposed block", "blockID", event.Id)
	metrics.ProverReceivedProposedBlockGauge.Update(event.Id.Int64())

	// Check whether the block has been verified.
	isVerified, err := p.isBlockVerified(event.Id)
	if err != nil {
		return err
	}

	if isVerified {
		log.Info("Block is verified", "blockID", event.Id)
		return nil
	}

	// Check whether the transactions list is valid.
	proposeBlockTx, err := p.rpc.L1.TransactionInBlock(ctx, event.Raw.BlockHash, event.Raw.TxIndex)
	if err != nil {
		return err
	}

	hint, invalidTxIndex, err := p.txListValidator.ValidateTxList(event.Id, proposeBlockTx.Data())
	if err != nil {
		return err
	}

	// Prove the proposed block is valid.
	if hint == txListValidator.HintOK {
		return p.proveBlockValid(ctx, event)
	}

	// Otherwise, prove the proposed block is invalid.
	return p.proveBlockInvalid(ctx, event, hint, invalidTxIndex)
}

// onBlockVerified update the lastVerified block in current state.
func (p *Prover) onBlockVerified(ctx context.Context, event *bindings.TaikoL1ClientBlockVerified) error {
	if event.BlockHash == (common.Hash{}) {
		log.Info("Ignore BlockVerified event of invalid transaction list", "blockID", event.Id)
		return nil
	}

	metrics.ProverLatestVerifiedIDGauge.Update(event.Id.Int64())

	l2BlockHeader, err := p.rpc.L2.HeaderByHash(ctx, event.BlockHash)
	if err != nil {
		return fmt.Errorf("failed to find L2 block with hash %s: %w", common.BytesToHash(event.BlockHash[:]), err)
	}

	log.Info(
		"New verified block",
		"blockID", event.Id,
		"height", l2BlockHeader.Number,
		"hash", common.BytesToHash(event.BlockHash[:]),
	)
	p.lastVerifiedHeader = l2BlockHeader
	p.lastVerifiedL1Height = event.Raw.BlockNumber

	return nil
}

// Name returns the application name.
func (p *Prover) Name() string {
	return "prover"
}

// getProveBlocksTxOpts creates a bind.TransactOpts instance using the given private key.
// Used for creating TaikoL1.proveBlock and TaikoL1.proveBlockInvalid transactions.
func (p *Prover) getProveBlocksTxOpts(ctx context.Context, cli *ethclient.Client) (*bind.TransactOpts, error) {
	opts, err := bind.NewKeyedTransactorWithChainID(p.cfg.L1ProverPrivKey, p.rpc.L1ChainID)
	if err != nil {
		return nil, err
	}
	gasTipCap, err := cli.SuggestGasTipCap(ctx)
	if err != nil {
		if rpc.IsMaxPriorityFeePerGasNotFoundError(err) {
			gasTipCap = rpc.FallbackGasTipCap
		} else {
			return nil, err
		}
	}

	opts.GasTipCap = gasTipCap
	opts.GasLimit = proveBlocksGasLimit

	return opts, nil
}

func (p *Prover) initL1Current() error {
	_, _, latestVerifiedID, _, err := p.rpc.TaikoL1.GetStateVariables(nil)
	if err != nil {
		return err
	}

	if latestVerifiedID == 0 {
		p.l1Current = 0
		return nil
	}

	latestVerifiedHeaderL1Origin, err := p.rpc.L2.L1OriginByID(p.ctx, new(big.Int).SetUint64(latestVerifiedID))
	if err != nil {
		return err
	}

	p.l1Current = latestVerifiedHeaderL1Origin.L1BlockHeight.Uint64()
	return nil
}

// isBlockVerified checks whether the given block has been verified by other provers.
func (p *Prover) isBlockVerified(id *big.Int) (bool, error) {
	_, _, latestVerifiedID, _, err := p.rpc.TaikoL1.GetStateVariables(nil)
	if err != nil {
		return false, err
	}

	return id.Uint64() <= latestVerifiedID, nil
}

// isWhitelisted checks whether the current prover is whitelisted.
func (p *Prover) isWhitelisted(proverAddress common.Address) (bool, error) {
	isWhitelisted, err := p.rpc.TaikoL1.IsProverWhitelisted(nil, proverAddress)
	if err != nil {
		if strings.Contains(err.Error(), "Assertion error") { // whitelist feature disabled, everyone can submit proofs
			return true, nil
		}

		return false, err
	}

	return isWhitelisted, nil
}
