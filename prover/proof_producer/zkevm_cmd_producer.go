package producer

import (
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/taikoxyz/taiko-client/bindings"
)

type ZkevmCmdProducer struct {
	CmdPath    string
	L2Endpoint string // a L2 execution engine's RPC endpoint
}

func NewZkevmCmdProducer(
	cmdPath string,
	l2Endpoint string,
) (*ZkevmCmdProducer, error) {
	return &ZkevmCmdProducer{cmdPath, l2Endpoint}, nil
}

// RequestProof implements the ProofProducer interface.
func (d *ZkevmCmdProducer) RequestProof(
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
		"cmd", d.CmdPath,
	)

	go func() {
		var (
			proof []byte
			err   error
		)
		backoff.Retry(func() error {
			if proof, err = d.ExecProverCmd(header.Number); err != nil {
				log.Error("Execute prover cmd error", "error", err)
				return err
			}

			return nil
		}, backoff.NewConstantBackOff(3*time.Second))

		resultCh <- &ProofWithHeader{
			BlockID: blockID,
			Header:  header,
			Meta:    meta,
			ZkProof: proof,
		}
	}()
	return nil
}

type ProverCmdOutput struct {
	Instances []string `json:"instances"`
	Proof     []byte   `json:"proof"`
}

func (d *ZkevmCmdProducer) ExecProverCmd(height *big.Int) ([]byte, error) {
	start := time.Now()
	cmd := exec.Command(d.CmdPath, d.L2Endpoint, height.String())

	if err := cmd.Run(); err != nil {
		return nil, err
	}

	outputPath := filepath.Join(d.CmdPath, fmt.Sprintf("../block-%s_proof.json", height))

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

	return proverCmdOutput.Proof, nil
}
