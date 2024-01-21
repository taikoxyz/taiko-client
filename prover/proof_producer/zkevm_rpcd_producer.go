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
	"github.com/taikoxyz/taiko-client/bindings/encoding"
	"github.com/taikoxyz/taiko-client/internal/metrics"
)

var (
	errProofGenerating   = errors.New("proof is generating")
	proofPollingInterval = 10 * time.Second
)

// ZkevmRpcdProducer is responsible for requesting zk proofs from the given proverd endpoint.
type ZkevmRpcdProducer struct {
	RpcdEndpoint        string                         // a proverd RPC endpoint
	Param               string                         // parameter file to use
	L1Endpoint          string                         // a L1Client node RPC endpoint
	L2Endpoint          string                         // a L2Client execution engine's RPC endpoint
	Retry               bool                           // retry proof computation if error
	ProofTimeTarget     uint64                         // used for calculating proof delay
	ProtocolConfig      *bindings.TaikoDataConfig      // protocol configurations
	CustomProofHook     func() ([]byte, uint64, error) // only for testing purposes
	*DummyProofProducer                                // only for testing purposes
}

// RequestProofBody represents the JSON body for requesting the proof.
type RequestProofBody struct {
	JsonRPC string                   `json:"jsonrpc"` //nolint:revive,stylecheck
	ID      *big.Int                 `json:"id"`
	Method  string                   `json:"method"`
	Params  []*RequestProofBodyParam `json:"params"`
}

// RequestProofBodyParam represents the JSON body of RequestProofBody's `param` field.
type RequestProofBodyParam struct {
	Circuit          string            `json:"circuit"`
	Block            *big.Int          `json:"block"`
	L2RPC            string            `json:"rpc"`
	Retry            bool              `json:"retry"`
	Param            string            `json:"param"`
	VerifyProof      bool              `json:"verify_proof"`
	Mock             bool              `json:"mock"`
	MockFeedback     bool              `json:"mock_feedback"`
	Aggregate        bool              `json:"aggregate"`
	ProtocolInstance *ProtocolInstance `json:"protocol_instance"`
}

type RequestMetaData struct {
	L1Hash           string   `json:"l1_hash"`
	Difficulty       string   `json:"difficulty"`
	BlobHash         string   `json:"blob_hash"`
	ExtraData        string   `json:"extra_data"`
	DepositsHash     string   `json:"deposits_hash"`
	Coinbase         string   `json:"coinbase"`
	ID               uint64   `json:"id"`
	GasLimit         uint32   `json:"gas_limit"`
	Timestamp        uint64   `json:"timestamp"`
	L1Height         uint64   `json:"l1_height"`
	TxListByteOffset *big.Int `json:"tx_list_byte_offset"`
	TxListByteSize   *big.Int `json:"tx_list_byte_size"`
	MinTier          uint16   `json:"min_tier"`
	BlobUsed         bool     `json:"blob_used"`
	ParentMetaHash   string   `json:"parent_metahash"`
}

// ProtocolInstance represents the JSON body of RequestProofBody.Param's `protocol_instance` field.
type ProtocolInstance struct {
	L1SignalService         string           `json:"l1_signal_service"`
	L2SignalService         string           `json:"l2_signal_service"`
	TaikoL2                 string           `json:"l2_contract"`
	MetaHash                string           `json:"meta_hash"`
	BlockHash               string           `json:"block_hash"`
	ParentHash              string           `json:"parent_hash"`
	SignalRoot              string           `json:"signal_root"`
	Graffiti                string           `json:"graffiti"`
	Prover                  string           `json:"prover"`
	Treasury                string           `json:"treasury"`
	GasUsed                 uint64           `json:"gas_used"`
	ParentGasUsed           uint64           `json:"parent_gas_used"`
	BlockMaxGasLimit        uint64           `json:"block_max_gas_limit"`
	MaxTransactionsPerBlock uint64           `json:"max_transactions_per_block"`
	MaxBytesPerTxList       uint64           `json:"max_bytes_per_tx_list"`
	AnchorGasLimit          uint64           `json:"anchor_gas_limit"`
	RequestMetaData         *RequestMetaData `json:"request_meta_data"`
}

// RequestProofBodyResponse represents the JSON body of the response of the proof requests.
type RequestProofBodyResponse struct {
	JsonRPC string      `json:"jsonrpc"` //nolint:revive,stylecheck
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
	Aggregation struct {
		Instances []string `json:"instance"`
		Proof     string   `json:"proof"`
		Degree    uint64   `json:"k"`
	} `json:"aggregation"`
}

// NewZkevmRpcdProducer creates a new `ZkevmRpcdProducer` instance.
func NewZkevmRpcdProducer(
	rpcdEndpoint string,
	param string,
	l1Endpoint string,
	l2Endpoint string,
	retry bool,
	protocolConfig *bindings.TaikoDataConfig,
) (*ZkevmRpcdProducer, error) {
	return &ZkevmRpcdProducer{
		RpcdEndpoint:   rpcdEndpoint,
		Param:          param,
		L1Endpoint:     l1Endpoint,
		L2Endpoint:     l2Endpoint,
		Retry:          retry,
		ProtocolConfig: protocolConfig,
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
		"coinbase", meta.Coinbase,
		"height", header.Number,
		"hash", header.Hash(),
	)

	if p.DummyProofProducer != nil {
		return p.DummyProofProducer.RequestProof(ctx, opts, blockID, meta, header, p.Tier(), resultCh)
	}

	var (
		proof  []byte
		degree uint64
		err    error
	)
	if p.CustomProofHook != nil {
		proof, degree, err = p.CustomProofHook()
	} else {
		proof, degree, err = p.callProverDaemon(ctx, opts, meta)
	}
	if err != nil {
		return err
	}

	if proof, err = encoding.EncodeZKEvmProof(proof); err != nil {
		return err
	}

	resultCh <- &ProofWithHeader{
		BlockID: blockID,
		Header:  header,
		Meta:    meta,
		Proof:   proof,
		Degree:  degree,
		Opts:    opts,
		Tier:    p.Tier(),
	}

	metrics.ProverPseProofGeneratedCounter.Inc(1)

	return nil
}

// callProverDaemon keeps polling the proverd service to get the requested proof.
func (p *ZkevmRpcdProducer) callProverDaemon(
	ctx context.Context,
	opts *ProofRequestOptions,
	meta *bindings.TaikoDataBlockMetadata,
) ([]byte, uint64, error) {
	var (
		proof  []byte
		degree uint64
		start  = time.Now()
	)
	if err := backoff.Retry(func() error {
		if ctx.Err() != nil {
			return nil
		}
		output, err := p.requestProof(opts, meta)
		if err != nil {
			log.Error("Failed to request proof", "height", opts.BlockID, "err", err, "endpoint", p.RpcdEndpoint)
			return err
		}

		if output == nil {
			log.Info(
				"Proof generating",
				"height", opts.BlockID,
				"time", time.Since(start),
				"producer", "ZkevmRpcdProducer",
			)
			return errProofGenerating
		}

		log.Debug("Proof generation output", "output", output)

		var proofOutput string
		for _, instance := range output.Aggregation.Instances {
			proofOutput += instance[2:]
		}
		proofOutput += output.Aggregation.Proof[2:]

		proof = common.Hex2Bytes(proofOutput)
		degree = output.Aggregation.Degree
		log.Info(
			"Proof generated",
			"height", opts.BlockID,
			"degree", degree,
			"time", time.Since(start),
			"producer", "ZkevmRpcdProducer",
		)
		return nil
	}, backoff.NewConstantBackOff(proofPollingInterval)); err != nil {
		return nil, 0, err
	}

	return proof, degree, nil
}

// requestProof sends a RPC request to proverd to try to get the requested proof.
func (p *ZkevmRpcdProducer) requestProof(
	opts *ProofRequestOptions,
	meta *bindings.TaikoDataBlockMetadata,
) (*RpcdOutput, error) {
	reqBody := RequestProofBody{
		JsonRPC: "2.0",
		ID:      common.Big1,
		Method:  "proof",
		Params: []*RequestProofBodyParam{{
			Circuit:      "super",
			Block:        opts.BlockID,
			L2RPC:        p.L2Endpoint,
			Retry:        true,
			Param:        p.Param,
			VerifyProof:  true,
			Mock:         false,
			MockFeedback: false,
			Aggregate:    true,
			ProtocolInstance: &ProtocolInstance{
				Prover:            opts.ProverAddress.Hex()[2:],
				Treasury:          opts.TaikoL2.Hex()[2:],
				L1SignalService:   opts.L1SignalService.Hex()[2:],
				L2SignalService:   opts.L2SignalService.Hex()[2:],
				TaikoL2:           opts.TaikoL2.Hex()[2:],
				MetaHash:          opts.MetaHash.Hex()[2:],
				BlockHash:         opts.BlockHash.Hex()[2:],
				ParentHash:        opts.ParentHash.Hex()[2:],
				SignalRoot:        opts.SignalRoot.Hex()[2:],
				Graffiti:          opts.Graffiti,
				GasUsed:           opts.GasUsed,
				ParentGasUsed:     opts.ParentGasUsed,
				BlockMaxGasLimit:  uint64(p.ProtocolConfig.BlockMaxGasLimit),
				MaxBytesPerTxList: p.ProtocolConfig.BlockMaxTxListBytes.Uint64(),
				AnchorGasLimit:    encoding.AnchorTxGasLimit,
				RequestMetaData: &RequestMetaData{
					L1Hash:           common.BytesToHash(meta.L1Hash[:]).Hex()[2:],
					Difficulty:       common.BytesToHash(meta.Difficulty[:]).Hex()[2:],
					BlobHash:         common.BytesToHash(meta.BlobHash[:]).Hex()[2:],
					ExtraData:        common.BytesToHash(meta.ExtraData[:]).Hex()[2:],
					DepositsHash:     common.BytesToHash(meta.DepositsHash[:]).Hex()[2:],
					Coinbase:         meta.Coinbase.Hex()[2:],
					ID:               meta.Id,
					GasLimit:         meta.GasLimit,
					Timestamp:        meta.Timestamp,
					L1Height:         meta.L1Height,
					TxListByteOffset: meta.TxListByteOffset,
					TxListByteSize:   meta.TxListByteSize,
					MinTier:          meta.MinTier,
					BlobUsed:         meta.BlobUsed,
					ParentMetaHash:   common.BytesToHash(meta.ParentMetaHash[:]).Hex()[2:],
				},
			},
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
		return nil, fmt.Errorf("failed to request proof, id: %d, statusCode: %d", opts.BlockID, res.StatusCode)
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

// Tier implements the ProofProducer interface.
func (p *ZkevmRpcdProducer) Tier() uint16 {
	return encoding.TierSgxAndPseZkevmID
}

// Cancellable implements the ProofProducer interface.
func (p *ZkevmRpcdProducer) Cancellable() bool {
	return false
}

// Cancel cancels an existing proof generation.
// Right now, it is just a stub that does nothing, because it is not possible to cancel the proof
// with the current zkevm software.
func (p *ZkevmRpcdProducer) Cancel(ctx context.Context, blockID *big.Int) error {
	log.Info("Cancel proof generation for block", "blockId", blockID)
	return nil
}
