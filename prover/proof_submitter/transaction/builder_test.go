package transaction

import (
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
	)(nil)
	s.NotNil(err)

	_, err = s.builder.Build(
		common.Big256,
		&bindings.TaikoDataBlockMetadata{TxListByteOffset: common.Big1, TxListByteSize: common.Big256},
		&bindings.TaikoDataTransition{},
		&bindings.TaikoDataTierProof{},
		true,
	)(nil)
	s.NotNil(err)
}
