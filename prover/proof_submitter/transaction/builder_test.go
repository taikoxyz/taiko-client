package transaction

import (
	"context"
	"github.com/ethereum/go-ethereum/common"

	"github.com/taikoxyz/taiko-client/bindings"
)

func (s *TransactionTestSuite) TestBuildTxs() {
	_, err := s.builder.Build(
		common.Big256,
		&bindings.TaikoDataBlockMetadata{TxListByteOffset: common.Big1, TxListByteSize: common.Big256},
		&bindings.TaikoDataTransition{},
		&bindings.TaikoDataTierProof{},
		false,
	)(s.sender.innerSender.GetOpts(context.TODO()))
	s.NotNil(err)

	_, err = s.builder.Build(
		common.Big256,
		&bindings.TaikoDataBlockMetadata{TxListByteOffset: common.Big1, TxListByteSize: common.Big256},
		&bindings.TaikoDataTransition{},
		&bindings.TaikoDataTierProof{},
		true,
	)(s.sender.innerSender.GetOpts(context.TODO()))
	s.NotNil(err)
}
