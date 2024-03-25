package handler

import (
	"testing"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/suite"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/internal/testutils"
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
