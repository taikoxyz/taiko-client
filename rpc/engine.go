package rpc

import (
	"context"
	"time"

	"github.com/taikochain/taiko-client/core/beacon"
	"github.com/taikochain/taiko-client/rpc"
)

// EngineRPCClient represents a RPC client connecting to an Ethereum Engine RPC
// endpoint.
type EngineRPCClient struct {
	*rpc.Client
}

// ForkchoiceUpdate updates the forkchoice on the execution client. If
// attributes is not nil, the engine client will also begin building a block
// based on attributes after the new head block and return the payload ID.
// May return an error in ForkChoiceResult, but the error is marshalled into
// the error return
func (c *EngineRPCClient) ForkchoiceUpdate(
	ctx context.Context,
	fc *beacon.ForkchoiceStateV1,
	attributes *beacon.PayloadAttributesV1,
) (*beacon.ForkChoiceResponse, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var result *beacon.ForkChoiceResponse
	if err := c.Client.CallContext(timeoutCtx, &result, "engine_forkchoiceUpdatedV1", fc, attributes); err != nil {
		return nil, err
	}

	return result, nil
}

// ExecutePayload executes a built block on the execution engine and returns an error if it was not successful.
func (c *EngineRPCClient) NewPayload(
	ctx context.Context,
	payload *beacon.ExecutableDataV1,
) (*beacon.PayloadStatusV1, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var result *beacon.PayloadStatusV1
	if err := c.Client.CallContext(timeoutCtx, &result, "engine_newPayloadV1", payload); err != nil {
		return nil, err
	}

	return result, nil
}

// GetPayload gets the execution payload associated with the payloadId.
func (c *EngineRPCClient) GetPayload(
	ctx context.Context,
	payloadID *beacon.PayloadID,
) (*beacon.ExecutableDataV1, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var result *beacon.ExecutableDataV1
	if err := c.Client.CallContext(timeoutCtx, &result, "engine_getPayloadV1", payloadID); err != nil {
		return nil, err
	}

	return result, nil
}
