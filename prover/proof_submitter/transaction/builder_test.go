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
		&bindings.TaikoDataBlockMetadata{TxListByteOffset: common.Big1, TxListByteSize: common.Big256},
		&bindings.TaikoDataTransition{},
		&bindings.TaikoDataTierProof{},
		s.sender.innerSender.GetOpts(),
		false,
	)()
	s.NotNil(err)

	_, err = s.builder.Build(
		context.Background(),
		common.Big256,
		&bindings.TaikoDataBlockMetadata{TxListByteOffset: common.Big1, TxListByteSize: common.Big256},
		&bindings.TaikoDataTransition{},
		&bindings.TaikoDataTierProof{},
		s.sender.innerSender.GetOpts(),
		true,
	)()
	s.NotNil(err)
}
