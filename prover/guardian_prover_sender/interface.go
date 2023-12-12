package guardianproversender

import (
	"context"
	"math/big"
)

// Sender defines an interface that communicates with a central Guardian Prover server,
// sending heartbeats and signed blocks (and in the future, contested blocks).
type Sender interface {
	SignAndSendBlock(ctx context.Context, blockID *big.Int) error
	SendHeartbeat(ctx context.Context) error
	Close() error
}
