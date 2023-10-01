package rpc

import (
	"context"

	"github.com/ethereum/go-ethereum/beacon/engine"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func (s *RpcTestSuite) TestL2EngineForbidden() {
	_, err := s.cli.L2Engine.ForkchoiceUpdate(
		context.Background(),
		&engine.ForkchoiceStateV1{},
		&engine.PayloadAttributes{},
	)
	s.ErrorContains(err, "Unauthorized")

	_, err = s.cli.L2Engine.NewPayload(
		context.Background(),
		&engine.ExecutableData{},
	)
	s.ErrorContains(err, "Unauthorized")

	_, err = s.cli.L2Engine.GetPayload(
		context.Background(),
		&engine.PayloadID{},
	)
	s.ErrorContains(err, "Unauthorized")

	_, err = s.cli.L2Engine.ExchangeTransitionConfiguration(
		context.Background(),
		&engine.TransitionConfigurationV1{
			TerminalTotalDifficulty: (*hexutil.Big)(common.Big0),
			TerminalBlockHash:       common.Hash{},
			TerminalBlockNumber:     0,
		})
	s.ErrorContains(err, "Unauthorized")
}
