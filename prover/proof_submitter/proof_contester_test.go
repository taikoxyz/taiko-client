package submitter

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/testutils"
)

func (s *ProofSubmitterTestSuite) TestSubmitContestNoTransition() {
	s.NotNil(
		s.contester.SubmitContest(
			context.Background(),
			&bindings.TaikoL1ClientBlockProposed{},
			&bindings.TaikoL1ClientTransitionProved{
				BlockId:    common.Big256,
				ParentHash: testutils.RandomHash(),
			},
		),
	)
}
