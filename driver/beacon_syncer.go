package driver

import (
	"bytes"
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/beacon"
	"github.com/ethereum/go-ethereum/log"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
	eventIterator "github.com/taikoxyz/taiko-client/pkg/chain_iterator/event_iterator"
)

// TriggerBeaconSync triggers the L2 node to start performing a beacon-sync.
func (s *L2ChainSyncer) TriggerBeaconSync() error {
	blockID, lastVerifiedHeadPayload, err := s.getVerifiedBlockPayload(s.ctx)
	if err != nil {
		return err
	}

	if s.beaconSyncTriggered && s.lastSyncedVerifiedBlockID != nil && s.lastSyncedVerifiedBlockID.Cmp(blockID) == 0 {
		log.Debug("Verified head not updated", "blockID", blockID, "hash", lastVerifiedHeadPayload.BlockHash)
		return nil
	}

	status, err := s.rpc.L2Engine.NewPayload(
		s.ctx,
		lastVerifiedHeadPayload,
	)
	if err != nil {
		return err
	}
	if status.Status != beacon.SYNCING {
		return fmt.Errorf("unexpected NewPayload response status: %s", status.Status)
	}

	fcRes, err := s.rpc.L2Engine.ForkchoiceUpdate(s.ctx, &beacon.ForkchoiceStateV1{
		HeadBlockHash:      lastVerifiedHeadPayload.BlockHash,
		SafeBlockHash:      lastVerifiedHeadPayload.BlockHash,
		FinalizedBlockHash: lastVerifiedHeadPayload.BlockHash,
	}, nil)
	if err != nil {
		return err
	}
	if fcRes.PayloadStatus.Status != beacon.SYNCING {
		return fmt.Errorf("unexpected ForkchoiceUpdate response status: %s", status.Status)
	}

	// Update sync status.
	s.beaconSyncTriggered = true
	s.lastSyncedVerifiedBlockHash = lastVerifiedHeadPayload.BlockHash
	s.lastSyncedVerifiedBlockID = blockID

	log.Info(
		"⛓️ Beacon-sync triggered",
		"newHeadID", blockID,
		"newHeadHeight", lastVerifiedHeadPayload.Number,
		"newHeadHash", s.lastSyncedVerifiedBlockHash,
	)

	return nil
}

// getVerifiedBlockPayload fetches the latest verified block's header, and converts it to an Engine API executable data,
// which will be used to let the node to start beacon-syncing.
func (s *L2ChainSyncer) getVerifiedBlockPayload(ctx context.Context) (*big.Int, *beacon.ExecutableDataV1, error) {
	var (
		proveBlockTxHash  common.Hash
		lastVerifiedBlock = s.state.getLastVerifiedBlock()
	)

	iter, err := eventIterator.NewBlockProvenIterator(s.ctx, &eventIterator.BlockProvenIteratorConfig{
		Client:      s.rpc.L1,
		TaikoL1:     s.rpc.TaikoL1,
		StartHeight: s.state.genesisL1Height,
		EndHeight:   s.state.GetL1Head().Number,
		FilterQuery: []*big.Int{lastVerifiedBlock.ID},
		Reverse:     true,
		OnBlockProvenEvent: func(
			ctx context.Context,
			e *bindings.TaikoL1ClientBlockProven,
			endIter eventIterator.EndBlockProvenEventIterFunc,
		) error {
			if bytes.Equal(e.BlockHash[:], lastVerifiedBlock.Hash.Bytes()) {
				log.Info(
					"Last verified block's BlockProven event found",
					"height", e.Raw.BlockNumber,
					"txHash", e.Raw.TxHash,
				)
				proveBlockTxHash = e.Raw.TxHash
				endIter()
			}
			return nil
		},
	})

	if err != nil {
		return nil, nil, err
	}

	if err := iter.Iter(); err != nil {
		return nil, nil, err
	}

	if proveBlockTxHash == (common.Hash{}) {
		return nil, nil, fmt.Errorf(
			"failed to find L1 height of last verified block's ProveBlock transaction, id: %s",
			lastVerifiedBlock.ID,
		)
	}

	proveBlockTx, _, err := s.rpc.L1.TransactionByHash(s.ctx, proveBlockTxHash)
	if err != nil {
		return nil, nil, err
	}

	evidenceHeader, err := encoding.UnpackEvidenceHeader(proveBlockTx.Data())
	if err != nil {
		return nil, nil, err
	}

	header := encoding.ToGethHeader(evidenceHeader)

	if header.Hash() != lastVerifiedBlock.Hash {
		return nil, nil, fmt.Errorf("last verified block hash mismatch: %s != %s", header.Hash(), lastVerifiedBlock.Hash)
	}

	log.Info("Last verified block header retrieved", "hash", header.Hash())

	return lastVerifiedBlock.ID, encoding.ToExecutableDataV1(header), nil
}
