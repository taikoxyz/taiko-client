package txlistdecoder

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
)

type BlobDecoder struct {
	rpc *rpc.Client
}

func NewBlobDecoder(rpc *rpc.Client) *BlobDecoder {
	return &BlobDecoder{rpc}
}

func (d *BlobDecoder) DecodeTxList(
	ctx context.Context,
	tx *types.Transaction,
	meta *bindings.TaikoDataBlockMetadata,
) ([]byte, error) {
	if !meta.BlobUsed {
		return nil, errBlobUnused
	}

	sidecars, err := d.rpc.GetBlobs(ctx, new(big.Int).SetUint64(meta.L1Height+1))
	if err != nil {
		return nil, err
	}

	for _, sidecar := range sidecars {
		if sidecar.KzgCommitment == common.Bytes2Hex(meta.BlobHash[:]) {
			return common.Hex2Bytes(sidecar.Blob), nil
		}
	}

	return nil, errSidecarNotFound
}
