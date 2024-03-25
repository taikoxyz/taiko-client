package handler

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/taikoxyz/taiko-client/internal/testutils"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/taikoxyz/taiko-client/bindings"
)

func (s *EventHandlerTestSuite) TestBlockVerifiedHandle() {
	handler := &BlockVerifiedEventHandler{}
	id := testutils.RandomHash().Big().Uint64()
	s.NotPanics(func() {
		handler.Handle(&bindings.TaikoL1ClientBlockVerified{
			BlockId: testutils.RandomHash().Big(),
			Raw: types.Log{
				BlockHash:   testutils.RandomHash(),
				BlockNumber: id,
			},
		})
	})
}

func TestBlockVerifiedEventHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(EventHandlerTestSuite))
}
