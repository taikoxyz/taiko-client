package testutils

import (
	"bufio"
	"context"
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/taikoxyz/taiko-client/testutils/docker"
)

const (
	gethHttpPort      uint64 = 8545
	gethWSPort        uint64 = 8546
	gethAuthPort      uint64 = 8551
	gethDiscoveryPort uint64 = 30303
	showDeployLog            = false
)

var (
	gethHttpNatPort         = natTcpPort(gethHttpPort)
	gethWSNatPort           = natTcpPort(gethWSPort)
	gethAuthNatPort         = natTcpPort(gethAuthPort)
	gethDiscoveryNatPort    = natTcpPort(gethDiscoveryPort)
	gethDiscoveryUdpNatPort = natUdpPort(gethDiscoveryPort)
)

type gethContainer struct {
	*docker.ReadyContainer
	isAnvil              bool
	TaikoL1Address       common.Address
	TaikoL1TokenAddress  common.Address
	TaikoL1SignalService common.Address
}

func natTcpPort(p uint64) nat.Port {
	return nat.Port(fmt.Sprintf("%d/tcp", p))
}

func natUdpPort(p uint64) nat.Port {
	return nat.Port(fmt.Sprintf("%d/udp", p))
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

func (e *gethContainer) deployTaikoL1() error {
	l2GenesisHash, err := e.getL2GenesisHash()
	if err != nil {
		return err
	}
	cmd := exec.Command("forge",
		"script",
		"script/DeployOnL1.s.sol:DeployOnL1",
		"--fork-url",
		e.HttpEndpoint(),
		"--broadcast",
		"--ffi",
		"-vvvvv",
		"--block-gas-limit",
		"100000000",
	)

	cmd.Env = []string{
		fmt.Sprintf("PRIVATE_KEY=%s", ProposerPrivateKey),
		fmt.Sprintf("ORACLE_PROVER=%s", OracleProverAddress.Hex()),
		fmt.Sprintf("OWNER=%s", ownerAddress),
		fmt.Sprintf("TAIKO_L2_ADDRESS=%s", TaikoL2Address.Hex()),
		fmt.Sprintf("L2_SIGNAL_SERVICE=%s", l2SignalService.Hex()),
		fmt.Sprintf("SHARED_SIGNAL_SERVICE=%s", sharedSignalService.Hex()),
		fmt.Sprintf("TAIKO_TOKEN_PREMINT_RECIPIENTS=%s,%s", ProposerAddress.Hex(), OracleProverAddress.Hex()),
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
	if err := e.initL1ContractAddress(out); err != nil {
		return err
	}
	return nil
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
		ExposedPorts: map[nat.Port]struct{}{
			gethHttpNatPort: {},
			gethWSNatPort:   {},
			gethAuthNatPort: {},
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
			gethWSNatPort: {
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
			gethDiscoveryNatPort: {
				{
					HostIP:   "0.0.0.0",
					HostPort: "0",
				},
			},
			gethDiscoveryUdpNatPort: {
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

func newAnvilContainer(ctx context.Context, isBase bool, name string) (*gethContainer, error) {
	cc := &container.Config{
		Image: "ghcr.io/foundry-rs/foundry:latest",
		ExposedPorts: map[nat.Port]struct{}{
			gethHttpNatPort: {},
		},
		Entrypoint: []string{"anvil", "--host", "0.0.0.0"},
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
	gc := &gethContainer{ReadyContainer: c, isAnvil: true}
	if err := gc.deployTaikoL1(); err != nil {
		return nil, err
	}
	if err := ensureProverBalance(gc); err != nil {
		return nil, err
	}
	return gc, nil
}

func (gc *gethContainer) getL2GenesisHash() (string, error) {
	c, err := newL2Container("genesis_" + gc.Name)
	if err != nil {
		return "", err
	}
	ctx := context.Background()
	cli, err := ethclient.DialContext(ctx, c.HttpEndpoint())
	if err != nil {
		return "", err
	}
	defer cli.Close()
	genesis, err := cli.HeaderByNumber(ctx, common.Big0)
	if err != nil {
		return "", err
	}
	if err := c.Stop(); err != nil {
		log.Warn("Can not stop genesis container: %v", err)
	}

	return genesis.Hash().Hex(), nil
}

func (gc *gethContainer) initL1ContractAddress(output []byte) error {
	re := regexp.MustCompile(`(\w+) \((\w+)\) -> (\w+)`)
	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		text := scanner.Text()
		matches := re.FindStringSubmatch(text)
		if len(matches) != 4 {
			continue
		}
		switch {
		case matches[1] == "taiko" && matches[2] == "proxy":
			gc.TaikoL1Address = common.HexToAddress(matches[3])
		case matches[1] == "taiko_token" && matches[2] == "proxy":
			gc.TaikoL1TokenAddress = common.HexToAddress(matches[3])
		case matches[1] == "signal_service" && matches[2] == "proxy":
			gc.TaikoL1SignalService = common.HexToAddress(matches[3])
		default:
			continue
		}
	}
	log.Info("Init", "taikoL1 address", gc.TaikoL1Address.Hex())
	log.Info("Init", "taikoL1Token address", gc.TaikoL1TokenAddress.Hex())
	log.Info("Init", "taikoL1SignalService address", gc.TaikoL1SignalService.Hex())
	return nil
}
