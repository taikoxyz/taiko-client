package driver

import (
	"errors"
	"fmt"
)

var (
	// errInvalidProposeBlockTx is returned when the given `proposeBlock` tx
	// is invalid.
	errInvalidProposeBlockTx = errors.New("invalid propose block tx")
	// errEmptyPayloadID is returned when the received payload ID is empty.
	errEmptyPayloadID = errors.New("empty payload ID")
)

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
