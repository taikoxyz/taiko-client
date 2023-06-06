package rpc

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/prysmaticlabs/prysm/v4/network"
	"github.com/prysmaticlabs/prysm/v4/network/authorization"
)

// DialClientWithBackoff connects a ethereum RPC client at the given URL with
// a backoff strategy.
func DialClientWithBackoff(ctx context.Context, url string) (*ethclient.Client, error) {
	var client *ethclient.Client
	if err := backoff.Retry(
		func() (err error) {
			client, err = ethclient.DialContext(ctx, url)
			log.Error("Dial ethclient error", "url", url, "error", err)
			return err
		},
		backoff.NewConstantBackOff(12*time.Second),
	); err != nil {
		return nil, err
	}

	return client, nil
}

// DialEngineClientWithBackoff connects an ethereum engine RPC client at the
// given URL with a backoff strategy.
func DialEngineClientWithBackoff(ctx context.Context, url string, jwtSecret string) (*EngineClient, error) {
	var engineClient *EngineClient
	if err := backoff.Retry(
		func() (err error) {
			client, err := DialEngineClient(ctx, url, jwtSecret)
			if err != nil {
				log.Error("Dial engine client error", "url", url, "error", err)
				return err
			}

			engineClient = &EngineClient{client}
			return nil
		},
		backoff.NewConstantBackOff(12*time.Second),
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
		client, err = rpc.DialOptions(ctx, endpoint.Url, rpc.WithHTTPClient(endpoint.HttpClient()))
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
