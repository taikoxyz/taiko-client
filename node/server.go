package node

import (
	"context"
)

// Service is the interface for the server that the node runs.
type Service interface {
	Name() string
	Start() error
	Close(context.Context)
}
