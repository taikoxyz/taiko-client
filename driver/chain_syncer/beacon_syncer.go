package chainsyncer

import (
	"bytes"
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/beacon"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
	eventiterator "github.com/taikoxyz/taiko-client/pkg/chain_iterator/event_iterator"
)

func (s *L2ChainSyncer) TriggerBeaconSync() error {
	lastVerifiedHead, err := s.getVerifiedBlockHeader(s.ctx)
	if err != nil {
		return err
	}

	status, err := s.rpc.L2Engine.NewPayload(
		context.Background(),
		beacon.BlockToExecutableData(types.NewBlockWithHeader(lastVerifiedHead)),
	)
	if err != nil {
		return err
	}
	if status != &beacon.STATUS_SYNCING.PayloadStatus {
		return fmt.Errorf("invalid new payload response status: %s", status.Status)
	}

	fcRes, err := s.rpc.L2Engine.ForkchoiceUpdate(s.ctx, &beacon.ForkchoiceStateV1{
		HeadBlockHash:      lastVerifiedHead.Hash(),
		SafeBlockHash:      lastVerifiedHead.Hash(),
		FinalizedBlockHash: lastVerifiedHead.Hash(),
	}, nil)
	if err != nil {
		return err
	}
	if fcRes.PayloadStatus != beacon.STATUS_SYNCING.PayloadStatus {
		return fmt.Errorf("invalid new forkchoiceUpdate response status: %s", status.Status)
	}

	return nil
}

func (s *L2ChainSyncer) getVerifiedBlockHeader(ctx context.Context) (*types.Header, error) {
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
		OnBlockProvenEvent: func(ctx context.Context, e *bindings.TaikoL1ClientBlockProven) error {
			if bytes.Compare(e.BlockHash[:], lastVerifiedBlock.Hash.Bytes()) == 0 {
				proveBlockTxHash = e.Raw.TxHash
				// TODO: stop the iterator
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

	header := &types.Header{
		ParentHash:  evidence.Header.ParentHash,
		UncleHash:   evidence.Header.OmmersHash,
		Coinbase:    evidence.Header.Beneficiary,
		Root:        evidence.Header.StateRoot,
		TxHash:      evidence.Header.TransactionsRoot,
		ReceiptHash: evidence.Header.ReceiptsRoot,
		Bloom:       encoding.BytesToBloom(evidence.Header.LogsBloom),
		Difficulty:  evidence.Header.Difficulty,
		Number:      evidence.Header.Height,
		GasLimit:    evidence.Header.GasLimit,
		GasUsed:     evidence.Header.GasUsed,
		Time:        evidence.Header.Timestamp,
		Extra:       evidence.Header.ExtraData,
		MixDigest:   evidence.Header.MixHash,
		Nonce:       types.EncodeNonce(evidence.Header.Nonce),
	}

	if header.Hash() != lastVerifiedBlock.Hash {
		return nil, fmt.Errorf("last verified block hash mismatch: %s != %s", header.Hash(), lastVerifiedBlock.Hash)
	}

	return header, nil
}
