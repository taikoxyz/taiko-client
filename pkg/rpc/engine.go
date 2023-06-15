package rpc

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/beacon/engine"
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
	fc *engine.ForkchoiceStateV1,
	attributes *engine.PayloadAttributes,
) (*engine.ForkChoiceResponse, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var result *engine.ForkChoiceResponse
	if err := c.Client.CallContext(timeoutCtx, &result, "engine_forkchoiceUpdatedV2", fc, attributes); err != nil {
		return nil, err
	}

	return result, nil
}

// ExecutePayload executes a built block on the execution engine.
func (c *EngineClient) NewPayload(
	ctx context.Context,
	payload *engine.ExecutableData,
) (*engine.PayloadStatusV1, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var result *engine.PayloadStatusV1
	if err := c.Client.CallContext(timeoutCtx, &result, "engine_newPayloadV2", payload); err != nil {
		return nil, err
	}

	return result, nil
}

// GetPayload gets the execution payload associated with the payload ID.
func (c *EngineClient) GetPayload(
	ctx context.Context,
	payloadID *engine.PayloadID,
) (*engine.ExecutableData, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var result *engine.ExecutionPayloadEnvelope
	if err := c.Client.CallContext(timeoutCtx, &result, "engine_getPayloadV2", payloadID); err != nil {
		return nil, err
	}

	return result.ExecutionPayload, nil
}

// ExchangeTransitionConfiguration exchanges transition configs with the L2 execution engine.
func (c *EngineClient) ExchangeTransitionConfiguration(
	ctx context.Context,
	cfg *engine.TransitionConfigurationV1,
) (*engine.TransitionConfigurationV1, error) {
	var result *engine.TransitionConfigurationV1
	if err := c.Client.CallContext(ctx, &result, "engine_exchangeTransitionConfigurationV1", cfg); err != nil {
		return nil, err
	}

	return result, nil
}
