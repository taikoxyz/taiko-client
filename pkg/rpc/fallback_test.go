package rpc

import (
	"errors"
)

func (s *RpcTestSuite) TestIsMaxPriorityFeePerGasNotFoundError() {
	s.False(IsMaxPriorityFeePerGasNotFoundError(errors.New("test")))
	s.True(IsMaxPriorityFeePerGasNotFoundError(errMaxPriorityFeePerGasNotFound))
}
