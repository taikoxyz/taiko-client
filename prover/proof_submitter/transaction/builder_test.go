package transaction

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/testutils"
)

func (s *TransactionTestSuite) TestGetProveBlocksTxOpts() {
	optsL1, err := getProveBlocksTxOpts(context.Background(), s.RpcClient.L1, s.RpcClient.L1ChainID, s.TestAddrPrivKey)
	s.Nil(err)
	s.Greater(optsL1.GasTipCap.Uint64(), uint64(0))

	optsL2, err := getProveBlocksTxOpts(context.Background(), s.RpcClient.L2, s.RpcClient.L2ChainID, s.TestAddrPrivKey)
	s.Nil(err)
	s.Greater(optsL2.GasTipCap.Uint64(), uint64(0))
}

func (s *TransactionTestSuite) TestBuildTxs() {
	s.NotNil(s.builder.BuildForNormalProofSubmission(context.Background(), common.Big256, testutils.RandomBytes(1024)))
	s.NotNil(s.builder.BuildForGuardianProofSubmission(
		context.Background(),
		common.Big256,
		&bindings.TaikoDataBlockEvidence{},
	))
}
