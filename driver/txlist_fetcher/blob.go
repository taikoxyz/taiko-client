package txlistdecoder

import (
	"context"
	"crypto/sha256"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto/kzg4844"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
)

var (
	blobCommitmentVersionKZG uint8 = 0x01 // Version byte for the point evaluation precompile.
	errBlobInvalid                 = errors.New("invalid blob encoding")
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

	log.Info("Fetch sidecars", "slot", meta.L1Height+1, "sidecars", len(sidecars))

	for i, sidecar := range sidecars {
		log.Info(
			"Block sidecar",
			"index", i,
			"KzgCommitment", sidecar.KzgCommitment,
			"blobHash", common.Bytes2Hex(meta.BlobHash[:]),
		)

		if KZGToVersionedHash(
			kzg4844.Commitment(common.FromHex(sidecar.KzgCommitment)),
		) == common.BytesToHash(meta.BlobHash[:]) {
			return DecodeBlob(common.FromHex(sidecar.Blob))
		}
	}

	return nil, errSidecarNotFound
}

// KZGToVersionedHash implements kzg_to_versioned_hash from EIP-4844.
func KZGToVersionedHash(kzg kzg4844.Commitment) common.Hash {
	h := sha256.Sum256(kzg[:])
	h[0] = blobCommitmentVersionKZG

	return h
}

// DecodeBlob decode blob data.
func DecodeBlob(blob []byte) ([]byte, error) {
	if len(blob) != params.BlobTxFieldElementsPerBlob*32 {
		return nil, errBlobInvalid
	}
	log.Info("OK")
	var data []byte
	for i, j := 0, 0; i < params.BlobTxFieldElementsPerBlob; i++ {
		data = append(data, blob[j:j+31]...)
		j += 32
	}

	i := len(data) - 1
	for ; i >= 0; i-- {
		if data[i] != 0x00 {
			break
		}
	}
	data = data[:i+1]
	return data, nil
}
