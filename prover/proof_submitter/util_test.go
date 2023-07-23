package submitter

import (
	"context"
	"errors"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/taikoxyz/taiko-client/bindings"
)

func (s *ProofSubmitterTestSuite) TestIsSubmitProofTxErrorRetryable() {
	s.True(isSubmitProofTxErrorRetryable(errors.New(testAddr.String()), common.Big0))
	s.False(isSubmitProofTxErrorRetryable(errors.New("L1_NOT_SPECIAL_PROVER"), common.Big0))
	s.False(isSubmitProofTxErrorRetryable(errors.New("L1_DUP_PROVERS"), common.Big0))
	s.False(isSubmitProofTxErrorRetryable(errors.New("L1_"+testAddr.String()), common.Big0))
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
	l1Head, err := s.RpcClient.L1.HeaderByNumber(context.Background(), nil)
	s.Nil(err)
	l1HeadChild, err := s.RpcClient.L1.HeaderByNumber(context.Background(), new(big.Int).Sub(l1Head.Number, common.Big1))
	s.Nil(err)
	meta := &bindings.TaikoDataBlockMetadata{L1Height: l1HeadChild.Number.Uint64(), L1Hash: l1HeadChild.Hash()}
	s.NotNil(sendTxWithBackoff(
		context.Background(),
		s.RpcClient,
		common.Big1,
		l1Head.Hash(),
		0,
		0,
		meta,
		func() (*types.Transaction, error) { return nil, errors.New("L1_TEST") },
		12*time.Second,
	))

	s.Nil(sendTxWithBackoff(
		context.Background(),
		s.RpcClient,
		common.Big1,
		l1Head.Hash(),
		0,
		0,
		meta,
		func() (*types.Transaction, error) {
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
		},
		12*time.Second,
	))
}
