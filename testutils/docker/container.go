package docker

import (
	"bufio"
	"context"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

type ReadyContainer struct {
	ID              string
	name            string
	containerConfig *container.Config
	hostConfig      *container.HostConfig
	readyHint       string
	cli             *client.Client
	IPAddress       string
	PortMap         nat.PortMap
}

func NewReadyContainer(name string, cc *container.Config,
	hc *container.HostConfig, hint string,
) (*ReadyContainer, error) {
	cli, err := client.NewClientWithOpts()
	if err != nil {
		return nil, err
	}
	return &ReadyContainer{
		name:            name,
		containerConfig: cc,
		hostConfig:      hc,
		readyHint:       hint,
		cli:             cli,
	}, nil
}

func (c *ReadyContainer) Start(ctx context.Context) error {
	r, err := c.cli.ContainerCreate(ctx, c.containerConfig, c.hostConfig, nil, nil, c.name)
	if err != nil {
		return err
	}
	if err := c.cli.ContainerStart(ctx, r.ID, types.ContainerStartOptions{}); err != nil {
		return err
	}
	c.ID = r.ID
	resp, err := c.cli.ContainerAttach(ctx, r.ID, types.ContainerAttachOptions{Stream: true, Stderr: true, Stdout: true})
	if err != nil {
		return err
	}
	defer resp.Conn.Close()
	scanner := bufio.NewScanner(resp.Reader)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), c.readyHint) {
			break
		}
	}
	info, err := c.cli.ContainerInspect(ctx, c.ID)
	if err != nil {
		return err
	}
	c.IPAddress = info.NetworkSettings.IPAddress
	c.PortMap = info.NetworkSettings.Ports
	return nil
}

func (c *ReadyContainer) Stop() error {
	if err := c.cli.ContainerStop(context.Background(), c.ID, container.StopOptions{}); err != nil {
		return err
	}
	if err := c.cli.Close(); err != nil {
		return err
	}
	return nil
}
