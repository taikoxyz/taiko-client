package producer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/taikoxyz/taiko-client/bindings"
)

// ZkevmCmdProducer is responsible for generating zk proofs from the given command line binary file.
type ZkevmCmdProducer struct {
	CmdPath    string
	L2Endpoint string // a L2 execution engine's RPC endpoint
}

// NewZkevmCmdProducer creates a new NewZkevmCmdProducer instance.
func NewZkevmCmdProducer(
	cmdPath string,
	l2Endpoint string,
) (*ZkevmCmdProducer, error) {
	return &ZkevmCmdProducer{cmdPath, l2Endpoint}, nil
}

// RequestProof implements the ProofProducer interface.
func (p *ZkevmCmdProducer) RequestProof(
	ctx context.Context,
	opts *ProofRequestOptions,
	blockID *big.Int,
	meta *bindings.TaikoDataBlockMetadata,
	header *types.Header,
	resultCh chan *ProofWithHeader,
) error {
	log.Info(
		"Request proof from ZKEVM CMD",
		"blockID", blockID,
		"beneficiary", meta.Beneficiary,
		"height", header.Number,
		"hash", header.Hash(),
		"cmd", p.CmdPath,
	)

	var (
		proof []byte
		err   error
	)
	if err := backoff.Retry(func() error {
		if proof, err = p.ExecProverCmd(opts.Height); err != nil {
			log.Error("Execute prover cmd error", "error", err)
			return err
		}

		return nil
	}, backoff.NewConstantBackOff(3*time.Second)); err != nil {
		log.Error("Failed to generate proof", "error", err)
	}

	resultCh <- &ProofWithHeader{
		BlockID: blockID,
		Header:  header,
		Meta:    meta,
		ZkProof: proof,
		Degree:  CircuitsDegree10Txs,
		Opts:    opts,
	}

	return nil
}

type ProverCmdOutput struct {
	Instances []string `json:"instances"`
	Proof     []byte   `json:"proof"`
}

func (p *ZkevmCmdProducer) ExecProverCmd(height *big.Int) ([]byte, error) {
	start := time.Now()
	cmd := exec.Command(p.CmdPath, p.L2Endpoint, height.String())

	var stdout, stderr bytes.Buffer

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		log.Info("Exec output", "stdout", stdout.String(), "stderr", stderr.String())
		return nil, err
	}

	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	outputPath := filepath.Join(wd, fmt.Sprintf("./block-%s_proof.json", height))

	log.Info("Exec prover cmd finished", "outputPath", outputPath, "time", time.Since(start))

	if _, err := os.Stat(outputPath); err != nil {
		return nil, err
	}

	defer func() {
		if err := os.Remove(outputPath); err != nil {
			log.Warn("Remove prover cmd output file error", "error", err)
		}
	}()

	outputJSONBytes, err := os.ReadFile(outputPath)
	if err != nil {
		return nil, err
	}

	var proverCmdOutput ProverCmdOutput
	if err := json.Unmarshal(outputJSONBytes, &proverCmdOutput); err != nil {
		return nil, err
	}

	return p.outputToCalldata(&proverCmdOutput), nil
}

func (p *ZkevmCmdProducer) outputToCalldata(output *ProverCmdOutput) []byte {
	calldata := []byte{}
	bufLen := len(output.Instances)*32 + len(output.Proof)

	for i := 0; i < len(output.Instances); i++ {
		uint256Bytes := [32]byte{}
		evenHexLen := len(output.Instances[i]) - 2 + (len(output.Instances[i]) % 2)
		instanceHex := output.Instances[i][2:]
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

	for i := 0; i < len(output.Proof); i++ {
		calldata = append(calldata, output.Proof...)
	}

	return calldata[:bufLen]
}

// Cancel cancels an existing proof generation.
// Right now, it is just a stub that does nothing, because it is not possible to cnacel the proof
// with the current zkevm software.
func (p *ZkevmCmdProducer) Cancel(ctx context.Context, blockID *big.Int) error {
	log.Info("Cancel proof generation for block", "blockId", blockID)
	return nil
}
