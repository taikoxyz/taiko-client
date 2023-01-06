package submitter

import (
	"context"
	"errors"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/suite"
	"github.com/taikoxyz/taiko-client/bindings"
	proofProducer "github.com/taikoxyz/taiko-client/prover/proof_producer"
	"github.com/taikoxyz/taiko-client/testutils"
)

type ProofSubmitterTestSuite struct {
	testutils.ClientTestSuite
	validProofSubmitter   *ValidProofSubmitter
	invalidProofSubmitter *InvalidProofSubmitter
	validProofCh          chan *proofProducer.ProofWithHeader
	invalidProofCh        chan *proofProducer.ProofWithHeader
}

func (s *ProofSubmitterTestSuite) SetupTest() {
	s.ClientTestSuite.SetupTest()

	l1ProverPrivKey, err := crypto.ToECDSA(common.Hex2Bytes(os.Getenv("L1_PROVER_PRIVATE_KEY")))
	s.Nil(err)

	s.validProofSubmitter = NewValidProofSubmitter(
		s.RpcClient,
		&proofProducer.DummyProofProducer{},
		s.validProofCh,
		common.HexToAddress(os.Getenv("TAIKO_L2_ADDRESS")),
		l1ProverPrivKey,
		1,
		&sync.Mutex{},
	)

	s.invalidProofSubmitter = NewInvalidProofSubmitter(
		s.RpcClient,
		&proofProducer.DummyProofProducer{},
		s.invalidProofCh,
		l1ProverPrivKey,
		1,
		100000,
		&sync.Mutex{},
	)
}

func (s *ProofSubmitterTestSuite) TestValidProofSubmitterRequestProof() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	s.ErrorContains(
		s.validProofSubmitter.RequestProof(
			ctx, &bindings.TaikoL1ClientBlockProposed{Id: common.Big256}), "context deadline exceeded",
	)
}

func (s *ProofSubmitterTestSuite) TestValidProofSubmitterSubmitProofMetadataNotFound() {
	s.Error(
		s.validProofSubmitter.SubmitProof(
			context.Background(), &proofProducer.ProofWithHeader{
				BlockID: common.Big256,
				Meta:    &bindings.TaikoDataBlockMetadata{},
				Header:  &types.Header{},
				ZkProof: []byte{0xff},
			},
		),
	)
}

func (s *ProofSubmitterTestSuite) TestIsSubmitProofTxErrorRetryable() {
	s.True(isSubmitProofTxErrorRetryable(errors.New(testAddr.String()), common.Big0))
	s.False(isSubmitProofTxErrorRetryable(errors.New("L1:proof:tooMany"), common.Big0))
	s.False(isSubmitProofTxErrorRetryable(errors.New("L1:tooLate"), common.Big0))
	s.False(isSubmitProofTxErrorRetryable(errors.New("L1:prover:dup"), common.Big0))
	s.False(isSubmitProofTxErrorRetryable(errors.New("L1:"+testAddr.String()), common.Big0))
}

func TestProofSubmitterTestSuite(t *testing.T) {
	suite.Run(t, new(ProofSubmitterTestSuite))
}
