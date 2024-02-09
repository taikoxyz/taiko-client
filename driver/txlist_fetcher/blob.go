package txlistdecoder

import (
	"context"
	"crypto/sha256"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto/kzg4844"
	"github.com/ethereum/go-ethereum/log"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
)

const (
	blobCommitmentVersionKZG uint8 = 0x01 // Version byte for the point evaluation precompile.
)

type BlobFetcher struct {
	rpc *rpc.Client
}

func NewBlobTxListFetcher(rpc *rpc.Client) *BlobFetcher {
	return &BlobFetcher{rpc}
}

func (d *BlobFetcher) Fetch(
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

	log.Info("Fetch sidecars", "sidecars", sidecars)

	for _, sidecar := range sidecars {
		log.Info("Found sidecar", "KzgCommitment", sidecar.KzgCommitment, "blobHash", common.Bytes2Hex(meta.BlobHash[:]))

		if kZGToVersionedHash(
			kzg4844.Commitment(common.Hex2Bytes(sidecar.KzgCommitment)[:]),
		) == common.BytesToHash(meta.BlobHash[:]) {
			return common.Hex2Bytes(sidecar.Blob), nil
		}
	}

	return nil, errSidecarNotFound
}

// kZGToVersionedHash implements kzg_to_versioned_hash from EIP-4844
func kZGToVersionedHash(kzg kzg4844.Commitment) common.Hash {
	h := sha256.Sum256(kzg[:])
	h[0] = blobCommitmentVersionKZG

	return h
}
