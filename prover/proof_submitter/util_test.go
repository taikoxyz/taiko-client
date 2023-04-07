package submitter

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (s *ProofSubmitterTestSuite) TestIsSubmitProofTxErrorRetryable() {
	s.True(isSubmitProofTxErrorRetryable(errors.New(testAddr.String()), common.Big0, false))
	s.True(isSubmitProofTxErrorRetryable(errors.New("L1_CANNOT_BE_FIRST_PROVER"), common.Big0, false))
	s.False(isSubmitProofTxErrorRetryable(errors.New("L1_DUP_PROVERS"), common.Big0, false))
	s.False(isSubmitProofTxErrorRetryable(errors.New("L1_"+testAddr.String()), common.Big0, false))
}

func (s *ProofSubmitterTestSuite) TestGetProveBlocksTxOpts() {
	optsL1, err := getProveBlocksTxOpts(context.Background(), s.RpcClient.L1, s.RpcClient.L1ChainID, s.TestAddrPrivKey)
	s.Nil(err)
	s.Greater(optsL1.GasTipCap.Uint64(), uint64(0))

	optsL2, err := getProveBlocksTxOpts(context.Background(), s.RpcClient.L2, s.RpcClient.L2ChainID, s.TestAddrPrivKey)
	s.Nil(err)
	s.Greater(optsL2.GasTipCap.Uint64(), uint64(0))
}

func (s *ProofSubmitterTestSuite) TestSendTxWithBackoff() {
	err := sendTxWithBackoff(context.Background(), s.RpcClient, common.Big1, func() (*types.Transaction, error) {
		return nil, errors.New("L1_TEST")
	}, false)

	s.NotNil(err)

	err = sendTxWithBackoff(context.Background(), s.RpcClient, common.Big1, func() (*types.Transaction, error) {
		height, err := s.RpcClient.L1.BlockNumber(context.Background())
		s.Nil(err)

		var block *types.Block
		for {
			block, err = s.RpcClient.L1.BlockByNumber(context.Background(), new(big.Int).SetUint64(height))
			s.Nil(err)
			if block.Transactions().Len() != 0 {
				break
			}
			height -= 1
		}

		return block.Transactions()[0], nil
	}, false)

	s.Nil(err)
}
