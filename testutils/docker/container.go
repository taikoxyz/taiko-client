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
	Name            string
	containerConfig *container.Config
	hostConfig      *container.HostConfig
	readyHint       string
	IPAddress       string
	PortMap         nat.PortMap
}

func NewReadyContainer(name string, cc *container.Config,
	hc *container.HostConfig, hint string,
) (*ReadyContainer, error) {
	return &ReadyContainer{
		Name:            name,
		containerConfig: cc,
		hostConfig:      hc,
		readyHint:       hint,
	}, nil
}

func (c *ReadyContainer) Start(ctx context.Context) error {
	cli, err := client.NewClientWithOpts()
	if err != nil {
		return err
	}
	r, err := cli.ContainerCreate(ctx, c.containerConfig, c.hostConfig, nil, nil, c.Name)
	if err != nil {
		return err
	}
	if err := cli.ContainerStart(ctx, r.ID, types.ContainerStartOptions{}); err != nil {
		return err
	}
	c.ID = r.ID
	resp, err := cli.ContainerAttach(ctx, r.ID, types.ContainerAttachOptions{Stream: true, Stderr: true, Stdout: true})
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
	c.IPAddress, c.PortMap, err = GetContainerInfo(ctx, cli, c.ID)
	if err != nil {
		return err
	}
	return cli.Close()
}

func GetContainerInfo(ctx context.Context, cli *client.Client, containerID string) (string, nat.PortMap, error) {
	info, err := cli.ContainerInspect(ctx, containerID)
	if err != nil {
		return "", nil, err
	}
	return info.NetworkSettings.IPAddress, info.NetworkSettings.Ports, nil
}

func (c *ReadyContainer) Stop() error {
	cli, err := client.NewClientWithOpts()
	if err != nil {
		return err
	}
	if err := cli.ContainerStop(context.Background(), c.ID, container.StopOptions{}); err != nil {
		return err
	}
	if err := cli.Close(); err != nil {
		return err
	}
	return cli.Close()
}
