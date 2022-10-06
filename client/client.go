package client

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/prysmaticlabs/prysm/network"
	"github.com/prysmaticlabs/prysm/network/authorization"
	"github.com/taikochain/taiko-client/core/beacon"
	"github.com/taikochain/taiko-client/ethclient"
	"github.com/taikochain/taiko-client/rpc"
)

// EngineRPCClient represents a RPC client connecting to an Ethereum Engine RPC
// endpoint.
type EngineRPCClient struct {
	*rpc.Client
}

// DialClientWithBackoff connects a ethereum RPC client at the given URL with
// a backoff strategy.
func DialClientWithBackoff(ctx context.Context, url string) (*ethclient.Client, error) {
	var client *ethclient.Client
	if err := backoff.Retry(
		func() (err error) {
			client, err = ethclient.DialContext(ctx, url)
			return err
		},
		backoff.NewExponentialBackOff(),
	); err != nil {
		return nil, err
	}
	return client, nil
}

// DialEngineClientWithBackoff connects an ethereum engine RPC client at the
// given URL with a backoff strategy.
func DialEngineClientWithBackoff(ctx context.Context, url string, jwtSecret string) (*EngineRPCClient, error) {
	var engineClient *EngineRPCClient
	if err := backoff.Retry(
		func() (err error) {
			client, err := DialEngineClient(ctx, url, jwtSecret)
			if err != nil {
				return err
			}

			engineClient = &EngineRPCClient{client}
			return nil
		},
		backoff.NewExponentialBackOff(),
	); err != nil {
		return nil, err
	}
	return engineClient, nil
}

// DialEngineClient initializes an RPC connection with authentication headers.
// Taken from https://github.com/prysmaticlabs/prysm/blob/v2.1.4/beacon-chain/execution/rpc_connection.go#L151
func DialEngineClient(ctx context.Context, endpointUrl string, jwtSecret string) (*rpc.Client, error) {
	endpoint := network.Endpoint{
		Url: endpointUrl,
		Auth: network.AuthorizationData{
			Method: authorization.Bearer,
			Value:  jwtSecret,
		},
	}

	// Need to handle ipc and http
	var client *rpc.Client
	u, err := url.Parse(endpoint.Url)
	if err != nil {
		return nil, err
	}
	switch u.Scheme {
	case "http", "https":
		client, err = rpc.DialHTTPWithClient(endpoint.Url, endpoint.HttpClient())
		if err != nil {
			return nil, err
		}
	case "":
		client, err = rpc.DialIPC(ctx, endpoint.Url)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("no known transport for URL scheme %q", u.Scheme)
	}
	if endpoint.Auth.Method != authorization.None {
		header, err := endpoint.Auth.ToHeaderValue()
		if err != nil {
			return nil, err
		}
		client.SetHeader("Authorization", header)
	}
	return client, nil
}

// ForkchoiceUpdate updates the forkchoice on the execution client. If attributes is not nil, the engine client will also begin building a block
// based on attributes after the new head block and return the payload ID.
// May return an error in ForkChoiceResult, but the error is marshalled into the error return
func (c *EngineRPCClient) ForkchoiceUpdate(ctx context.Context, fc *beacon.ForkchoiceStateV1, attributes *beacon.PayloadAttributesV1) (*beacon.ForkChoiceResponse, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var result *beacon.ForkChoiceResponse
	if err := c.Client.CallContext(timeoutCtx, &result, "engine_forkchoiceUpdatedV1", fc, attributes); err != nil {
		return nil, err
	}

	return result, nil
}

// ExecutePayload executes a built block on the execution engine and returns an error if it was not successful.
func (c *EngineRPCClient) NewPayload(ctx context.Context, payload *beacon.ExecutableDataV1) (*beacon.PayloadStatusV1, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var result *beacon.PayloadStatusV1
	if err := c.Client.CallContext(timeoutCtx, &result, "engine_newPayloadV1", payload); err != nil {
		return nil, err
	}

	return result, nil
}

// GetPayload gets the execution payload associated with the payloadId.
func (c *EngineRPCClient) GetPayload(ctx context.Context, payloadID *beacon.PayloadID) (*beacon.ExecutableDataV1, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var result *beacon.ExecutableDataV1
	if err := c.Client.CallContext(timeoutCtx, &result, "engine_getPayloadV1", payloadID); err != nil {
		return nil, err
	}

	return result, nil
}
