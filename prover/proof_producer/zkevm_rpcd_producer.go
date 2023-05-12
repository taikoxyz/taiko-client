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
	Circuit         string   `json:"circuit"`
	Block           *big.Int `json:"block"`
	L2RPC           string   `json:"l2_rpc"`
	Retry           bool     `json:"retry"`
	Param           string   `json:"param"`
	VerifyProof     bool     `json:"verify_proof"`
	Mock            bool     `json:"mock"`
	Aggregate       bool     `json:"aggregate"`
	Prover          string   `json:"prover"`
	L1SignalService string   `json:"l1_signal_service"`
	L2SignalService string   `json:"l2_signal_service"`
	TaikoL2         string   `json:"l2_contract"`
	MetaHash        string   `json:"meta_hash"`
	BlockHash       string   `json:"block_hash"`
	ParentHash      string   `json:"parent_hash"`
	SignalRoot      string   `json:"signal_root"`
	Graffiti        string   `json:"graffiti"`
	GasUsed         uint64   `json:"gas_used"`
	ParentGasUsed   uint64   `json:"parent_gas_used"`
}

// RequestProofBodyResponse represents the JSON body of the response of the proof requests.
type RequestProofBodyResponse struct {
	JsonRPC string      `json:"jsonrpc"`
	ID      *big.Int    `json:"id"`
	Result  *RpcdOutput `json:"result"`
	Error   *struct {
		Code    *big.Int `json:"code"`
		Message string   `json:"message"`
	} `json:"error,omitempty"`
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
func (p *ZkevmRpcdProducer) RequestProof(
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
	if p.CustomProofHook != nil {
		proof, degree, err = p.CustomProofHook()
	} else {
		proof, degree, err = p.callProverDaemon(ctx, opts)
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
func (p *ZkevmRpcdProducer) callProverDaemon(ctx context.Context, opts *ProofRequestOptions) ([]byte, uint64, error) {
	var (
		proof  []byte
		degree uint64
		start  = time.Now()
	)
	if err := backoff.Retry(func() error {
		if ctx.Err() != nil {
			return nil
		}
		output, err := p.requestProof(opts)
		if err != nil {
			log.Error("Failed to request proof", "height", opts.Height, "err", err, "endpoint", p.RpcdEndpoint)
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
func (p *ZkevmRpcdProducer) requestProof(opts *ProofRequestOptions) (*RpcdOutput, error) {
	reqBody := RequestProofBody{
		JsonRPC: "2.0",
		ID:      common.Big1,
		Method:  "proof",
		Params: []*RequestProofBodyParam{{
			Circuit:         "pi",
			Block:           opts.Height,
			L2RPC:           p.L2Endpoint,
			Retry:           true,
			Param:           p.Param,
			VerifyProof:     true,
			Mock:            false,
			Aggregate:       false,
			Prover:          opts.ProverAddress.Hex()[2:],
			L1SignalService: opts.L1SignalService.Hex()[2:],
			L2SignalService: opts.L2SignalService.Hex()[2:],
			TaikoL2:         opts.TaikoL2.Hex()[2:],
			MetaHash:        opts.MetaHash.Hex()[2:],
			BlockHash:       opts.BlockHash.Hex()[2:],
			ParentHash:      opts.ParentHash.Hex()[2:],
			SignalRoot:      opts.SignalRoot.Hex()[2:],
			Graffiti:        opts.Graffiti,
			GasUsed:         opts.GasUsed,
			ParentGasUsed:   opts.ParentGasUsed,
		}},
	}

	jsonValue, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	res, err := http.Post(p.RpcdEndpoint, "application/json", bytes.NewBuffer(jsonValue))
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

	if output.Error != nil {
		return nil, errors.New(output.Error.Message)
	}

	return output.Result, nil
}

// Cancel cancels an existing proof generation.
// Right now, it is just a stub that does nothing, because it is not possible to cnacel the proof
// with the current zkevm software.
func (p *ZkevmRpcdProducer) Cancel(ctx context.Context, blockID *big.Int) error {
	log.Info("Cancel proof generation for block ", "blockId", blockID)
	return nil
}
