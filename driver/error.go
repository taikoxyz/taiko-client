package driver

import (
	"errors"
	"fmt"

	"github.com/taikochain/taiko-client/common"
)

var (
	// errInvalidProposeBlockTx is returned when the given `proposeBlock` tx
	// is invalid.
	errInvalidProposeBlockTx = errors.New("invalid propose block tx")
	// errGenesisNotFound is returned when the L2 genesis has not initialized yet
	// in TaikoL1 contract.
	errGenesisNotFound = errors.New("genesis block not found in TaikoL1 contract")
	// errEmptyPayloadID is returned when the received payload ID is empty.
	errEmptyPayloadID = errors.New("empty payload ID")
)

// errGenesisHashMismatch is returned when local genesis is not matched
// with the remote genesis in TaikoL1 contract.
type errGenesisHashMismatch struct {
	Node    common.Hash
	TaikoL1 common.Hash
}

// Error implements the error interface.
func (e errGenesisHashMismatch) Error() string {
	return fmt.Sprintf(
		"genesis header hash mismatch, node: %s, TaikoL1 contract: %s",
		e.Node,
		e.TaikoL1,
	)
}

type errForkchoiceUpdated struct {
	Status string
}

// Error implements the error interface.
func (e errForkchoiceUpdated) Error() string {
	return fmt.Sprintf("failed to update forkchoice, status: %s", e.Status)
}

type errExecPayload struct {
	Status string
}

// Error implements the error interface.
func (e errExecPayload) Error() string {
	return fmt.Sprintf("failed to execute the new built block, status: %s", e.Status)
}
