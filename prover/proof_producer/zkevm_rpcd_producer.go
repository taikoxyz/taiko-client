package producer

import (
	"errors"
	"math/big"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/taikoxyz/taiko-client/bindings"
)

var (
	errRpcdUnhealthy = errors.New("ZKEVM RPCD endpoint is unhealthy")
)

type ZkevmRpcdProducer struct {
	RpcdEndpoint string
	Param        string // parameter file to use
	L2Endpoint   string // a L2 execution engine's RPC endpoint
	Retry        bool   // retry proof computation if error
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
		"Request proof from ZKEVM RPCD service",
		"blockID", blockID,
		"beneficiary", meta.Beneficiary,
		"height", header.Number,
		"hash", header.Hash(),
	)

	// TODO: call zkevm RPCD to get a proof.
	go func() {
		resultCh <- &ProofWithHeader{
			BlockID: blockID,
			Header:  header,
			Meta:    meta,
			ZkProof: []byte{0x00},
		}
	}()
	return nil
}

type RpcdOutput struct {
	Result struct {
		Circuit struct {
			Instances []string `json:"instances"`
			Proof     string   `json:"proof"`
		} `json:"circuit"`
	} `json:"result"`
}

func (d *ZkevmRpcdProducer) outputToCalldata(output *RpcdOutput) []byte {
	calldata := []byte{}
	data := output.Result.Circuit
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
	for i := 0; i < len(proofBytes); i++ {
		calldata = append(calldata, proofBytes...)
	}

	return calldata[:bufLen]
}
