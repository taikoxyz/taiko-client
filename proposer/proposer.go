package proposer

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/txpool"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
	"github.com/taikoxyz/taiko-client/metrics"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	"github.com/urfave/cli/v2"
	"golang.org/x/sync/errgroup"
)

var (
	errNoNewTxs                = errors.New("no new transactions")
	maxSendProposeBlockTxRetry = 10
	retryInterval              = 12 * time.Second
)

type proposeBlockRequest struct {
	Input  *encoding.TaikoL1BlockMetadataInput `json:"input"`
	Fee    *big.Int                            `json:"fee"`
	Expiry uint64                              `json:"expiry"`
}

type proposeBlockResponse struct {
	SignedPayload []byte         `json:"signedPayload"`
	Prover        common.Address `json:"prover"`
}

// Proposer keep proposing new transactions from L2 execution engine's tx pool at a fixed interval.
type Proposer struct {
	// RPC clients
	rpc *rpc.Client

	// Private keys and account addresses
	l1ProposerPrivKey       *ecdsa.PrivateKey
	l1ProposerAddress       common.Address
	l2SuggestedFeeRecipient common.Address

	// Proposing configurations
	proposingInterval          *time.Duration
	proposeEmptyBlocksInterval *time.Duration
	proposingTimer             *time.Timer
	locals                     []common.Address
	localsOnly                 bool
	maxProposedTxListsPerEpoch uint64
	proposeBlockTxGasLimit     *uint64
	txReplacementTipMultiplier uint64
	proposeBlockTxGasTipCap    *big.Int
	proverEndpoints            []string

	// Protocol configurations
	protocolConfigs *bindings.TaikoDataConfig

	// Only for testing purposes
	CustomProposeOpHook func() error
	AfterCommitHook     func() error

	ctx context.Context
	wg  sync.WaitGroup

	waitReceiptTimeout time.Duration
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
	p.l1ProposerAddress = crypto.PubkeyToAddress(cfg.L1ProposerPrivKey.PublicKey)
	p.l2SuggestedFeeRecipient = cfg.L2SuggestedFeeRecipient
	p.proposingInterval = cfg.ProposeInterval
	p.proposeEmptyBlocksInterval = cfg.ProposeEmptyBlocksInterval
	p.proposeBlockTxGasLimit = cfg.ProposeBlockTxGasLimit
	p.wg = sync.WaitGroup{}
	p.locals = cfg.LocalAddresses
	p.localsOnly = cfg.LocalAddressesOnly
	p.maxProposedTxListsPerEpoch = cfg.MaxProposedTxListsPerEpoch
	p.txReplacementTipMultiplier = cfg.ProposeBlockTxReplacementMultiplier
	p.proposeBlockTxGasTipCap = cfg.ProposeBlockTxGasTipCap
	p.ctx = ctx
	p.waitReceiptTimeout = cfg.WaitReceiptTimeout
	p.proverEndpoints = cfg.ProverEndpoints

	// RPC clients
	if p.rpc, err = rpc.NewClient(p.ctx, &rpc.ClientConfig{
		L1Endpoint:     cfg.L1Endpoint,
		L2Endpoint:     cfg.L2Endpoint,
		TaikoL1Address: cfg.TaikoL1Address,
		TaikoL2Address: cfg.TaikoL2Address,
		RetryInterval:  cfg.BackOffRetryInterval,
		Timeout:        cfg.RPCTimeout,
	}); err != nil {
		return fmt.Errorf("initialize rpc clients error: %w", err)
	}

	// Protocol configs
	protocolConfigs, err := p.rpc.TaikoL1.GetConfig(&bind.CallOpts{Context: ctx})
	if err != nil {
		return fmt.Errorf("failed to get protocol configs: %w", err)
	}
	p.protocolConfigs = &protocolConfigs

	log.Info("Protocol configs", "configs", p.protocolConfigs)

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

	var lastNonEmptyBlockProposedAt = time.Now()
	for {
		p.updateProposingTicker()

		select {
		case <-p.ctx.Done():
			return
		// proposing interval timer has been reached
		case <-p.proposingTimer.C:
			metrics.ProposerProposeEpochCounter.Inc(1)
			// attempt propose operation
			if err := p.ProposeOp(p.ctx); err != nil {
				if !errors.Is(err, errNoNewTxs) {
					log.Error("Proposing operation error", "error", err)
					continue
				}
				// if no new transactions and empty block interval has passed, propose an empty block
				if p.proposeEmptyBlocksInterval != nil {
					if time.Now().Before(lastNonEmptyBlockProposedAt.Add(*p.proposeEmptyBlocksInterval)) {
						continue
					}

					if err := p.ProposeEmptyBlockOp(p.ctx); err != nil {
						log.Error("Proposing an empty block operation error", "error", err)
					}

					lastNonEmptyBlockProposedAt = time.Now()
				}

				continue
			}

			lastNonEmptyBlockProposedAt = time.Now()
		}
	}
}

// Close closes the proposer instance.
func (p *Proposer) Close(ctx context.Context) {
	p.wg.Wait()
}

// ProposeOp performs a proposing operation, fetching transactions
// from L2 execution engine's tx pool, splitting them by proposing constraints,
// and then proposing them to TaikoL1 contract.
func (p *Proposer) ProposeOp(ctx context.Context) error {
	if p.CustomProposeOpHook != nil {
		return p.CustomProposeOpHook()
	}

	// Wait until L2 execution engine is synced at first.
	if err := p.rpc.WaitTillL2ExecutionEngineSynced(ctx); err != nil {
		return fmt.Errorf("failed to wait until L2 execution engine synced: %w", err)
	}

	log.Info("Start fetching L2 execution engine's transaction pool content")

	l2Head, err := p.rpc.L2.HeaderByNumber(ctx, nil)
	if err != nil {
		return err
	}

	baseFee, err := p.rpc.TaikoL2.GetBasefee(
		&bind.CallOpts{Context: ctx},
		uint64(time.Now().Unix())-l2Head.Time,
		uint32(l2Head.GasUsed),
	)
	if err != nil {
		return err
	}

	log.Info("Current base fee", "fee", baseFee)

	txLists, err := p.rpc.GetPoolContent(
		ctx,
		p.L2SuggestedFeeRecipient(),
		baseFee,
		p.protocolConfigs.BlockMaxGasLimit,
		p.protocolConfigs.BlockMaxTxListBytes.Uint64(),
		p.locals,
		p.maxProposedTxListsPerEpoch,
	)
	if err != nil {
		return fmt.Errorf("failed to fetch transaction pool content: %w", err)
	}

	if p.localsOnly {
		var (
			localTxsLists []types.Transactions
			signer        = types.LatestSignerForChainID(p.rpc.L2ChainID)
		)
		for _, txs := range txLists {
			var filtered types.Transactions
			for _, tx := range txs {
				sender, err := types.Sender(signer, tx)
				if err != nil {
					return err
				}

				for _, localAddress := range p.locals {
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

	if len(txLists) == 0 {
		return errNoNewTxs
	}

	head, err := p.rpc.L1.BlockNumber(ctx)
	if err != nil {
		return err
	}
	nonce, err := p.rpc.L1.NonceAt(
		ctx,
		crypto.PubkeyToAddress(p.l1ProposerPrivKey.PublicKey),
		new(big.Int).SetUint64(head),
	)
	if err != nil {
		return err
	}

	log.Info("Proposer account information", "chainHead", head, "nonce", nonce)

	g := new(errgroup.Group)
	for i, txs := range txLists {
		func(i int, txs types.Transactions) {
			g.Go(func() error {
				if i >= int(p.maxProposedTxListsPerEpoch) {
					return nil
				}

				txListBytes, err := rlp.EncodeToBytes(txs)
				if err != nil {
					return fmt.Errorf("failed to encode transactions: %w", err)
				}

				txNonce := nonce + uint64(i)
				if err := p.ProposeTxList(ctx, &encoding.TaikoL1BlockMetadataInput{
					Beneficiary:     p.l2SuggestedFeeRecipient,
					TxListHash:      crypto.Keccak256Hash(txListBytes),
					TxListByteStart: common.Big0,
					TxListByteEnd:   new(big.Int).SetUint64(uint64(len(txListBytes))),
					CacheTxListInfo: false,
				}, txListBytes, uint(txs.Len()), &txNonce); err != nil {
					return fmt.Errorf("failed to propose transactions: %w", err)
				}

				return nil
			})
		}(i, txs)
	}

	if err := g.Wait(); err != nil {
		return fmt.Errorf("failed to propose transactions: %w", err)
	}

	if p.AfterCommitHook != nil {
		if err := p.AfterCommitHook(); err != nil {
			log.Error("Run AfterCommitHook error", "error", err)
		}
	}

	return nil
}

// sendProposeBlockTx tries to send a TaikoL1.proposeBlock transaction.
func (p *Proposer) sendProposeBlockTx(
	ctx context.Context,
	meta *encoding.TaikoL1BlockMetadataInput,
	txListBytes []byte,
	nonce *uint64,
	assignment []byte,
	fee *big.Int,
	isReplacement bool,
) (*types.Transaction, error) {
	// Propose the transactions list
	inputs, err := encoding.EncodeProposeBlockInput(meta)
	if err != nil {
		return nil, err
	}
	opts, err := getTxOpts(ctx, p.rpc.L1, p.l1ProposerPrivKey, p.rpc.L1ChainID, fee)
	if err != nil {
		return nil, err
	}
	if nonce != nil {
		opts.Nonce = new(big.Int).SetUint64(*nonce)
	}
	if p.proposeBlockTxGasLimit != nil {
		opts.GasLimit = *p.proposeBlockTxGasLimit
	}
	if isReplacement {
		log.Info("Try replacing a transaction with same nonce", "sender", p.l1ProposerAddress, "nonce", nonce)
		originalTx, err := rpc.GetPendingTxByNonce(ctx, p.rpc, p.l1ProposerAddress, *nonce)
		if err != nil || originalTx == nil {
			log.Warn(
				"Original transaction not found",
				"sender", p.l1ProposerAddress,
				"nonce", nonce,
				"error", err,
			)

			opts.GasTipCap = new(big.Int).Mul(opts.GasTipCap, new(big.Int).SetUint64(p.txReplacementTipMultiplier))
		} else {
			log.Info(
				"Original transaction to replace",
				"sender", p.l1ProposerAddress,
				"nonce", nonce,
				"gasTipCap", originalTx.GasTipCap(),
				"gasFeeCap", originalTx.GasFeeCap(),
			)

			opts.GasTipCap = new(big.Int).Mul(
				originalTx.GasTipCap(),
				new(big.Int).SetUint64(p.txReplacementTipMultiplier),
			)
		}

		if p.proposeBlockTxGasTipCap != nil && opts.GasTipCap.Cmp(p.proposeBlockTxGasTipCap) > 0 {
			log.Info(
				"New gasTipCap exceeds limit, keep waiting",
				"multiplier", p.txReplacementTipMultiplier,
				"newGasTipCap", opts.GasTipCap,
				"maxTipCap", p.proposeBlockTxGasTipCap,
			)
			return nil, txpool.ErrReplaceUnderpriced
		}
	}

	proposeTx, err := p.rpc.TaikoL1.ProposeBlock(opts, inputs, assignment, txListBytes)
	if err != nil {
		return nil, encoding.TryParsingCustomError(err)
	}

	return proposeTx, nil
}

// ProposeTxList proposes the given transactions list to TaikoL1 smart contract.
func (p *Proposer) ProposeTxList(
	ctx context.Context,
	meta *encoding.TaikoL1BlockMetadataInput,
	txListBytes []byte,
	txNum uint,
	nonce *uint64,
) error {
	assignment, fee, err := p.assignProver(ctx, meta)
	if err != nil {
		return err
	}

	var (
		isReplacement bool
		tx            *types.Transaction
	)
	if err := backoff.Retry(
		func() error {
			if ctx.Err() != nil {
				return nil
			}
			if tx, err = p.sendProposeBlockTx(ctx, meta, txListBytes, nonce, assignment, fee, isReplacement); err != nil {
				log.Warn("Failed to send propose block transaction, retrying", "error", encoding.TryParsingCustomError(err))
				if strings.Contains(err.Error(), txpool.ErrReplaceUnderpriced.Error()) {
					isReplacement = true
				} else {
					isReplacement = false
				}
				return err
			}

			return nil
		},
		backoff.WithMaxRetries(
			backoff.NewConstantBackOff(retryInterval),
			uint64(maxSendProposeBlockTxRetry),
		),
	); err != nil {
		return err
	}
	if ctx.Err() != nil {
		return ctx.Err()
	}
	if err != nil {
		return err
	}

	ctxWithTimeout, cancel := context.WithTimeout(ctx, p.waitReceiptTimeout)
	defer cancel()

	if _, err := rpc.WaitReceipt(ctxWithTimeout, p.rpc.L1, tx); err != nil {
		return err
	}

	log.Info("📝 Propose transactions succeeded", "txs", txNum)

	metrics.ProposerProposedTxListsCounter.Inc(1)
	metrics.ProposerProposedTxsCounter.Inc(int64(txNum))

	return nil
}

// ProposeEmptyBlockOp performs a proposing one empty block operation.
func (p *Proposer) ProposeEmptyBlockOp(ctx context.Context) error {
	return p.ProposeTxList(ctx, &encoding.TaikoL1BlockMetadataInput{
		TxListHash:      crypto.Keccak256Hash([]byte{}),
		Beneficiary:     p.L2SuggestedFeeRecipient(),
		TxListByteStart: common.Big0,
		TxListByteEnd:   common.Big0,
		CacheTxListInfo: false,
	}, []byte{}, 0, nil)
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
		// Random number between 12 - 120
		randomSeconds := rand.Intn(120-11) + 12
		duration = time.Duration(randomSeconds) * time.Second
	}

	p.proposingTimer = time.NewTimer(duration)
}

// assignProver attempts to find a prover who wants to win the block.
// if no provers want to do it for the price we set, we increase the price, and try again.
// TODO(jeff): increase fee on each iteration
// TODO(jeff): calculate fee
func (p *Proposer) assignProver(
	ctx context.Context,
	meta *encoding.TaikoL1BlockMetadataInput,
) ([]byte, *big.Int, error) {
	fee := big.NewInt(100000)
	expiry := uint64(time.Now().Add(90 * time.Minute).Unix())
	// TODO(jeff): fee and expiry dynamically
	proposeBlockReq := proposeBlockRequest{
		Expiry: expiry,
		Input:  meta,
		Fee:    fee,
	}

	jsonBody, err := json.Marshal(proposeBlockReq)
	if err != nil {
		return nil, nil, err
	}

	r := bytes.NewReader(jsonBody)

	// iterate over each configured endpoint, and see if someone wants to accept your block.
	// if it is denied, we continue on to the next endpoint.
	// if we do not find a prover, we can increase the fee up to a point, or give up.
	for _, endpoint := range p.proverEndpoints {
		req, err := http.NewRequestWithContext(
			ctx,
			"POST",
			fmt.Sprintf("%v/%v", endpoint, "proposeBlock"),
			r,
		)
		if err != nil {
			continue
		}

		req.Header.Add("Content-Type", "application/json")

		// TODO(jeff): custom client with timeout
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			continue
		}

		if res.StatusCode != http.StatusOK {
			continue
		}

		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			continue
		}

		resp := &proposeBlockResponse{}

		if err := json.Unmarshal(resBody, resp); err != nil {
			continue
		}

		log.Info("Prover assigned for block",
			"prover", resp.Prover.Hex(),
			"signedPayload", common.Bytes2Hex(resp.SignedPayload),
		)

		// TODO(jeff): do an ecrecover here, and make sure data signed is what we sent,
		// and signed by prover we received,
		// to ensure it will not revert onchain
		encoded, err := encoding.EncodeProverAssignment(&encoding.ProverAssignment{
			Prover: resp.Prover,
			Expiry: proposeBlockReq.Expiry,
			Data:   resp.SignedPayload,
		})
		if err != nil {
			return nil, nil, err
		}

		return encoded, fee, nil
	}

	return nil, nil, errors.New("unable to find prover")
}

// Name returns the application name.
func (p *Proposer) Name() string {
	return "proposer"
}

// L2SuggestedFeeRecipient returns the L2 suggested fee recipient of the current proposer.
func (p *Proposer) L2SuggestedFeeRecipient() common.Address {
	return p.l2SuggestedFeeRecipient
}

// getTxOpts creates a bind.TransactOpts instance using the given private key.
func getTxOpts(
	ctx context.Context,
	cli *rpc.EthClient,
	privKey *ecdsa.PrivateKey,
	chainID *big.Int,
	fee *big.Int,
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

	opts.GasTipCap = gasTipCap

	opts.Value = fee

	return opts, nil
}
