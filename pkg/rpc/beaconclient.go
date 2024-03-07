package rpc

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto/kzg4844"
	"github.com/prysmaticlabs/prysm/v4/api/client"
	"github.com/prysmaticlabs/prysm/v4/api/client/beacon"
	"github.com/prysmaticlabs/prysm/v4/beacon-chain/rpc/eth/blob"
)

var (
	// Request urls.
	sidecarsRequestURL = "eth/v1/beacon/blob_sidecars/%d"
)

type BeaconClient struct {
	*beacon.Client
}

// NewBeaconClient returns a new beacon client.
func NewBeaconClient(endpoint string, timeout time.Duration) (*BeaconClient, error) {
	cli, err := beacon.NewClient(endpoint, client.WithTimeout(timeout))
	if err != nil {
		return nil, err
	}
	return &BeaconClient{cli}, nil
}

// GetBlobs returns the sidecars for a given slot.
func (c *BeaconClient) GetBlobs(ctx context.Context, slot *big.Int) ([]*blob.Sidecar, error) {
	var sidecars *blob.SidecarsResponse
	resBytes, err := c.Get(ctx, fmt.Sprintf(sidecarsRequestURL, slot))
	if err != nil {
		return nil, err
	}

	return sidecars.Data, json.Unmarshal(resBytes, &sidecars)
}

// GetBlobByHash returns the sidecars for a given slot.
func (c *BeaconClient) GetBlobByHash(ctx context.Context, slot *big.Int, blobHash common.Hash) ([]byte, error) {
	sidecars, err := c.GetBlobs(ctx, slot)
	if err != nil {
		return nil, err
	}

	for _, sidecar := range sidecars {
		commitment := kzg4844.Commitment(common.FromHex(sidecar.KzgCommitment))
		if kzg4844.CalcBlobHashV1(
			sha256.New(),
			&commitment,
		) == blobHash {
			return DecodeBlob(common.FromHex(sidecar.Blob))
		}
	}

	return nil, errors.New("sidecar not found")
}
