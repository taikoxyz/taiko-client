package transaction

import (
	"context"

	"github.com/ethereum/go-ethereum/common"

	"github.com/taikoxyz/taiko-client/bindings"
)

func (s *TransactionTestSuite) TestBuildTxs() {
	_, err := s.builder.Build(
		context.Background(),
		common.Big256,
		&bindings.TaikoDataBlockMetadata{},
		&bindings.TaikoDataTransition{},
		&bindings.TaikoDataTierProof{},
		s.sender.innerSender.Opts,
		false,
	)()
	s.NotNil(err)

	_, err = s.builder.Build(
		context.Background(),
		common.Big256,
		&bindings.TaikoDataBlockMetadata{},
		&bindings.TaikoDataTransition{},
		&bindings.TaikoDataTierProof{},
		s.sender.innerSender.Opts,
		true,
	)()
	s.NotNil(err)
}
