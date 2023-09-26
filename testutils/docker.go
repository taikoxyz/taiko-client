package testutils

import (
	"context"
	"errors"
	"os/exec"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

var (
	composeFile     = "integration_test/nodes/docker-compose.yml"
	l1ContainerName = "l1_node-1"
	gethHttpPort    = "8545"
	gethWSPort      = "8546"
)

func (s *ClientTestSuite) SetupSuit() {
	c, err := client.NewClientWithOpts()
	s.NoError(err)
	s.dockerCli = c
	s.NoError(s.compose("up"))
	id, err := s.l1ContainerID()
	s.NoError(err)
	ctx := context.Background()
	s.taikoL1ImageID, err = s.buildL1Image(ctx, id)
	s.NoError(err)
	s.T().Logf("l1ImageID: %s\n", s.taikoL1ImageID)
}

func (s *ClientTestSuite) TearDownSuit() {
	s.dockerCli.Close()
}

func (s *ClientTestSuite) compose(action string) error {
	return exec.Command("docker-compose", action, "-f "+composeFile).Run()
}

func (s *ClientTestSuite) l1ContainerID() (string, error) {
	containers, err := s.dockerCli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		return "", err
	}
	for _, c := range containers {
		for _, n := range c.Names {
			if n == "/"+composeFile {
				return c.ID, nil
			}
		}
	}
	return "", errors.New("not found")
}

func (s *ClientTestSuite) buildL1Image(ctx context.Context, containerID string) (string, error) {
	resp, err := s.dockerCli.ContainerCommit(ctx, containerID, types.ContainerCommitOptions{})
	if err != nil {
		return "", err
	}
	return resp.ID, nil
}

func (s *ClientTestSuite) Counter() uint64 {
	s.testCounterLock.Lock()
	defer s.testCounterLock.Unlock()
	s.testCounter++
	return s.testCounter
}

type EndpointPorts struct {
	L1HTTP   uint
	L1WS     uint
	L2HTTP   uint
	L2WSPort uint
	L2Auth   uint
}

func (s *ClientTestSuite) eePorts(testID uint) *EndpointPorts {
	return &EndpointPorts{
		L1HTTP:   18545 + testID,
		L1WS:     18546 + testID,
		L2HTTP:   28545 + testID,
		L2WSPort: 28546 + testID,
		L2Auth:   28551 + testID,
	}
}

func (s *ClientTestSuite) l1ContainerConfig(p *EndpointPorts) (*container.Config, *container.HostConfig) {
	config := &container.Config{
		Image: s.taikoL1ImageID,
	}
	hConfig := &container.HostConfig{
		AutoRemove: true,
	}
	return config, hConfig
}

func (s *ClientTestSuite) StartL1L2(ctx context.Context, testID uint64) error {
	// ports := s.eePorts(testID)
	// config := container.Config{}
	// l1 := s.dockerCli.ContainerCreate(ctx)
	return nil
}

func (s *ClientTestSuite) TestDocker(t *testing.T) {
}
