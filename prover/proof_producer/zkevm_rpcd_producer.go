package producer

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/taikoxyz/taiko-client/bindings"
)

var (
	errRpcdUnhealthy   = errors.New("ZKEVM RPCD endpoint is unhealthy")
	errProofGenerating = errors.New("proof is generating")
)

type ZkevmRpcdProducer struct {
	RpcdEndpoint string
	Param        string // parameter file to use
	L2Endpoint   string // a L2 execution engine's RPC endpoint
	Retry        bool   // retry proof computation if error
}

type RequestProofBody struct {
	JsonRPC string                   `json:"jsonrpc"`
	ID      *big.Int                 `json:"id"`
	Method  string                   `json:"method"`
	Params  []*RequestProofBodyParam `json:"params"`
}

type RequestProofBodyParam struct {
	Circuit     string   `json:"circuit"`
	Block       *big.Int `json:"block"`
	RPC         string   `json:"rpc"`
	Retry       bool     `json:"retry"`
	Param       string   `json:"param"`
	VerifyProof bool     `json:"verify_proof"`
	Mock        bool     `json:"mock"`
	Aggregate   bool     `json:"aggregate"`
}

type RequestProofBodyResponse struct {
	JsonRPC string      `json:"jsonrpc"`
	ID      *big.Int    `json:"id"`
	Result  *RpcdOutput `json:"result"`
}

type RpcdOutput struct {
	Circuit struct {
		Instances []string `json:"instance"`
		Proof     string   `json:"proof"`
	} `json:"circuit"`
}

func NewZkevmRpcdProducer(
	rpcdEndpoint string,
	param string,
	l2Endpoint string,
	retry bool,
) (*ZkevmRpcdProducer, error) {
	resp, err := http.Get(rpcdEndpoint + "/health")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errRpcdUnhealthy
	}

	return &ZkevmRpcdProducer{RpcdEndpoint: rpcdEndpoint, Param: param, L2Endpoint: l2Endpoint, Retry: retry}, nil
}

// RequestProof implements the ProofProducer interface.
func (d *ZkevmRpcdProducer) RequestProof(
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

	go func() {
		resultCh <- &ProofWithHeader{
			BlockID: blockID,
			Header:  header,
			Meta:    meta,
			ZkProof: d.callProverDeamon(opts),
		}
	}()
	return nil
}

func (d *ZkevmRpcdProducer) callProverDeamon(opts *ProofRequestOptions) []byte {
	var (
		proof []byte
		start = time.Now()
	)
	backoff.Retry(func() error {
		output, err := d.requestProof(opts)
		if err != nil {
			log.Error("Failed to request proof", "height", opts.Height, "err", err, "endpoint", d.RpcdEndpoint)
			return err
		}

		log.Info("Request proof", "height", opts.Height, "output", output)

		if output == nil {
			return errProofGenerating
		}
		proof = d.outputToCalldata(output)
		log.Info("Proof generated", "heigth", opts.Height, "time", time.Since(start))
		return nil
	}, backoff.NewConstantBackOff(10*time.Second))
	return proof
}

func (d *ZkevmRpcdProducer) requestProof(opts *ProofRequestOptions) (*RpcdOutput, error) {
	reqBody := RequestProofBody{
		JsonRPC: "2.0",
		ID:      common.Big1,
		Method:  "proof",
		Params: []*RequestProofBodyParam{{
			Circuit:     "pi",
			Block:       opts.Height,
			RPC:         d.L2Endpoint,
			Retry:       true,
			Param:       d.Param,
			VerifyProof: true,
			Mock:        false,
			Aggregate:   false,
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

func (d *ZkevmRpcdProducer) outputToCalldata(output *RpcdOutput) []byte {
	calldata := []byte{}
	data := output.Circuit
	bufLen := len(data.Instances)*32 + len(data.Proof)

	for i := 0; i < len(data.Instances); i++ {
		uint256Bytes := [32]byte{}
		evenHexLen := len(data.Instances[i]) - 2 + (len(data.Instances[i]) % 2)
		instanceHex := data.Instances[i][2:]
		if len(instanceHex) < evenHexLen {
			instanceHex = strings.Repeat("0", evenHexLen-len(instanceHex)) + instanceHex
		}
		instanceBytes := common.Hex2Bytes(instanceHex)

		for j := 0; j < len(instanceBytes); j++ {
			uint256Bytes[31-j] = instanceBytes[len(instanceBytes)-1-j]
		}
		for k := 0; k < 32; k++ {
			calldata = append(calldata, uint256Bytes[k])
		}
	}

	evenHexLen := len(data.Proof) - 2 + (len(data.Proof) % 2)
	proofBytesHex := data.Proof[2:]
	if len(proofBytesHex) < evenHexLen {
		proofBytesHex = strings.Repeat("0", evenHexLen-len(proofBytesHex)) + proofBytesHex
	}
	proofBytes := common.Hex2Bytes(proofBytesHex)
	calldata = append(calldata, proofBytes...)

	for i := len(calldata); i < bufLen; i++ {
		calldata = append(calldata, byte(0))
	}

	return calldata[:bufLen]
}
