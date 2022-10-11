package rpc

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/core/beacon"
	"github.com/ethereum/go-ethereum/rpc"
)

// EngineClient represents a RPC client connecting to an Ethereum Engine API
// endpoint.
// ref: https://github.com/ethereum/execution-apis/blob/main/src/engine/specification.md
type EngineClient struct {
	*rpc.Client
}

// ForkchoiceUpdate updates the forkchoice on the execution client.
func (c *EngineClient) ForkchoiceUpdate(
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

// ExecutePayload executes a built block on the execution engine.
func (c *EngineClient) NewPayload(
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

// GetPayload gets the execution payload associated with the payload ID.
func (c *EngineClient) GetPayload(
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
