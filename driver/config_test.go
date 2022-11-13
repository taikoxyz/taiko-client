package driver

import (
	"flag"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/taikochain/taiko-client/cmd/flags"
	"github.com/urfave/cli/v2"
)

func TestNewConfigFromCliContext(t *testing.T) {
	l1Endpoint := randomHash().Hex()
	fmt.Println("222", l1Endpoint)
	set := flag.NewFlagSet("TestNewConfigFromCliContext", flag.PanicOnError)
	flag.StringVar(&l1Endpoint, flags.L1NodeEndpoint.Name, l1Endpoint, "")

	_, err := NewConfigFromCliContext(cli.NewContext(cli.NewApp(), set, nil))

	require.Nil(t, err)
}
