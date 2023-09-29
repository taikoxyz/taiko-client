package testutils

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/taiko-client/pkg/jwt"
	"github.com/taikoxyz/taiko-client/testutils/docker"
)

const (
	gethHttpPort      uint64 = 8545
	gethWSPort        uint64 = 8546
	gethAuthPort      uint64 = 8551
	baseContainerName        = "L1Base"
	showDeployLog            = false
)

var (
	gethHttpNatPort = natTcpPort(gethHttpPort)
	gethWSNatPort   = natTcpPort(gethWSPort)
	gethAuthNatPort = natTcpPort(gethAuthPort)
)

// variables need to be initialized
var (
	jwtSecret         []byte
	JwtSecretFile     string
	monoPath          string
	l1BaseContainer   = baseContainer{delExisted: true}
	TaikoL1Address    common.Address
	TaikoTokenAddress common.Address
)

type gethContainer struct {
	*docker.ReadyContainer
	isAnvil bool
}

type baseContainer struct {
	*gethContainer
	delExisted bool
}

func natTcpPort(p uint64) nat.Port {
	return nat.Port(fmt.Sprintf("%d/tcp", p))
}

func (e *gethContainer) HttpEndpoint() string {
	for k, v := range e.PortMap {
		if k == gethHttpNatPort {
			return fmt.Sprintf("http://localhost:%s", v[0].HostPort)
		}
	}
	return ""
}

func (e *gethContainer) InnerHttpEndpoint() string {
	return fmt.Sprintf("http://%s:%d", e.IPAddress, gethHttpPort)
}

func (e *gethContainer) WsEndpoint() string {
	p := gethWSNatPort
	if e.isAnvil {
		p = gethHttpNatPort
	}
	for k, v := range e.PortMap {
		if k == p {
			return fmt.Sprintf("ws://localhost:%s", v[0].HostPort)
		}
	}
	return ""
}

func (e *gethContainer) AuthEndpoint() string {
	for k, v := range e.PortMap {
		if k == gethAuthNatPort {
			return fmt.Sprintf("http://localhost:%s", v[0].HostPort)
		}
	}
	return ""
}

func newL1Container(name string) (*gethContainer, error) {
	c, err := newAnvilContainer(context.Background(), false, name)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func newL2Container(name string) (*gethContainer, error) {
	cc := &container.Config{
		Image: "gcr.io/evmchain/taiko-geth:taiko",
		Cmd: []string{
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
	}
	hc := &container.HostConfig{
		AutoRemove: true,
		Binds:      []string{fmt.Sprintf("%s:/host/jwt.hex", JwtSecretFile)},
		PortBindings: map[nat.Port][]nat.PortBinding{
			gethHttpNatPort: {
				{
					HostIP:   "0.0.0.0",
					HostPort: "0",
				},
			},
			natTcpPort(gethWSPort): {
				{
					HostIP:   "0.0.0.0",
					HostPort: "0",
				},
			},
			gethAuthNatPort: {
				{
					HostIP:   "0.0.0.0",
					HostPort: "0",
				},
			},
		},
	}
	c, err := docker.NewReadyContainer(name, cc, hc, "HTTP server started")
	if err != nil {
		return nil, err
	}
	if err := c.Start(context.Background()); err != nil {
		return nil, err
	}
	return &gethContainer{
		ReadyContainer: c,
	}, nil
}

func delExistedBaseContainer(ctx context.Context) error {
	cli, err := client.NewClientWithOpts()
	if err != nil {
		return err
	}
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		return err
	}
	for _, c := range containers {
		for _, n := range c.Names {
			if n[1:] == baseContainerName {
				if err := cli.ContainerRemove(ctx, c.ID, types.ContainerRemoveOptions{Force: true}); err != nil {
					return err
				}
				break
			}
		}
	}
	return nil
}

func startBaseContainer(ctx context.Context) (err error) {
	if l1BaseContainer.delExisted {
		if err := delExistedBaseContainer(ctx); err != nil {
			return err
		}
	}
	l1BaseContainer.gethContainer, err = newAnvilContainer(ctx, true, baseContainerName)
	if err != nil {
		return err
	}
	return deployTaikoL1(l1BaseContainer.HttpEndpoint())
}

func newAnvilContainer(ctx context.Context, isBase bool, name string) (*gethContainer, error) {
	cc := &container.Config{
		Image: "ghcr.io/foundry-rs/foundry:latest",
		ExposedPorts: map[nat.Port]struct{}{
			gethHttpNatPort: {},
		},
		Entrypoint: []string{"anvil", "--host", "0.0.0.0"},
	}
	if !isBase {
		cc.Entrypoint = append(cc.Entrypoint, "--fork-url", l1BaseContainer.InnerHttpEndpoint())
	}
	hc := &container.HostConfig{
		AutoRemove: true,
		PortBindings: map[nat.Port][]nat.PortBinding{
			gethHttpNatPort: {
				{
					HostIP:   "0.0.0.0",
					HostPort: "0",
				},
				{
					HostIP:   "0.0.0.0",
					HostPort: "0",
				},
			},
		},
	}
	c, err := docker.NewReadyContainer(name, cc, hc, "Listening on 0.0.0.0:8545")
	if err != nil {
		return nil, err
	}
	if err := c.Start(ctx); err != nil {
		return nil, err
	}

	return &gethContainer{ReadyContainer: c, isAnvil: true}, nil
}

func getL2GenesisHash() (string, error) {
	c, err := newL2Container("genesis")
	if err != nil {
		return "", err
	}
	ctx := context.Background()
	cli, err := ethclient.DialContext(ctx, c.HttpEndpoint())
	if err != nil {
		return "", err
	}
	defer cli.Close()
	genesis, err := cli.BlockByNumber(ctx, big.NewInt(0))
	if err != nil {
		return "", err
	}
	if err := c.Stop(); err != nil {
		fmt.Printf("Can not stop genesis container: %v", err)
	}
	return genesis.Hash().String(), nil
}

func deployTaikoL1(endpoint string) error {
	l2GenesisHash, err := getL2GenesisHash()
	if err != nil {
		return err
	}
	cmd := exec.Command("forge",
		"script",
		"script/DeployOnL1.s.sol:DeployOnL1",
		"--fork-url",
		endpoint,
		"--broadcast",
		"--ffi",
		"-vvvvv",
		"--block-gas-limit",
		"100000000",
	)

	cmd.Env = []string{
		fmt.Sprintf("PRIVATE_KEY=%s", proposerPrivKey),
		fmt.Sprintf("ORACLE_PROVER=%s", oracleProverAddress.Hex()),
		"OWNER=0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC",
		fmt.Sprintf("TAIKO_L2_ADDRESS=%s", taikoL2Address.Hex()),
		"L2_SIGNAL_SERVICE=0x1000777700000000000000000000000000000007",
		"SHARED_SIGNAL_SERVICE=0x0000000000000000000000000000000000000000",
		fmt.Sprintf("TAIKO_TOKEN_PREMINT_RECIPIENTS=%s,%s", TestAddr.Hex(), oracleProverAddress.Hex()),
		fmt.Sprintf("TAIKO_TOKEN_PREMINT_AMOUNTS=%s,%s", premintTokenAmount, premintTokenAmount),
		fmt.Sprintf("L2_GENESIS_HASH=%s", l2GenesisHash),
	}
	cmd.Dir = monoPath + "/packages/protocol"
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("out=%s,err=%w", string(out), err)
	}
	if showDeployLog {
		scanner := bufio.NewScanner(strings.NewReader(string(out)))
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}
	return taikoL1DeployAddress()
}

func taikoL1DeployAddress() error {
	data, err := os.ReadFile(monoPath + "/packages/protocol/deployments/deploy_l1.json")
	if err != nil {
		return err
	}
	v := struct {
		TaikoL1    string `json:"taiko"`
		TaikoToken string `json:"taiko_token"`
	}{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	TaikoL1Address = common.HexToAddress(v.TaikoL1)
	TaikoTokenAddress = common.HexToAddress(v.TaikoToken)
	return nil
}

func initJwtSecret() (err error) {
	path := os.Getenv("JWT_SECRET")
	if path == "" {
		path = "../integration_test/nodes/jwt.hex"
	}
	JwtSecretFile, err = filepath.Abs(path)
	if err != nil {
		return err
	}
	if jwtSecret, err = jwt.ParseSecretFromFile(path); err != nil {
		return err
	}
	return nil
}

func initMonoPath() (err error) {
	path := os.Getenv("TAIKO_MONO")
	if path == "" {
		path = "../../taiko-mono/"
	}
	monoPath, err = filepath.Abs(path)
	if err != nil {
		return err
	}
	return nil
}
