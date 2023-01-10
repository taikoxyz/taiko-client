package submitter

import (
	"context"
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (s *ProofSubmitterTestSuite) TestIsSubmitProofTxErrorRetryable() {
	s.True(isSubmitProofTxErrorRetryable(errors.New(testAddr.String()), common.Big0))
	s.False(isSubmitProofTxErrorRetryable(errors.New("L1:proof:tooMany"), common.Big0))
	s.False(isSubmitProofTxErrorRetryable(errors.New("L1:tooLate"), common.Big0))
	s.False(isSubmitProofTxErrorRetryable(errors.New("L1:prover:dup"), common.Big0))
	s.False(isSubmitProofTxErrorRetryable(errors.New("L1:"+testAddr.String()), common.Big0))
}

func (s *ProofSubmitterTestSuite) TestGetProveBlocksTxOpts() {
	optsL1, err := getProveBlocksTxOpts(context.Background(), s.RpcClient.L1, s.RpcClient.L1ChainID, s.TestAddrPrivKey)
	s.Nil(err)
	s.Greater(optsL1.GasTipCap.Uint64(), 0)

	optsL2, err := getProveBlocksTxOpts(context.Background(), s.RpcClient.L2, s.RpcClient.L2ChainID, s.TestAddrPrivKey)
	s.Nil(err)
	s.Greater(optsL2.GasTipCap.Uint64(), 0)
}

func (s *ProofSubmitterTestSuite) TestSendTxWithBackoff() {
	err := sendTxWithBackoff(context.Background(), s.RpcClient, common.Big1, func() (*types.Transaction, error) {
		return nil, errors.New("L1:test")
	})

	s.NotNil(err)

	err = sendTxWithBackoff(context.Background(), s.RpcClient, common.Big1, func() (*types.Transaction, error) {
		block, err := s.RpcClient.L1.BlockByNumber(context.Background(), nil)
		s.Nil(err)
		s.NotEmpty(block.Transactions())

		return block.Transactions()[0], nil
	})

	s.Nil(err)
}
