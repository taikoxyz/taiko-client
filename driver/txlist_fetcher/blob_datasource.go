package txlistdecoder

import (
	"context"
	"fmt"
	"net/url"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/go-resty/resty/v2"
	"github.com/prysmaticlabs/prysm/v4/beacon-chain/rpc/eth/blob"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
)

type BlobDataSource struct {
	ctx                context.Context
	rpc                *rpc.Client
	blobServerEndpoint *url.URL
}

type BlobData struct {
	// TODO: wait /getBlob add column
	BlobHash      string `json:"blob_hash"`
	KzgCommitment string `json:"kzg_commitment"`
}

type BlobDataSeq struct {
	Data []*BlobData `json:"data"`
}

func NewBlobDataSource(
	ctx context.Context,
	rpc *rpc.Client,
	blobServerEndpoint *url.URL,
) *BlobDataSource {
	return &BlobDataSource{
		ctx:                ctx,
		rpc:                rpc,
		blobServerEndpoint: blobServerEndpoint,
	}
}

// GetBlobs get blob sidecar by meta
func (ds *BlobDataSource) GetBlobs(
	ctx context.Context,
	meta *bindings.TaikoDataBlockMetadata,
) ([]*blob.Sidecar, error) {
	if !meta.BlobUsed {
		return nil, errBlobUnused
	}

	sidecars, err := ds.rpc.L1Beacon.GetBlobs(ctx, meta.Timestamp)
	if err != nil {
		log.Info("Failed to get blobs from beacon, try to use blob server.", "err", err.Error())
		if ds.blobServerEndpoint == nil {
			log.Info("No blob server endpoint set")
			return nil, err
		}
		blobs, err := ds.getBlobFromServer(ctx, common.Bytes2Hex(meta.BlobHash[:]))
		if err != nil {
			return nil, err
		}
		for index, value := range blobs.Data {
			sidecars[index] = &blob.Sidecar{
				// TODO: wait /getBlob add column
				KzgCommitment: value.KzgCommitment,
			}
		}
	}
	return sidecars, err
}

// getBlobFromServer get blob data from server path `/getBlob`.
func (ds *BlobDataSource) getBlobFromServer(ctx context.Context, blobHash string) (*BlobDataSeq, error) {
	var (
		route  = "/getBlob"
		param  = map[string]string{"blobHash": blobHash}
		result = &BlobDataSeq{}
	)
	err := ds.get(ctx, route, param, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// get send the given GET request to the blob server.
func (ds *BlobDataSource) get(ctx context.Context, route string, param map[string]string, result interface{}) error {
	resp, err := resty.New().R().
		SetResult(result).
		SetQueryParams(param).
		SetContext(ctx).
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		Get(fmt.Sprintf("%v/%v", ds.blobServerEndpoint.String(), route))
	if err != nil {
		return err
	}

	if !resp.IsSuccess() {
		return fmt.Errorf(
			"unable to contect blob server endpoint, status code: %v",
			resp.StatusCode(),
		)
	}

	return nil
}
