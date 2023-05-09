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
	"github.com/taikoxyz/taiko-client/bindings"
)

var (
	errProofGenerating = errors.New("proof is generating")
)

// ZkevmRpcdProducer is responsible for requesting zk proofs from the given proverd endpoint.
type ZkevmRpcdProducer struct {
	RpcdEndpoint    string                         // a proverd RPC endpoint
	Param           string                         // parameter file to use
	L1Endpoint      string                         // a L1 node RPC endpoint
	L2Endpoint      string                         // a L2 execution engine's RPC endpoint
	Retry           bool                           // retry proof computation if error
	CustomProofHook func() ([]byte, uint64, error) // only for testing purposes
}

// RequestProofBody represents the JSON body for requesting the proof.
type RequestProofBody struct {
	JsonRPC string                   `json:"jsonrpc"`
	ID      *big.Int                 `json:"id"`
	Method  string                   `json:"method"`
	Params  []*RequestProofBodyParam `json:"params"`
}

// RequestProofBody represents the JSON body of RequestProofBody's `param` field.
type RequestProofBodyParam struct {
	Circuit            string   `json:"circuit"`
	Block              *big.Int `json:"block"`
	L1RPC              string   `json:"l1_rpc"`
	L2RPC              string   `json:"l2_rpc"`
	ProposeBlockTxHash string   `json:"propose_tx_hash"`
	Retry              bool     `json:"retry"`
	Param              string   `json:"param"`
	VerifyProof        bool     `json:"verify_proof"`
	Mock               bool     `json:"mock"`
	Aggregate          bool     `json:"aggregate"`
	Prover             string   `json:"prover"`
}

// RequestProofBodyResponse represents the JSON body of the response of the proof requests.
type RequestProofBodyResponse struct {
	JsonRPC string      `json:"jsonrpc"`
	ID      *big.Int    `json:"id"`
	Result  *RpcdOutput `json:"result"`
}

// RpcdOutput represents the JSON body of RequestProofBodyResponse's `result` field.
type RpcdOutput struct {
	Circuit struct {
		Instances []string `json:"instance"`
		Proof     string   `json:"proof"`
		Degree    uint64   `json:"k"`
	} `json:"circuit"`
}

// NewZkevmRpcdProducer creates a new `ZkevmRpcdProducer` instance.
func NewZkevmRpcdProducer(
	rpcdEndpoint string,
	param string,
	l1Endpoint string,
	l2Endpoint string,
	retry bool,
) (*ZkevmRpcdProducer, error) {
	return &ZkevmRpcdProducer{
		RpcdEndpoint: rpcdEndpoint,
		Param:        param,
		L1Endpoint:   l1Endpoint,
		L2Endpoint:   l2Endpoint,
		Retry:        retry,
	}, nil
}

// RequestProof implements the ProofProducer interface.
func (d *ZkevmRpcdProducer) RequestProof(
	ctx context.Context,
	opts *ProofRequestOptions,
	blockID *big.Int,
	meta *bindings.TaikoDataBlockMetadata,
	header *types.Header,
	resultCh chan *ProofWithHeader,
) error {
	log.Info(
		"Request proof from zkevm-chain proverd service",
		"blockID", blockID,
		"beneficiary", meta.Beneficiary,
		"height", header.Number,
		"hash", header.Hash(),
	)

	var (
		proof  []byte
		degree uint64
		err    error
	)
	if d.CustomProofHook != nil {
		proof, degree, err = d.CustomProofHook()
	} else {
		proof, degree, err = d.callProverDaemon(ctx, opts)
	}
	if err != nil {
		return err
	}

	resultCh <- &ProofWithHeader{
		BlockID: blockID,
		Header:  header,
		Meta:    meta,
		ZkProof: proof,
		Degree:  degree,
	}

	return nil
}

// callProverDaemon keeps polling the proverd service to get the requested proof.
func (d *ZkevmRpcdProducer) callProverDaemon(ctx context.Context, opts *ProofRequestOptions) ([]byte, uint64, error) {
	var (
		proof  []byte
		degree uint64
		start  = time.Now()
	)
	if err := backoff.Retry(func() error {
		if ctx.Err() != nil {
			return nil
		}
		output, err := d.requestProof(opts)
		if err != nil {
			log.Error("Failed to request proof", "height", opts.Height, "err", err, "endpoint", d.RpcdEndpoint)
			return err
		}

		log.Info("Request proof", "height", opts.Height, "output", output)

		if output == nil {
			return errProofGenerating
		}
		proof = common.Hex2Bytes(output.Circuit.Proof[2:])
		degree = output.Circuit.Degree
		log.Info("Proof generated", "height", opts.Height, "degree", degree, "time", time.Since(start))
		return nil
	}, backoff.NewConstantBackOff(10*time.Second)); err != nil {
		return nil, 0, err
	}
	return proof, degree, nil
}

// requestProof sends a RPC request to proverd to try to get the requested proof.
func (d *ZkevmRpcdProducer) requestProof(opts *ProofRequestOptions) (*RpcdOutput, error) {
	reqBody := RequestProofBody{
		JsonRPC: "2.0",
		ID:      common.Big1,
		Method:  "proof",
		Params: []*RequestProofBodyParam{{
			Circuit:            "pi",
			Block:              opts.Height,
			L1RPC:              d.L1Endpoint,
			L2RPC:              d.L2Endpoint,
			Retry:              true,
			Param:              d.Param,
			VerifyProof:        true,
			Mock:               false,
			Aggregate:          false,
			Prover:             opts.ProverAddress.Hex()[2:],
			ProposeBlockTxHash: opts.ProposeBlockTxHash.Hex()[2:],
		}},
	}

	jsonValue, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	res, err := http.Post(d.RpcdEndpoint, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to request proof, id: %d, statusCode: %d", opts.Height, res.StatusCode)
	}

	resBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var output RequestProofBodyResponse
	if err := json.Unmarshal(resBytes, &output); err != nil {
		return nil, err
	}

	return output.Result, nil
}

// Cancel cancels an existing proof generation.
// Right now, it is just a stub that does nothing, because it is not possible to cnacel the proof
// with the current zkevm software.
func (d *ZkevmRpcdProducer) Cancel(ctx context.Context, blockID *big.Int) error {
	log.Info("Cancel proof generation for block ", "blockId", blockID)
	return nil
}
