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

// ProofContester is responsible for contesting wrong L2Client transitions.
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
	l2SignalService, err := rpcClient.TaikoL2.Resolve0(
		nil,
		rpc.StringToBytes32("signal_service"),
		false,
	)
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

// SubmitContest submits a taikoL1.proveBlock transaction to contest a L2Client block transition.
func (c *ProofContester) SubmitContest(
	ctx context.Context,
	blockID *big.Int,
	proposedIn *big.Int,
	parentHash common.Hash,
	meta *bindings.TaikoDataBlockMetadata,
	tier uint16,
) error {
	// Ensure the transition has not been contested yet.
	transition, err := c.rpc.TaikoL1.GetTransition(
		&bind.CallOpts{Context: ctx},
		blockID.Uint64(),
		parentHash,
	)
	if err != nil {
		if !strings.Contains(encoding.TryParsingCustomError(err).Error(), "L1_") {
			log.Warn(
				"Failed to get transition",
				"blockID", blockID,
				"parentHash", parentHash,
				"error", encoding.TryParsingCustomError(err),
			)
			return nil
		}
		return err
	}
	if transition.Contester != (common.Address{}) {
		log.Info(
			"Transaction has already been contested",
			"blockID", blockID,
			"parentHash", parentHash,
			"contester", transition.Contester,
		)
		return nil
	}

	header, err := c.rpc.L2Client.HeaderByNumber(ctx, blockID)
	if err != nil {
		return err
	}

	signalRoot, err := c.rpc.GetStorageRoot(ctx, c.rpc.L2Client, c.l2SignalService, blockID)
	if err != nil {
		return fmt.Errorf("failed to get L2Client signal service storage root: %w", err)
	}

	l1HeaderProposedIn, err := c.rpc.L1Client.HeaderByNumber(ctx, proposedIn)
	if err != nil {
		return err
	}

	if err := c.txSender.Send(
		ctx,
		&proofProducer.ProofWithHeader{
			BlockID: blockID,
			Meta:    meta,
			Header:  header,
			Proof:   []byte{},
			Opts: &proofProducer.ProofRequestOptions{
				EventL1Hash: l1HeaderProposedIn.Hash(),
				SignalRoot:  signalRoot,
			},
			Tier: tier,
		},
		c.txBuilder.Build(
			ctx,
			blockID,
			meta,
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
