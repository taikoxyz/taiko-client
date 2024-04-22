package producer

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"golang.org/x/sync/errgroup"

	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
	"github.com/taikoxyz/taiko-client/internal/metrics"
)

// R0ProofParam represents the JSON body of SGXRequestProofBodyParam's `risc0` field.
type R0ProofParam struct {
	Bonsai       bool     `json:"bonsai"`
	Snark        bool     `json:"snark"`
	Profile      bool     `json:"profile"`
	ExecutionPo2 *big.Int `json:"execution_po2"`
}

// SGXAndZKProofProducer generates an SGX + ZK proof for the given block.
type SGXAndZKProofProducer struct {
	RaikoHostEndpoint string // a proverd RPC endpoint
	L1Endpoint        string // a L1 node RPC endpoint
	L1BeaconEndpoint  string // a L1 beacon node RPC endpoint
	L2Endpoint        string // a L2 execution engine's RPC endpoint
	Dummy             bool
	SGXProducer       *SGXProofProducer
	DummyProofProducer
}

// RequestProof implements the ProofProducer interface.
func (o *SGXAndZKProofProducer) RequestProof(
	ctx context.Context,
	opts *ProofRequestOptions,
	blockID *big.Int,
	meta *bindings.TaikoDataBlockMetadata,
	header *types.Header,
) (*ProofWithHeader, error) {
	log.Info(
		"Request SGX+ZK proof",
		"blockID", blockID,
		"coinbase", meta.Coinbase,
		"height", header.Number,
		"hash", header.Hash(),
	)

	proofs := make([][]byte, 2)
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		res, err := o.SGXProducer.RequestProof(ctx, opts, blockID, meta, header)
		if err == nil {
			proofs[0] = res.Proof
		}
		return err
	})
	g.Go(func() error {
		res, err := o.requestZKProof(ctx, opts, blockID, meta, header)
		if err == nil {
			proofs[1] = res.Proof
		}
		return err
	})
	if err := g.Wait(); err != nil {
		return nil, err
	}

	return &ProofWithHeader{
		BlockID: blockID,
		Meta:    meta,
		Header:  header,
		Proof:   append(proofs[0], proofs[1]...),
		Opts:    opts,
		Tier:    o.Tier(),
	}, nil
}

func (o *SGXAndZKProofProducer) requestZKProof(
	ctx context.Context,
	opts *ProofRequestOptions,
	blockID *big.Int,
	meta *bindings.TaikoDataBlockMetadata,
	header *types.Header,
) (*ProofWithHeader, error) {
	log.Info(
		"Request ZK proof from raiko-host service",
		"blockID", blockID,
		"coinbase", meta.Coinbase,
		"height", header.Number,
		"hash", header.Hash(),
	)

	if o.Dummy {
		return o.DummyProofProducer.RequestProof(opts, blockID, meta, header, o.Tier())
	}

	proof, err := o.callProverDaemon(ctx, opts)
	if err != nil {
		return nil, err
	}

	metrics.ProverZkProofGeneratedCounter.Add(1)

	return &ProofWithHeader{
		BlockID: blockID,
		Header:  header,
		Meta:    meta,
		Proof:   proof,
		Opts:    opts,
		Tier:    o.Tier(),
	}, nil
}

// callProverDaemon keeps polling the proverd service to get the requested proof.
func (o *SGXAndZKProofProducer) callProverDaemon(ctx context.Context, opts *ProofRequestOptions) ([]byte, error) {
	var (
		proof []byte
		start = time.Now()
	)
	if err := backoff.Retry(func() error {
		if ctx.Err() != nil {
			return nil
		}
		output, err := o.requestProof(opts)
		if err != nil {
			log.Error("Failed to request proof", "height", opts.BlockID, "error", err, "endpoint", o.RaikoHostEndpoint)
			return err
		}

		if output == nil {
			log.Info(
				"ZK proof generating",
				"height", opts.BlockID,
				"time", time.Since(start),
				"producer", "SGXAndZKProofProducer",
			)
			return errProofGenerating
		}

		log.Debug("ZK proof generation output", "output", output)

		proof = common.Hex2Bytes(output.Proof[2:])
		log.Info(
			"ZK proof generated",
			"height", opts.BlockID,
			"time", time.Since(start),
			"producer", "SGXAndZKProofProducer",
		)
		return nil
	}, backoff.WithContext(backoff.NewConstantBackOff(proofPollingInterval), ctx)); err != nil {
		return nil, err
	}

	return proof, nil
}

// requestProof sends an RPC request to proverd to try to get the requested proof.
func (o *SGXAndZKProofProducer) requestProof(opts *ProofRequestOptions) (*RaikoHostOutput, error) {
	reqBody := RaikoRequestProofBody{
		JsonRPC: "2.0",
		ID:      common.Big1,
		Method:  "proof",
		Params: []*RaikoRequestProofBodyParam{{
			Type:        "risc0",
			Block:       opts.BlockID,
			L2RPC:       o.L2Endpoint,
			L1RPC:       o.L1Endpoint,
			L1BeaconRPC: o.L1BeaconEndpoint,
			Prover:      opts.ProverAddress.Hex()[2:],
			Graffiti:    opts.Graffiti,
			R0ProofParam: &R0ProofParam{
				Bonsai:       true,
				Snark:        true,
				Profile:      false,
				ExecutionPo2: new(big.Int).SetUint64(20),
			},
		}},
	}

	jsonValue, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	res, err := http.Post(o.RaikoHostEndpoint, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to request proof, id: %d, statusCode: %d", opts.BlockID, res.StatusCode)
	}

	resBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var output RaikoRequestProofBodyResponse
	if err := json.Unmarshal(resBytes, &output); err != nil {
		return nil, err
	}

	if output.Error != nil {
		return nil, errors.New(output.Error.Message)
	}

	return output.Result, nil
}

// Tier implements the ProofProducer interface.
func (o *SGXAndZKProofProducer) Tier() uint16 {
	return encoding.TierSgxAndZkVMID
}
