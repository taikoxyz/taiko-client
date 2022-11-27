package chainsyncer

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/beacon"
	"github.com/ethereum/go-ethereum/log"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
	eventiterator "github.com/taikoxyz/taiko-client/pkg/chain_iterator/event_iterator"
)

var (
	checkBeaconSyncProgressInterval = 120 * time.Second
)

// TriggerBeaconSync triggers the L2 node to start performing a beacon-sync.
func (s *L2ChainSyncer) TriggerBeaconSync() error {
	lastVerifiedHead, err := s.getVerifiedBlockPayload(s.ctx)
	if err != nil {
		return err
	}

	headBeofreTriggerSync, err := s.rpc.L2.HeaderByNumber(s.ctx, nil)
	if err != nil {
		return err
	}

	status, err := s.rpc.L2Engine.NewPayload(
		context.Background(),
		lastVerifiedHead,
	)
	if err != nil {
		return err
	}
	if status != &beacon.STATUS_SYNCING.PayloadStatus {
		return fmt.Errorf("invalid new payload response status: %s", status.Status)
	}

	fcRes, err := s.rpc.L2Engine.ForkchoiceUpdate(s.ctx, &beacon.ForkchoiceStateV1{
		HeadBlockHash:      lastVerifiedHead.BlockHash,
		SafeBlockHash:      lastVerifiedHead.BlockHash,
		FinalizedBlockHash: lastVerifiedHead.BlockHash,
	}, nil)
	if err != nil {
		return err
	}
	if fcRes.PayloadStatus != beacon.STATUS_SYNCING.PayloadStatus {
		return fmt.Errorf("invalid new forkchoiceUpdate response status: %s", status.Status)
	}

	s.headBeforeLastBeaconSync = headBeofreTriggerSync
	s.checkBeaconSyncProgressTimer = time.AfterFunc(checkBeaconSyncProgressInterval, s.checkBeaconSyncProgress)

	return nil
}

// getVerifiedBlockPayload fetches the latest verfied block's header, and converts it to an Engine API executable data,
// which will be used to let the node to start beacon-syncing.
func (s *L2ChainSyncer) getVerifiedBlockPayload(ctx context.Context) (*beacon.ExecutableDataV1, error) {
	var (
		proveBlockTxHash  common.Hash
		lastVerifiedBlock = s.state.GetLastVerifiedBlock()
	)

	iter, err := eventiterator.NewBlockProvenIterator(s.ctx, &eventiterator.BlockProvenIteratorConfig{
		Client:      s.rpc.L1,
		TaikoL1:     s.rpc.TaikoL1,
		StartHeight: common.Big0,   // TODO: change this number
		EndHeight:   common.Big256, // TODO: change this number
		FilterQuery: []*big.Int{lastVerifiedBlock.ID},
		OnBlockProvenEvent: func(
			ctx context.Context,
			e *bindings.TaikoL1ClientBlockProven,
			endIter eventiterator.EndBlockProvenEventIterFunc,
		) error {
			if bytes.Compare(e.BlockHash[:], lastVerifiedBlock.Hash.Bytes()) == 0 {
				proveBlockTxHash = e.Raw.TxHash
				endIter()
			}
			return nil
		},
	})

	if err != nil {
		return nil, err
	}

	if err := iter.Iter(); err != nil {
		return nil, err
	}

	if proveBlockTxHash == (common.Hash{}) {
		return nil, fmt.Errorf(
			"failed to find L1 height of last verified block's ProveBlock transaction, id: %s",
			lastVerifiedBlock.ID,
		)
	}

	proveBlockTx, _, err := s.rpc.L1.TransactionByHash(s.ctx, proveBlockTxHash)
	if err != nil {
		return nil, err
	}

	evidence, err := encoding.UnpackEvidence(proveBlockTx.Data())
	if err != nil {
		return nil, err
	}

	header := encoding.ToGethHeader(&evidence.Header)

	if header.Hash() != lastVerifiedBlock.Hash {
		return nil, fmt.Errorf("last verified block hash mismatch: %s != %s", header.Hash(), lastVerifiedBlock.Hash)
	}

	return encoding.ToExecutableDataV1(header), nil
}

func (s *L2ChainSyncer) checkBeaconSyncProgress() {
	if s.headBeforeLastBeaconSync == nil {
		log.Warn("Head before last beacon-sync not recorded")
		return
	}

	syncProgress, err := s.rpc.L2.SyncProgress(s.ctx)
	if err != nil {
		log.Error("Get chain sync progress error", "error", err)
		return
	}

	if syncProgress == nil {
		log.Info("Beacon-sync finished")
		return
	}
}
