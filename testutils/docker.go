package testutils

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"sync"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/stretchr/testify/suite"
)

var (
	composeFile            = "integration_test/nodes/docker-compose.yml"
	l1ContainerName        = "l1_node-1"
	gethHttpPort    uint64 = 8545
	gethWSPort      uint64 = 8546
	gethAuthPort    uint64 = 8551
	counter         uint64
	counterLock     sync.Mutex
)

type ExampleTestSuite struct {
	suite.Suite
	dockerCli *client.Client
	l1ImageID string
}

func (s *ExampleTestSuite) SetupSuite() {
	c, err := client.NewClientWithOpts()
	s.NoError(err)
	s.dockerCli = c
	// s.NoError(s.startDevNet())
	id, err := s.l1ContainerID()
	s.NoError(err)
	ctx := context.Background()
	s.l1ImageID, err = s.buildL1Image(ctx, id)
	s.NoError(err)
	s.T().Logf("l1ImageID: %s\n", s.l1ImageID)
}

func (s *ExampleTestSuite) TearDownSuite() {
	s.dockerCli.Close()
}

func (s *ExampleTestSuite) compose(action string) error {
	return exec.Command("docker-compose", action, "-f "+composeFile).Run()
}

func (s *ExampleTestSuite) startDevNet() error {
	// cmd := exec.Command("make", "dev_net")
	// cmd.Env = []string{
	// 	"TAIKO_MONO_DIR=../taiko-mono", "COMPILE_PROTOCOL=false",
	// }
	// return cmd.Run()

	cmd := exec.Command("pwd")
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}
	d, err := cmd.Output()
	fmt.Print(d)
	return err
}

func (s *ExampleTestSuite) l1ContainerID() (string, error) {
	containers, err := s.dockerCli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		return "", err
	}
	for _, c := range containers {
		for _, n := range c.Names {
			if n == l1ContainerName {
				return c.ID, nil
			}
		}
	}
	return "", errors.New("not found")
}

func (s *ExampleTestSuite) buildL1Image(ctx context.Context, containerID string) (string, error) {
	resp, err := s.dockerCli.ContainerCommit(ctx, containerID, types.ContainerCommitOptions{})
	if err != nil {
		return "", err
	}
	return resp.ID, nil
}

func (s *ExampleTestSuite) Counter() uint64 {
	counterLock.Lock()
	defer counterLock.Unlock()
	counter++
	return counter
}

type EndpointPorts struct {
	L1HTTP   uint64
	L1WS     uint64
	L2HTTP   uint64
	L2WSPort uint64
	L2Auth   uint64
}

func (s *ExampleTestSuite) eePorts(testID uint64) *EndpointPorts {
	return &EndpointPorts{
		L1HTTP:   18545 + testID,
		L1WS:     28546 + testID,
		L2HTTP:   38545 + testID,
		L2WSPort: 48546 + testID,
		L2Auth:   58551 + testID,
	}
}

func (s *ExampleTestSuite) l1ContainerConfig(p *EndpointPorts) (*container.Config, *container.HostConfig) {
	config := &container.Config{
		Image: s.l1ImageID,
	}
	hConfig := &container.HostConfig{
		AutoRemove: true,
		PortBindings: map[nat.Port][]nat.PortBinding{
			nat.Port(strconv.FormatUint(gethHttpPort, 10)): {
				{
					HostIP:   "localhost",
					HostPort: strconv.FormatUint(p.L1HTTP, 10),
				},
			},
			nat.Port(strconv.FormatUint(gethWSPort, 10)): {
				{
					HostIP:   "localhost",
					HostPort: strconv.FormatUint(p.L1WS, 10),
				},
			},
		},
	}
	return config, hConfig
}

func (s *ExampleTestSuite) l2ContainerConfig(p *EndpointPorts) (*container.Config, *container.HostConfig) {
	config := &container.Config{
		Image: "gcr.io/evmchain/taiko-geth:taiko",
	}
	hConfig := &container.HostConfig{
		AutoRemove: true,
		PortBindings: map[nat.Port][]nat.PortBinding{
			nat.Port(strconv.FormatUint(gethHttpPort, 10)): {
				{
					HostIP:   "localhost",
					HostPort: strconv.FormatUint(p.L2HTTP, 10),
				},
			},
			nat.Port(strconv.FormatUint(gethWSPort, 10)): {
				{
					HostIP:   "localhost",
					HostPort: strconv.FormatUint(p.L2WSPort, 10),
				},
			},
			nat.Port(strconv.FormatUint(gethAuthPort, 10)): {
				{
					HostIP:   "localhost",
					HostPort: strconv.FormatUint(p.L2Auth, 10),
				},
			},
		},
	}
	return config, hConfig
}

func (s *ExampleTestSuite) StartL1L2(ctx context.Context, testID uint64) (string, string, error) {
	ports := s.eePorts(testID)
	config, hConfig := s.l1ContainerConfig(ports)
	name := fmt.Sprintf("l1_%d", testID)
	l1ID, err := s.startContainer(ctx, name, config, hConfig)
	if err != nil {
		return "", "", err
	}

	config, hConfig = s.l2ContainerConfig(ports)
	name = fmt.Sprintf("l2_%d", testID)
	l2ID, err := s.startContainer(ctx, name, config, hConfig)
	if err != nil {
		return "", "", err
	}
	return l1ID, l2ID, nil
}

func (s *ExampleTestSuite) startContainer(ctx context.Context, name string, config *container.Config, hConfig *container.HostConfig) (string, error) {
	l1, err := s.dockerCli.ContainerCreate(ctx, config, hConfig, nil, nil, name)
	if err != nil {
		return "", err
	}
	if err := s.dockerCli.ContainerStart(ctx, l1.ID, types.ContainerStartOptions{}); err != nil {
		return "", err
	}
	return l1.ID, nil
}

func (s *ExampleTestSuite) stopContainer(ctx context.Context, containerID string) error {
	return s.dockerCli.ContainerStop(ctx, containerID, container.StopOptions{})
}
