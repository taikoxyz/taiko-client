package bindings

import (
	"math/big"
)

// ProtocolStateVariables contains some state variables used by Taiko protocol, defined in protocol's LibData.
// NOTE: this struct *MUST* match the return values of TaikoL1.getStateVariables method.
// ref: https://github.com/taikoxyz/taiko-mono/blob/main/packages/protocol/contracts/L1/LibData.sol
type ProtocolStateVariables struct {
	GenesisHeight        uint64
	GenesisTimestamp     uint64
	StatusBits           uint64
	FeeBase              *big.Int
	NextBlockID          uint64
	LastProposedAt       uint64
	AvgBlockTime         uint64
	LatestVerifiedHeight uint64
	LatestVerifiedID     uint64
	AvgProofTime         uint64
}
