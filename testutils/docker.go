package testutils

import (
	"bufio"
	"context"
	"fmt"
	"math/big"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/suite"
)

const (
	gethHttpPort       uint64 = 8545
	gethWSPort         uint64 = 8546
	gethAuthPort       uint64 = 8551
	premintTokenAmount        = "92233720368547758070000000000000"
	anvilReady                = "Listening on 0.0.0.0:8545"
	gethReady                 = "HTTP server started"
)

var (
	jwtFile     string
	l1ImageID   string
	counter     uint64
	counterLock sync.Mutex
)

type ExampleTestSuite struct {
	suite.Suite
	l1ContainerConf *nodeConfig
	l2ContainerConf *nodeConfig
}

func (s *ExampleTestSuite) SetupSuite() {
	packageID := incrCounter()

	ctx := context.Background()
	s.l1ContainerConf = l1Config(packageID)
	s.NoError(startGethContainer(ctx, s.l1ContainerConf, anvilReady))

	s.l2ContainerConf = l2Config(packageID)
	s.NoError(startGethContainer(ctx, s.l2ContainerConf, gethReady))
}

func (s *ExampleTestSuite) TearDownSuite() {
	stopContainer(context.Background(), s.l1ContainerConf.ContainerID)
	stopContainer(context.Background(), s.l2ContainerConf.ContainerID)
}

type endpointPorts struct {
	HTTP uint64
	WS   uint64
	Auth uint64
}

func (e *endpointPorts) HttpEndpoint() string {
	return fmt.Sprintf("http://localhost:%d", e.HTTP)
}

func (e *endpointPorts) WsEndpoint() string {
	return fmt.Sprintf("ws://localhost:%d", e.WS)
}

func (e *endpointPorts) AuthEndpoint() string {
	return fmt.Sprintf("ws://localhost:%d", e.Auth)
}

func getL1Ports(testID uint64) *endpointPorts {
	return &endpointPorts{
		HTTP: 11545 + testID,
		WS:   12546 + testID,
		Auth: 13551 + testID,
	}
}

func getL2Ports(testID uint64) *endpointPorts {
	return &endpointPorts{
		HTTP: 21545 + testID,
		WS:   22546 + testID,
		Auth: 23551 + testID,
	}
}

func l1Config(id uint64) *nodeConfig {
	p := getL1Ports(id)
	nc := &nodeConfig{
		ContainerName: fmt.Sprintf("L1_%d", id),
		ContainerConfig: &container.Config{
			Image: l1ImageID,
		},
		Ports: p,
		HostConfig: &container.HostConfig{
			AutoRemove: true,
			PortBindings: map[nat.Port][]nat.PortBinding{
				nat.Port(tcpPortString(gethHttpPort)): {
					{
						HostIP:   "0.0.0.0",
						HostPort: strconv.FormatUint(p.HTTP, 10),
					},
				},
				nat.Port(tcpPortString(gethWSPort)): {
					{
						HostIP:   "0.0.0.0",
						HostPort: strconv.FormatUint(p.WS, 10),
					},
				},
			},
		},
	}
	return nc
}

func l2Config(id uint64) *nodeConfig {
	p := getL2Ports(id)
	nc := &nodeConfig{
		ContainerName: fmt.Sprintf("L2_%d", id),
		Ports:         p,
		ContainerConfig: &container.Config{
			Image: "gcr.io/evmchain/taiko-geth:taiko",
			Cmd: []string{
				"--nodiscover",
				"--gcmode",
				"archive",
				"--syncmode",
				"full",
				"--datadir",
				"/data/taiko-geth",
				"--networkid",
				"167001",
				"--metrics",
				"--metrics.expensive",
				"--metrics.addr",
				"0.0.0.0",
				"--http",
				"--http.addr",
				"0.0.0.0",
				"--http.vhosts",
				"*",
				"--http.corsdomain",
				"*",
				"--ws",
				"--ws.addr",
				"0.0.0.0",
				"--ws.origins",
				"*",
				"--authrpc.addr",
				"0.0.0.0",
				"--authrpc.port",
				"8551",
				"--authrpc.vhosts",
				"*",
				"--authrpc.jwtsecret",
				"/host/jwt.hex",
				"--allow-insecure-unlock",
				"--http.api",
				"admin,debug,eth,net,web3,txpool,miner,taiko",
				"--ws.api",
				"admin,debug,eth,net,web3,txpool,miner,taiko",
				"--taiko",
			},
		},
		HostConfig: &container.HostConfig{
			Binds: []string{fmt.Sprintf("%s:/host/jwt.hex", jwtFile)},
			PortBindings: map[nat.Port][]nat.PortBinding{
				nat.Port(tcpPortString(gethHttpPort)): {
					{
						HostIP:   "0.0.0.0",
						HostPort: strconv.FormatUint(p.HTTP, 10),
					},
				},
				nat.Port(tcpPortString(gethWSPort)): {
					{
						HostIP:   "0.0.0.0",
						HostPort: strconv.FormatUint(p.WS, 10),
					},
				},
				nat.Port(tcpPortString(gethAuthPort)): {
					{
						HostIP:   "0.0.0.0",
						HostPort: strconv.FormatUint(p.Auth, 10),
					},
				},
			},
			AutoRemove: true,
		},
	}
	return nc
}

func tcpPortString(port uint64) string {
	return fmt.Sprintf("%d/tcp", port)
}

type nodeConfig struct {
	ContainerName   string
	ContainerID     string
	Ports           *endpointPorts
	ContainerConfig *container.Config
	HostConfig      *container.HostConfig
}

func startGethContainer(ctx context.Context, conf *nodeConfig, message string) error {
	cli, err := client.NewClientWithOpts()
	if err != nil {
		return err
	}
	defer cli.Close()
	c, err := cli.ContainerCreate(ctx, conf.ContainerConfig, conf.HostConfig, nil, nil, conf.ContainerName)
	if err != nil {
		return err
	}
	if err := cli.ContainerStart(ctx, c.ID, types.ContainerStartOptions{}); err != nil {
		return err
	}
	conf.ContainerID = c.ID
	r, err := cli.ContainerAttach(ctx, c.ID, types.ContainerAttachOptions{Stream: true, Stderr: true, Stdout: true})
	if err != nil {
		return err
	}
	defer r.Conn.Close()
	scanner := bufio.NewScanner(r.Reader)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), message) {
			break
		}
	}
	return nil
}

func stopContainer(ctx context.Context, containerID string) error {
	cli, err := client.NewClientWithOpts()
	if err != nil {
		return err
	}
	defer cli.Close()
	return cli.ContainerStop(ctx, containerID, container.StopOptions{})
}

func buildL1Image() (string, error) {
	cli, err := client.NewClientWithOpts()
	if err != nil {
		return "", err
	}
	defer cli.Close()
	ctx := context.Background()
	nc := AnvilContainerConf()
	if err := startGethContainer(ctx, nc, anvilReady); err != nil {
		return "", err
	}
	if err := deployTaikoL1(); err != nil {
		return "", err
	}
	r, err := cli.ContainerCommit(ctx, nc.ContainerID, types.ContainerCommitOptions{})
	if err != nil {
		return "", err
	}
	if err := cli.ContainerStop(ctx, nc.ContainerID, container.StopOptions{}); err != nil {
		return "", err
	}
	return r.ID, nil
}

func AnvilContainerConf() *nodeConfig {
	p := &endpointPorts{
		HTTP: gethHttpPort,
	}
	nc := &nodeConfig{
		ContainerName: "anvil",
		Ports:         p,
		ContainerConfig: &container.Config{
			Image: "ghcr.io/foundry-rs/foundry:latest",
			ExposedPorts: map[nat.Port]struct{}{
				nat.Port(tcpPortString(gethHttpPort)): {},
			},
			Entrypoint: []string{"anvil", "--host", "0.0.0.0"},
		},
		HostConfig: &container.HostConfig{
			AutoRemove: true,
			PortBindings: map[nat.Port][]nat.PortBinding{
				nat.Port(tcpPortString(gethHttpPort)): {
					{
						HostIP:   "0.0.0.0",
						HostPort: strconv.FormatUint(p.HTTP, 10),
					},
				},
			},
		},
	}
	return nc
}

func getL2GenesisHash() (string, error) {
	ctx, l2ContainerConf := context.Background(), l2Config(0)
	if err := startGethContainer(ctx, l2ContainerConf, gethReady); err != nil {
		return "", err
	}
	cli, err := ethclient.DialContext(ctx, l2ContainerConf.Ports.HttpEndpoint())
	if err != nil {
		return "", err
	}
	defer cli.Close()
	genesis, err := cli.BlockByNumber(ctx, big.NewInt(0))
	if err != nil {
		return "", err
	}
	if err := stopContainer(ctx, l2ContainerConf.ContainerID); err != nil {
		return "", err
	}
	return genesis.Hash().String(), nil
}

func deployTaikoL1() error {
	l2GenesisHash, err := getL2GenesisHash()
	if err != nil {
		return err
	}
	cmd := exec.Command("forge",
		"script",
		"script/DeployOnL1.s.sol:DeployOnL1",
		"--fork-url",
		"http://localhost:8545",
		"--broadcast",
		"--ffi",
		"-vvvvv",
		"--block-gas-limit",
		"100000000",
	)
	cmd.Env = []string{
		"PRIVATE_KEY=ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80",
		"ORACLE_PROVER=0x70997970C51812dc3A010C7d01b50e0d17dc79C8",
		"OWNER=0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC",
		"TAIKO_L2_ADDRESS=0x1000777700000000000000000000000000000001",
		"L2_SIGNAL_SERVICE=0x1000777700000000000000000000000000000007",
		"SHARED_SIGNAL_SERVICE=0x0000000000000000000000000000000000000000",
		"TAIKO_TOKEN_PREMINT_RECIPIENTS=0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266,0x70997970C51812dc3A010C7d01b50e0d17dc79C8",
		fmt.Sprintf("TAIKO_TOKEN_PREMINT_AMOUNTS=%s,%s", premintTokenAmount, premintTokenAmount),
		fmt.Sprintf("L2_GENESIS_HASH=%s", l2GenesisHash),
	}
	monoPath, err := filepath.Abs("../../taiko-mono/packages/protocol")
	if err != nil {
		return err
	}
	cmd.Dir = monoPath
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		return fmt.Errorf("out=%s,err=%v", string(out), err)
	}
	return nil
}

func incrCounter() uint64 {
	counterLock.Lock()
	defer counterLock.Unlock()
	counter++
	return counter
}

func init() {
	f, err := filepath.Abs("../integration_test/nodes/jwt.hex")
	if err != nil {
		panic(err)
	}
	jwtFile = f

	l1ImageID, err = buildL1Image()
	if err != nil {
		panic(err)
	}
}
