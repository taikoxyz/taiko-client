package guardianproversender

import (
	"context"
	"math/big"
)

// DummySender does not send a signed block or health check to the health check endpoint,
// used for tests to not have to create a mock test server.
type DummySender struct {
}

func (s *DummySender) SignAndSendBlock(ctx context.Context, blockID *big.Int) error {
	return nil
}

func (s *DummySender) Close() error {
	return nil
}

func (s *DummySender) SendHeartbeat(ctx context.Context) error {
	return nil
}
