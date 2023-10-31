package submitter

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	proofProducer "github.com/taikoxyz/taiko-client/prover/proof_producer"
	"github.com/taikoxyz/taiko-client/prover/proof_submitter/transaction"
)

var _ Contester = (*ProofContester)(nil)

// ProofContester is responsible for contesting wrong L2 transitions.
type ProofContester struct {
	rpc             *rpc.Client
	txBuilder       *transaction.ProveBlockTxBuilder
	txSender        *transaction.Sender
	l2SignalService common.Address
	graffiti        [32]byte
}

// NewProofContester creates a new ProofContester instance.
func NewProofContester(
	rpcClient *rpc.Client,
	proverPrivKey *ecdsa.PrivateKey,
	proveBlockTxGasLimit *uint64,
	txReplacementTipMultiplier uint64,
	proveBlockMaxTxGasTipCap *big.Int,
	submissionMaxRetry uint64,
	retryInterval time.Duration,
	waitReceiptTimeout time.Duration,
	graffiti string,
) (*ProofContester, error) {
	l2SignalService, err := rpcClient.TaikoL2.SignalService(nil)
	if err != nil {
		return nil, err
	}

	var txGasLimit *big.Int
	if proveBlockTxGasLimit != nil {
		txGasLimit = new(big.Int).SetUint64(*proveBlockTxGasLimit)
	}

	return &ProofContester{
		rpc: rpcClient,
		txBuilder: transaction.NewProveBlockTxBuilder(
			rpcClient,
			proverPrivKey,
			txGasLimit,
			proveBlockMaxTxGasTipCap,
			new(big.Int).SetUint64(txReplacementTipMultiplier),
		),
		txSender:        transaction.NewSender(rpcClient, retryInterval, &submissionMaxRetry, waitReceiptTimeout),
		l2SignalService: l2SignalService,
		graffiti:        rpc.StringToBytes32(graffiti),
	}, nil
}

// SubmitContest submits a taikoL1.proveBlock transaction to contest a L2 block transition.
func (c *ProofContester) SubmitContest(
	ctx context.Context,
	blockProposedEvent *bindings.TaikoL1ClientBlockProposed,
	transitionProvedEvent *bindings.TaikoL1ClientTransitionProved,
) error {
	// Ensure the transition has not been contested yet.
	transition, err := c.rpc.TaikoL1.GetTransition(
		&bind.CallOpts{Context: ctx},
		transitionProvedEvent.BlockId.Uint64(),
		transitionProvedEvent.Tran.ParentHash,
	)
	if err != nil {
		if !strings.Contains(encoding.TryParsingCustomError(err).Error(), "L1_") {
			log.Warn(
				"Failed to get transition",
				"blockID", transitionProvedEvent.BlockId,
				"parentHash", transitionProvedEvent.Tran.ParentHash,
				"error", encoding.TryParsingCustomError(err),
			)
			return nil
		}
		return err
	}
	if transition.Contester != (common.Address{}) {
		log.Info(
			"Transaction has already been contested",
			"blockID", transitionProvedEvent.BlockId,
			"parentHash", transitionProvedEvent.Tran.ParentHash,
			"contester", transition.Contester,
		)
		return nil
	}

	header, err := c.rpc.L2.HeaderByNumber(ctx, transitionProvedEvent.BlockId)
	if err != nil {
		return err
	}

	signalRoot, err := c.rpc.GetStorageRoot(ctx, c.rpc.L2GethClient, c.l2SignalService, transitionProvedEvent.BlockId)
	if err != nil {
		return fmt.Errorf("failed to get L2 signal service storage root: %w", err)
	}

	if err := c.txSender.Send(
		ctx,
		&proofProducer.ProofWithHeader{
			BlockID: transitionProvedEvent.BlockId,
			Meta:    &blockProposedEvent.Meta,
			Header:  header,
			Proof:   []byte{},
			Opts: &proofProducer.ProofRequestOptions{
				EventL1Hash: blockProposedEvent.Raw.BlockHash,
				SignalRoot:  signalRoot,
			},
			Tier: transitionProvedEvent.Tier,
		},
		c.txBuilder.BuildForNormalProofSubmission(
			ctx,
			transitionProvedEvent.BlockId,
			&blockProposedEvent.Meta,
			&bindings.TaikoDataTransition{
				ParentHash: header.ParentHash,
				BlockHash:  header.Hash(),
				SignalRoot: signalRoot,
				Graffiti:   c.graffiti,
			},
			&bindings.TaikoDataTierProof{
				Tier: transition.Tier,
				Data: []byte{},
			},
			false,
		),
	); err != nil {
		if errors.Is(err, transaction.ErrUnretryable) {
			return nil
		}

		return err
	}
	return nil
}
