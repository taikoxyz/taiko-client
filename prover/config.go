package prover

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/taikoxyz/taiko-client/cmd/flags"
	"github.com/taikoxyz/taiko-client/prover/auction"
	"github.com/urfave/cli/v2"
)

// Config contains the configurations to initialize a Taiko prover.
type Config struct {
	L1WsEndpoint             string
	L1HttpEndpoint           string
	L2WsEndpoint             string
	L2HttpEndpoint           string
	TaikoL1Address           common.Address
	TaikoL2Address           common.Address
	L1ProverPrivKey          *ecdsa.PrivateKey
	ZKEvmRpcdEndpoint        string
	ZkEvmRpcdParamsPath      string
	StartingBlockID          *big.Int
	MaxConcurrentProvingJobs uint
	Dummy                    bool
	OracleProver             bool
	SystemProver             bool
	OracleProverPrivateKey   *ecdsa.PrivateKey
	SystemProverPrivateKey   *ecdsa.PrivateKey
	Graffiti                 string
	ExpectedReward           uint64
	BidConfig
}

// Config contains the configurations to initialize a bidder for a Taiko block batch auction
type BidConfig struct {
	BidStrategyOption   auction.Option
	MinimumBidFeePerGas *big.Int
	BidDeposit          *big.Int
}

// NewConfigFromCliContext creates a new config instance from command line flags.
func NewConfigFromCliContext(c *cli.Context) (*Config, error) {
	l1ProverPrivKeyStr := c.String(flags.L1ProverPrivKey.Name)

	l1ProverPrivKey, err := crypto.ToECDSA(common.Hex2Bytes(l1ProverPrivKeyStr))
	if err != nil {
		return nil, fmt.Errorf("invalid L1 prover private key: %w", err)
	}

	bidStrategyOption := auction.Option(c.String(flags.BidStrategy.Name))

	if !auction.IsValidStrategy(bidStrategyOption) {
		return nil, fmt.Errorf("unsupported bid strategy")
	}

	var minimumFeePerGas *big.Int
	var ok bool

	if bidStrategyOption == auction.StrategyMinimumBidFeePerGas {
		if !c.IsSet(flags.MinimumBidFeePerGas.Name) {
			return nil, fmt.Errorf("MinimumBidFeePerGas flag is required with MinimumBidFeePerGas bid strategy")
		}
		minimumFeePerGas, ok = new(big.Int).SetString(c.String(flags.MinimumBidFeePerGas.Name), 10)
		if !ok {
			return nil, fmt.Errorf("could not convert MinimumBidFeePerGas to big int")
		}
	}

	bidDeposit, ok := new(big.Int).SetString(c.String(flags.BidDeposit.Name), 10)
	if !ok {
		return nil, fmt.Errorf("could not convert bid deposit to big int")
	}

	oracleProverSet := c.IsSet(flags.OracleProver.Name)
	systemProverSet := c.IsSet(flags.SystemProver.Name)

	if oracleProverSet && systemProverSet {
		return nil, fmt.Errorf("cannot set both oracleProver and systemProver")
	}

	var oracleProverPrivKey *ecdsa.PrivateKey
	if oracleProverSet {
		if !c.IsSet(flags.OracleProverPrivateKey.Name) {
			return nil, fmt.Errorf("oracleProver flag set without oracleProverPrivateKey set")
		}

		oracleProverPrivKeyStr := c.String(flags.OracleProverPrivateKey.Name)

		oracleProverPrivKey, err = crypto.ToECDSA(common.Hex2Bytes(oracleProverPrivKeyStr))
		if err != nil {
			return nil, fmt.Errorf("invalid oracle private key: %w", err)
		}
	}

	var systemProverPrivKey *ecdsa.PrivateKey
	if systemProverSet {
		if !c.IsSet(flags.SystemProverPrivateKey.Name) {
			return nil, fmt.Errorf("systemProver flag set without systemProverPrivateKey set")
		}

		systemProverPrivKeyStr := c.String(flags.SystemProverPrivateKey.Name)

		systemProverPrivKey, err = crypto.ToECDSA(common.Hex2Bytes(systemProverPrivKeyStr))
		if err != nil {
			return nil, fmt.Errorf("invalid system private key: %w", err)
		}
	}

	var startingBlockID *big.Int
	if c.IsSet(flags.StartingBlockID.Name) {
		startingBlockID = new(big.Int).SetUint64(c.Uint64(flags.StartingBlockID.Name))
	}

	return &Config{
		L1WsEndpoint:             c.String(flags.L1WSEndpoint.Name),
		L1HttpEndpoint:           c.String(flags.L1HTTPEndpoint.Name),
		L2WsEndpoint:             c.String(flags.L2WSEndpoint.Name),
		L2HttpEndpoint:           c.String(flags.L2HTTPEndpoint.Name),
		TaikoL1Address:           common.HexToAddress(c.String(flags.TaikoL1Address.Name)),
		TaikoL2Address:           common.HexToAddress(c.String(flags.TaikoL2Address.Name)),
		L1ProverPrivKey:          l1ProverPrivKey,
		ZKEvmRpcdEndpoint:        c.String(flags.ZkEvmRpcdEndpoint.Name),
		ZkEvmRpcdParamsPath:      c.String(flags.ZkEvmRpcdParamsPath.Name),
		StartingBlockID:          startingBlockID,
		MaxConcurrentProvingJobs: c.Uint(flags.MaxConcurrentProvingJobs.Name),
		Dummy:                    c.Bool(flags.Dummy.Name),
		OracleProver:             c.Bool(flags.OracleProver.Name),
		OracleProverPrivateKey:   oracleProverPrivKey,
		SystemProver:             c.Bool(flags.SystemProver.Name),
		SystemProverPrivateKey:   systemProverPrivKey,
		Graffiti:                 c.String(flags.Graffiti.Name),
		ExpectedReward:           c.Uint64(flags.ExpectedReward.Name),
		BidConfig: BidConfig{
			BidStrategyOption:   bidStrategyOption,
			MinimumBidFeePerGas: minimumFeePerGas,
			BidDeposit:          bidDeposit,
		},
	}, nil
}
