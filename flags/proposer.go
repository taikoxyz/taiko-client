package flags

import (
	"github.com/urfave/cli/v2"
)

// Flags used by Proposer.
var (
	L1NodeFlag = cli.StringFlag{
		Name:     "l1",
		Usage:    "RPC endpoint of a L1 ethereum node",
		Value:    "http://127.0.0.1:18545",
		Required: true,
	}
	L2NodeFlag = cli.StringFlag{
		Name:     "l2",
		Usage:    "RPC endpoint of a L2 ethereum node",
		Value:    "http://127.0.0.1:28545",
		Required: true,
	}
	TaikoL1AddressFlag = cli.StringFlag{
		Name:     "taikoL1",
		Usage:    "TaikoL1 contract address",
		Required: true,
	}
	TaikoL2AddressFlag = cli.StringFlag{
		Name:     "taikoL2",
		Usage:    "TaikoL2 contract address",
		Required: true,
	}
	L1TransactorPrivKeyFlag = cli.StringFlag{
		Name:     "l1.transactorPrivkey",
		Usage:    "Private key for L1 transactor",
		Required: true,
	}
	L2ProposerPrivKeyFlag = cli.StringFlag{
		Name:     "l2.proposerPrivKey",
		Usage:    "Private key for L2 proposer",
		Required: true,
	}
	ProposeIntervalFlag = cli.StringFlag{
		Name:  "proposeInterval",
		Usage: "Interval for proposing new transactions",
	}
	// Special flags for testing
	ProduceInvalidBlocksFlag = cli.BoolFlag{
		Name:  "produceInvalidBlocks",
		Usage: "Special flag for testnet testing, if activated, the proposer will start producing bad blocks",
	}
	ProduceInvalidBlocksInterval = cli.Uint64Flag{
		Name:  "produceInvalidBlocksInterval",
		Usage: "Special flag for testnet testing, if activated, bad blocks will be produced every N valid blocks",
	}
)
