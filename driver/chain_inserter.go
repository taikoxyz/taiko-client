package driver

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/beacon"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/taikochain/taiko-client/bindings"
	"github.com/taikochain/taiko-client/bindings/encoding"
	"github.com/taikochain/taiko-client/metrics"
	"github.com/taikochain/taiko-client/pkg/rpc"
)

const (
	MaxL1BlocksRead = 1000
)

// InvalidTxListReason represents a reason why a transactions list is invalid,
// must match the definitions in LibInvalidTxList.sol:
//
//	enum Reason {
//		OK,
//		BINARY_TOO_LARGE,
//		BINARY_NOT_DECODABLE,
//		BLOCK_TOO_MANY_TXS,
//		BLOCK_GAS_LIMIT_TOO_LARGE,
//		TX_INVALID_SIG,
//		TX_GAS_LIMIT_TOO_SMALL
//	}
type InvalidTxListReason uint8

// All invalid transactions list reasons.
const (
	HintOK InvalidTxListReason = iota
	HintBinaryTooLarge
	HintBinaryNotDecodable
	HintBlockTooManyTxs
	HintBlockGasLimitTooLarge
	HintTxInvalidSig
	HintTxGasLimitTooSmall
)

type L2ChainInserter struct {
	state                         *State            // Driver's state
	rpc                           *rpc.Client       // L1/L2 RPC clients
	throwawayBlocksBuilderPrivKey *ecdsa.PrivateKey // Private key of L2 throwaway blocks builder
}

// NewL2ChainInserter creates a new block inserter instance.
func NewL2ChainInserter(
	ctx context.Context,
	rpc *rpc.Client,
	state *State,
	throwawayBlocksBuilderPrivKey *ecdsa.PrivateKey,
) (*L2ChainInserter, error) {
	return &L2ChainInserter{
		rpc:                           rpc,
		state:                         state,
		throwawayBlocksBuilderPrivKey: throwawayBlocksBuilderPrivKey,
	}, nil
}

// ProcessL1Blocks fetches all `TaikoL1.BlockProposed` events between given
// L1 block heights, and then tries inserting them into L2 node's block chain.
func (b *L2ChainInserter) ProcessL1Blocks(ctx context.Context, l1End *types.Header) error {
	l1Start, err := b.state.ConfirmL1Current(ctx)
	if err != nil {
		return err
	}

	for l1Start.Number.Uint64() < l1End.Number.Uint64() {
		if l1Start.Number.Uint64()+MaxL1BlocksRead > l1End.Number.Uint64() {
			return b.processL1Blocks(ctx, l1Start, l1End)
		}

		endHeight := new(big.Int).Add(l1Start.Number, big.NewInt(MaxL1BlocksRead))
		currentEndBlock, err := b.rpc.L1.HeaderByNumber(ctx, endHeight)
		if err != nil {
			return fmt.Errorf("fetch L1 header by number (%d) error: %w", endHeight, err)
		}

		if err := b.processL1Blocks(ctx, l1Start, currentEndBlock); err != nil {
			return fmt.Errorf("process L1 blocks error: %w", err)
		}

		l1Start = currentEndBlock
	}

	return nil
}

func (b *L2ChainInserter) processL1Blocks(ctx context.Context, l1Start *types.Header, l1End *types.Header) error {
	log.Info(
		"New synchronising operation",
		"l1StartHeight", l1Start.Number,
		"l1StartHash", l1Start.Hash(),
		"l1End", l1End.Number,
		"l1EndHash", l1End.Hash(),
	)

	end := l1End.Number.Uint64()
	iter, err := b.rpc.TaikoL1.FilterBlockProposed(&bind.FilterOpts{
		Start: l1Start.Number.Uint64(),
		End:   &end,
	}, nil)
	if err != nil {
		return err
	}

	for iter.Next() {
		if ctx.Err() != nil {
			return nil
		}

		event := iter.Event

		// Since we are not using eth_subscribe, this should not happen,
		// only check for safety.
		if event.Raw.Removed {
			continue
		}

		// No need to insert genesis again, its already in L2 block chain.
		if event.Id.Cmp(common.Big0) == 0 {
			continue
		}

		// Fetch the L2 parent block.
		parent, err := b.rpc.L2ParentByBlockId(ctx, event.Id)
		if err != nil {
			return fmt.Errorf("failed to fetch L2 parent block: %w", err)
		}

		log.Debug("Parent block", "height", parent.Number, "hash", parent.Hash())

		tx, err := b.rpc.L1.TransactionInBlock(
			ctx,
			event.Raw.BlockHash,
			event.Raw.TxIndex,
		)
		if err != nil {
			return fmt.Errorf("failed to fetch original TaikoL1.proposeBlock transaction: %w", err)
		}

		txListBytes, err := encoding.UnpackTxListBytes(tx.Data())
		if err != nil {
			log.Warn("Unpack transactions bytes error", "error", err)
			continue
		}

		// Check whether the transactions list is valid.
		hint, invalidTxIndex := b.isTxListValid(txListBytes)

		log.Info(
			"Validate transactions list",
			"blockID", event.Id,
			"hint", hint,
			"invalidTxIndex", invalidTxIndex,
		)

		l1Origin := &rawdb.L1Origin{
			BlockID:       event.Id,
			L2BlockHash:   common.Hash{}, // Will be set by taiko-geth.
			L1BlockHeight: new(big.Int).SetUint64(event.Raw.BlockNumber),
			L1BlockHash:   event.Raw.BlockHash,
			Throwaway:     hint != HintOK,
		}

		var (
			payloadData  *beacon.ExecutableDataV1
			rpcError     error
			payloadError error
		)
		if hint == HintOK {
			payloadData, rpcError, payloadError = b.insertNewHead(
				ctx,
				event,
				parent,
				b.state.getHeadBlockID(),
				txListBytes,
				l1Origin,
			)
		} else {
			payloadData, rpcError, payloadError = b.insertThrowAwayBlock(
				ctx,
				event,
				parent,
				uint8(hint),
				new(big.Int).SetInt64(int64(invalidTxIndex)),
				b.state.getHeadBlockID(),
				txListBytes,
				l1Origin,
			)
		}

		// RPC errors are recoverable.
		if rpcError != nil {
			return fmt.Errorf("failed to insert new head to L2 node: %w", rpcError)
		}

		if payloadError != nil {
			log.Warn("Ignore invalid block context", "blockID", event.Id, "payloadError", payloadError)
			continue
		}

		log.Debug("Payload data", "payload", payloadData)

		if b.state.l1Current, err = b.rpc.L1.HeaderByHash(ctx, event.Raw.BlockHash); err != nil {
			return fmt.Errorf("failed to update L1 current sync cursor: %w", err)
		}

		metrics.DriverL1CurrentHeightGauge.Update(b.state.l1Current.Number.Int64())

		log.Info(
			"🔗 New L2 block inserted",
			"throwaway", l1Origin.Throwaway,
			"blockID", event.Id,
			"height", payloadData.Number,
			"hash", payloadData.BlockHash,
			"lastFinalizedBlockHash", b.state.getLastFinalizedBlockHash(),
			"transactions", len(payloadData.Transactions),
		)
	}

	if b.state.l1Current, err = b.rpc.L1.HeaderByHash(ctx, l1End.Hash()); err != nil {
		return fmt.Errorf("failed to update L1 current sync cursor: %w", err)
	}

	metrics.DriverL1CurrentHeightGauge.Update(b.state.l1Current.Number.Int64())

	return nil
}

// insertNewHead tries to insert a new head block to the L2 node's local
// block chain through Engine APIs.
func (b *L2ChainInserter) insertNewHead(
	ctx context.Context,
	event *bindings.TaikoL1ClientBlockProposed,
	parent *types.Header,
	headBlockID *big.Int,
	txListBytes []byte,
	l1Origin *rawdb.L1Origin,
) (payloadData *beacon.ExecutableDataV1, rpcError error, payloadError error) {
	log.Debug(
		"Try to insert a new L2 head block",
		"parentNumber", parent.Number,
		"parentHash", parent.Hash(),
		"headBlockID", headBlockID,
		"l1Origin", l1Origin,
	)

	// Insert a TaikoL2.anchor transaction at transactions list head
	var txList []*types.Transaction
	if err := rlp.DecodeBytes(txListBytes, &txList); err != nil {
		log.Info("Ignore invalid txList bytes", "blockID", event.Id)
		return nil, nil, err
	}

	// Assemble a TaikoL2.anchor transaction
	anchorTx, err := b.assembleAnchorTx(
		ctx,
		event.Meta.L1Height,
		event.Meta.L1Hash,
		parent.Number,
	)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create V1TaikoL2.anchor transaction: %w", err)
	}

	txList = append([]*types.Transaction{anchorTx}, txList...)

	if txListBytes, err = rlp.EncodeToBytes(txList); err != nil {
		log.Warn("Encode txList error", "blockID", event.Id, "error", err)
		return nil, nil, err
	}

	payload, rpcErr, payloadErr := b.createExecutionPayloads(
		ctx,
		event,
		parent.Hash(),
		l1Origin,
		headBlockID,
		txListBytes,
	)

	if rpcErr != nil || payloadErr != nil {
		return nil, rpcError, payloadErr
	}

	fc := &beacon.ForkchoiceStateV1{HeadBlockHash: parent.Hash()}

	// Update the fork choice
	fc.HeadBlockHash = payload.BlockHash
	fc.SafeBlockHash = payload.BlockHash
	fcRes, err := b.rpc.L2Engine.ForkchoiceUpdate(ctx, fc, nil)
	if err != nil {
		return nil, err, nil
	}
	if fcRes.PayloadStatus.Status != beacon.VALID {
		return nil, nil, fmt.Errorf("failed to update forkchoice, status: %s", fcRes.PayloadStatus.Status)
	}

	return payload, nil, nil
}

// insertNewHead tries to insert a throw away block to the L2 node's local
// block chain through Engine APIs.
func (b *L2ChainInserter) insertThrowAwayBlock(
	ctx context.Context,
	event *bindings.TaikoL1ClientBlockProposed,
	parent *types.Header,
	hint uint8,
	invalidTxIndex *big.Int,
	headBlockID *big.Int,
	txListBytes []byte,
	l1Origin *rawdb.L1Origin,
) (payloadData *beacon.ExecutableDataV1, rpcError error, payloadError error) {
	log.Info(
		"Try to insert a new L2 throwaway block",
		"parentHash", parent.Hash(),
		"headBlockID", headBlockID,
		"l1Origin", l1Origin,
	)

	// Assemble a TaikoL2.invalidateBlock transaction
	opts, err := b.getInvalidateBlockTxOpts(ctx, parent.Number)
	if err != nil {
		return nil, nil, err
	}

	invalidateBlockTx, err := b.rpc.TaikoL2.InvalidateBlock(
		opts,
		txListBytes,
		hint,
		invalidTxIndex,
	)
	if err != nil {
		return nil, nil, err
	}

	throwawayBlockTxListBytes, err := rlp.EncodeToBytes(
		types.Transactions{invalidateBlockTx},
	)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to encode TaikoL2.InvalidateBlock transaction bytes, err: %w", err)
	}

	return b.createExecutionPayloads(
		ctx,
		event,
		parent.Hash(),
		l1Origin,
		headBlockID,
		throwawayBlockTxListBytes,
	)
}

// createExecutionPayloads creates a new execution payloads through
// Engine APIs.
func (b *L2ChainInserter) createExecutionPayloads(
	ctx context.Context,
	event *bindings.TaikoL1ClientBlockProposed,
	parentHash common.Hash,
	l1Origin *rawdb.L1Origin,
	headBlockID *big.Int,
	txListBytes []byte,
) (payloadData *beacon.ExecutableDataV1, rpcError error, payloadError error) {
	fc := &beacon.ForkchoiceStateV1{HeadBlockHash: parentHash}
	attributes := &beacon.PayloadAttributesV1{
		Timestamp:             event.Meta.Timestamp,
		Random:                event.Meta.MixHash,
		SuggestedFeeRecipient: event.Meta.Beneficiary,
		BlockMetadata: &beacon.BlockMetadata{
			HighestBlockID: headBlockID,
			Beneficiary:    event.Meta.Beneficiary,
			GasLimit:       event.Meta.GasLimit + b.state.anchorTxGasLimit.Uint64(),
			Timestamp:      event.Meta.Timestamp,
			TxList:         txListBytes,
			MixHash:        event.Meta.MixHash,
			ExtraData:      event.Meta.ExtraData,
		},
		L1Origin: l1Origin,
	}

	// TODO: handle payload error more precisely
	// Step 1, prepare a payload
	fcRes, err := b.rpc.L2Engine.ForkchoiceUpdate(ctx, fc, attributes)
	if err != nil {
		return nil, err, nil
	}
	if fcRes.PayloadStatus.Status != beacon.VALID {
		return nil, nil, fmt.Errorf("failed to update forkchoice, status: %s", fcRes.PayloadStatus.Status)
	}
	if fcRes.PayloadID == nil {
		return nil, nil, errors.New("empty payload ID")
	}

	// Step 2, get the payload
	payload, err := b.rpc.L2Engine.GetPayload(ctx, fcRes.PayloadID)
	if err != nil {
		return nil, err, nil
	}

	// Step 3, execute the payload
	execStatus, err := b.rpc.L2Engine.NewPayload(ctx, payload)
	if err != nil {
		return nil, err, nil
	}
	if execStatus.Status != beacon.VALID {
		return nil, nil, fmt.Errorf("failed to execute the newly built block, status: %s", execStatus.Status)
	}

	return payload, nil, nil
}

// isTxListValid checks whether the transaction list is valid, must match
// the validation rule defined in LibInvalidTxList.sol.
// ref: https://github.com/taikochain/taiko-mono/blob/main/packages/protocol/contracts/libs/LibInvalidTxList.sol
func (b *L2ChainInserter) isTxListValid(txListBytes []byte) (hint InvalidTxListReason, txIdx int) {
	if len(txListBytes) > int(b.state.maxTxlistBytes.Uint64()) {
		log.Warn("Transactions list binary too large", "length", len(txListBytes))
		return HintBinaryTooLarge, 0
	}

	var txs types.Transactions
	if err := rlp.DecodeBytes(txListBytes, &txs); err != nil {
		log.Debug("Failed to decode transactions list bytes", "error", err)
		return HintBinaryNotDecodable, 0
	}

	log.Debug("Transactions list decoded", "length", len(txs))

	if txs.Len() > int(b.state.maxBlockNumTxs.Uint64()) {
		log.Debug("Too many transactions", "count", txs.Len())
		return HintBlockTooManyTxs, 0
	}

	sumGasLimit := uint64(0)
	for _, tx := range txs {
		sumGasLimit += tx.Gas()
	}

	if sumGasLimit > b.state.maxBlocksGasLimit.Uint64() {
		log.Debug("Accumulate gas limit too large", "sumGasLimit", sumGasLimit)
		return HintBlockGasLimitTooLarge, 0
	}

	signer := types.LatestSignerForChainID(b.rpc.L2ChainID)

	for i, tx := range txs {
		sender, err := types.Sender(signer, tx)
		if err != nil || sender == (common.Address{}) {
			log.Debug("Invalid transaction signature", "error", err)
			return HintTxInvalidSig, i
		}

		if tx.Gas() < b.state.minTxGasLimit.Uint64() {
			log.Debug("Transaction gas limit too small", "gasLimit", tx.Gas())
			return HintTxGasLimitTooSmall, i
		}
	}

	return HintOK, 0
}

// getInvalidateBlockTxOpts signs the transaction with a the
// throwaway blocks builder private key.
func (b *L2ChainInserter) getInvalidateBlockTxOpts(ctx context.Context, height *big.Int) (*bind.TransactOpts, error) {
	opts, err := bind.NewKeyedTransactorWithChainID(
		b.throwawayBlocksBuilderPrivKey,
		b.rpc.L2ChainID,
	)
	if err != nil {
		return nil, err
	}

	nonce, err := b.rpc.L2AccountNonce(
		ctx,
		crypto.PubkeyToAddress(b.throwawayBlocksBuilderPrivKey.PublicKey),
		height,
	)
	if err != nil {
		return nil, err
	}

	opts.Nonce = new(big.Int).SetUint64(nonce)
	opts.NoSend = true

	return opts, nil
}
