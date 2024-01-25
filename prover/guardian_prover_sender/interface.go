package guardianproversender

import (
	"context"
	"math/big"
)

type BlockSigner interface {
	SignAndSendBlock(ctx context.Context, blockID *big.Int) error
	SendStartup(ctx context.Context, revision string, version string) error
}

type Heartbeater interface {
	SendHeartbeat(ctx context.Context) error
}

// BlockSenderHeartbeater defines an interface that communicates with a central Guardian Prover server,
// sending heartbeats and signed blocks (and in the future, contested blocks).
type BlockSenderHeartbeater interface {
	BlockSigner
	Heartbeater
	Close() error
}
