package bindings

import "github.com/ethereum/go-ethereum/accounts/abi/bind"

// L1State contains some variables used by L1 Taiko protocol, defined in protocol's LibData.
// NOTE: this struct *MUST* match the return values of TaikoL1.getStateVariables method.
// ref: https://github.com/taikoxyz/taiko-mono/blob/main/packages/protocol/contracts/L1/LibData.sol
type L1State struct {
	GenesisHeight        uint64
	LatestVerifiedHeight uint64
	LatestVerifiedId     uint64
	NextBlockId          uint64
}

// GetL1State gets the L1 state variables from TaikoL1 contract.
func GetL1State(taikoL1 *TaikoL1Client, opts *bind.CallOpts) (*L1State, error) {
	var (
		s   = new(L1State)
		err error
	)
	s.GenesisHeight,
		s.LatestVerifiedHeight,
		s.LatestVerifiedId,
		s.NextBlockId,
		err = taikoL1.GetStateVariables(opts)
	return s, err
}
