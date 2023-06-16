package auction

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/log"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
	"github.com/taikoxyz/taiko-client/metrics"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
)

type Bidder struct {
	strategy       Strategy
	rpc            *rpc.Client
	privateKey     *ecdsa.PrivateKey
	proverAddress  common.Address
	protocolConfig *bindings.TaikoDataConfig
}

func NewBidder(
	strategy Strategy,
	rpc *rpc.Client,
	privateKey *ecdsa.PrivateKey,
	proverAddress common.Address,
	protocolConfig *bindings.TaikoDataConfig,
) *Bidder {
	return &Bidder{
		strategy:       strategy,
		rpc:            rpc,
		privateKey:     privateKey,
		proverAddress:  proverAddress,
		protocolConfig: protocolConfig,
	}
}

func (b *Bidder) SubmitBid(ctx context.Context, batchID *big.Int) error {
	isBatchAuctionable, err := b.rpc.TaikoL1.IsBatchAuctionable(nil, batchID)
	if err != nil {
		return fmt.Errorf("error checking whether batch is auctionable: %w", err)
	}

	if !isBatchAuctionable {
		return fmt.Errorf("trying to submit bid for unauctionable batchID: %v", batchID.Uint64())
	}

	auctions, err := b.rpc.TaikoL1.GetAuctions(nil, batchID, new(big.Int).SetInt64(1))
	if err != nil {
		return fmt.Errorf("error getting auctions for bid: %w", err)
	}

	currentBid := auctions.Auctions[0].Bid

	if currentBid.Prover == b.proverAddress {
		log.Info("not bidding for batch, already current winner", "batchId", batchID.Uint64())
		return nil
	}

	log.Info("Current bid for batch ID ",
		batchID.Uint64(),
		"currentBidDeposit",
		currentBid.Deposit,
		"currentBidAmount",
		currentBid.FeePerGas,
		"prover",
		currentBid.Prover,
		"proofWindow",
		currentBid.ProofWindow,
	)

	shouldBid, err := b.strategy.ShouldBid(ctx, currentBid)
	if err != nil {
		return fmt.Errorf("error determining if should bid on current auction: %w", err)
	}

	if !shouldBid {
		log.Info("Bid strategy determined to not bid on current auction", "batchID",
			batchID.Uint64())
	} else {
		log.Info("Bid strategy determined to bid on current auction", "batchID",
			batchID.Uint64())
	}

	bid, err := b.strategy.NextBid(ctx, b.proverAddress, currentBid)
	if err != nil {
		return fmt.Errorf("error crafting next bid: %w", err)
	}

	stateVars, _ := b.rpc.TaikoL1.State(nil)
	log.Info("state vars", "avgProofTime", stateVars.AvgProofWindow)

	log.Info("Next bid",
		"batchID",
		batchID.Uint64(),
		"currentBidDeposit",
		bid.Deposit,
		"currentBidAmount",
		bid.FeePerGas,
		"proofWindow",
		bid.ProofWindow,
	)

	// if there is an eixsting bid, we need to see if ours is better
	if currentBid.Prover != common.HexToAddress("0x0000000000000000000000000000000000000000") {
		isBetter, err := b.rpc.TaikoL1.IsBidBetter(nil, bid, currentBid)
		if err != nil {
			return fmt.Errorf("error determining if bid is better than existing bid: %w", err)
		}

		if !isBetter {
			return fmt.Errorf("crafted a bid that is not better than existing bid: %w", err)
		}
	}

	requiredDepositAmount, err := b.getRequiredDepositAmount(ctx, bid)
	if err != nil {
		return fmt.Errorf("error getting required deposit amount: %w", err)
	}

	log.Info("required deposit amount", "batchID", batchID.Uint64(), "amount", requiredDepositAmount.String())

	if requiredDepositAmount.Cmp(big.NewInt(0)) > 0 {
		if err := b.deposit(ctx, requiredDepositAmount); err != nil {
			return fmt.Errorf("error depositing taiko token: %w", err)
		}
	}

	if err := b.submitBid(ctx, batchID, bid); err != nil {
		return fmt.Errorf("error submitting bid: %w", encoding.TryParsingCustomError(err))
	}

	metrics.ProverAuctionableBatchBidGauge.Update(int64(batchID.Uint64()))

	return nil
}

// getRequiredDepositAmount calculates the difference between the prover's current state taikoTokenBalance,
// and what will be required for the bid.
func (b *Bidder) getRequiredDepositAmount(ctx context.Context, bid bindings.TaikoDataBid) (*big.Int, error) {
	balance, err := b.rpc.TaikoL1.GetTaikoTokenBalance(nil, b.proverAddress)
	if err != nil {
		return big.NewInt(0), fmt.Errorf("error getting taiko token balance: %w", err)
	}

	deposit := new(big.Int).SetUint64(bid.Deposit)

	// deposit is per block, so to calculate total deposit, we multiply the per-block deposit by the batch size.
	totalDeposit := new(big.Int).Mul(deposit, new(big.Int).SetUint64(uint64(b.protocolConfig.AuctionBatchSize)))

	if balance.Cmp(totalDeposit) >= 0 {
		return big.NewInt(0), nil
	} else {
		return new(big.Int).Sub(totalDeposit, balance), nil
	}
}

func (b *Bidder) deposit(ctx context.Context, amount *big.Int) error {
	opts, err := getTxOpts(ctx, b.rpc.L1, b.privateKey, b.rpc.L1ChainID)
	if err != nil {
		return err
	}

	log.Info("depositing taiko tokens", "amount", amount.String())

	tx, err := b.rpc.TaikoL1.DepositTaikoToken(opts, amount.Uint64())
	if err != nil {
		return err
	}

	if _, err := rpc.WaitReceipt(ctx, b.rpc.L1, tx); err != nil {
		return err
	}

	log.Info("📝 Deposited Taiko Token", "txHash", tx.Hash())

	return nil
}

func (b *Bidder) submitBid(ctx context.Context, batchID *big.Int, bid bindings.TaikoDataBid) error {
	opts, err := getTxOpts(ctx, b.rpc.L1, b.privateKey, b.rpc.L1ChainID)
	if err != nil {
		return err
	}

	log.Info("Sending bid for batch",
		"batchID",
		batchID.Uint64(),
		"bidFeePerGas",
		bid.FeePerGas,
		"deposit",
		bid.Deposit,
		"proofWindow",
		bid.ProofWindow,
	)

	tx, err := b.rpc.TaikoL1.BidForBatch(opts, batchID.Uint64(), bid)
	if err != nil {
		return fmt.Errorf("error submitting bid for batch: %w", err)
	}

	if _, err := rpc.WaitReceipt(ctx, b.rpc.L1, tx); err != nil {
		return err
	}

	log.Info("📝 Bid for batch tx succeeded", "txHash", tx.Hash(), "batchID", batchID)

	return nil
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
