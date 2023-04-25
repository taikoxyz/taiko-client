package submitter

import (
	"os"
	"sync"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/suite"
	"github.com/taikoxyz/taiko-client/driver/chain_syncer/calldata"
	"github.com/taikoxyz/taiko-client/proposer"
	proofProducer "github.com/taikoxyz/taiko-client/prover/proof_producer"
	"github.com/taikoxyz/taiko-client/testutils"
)

type OracleProofSubmitterTestSuite struct {
	testutils.ClientTestSuite
	oracleProofSubmitter *OracleProofSubmitter
	calldataSyncer       *calldata.Syncer
	proposer             *proposer.Proposer
	validProofCh         chan *proofProducer.ProofWithHeader
	invalidProofCh       chan *proofProducer.ProofWithHeader
}

func (s *OracleProofSubmitterTestSuite) SetupTest() {
	s.ClientTestSuite.SetupTest()

	l1ProverPrivKey, err := crypto.ToECDSA(common.Hex2Bytes("1acb95df9ff6e93035ca3b8afce58273ac880d7b8bcb8a26b0be5a84be3a879d"))
	s.Nil(err)

	s.validProofCh = make(chan *proofProducer.ProofWithHeader, 1024)
	s.invalidProofCh = make(chan *proofProducer.ProofWithHeader, 1024)

	s.oracleProofSubmitter, err = NewOracleProofSubmitter(
		s.RpcClient,
		s.validProofCh,
		common.HexToAddress(os.Getenv("TAIKO_L2_ADDRESS")),
		l1ProverPrivKey,
		&sync.Mutex{},
	)
	s.Nil(err)

}

func TestOracleProofSubmitterTestSuite(t *testing.T) {
	suite.Run(t, new(ProofSubmitterTestSuite))
}
