package builder

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/taikoxyz/taiko-client/internal/testutils"
)

type TransactionBuilderTestSuite struct {
	testutils.ClientTestSuite
}

func (s *TransactionBuilderTestSuite) TestGetParentMetaHash() {
	metahash, err := getParentMetaHash(context.Background(), s.RPCClient)
	s.Nil(err)
	s.Empty(metahash)
}

func TestTransactionBuilderTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionBuilderTestSuite))
}
