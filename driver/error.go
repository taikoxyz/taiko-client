package driver

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
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

// genesisHashMismatchError is returned when local genesis is not matched
// with the remote genesis in TaikoL1 contract.
type genesisHashMismatchError struct {
	Node    common.Hash
	TaikoL1 common.Hash
}

// Error implements the error interface.
func (e genesisHashMismatchError) Error() string {
	return fmt.Sprintf(
		"genesis header hash mismatch, node: %s, TaikoL1 contract: %s",
		e.Node,
		e.TaikoL1,
	)
}

type forkchoiceUpdatedError struct {
	Status string
}

// Error implements the error interface.
func (e forkchoiceUpdatedError) Error() string {
	return fmt.Sprintf("failed to update forkchoice, status: %s", e.Status)
}

type execPayloadError struct {
	Status string
}

// Error implements the error interface.
func (e execPayloadError) Error() string {
	return fmt.Sprintf("failed to execute the new built block, status: %s", e.Status)
}
