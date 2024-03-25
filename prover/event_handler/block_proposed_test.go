package handler

import (
	"context"
	"testing"
	"time"

	proofProducer "github.com/taikoxyz/taiko-client/prover/proof_producer"
	state "github.com/taikoxyz/taiko-client/prover/shared_state"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/suite"
	"github.com/taikoxyz/taiko-client/bindings"
)

func (s *EventHandlerTestSuite) TestBlockProposedHandle() {
	opts := &NewBlockProposedEventHandlerOps{
		SharedState:           &state.SharedState{},
		ProverAddress:         common.Address{},
		GenesisHeightL1:       9,
		RPC:                   s.RPCClient,
		ProofGenerationCh:     make(chan *proofProducer.ProofWithHeader),
		AssignmentExpiredCh:   make(chan *bindings.TaikoL1ClientBlockProposed),
		ProofSubmissionCh:     make(chan *proofProducer.ProofRequestBody),
		ProofContestCh:        make(chan *proofProducer.ContestRequestBody),
		BackOffRetryInterval:  1 * time.Minute,
		BackOffMaxRetrys:      5,
		ContesterMode:         true,
		ProveUnassignedBlocks: true,
	}
	handler := NewBlockProposedEventHandler(
		opts,
	)
	e := s.ProposeAndInsertValidBlock(s.proposer, s.d.ChainSyncer().CalldataSyncer())
	err := handler.Handle(context.Background(), e, func() {})
	s.Nil(err)
}

func TestBlockProposedEventHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(EventHandlerTestSuite))
}
