package transaction

import (
	"context"

	"github.com/ethereum/go-ethereum/common"

	"github.com/taikoxyz/taiko-client/bindings"
)

func (s *TransactionTestSuite) TestBuildTxs() {
	opts := s.sender.GetOpts()
	_, err := s.builder.Build(
		context.Background(),
		common.Big256,
		&bindings.TaikoDataBlockMetadata{TxListByteOffset: common.Big1, TxListByteSize: common.Big256},
		&bindings.TaikoDataTransition{},
		&bindings.TaikoDataTierProof{},
		opts,
		false,
	)()
	s.NotNil(err)

	_, err = s.builder.Build(
		context.Background(),
		common.Big256,
		&bindings.TaikoDataBlockMetadata{TxListByteOffset: common.Big1, TxListByteSize: common.Big256},
		&bindings.TaikoDataTransition{},
		&bindings.TaikoDataTierProof{},
		opts,
		true,
	)()
	s.NotNil(err)
}
