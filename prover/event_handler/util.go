package handler

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
)

var (
	errTierNotFound = errors.New("tier not found")
)

// isBlockVerified checks whether the given L2 block has been verified.
func isBlockVerified(ctx context.Context, rpc *rpc.Client, id *big.Int) (bool, error) {
	stateVars, err := rpc.GetProtocolStateVariables(&bind.CallOpts{Context: ctx})
	if err != nil {
		return false, err
	}

	return id.Uint64() <= stateVars.B.LastVerifiedBlockId, nil
}

// isValidProof checks if the given proof is a valid one, comparing to current L2 node canonical chain.
func isValidProof(
	ctx context.Context,
	rpc *rpc.Client,
	blockID *big.Int,
	parentHash common.Hash,
	blockHash common.Hash,
	stateRoot common.Hash,
) (bool, error) {
	parent, err := rpc.L2ParentByBlockID(ctx, blockID)
	if err != nil {
		return false, err
	}

	l2Header, err := rpc.L2.HeaderByNumber(ctx, blockID)
	if err != nil {
		return false, err
	}

	l1Origin, err := rpc.L2.L1OriginByID(ctx, blockID)
	if err != nil {
		return false, err
	}

	l1Header, err := rpc.L1.HeaderByNumber(ctx, new(big.Int).Sub(l1Origin.L1BlockHeight, common.Big1))
	if err != nil {
		return false, err
	}

	return parent.Hash() == parentHash &&
		l2Header.Hash() == blockHash &&
		l1Header.Root == stateRoot, nil
}

// getProvingWindow returns the provingWindow of the given proposed block.
func getProvingWindow(
	e *bindings.TaikoL1ClientBlockProposed,
	tiers []*rpc.TierProviderTierWithID,
) (time.Duration, error) {
	for _, t := range tiers {
		if e.Meta.MinTier == t.ID {
			return time.Duration(t.ProvingWindow) * time.Minute, nil
		}
	}

	return 0, errTierNotFound
}

// isProvingWindowExpired returns true if the assigned prover proving window of
// the given proposed block is expired.
func isProvingWindowExpired(
	e *bindings.TaikoL1ClientBlockProposed,
	tiers []*rpc.TierProviderTierWithID,
) (bool, time.Duration, error) {
	provingWindow, err := getProvingWindow(e, tiers)
	if err != nil {
		return false, 0, fmt.Errorf("failed to get proving window: %w", err)
	}

	var (
		now       = uint64(time.Now().Unix())
		exipredAt = e.Meta.Timestamp + uint64(provingWindow.Seconds())
	)

	return now > exipredAt, time.Duration(exipredAt-now) * time.Second, nil
}
