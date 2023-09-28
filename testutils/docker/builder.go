package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type CommitBuilder struct {
	container *ReadyContainer
	action    func() error
	cli       *client.Client
	imageName string
}

func NewCommitBuilder(c *ReadyContainer, do func() error, imageName string) (*CommitBuilder, error) {
	cli, err := client.NewClientWithOpts()
	if err != nil {
		return nil, err
	}
	b := &CommitBuilder{
		container: c,
		action:    do,
		cli:       cli,
		imageName: imageName,
	}
	return b, nil
}

func (b *CommitBuilder) Build(ctx context.Context) error {
	if err := b.action(); err != nil {
		return err
	}
	o := types.ContainerCommitOptions{Reference: b.imageName}
	if _, err := b.cli.ContainerCommit(ctx, b.container.ID, o); err != nil {
		return err
	}
	return nil
}

func (b *CommitBuilder) Stop() {
	if err := b.cli.Close(); err != nil {
		panic(err)
	}
}
