package node

import (
	"context"

	"github.com/urfave/cli/v2"
)

// Service is the interface for the server that the node runs.
type Service interface {
	InitFromCli(context.Context, *cli.Context) error
	Name() string
	Start() error
	Close(context.Context)
}
