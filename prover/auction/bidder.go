package auction

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/log"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
)

type Bidder struct {
	strategy      Strategy
	rpc           *rpc.Client
	privateKey    *ecdsa.PrivateKey
	proverAddress common.Address
}

func NewBidder(strategy Strategy, rpc *rpc.Client, privateKey *ecdsa.PrivateKey, proverAddress common.Address) (*Bidder, error) {
	return &Bidder{
		strategy:      strategy,
		rpc:           rpc,
		privateKey:    privateKey,
		proverAddress: proverAddress,
	}, nil
}

func (b *Bidder) SubmitBid(ctx context.Context, batchID *big.Int, amount *big.Int) (*types.Transaction, error) {
	isBatchAuctionable, err := b.rpc.TaikoL1.IsBatchAuctionable(nil, batchID)
	if err != nil {
		return nil, fmt.Errorf("error checking whether batch is auctionable: %w", err)
	}

	if !isBatchAuctionable {
		return nil, fmt.Errorf("trying to submit bid for unauctionable batchID: %w", err)
	}

	auctions, err := b.rpc.TaikoL1.GetAuctions(nil, batchID, new(big.Int).SetInt64(1))
	if err != nil {
		return nil, fmt.Errorf("error getting auctions for bid: %w", err)
	}

	currentBid := auctions.Auctions[0].Bid

	log.Info("Current bid for batch ID",
		batchID,
		"currentBidDeposit",
		currentBid.Deposit,
		"currentBidAmount",
		currentBid.FeePerGas,
		"blockMaxGasLimit",
		currentBid.BlockMaxGasLimit,
		"prover",
		currentBid.Prover,
	)

	shouldBid, err := b.strategy.ShouldBid(ctx, currentBid)
	if err != nil {
		return nil, fmt.Errorf("error determing if should bid on current auction: %w", err)
	}

	if !shouldBid {
		log.Info("Bid strategy determined to not bid on current auction for batch ID",
			batchID)
	}

	bid, err := b.strategy.NextBid(ctx, b.proverAddress, currentBid)
	if err != nil {
		return nil, fmt.Errorf("error crafting next bid: %w", err)
	}

	isBetter, err := b.rpc.TaikoL1.IsBidBetter(nil, bid, currentBid)
	if err != nil {
		return nil, fmt.Errorf("error determing if bid is better than existing bid: %w", err)
	}

	if !isBetter {
		return nil, fmt.Errorf("crafted a bid that is not better than existing bid: %w", err)
	}

	opts, err := getTxOpts(ctx, b.rpc.L1, b.privateKey, b.rpc.L1ChainID)
	if err != nil {
		return nil, err
	}
	tx, err := b.rpc.TaikoL1.BidForBatch(opts, batchID.Uint64(), bid)

	if _, err := rpc.WaitReceipt(ctx, b.rpc.L1, tx); err != nil {
		return nil, err
	}

	log.Info("📝 Bid for batch tx succeeded", "txHash", tx.Hash(), "batchID", batchID)

	return nil, nil
}

func getTxOpts(
	ctx context.Context,
	cli *ethclient.Client,
	privKey *ecdsa.PrivateKey,
	chainID *big.Int,
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

	return opts, nil
}
