package prover

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
)

// Config contains the configurations to initialize a Taiko prover.
type Config struct {
	L1WsEndpoint                    string
	L1HttpEndpoint                  string
	L2WsEndpoint                    string
	L2HttpEndpoint                  string
	TaikoL1Address                  common.Address
	TaikoL2Address                  common.Address
	L1ProverPrivKey                 *ecdsa.PrivateKey
	ZKEvmRpcdEndpoint               string
	ZkEvmRpcdParamsPath             string
	StartingBlockID                 *big.Int
	MaxConcurrentProvingJobs        uint
	Dummy                           bool
	OracleProver                    bool
	OracleProverPrivateKey          *ecdsa.PrivateKey
	OracleProofSubmissionDelay      time.Duration
	ProofSubmissionMaxRetry         uint64
	Graffiti                        string
	RandomDummyProofDelayLowerBound *time.Duration
	RandomDummyProofDelayUpperBound *time.Duration
	BackOffMaxRetrys                uint64
	BackOffRetryInterval            time.Duration
	CheckProofWindowExpiredInterval time.Duration
	ProveUnassignedBlocks           bool
	RPCTimeout                      *time.Duration
	WaitReceiptTimeout              time.Duration
	ProveBlockGasLimit              *uint64
	HTTPServerPort                  uint64
	Capacity                        uint64
	MinProofFee                     *big.Int
	MaxExpiry                       time.Duration
}

func (c *Config) Validate() error {
	if err := rpc.CheckURLScheme(c.L1WsEndpoint, "ws"); err != nil {
		return err
	}
	if err := rpc.CheckURLScheme(c.L2WsEndpoint, "ws"); err != nil {
		return err
	}
	if err := rpc.CheckURLScheme(c.L1HttpEndpoint, "http"); err != nil {
		return err
	}
	if err := rpc.CheckURLScheme(c.L2HttpEndpoint, "http"); err != nil {
		return err
	}

	if c.OracleProver {
		if c.OracleProverPrivateKey == nil {
			return fmt.Errorf("oracleProver flag set without oracleProverPrivateKey set")
		}
	} else {
		if c.Capacity == 0 {
			return fmt.Errorf("capacity is required if oracleProver is not set to true")
		}
	}
	return nil
}
