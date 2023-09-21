package main

import (
	"github.com/taikoxyz/taiko-client/driver"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	"github.com/taikoxyz/taiko-client/proposer"
	"github.com/taikoxyz/taiko-client/prover"
)

var (
	endpointConf = &rpc.ClientConfig{}
	driverConf   = &driver.Config{}
	proposerConf = &proposer.Config{}
	proverConf   = &prover.Config{}
)
