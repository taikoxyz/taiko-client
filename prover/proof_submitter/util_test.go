package submitter

import (
	"context"
	"errors"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/testutils"
)

var (
	testKey, _          = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	testAddr            = crypto.PubkeyToAddress(testKey.PublicKey)
	testMaxRetry uint64 = 1
)

func (s *ProofSubmitterTestSuite) TestIsSubmitProofTxErrorRetryable() {
	s.True(isSubmitProofTxErrorRetryable(errors.New(testAddr.String()), common.Big0))
	s.False(isSubmitProofTxErrorRetryable(errors.New("L1_NOT_SPECIAL_PROVER"), common.Big0))
	s.False(isSubmitProofTxErrorRetryable(errors.New("L1_DUP_PROVERS"), common.Big0))
	s.False(isSubmitProofTxErrorRetryable(errors.New("L1_"+testAddr.String()), common.Big0))
}

func (s *ProofSubmitterTestSuite) TestGetProveBlocksTxOpts() {
	optsL1, err := getProveBlocksTxOpts(context.Background(),
		s.rpcClient.L1, s.rpcClient.L1ChainID, testutils.ProposerPrivKey)
	s.Nil(err)
	s.Greater(optsL1.GasTipCap.Uint64(), uint64(0))

	optsL2, err := getProveBlocksTxOpts(context.Background(),
		s.rpcClient.L2, s.rpcClient.L2ChainID, testutils.ProposerPrivKey)
	s.Nil(err)
	s.Greater(optsL2.GasTipCap.Uint64(), uint64(0))
}

func (s *ProofSubmitterTestSuite) TestSendTxWithBackoff() {
	l1Head, err := s.rpcClient.L1.HeaderByNumber(context.Background(), nil)
	s.Nil(err)
	l1HeadChild, err := s.rpcClient.L1.HeaderByNumber(context.Background(), new(big.Int).Sub(l1Head.Number, common.Big1))
	s.Nil(err)
	meta := &bindings.TaikoDataBlockMetadata{L1Height: l1HeadChild.Number.Uint64(), L1Hash: l1HeadChild.Hash()}
	s.NotNil(sendTxWithBackoff(
		context.Background(),
		s.rpcClient,
		common.Big1,
		l1Head.Hash(),
		0,
		meta,
		func(nonce *big.Int) (*types.Transaction, error) { return nil, errors.New("L1_TEST") },
		12*time.Second,
		&testMaxRetry,
		5*time.Second,
	))

	s.Nil(sendTxWithBackoff(
		context.Background(),
		s.rpcClient,
		common.Big1,
		l1Head.Hash(),
		0,
		meta,
		func(nonce *big.Int) (*types.Transaction, error) {
			height, err := s.rpcClient.L1.BlockNumber(context.Background())
			s.Nil(err)

			var block *types.Block
			for {
				block, err = s.rpcClient.L1.BlockByNumber(context.Background(), new(big.Int).SetUint64(height))
				s.Nil(err)
				if block.Transactions().Len() != 0 {
					break
				}
				height -= 1
			}

			return block.Transactions()[0], nil
		},
		12*time.Second,
		&testMaxRetry,
		5*time.Second,
	))
}
