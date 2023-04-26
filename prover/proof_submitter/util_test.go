package submitter

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"errors"
	"math/big"
	"math/rand"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
)

func (s *ProofSubmitterTestSuite) TestIsSubmitProofTxErrorRetryable() {
	s.True(isSubmitProofTxErrorRetryable(errors.New(testAddr.String()), common.Big0))
	s.True(isSubmitProofTxErrorRetryable(errors.New("L1_NOT_ORACLE_PROVER"), common.Big0))
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
	err := sendTxWithBackoff(context.Background(), s.RpcClient, common.Big1, func() (*types.Transaction, error) {
		return nil, errors.New("L1_TEST")
	})

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
	})

	s.Nil(err)
}

// randomHash generates a random blob of data and returns it as a hash.
func randomHash() common.Hash {
	var hash common.Hash
	if n, err := rand.Read(hash[:]); n != common.HashLength || err != nil {
		panic(err)
	}
	return hash
}

func (s *ProofSubmitterTestSuite) TestHashAndSignEvidenceForOracleProof() {
	evidence := &encoding.TaikoL1Evidence{
		MetaHash:      randomHash(),
		BlockHash:     randomHash(),
		ParentHash:    randomHash(),
		SignalRoot:    randomHash(),
		Graffiti:      randomHash(),
		Prover:        common.BigToAddress(new(big.Int).SetUint64(rand.Uint64())),
		ParentGasUsed: 1024,
		GasUsed:       1024,
		VerifierId:    0,
		Proof:         nil,
	}

	privateKey, err := crypto.HexToECDSA(os.Getenv("L1_PROVER_PRIVATE_KEY"))
	s.Nil(err)

	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	s.True(ok)

	input, err := encoding.EncodeProveBlockInput(evidence)
	s.Nil(err)

	hash := crypto.Keccak256Hash(input)

	sig, _, err := hashSignAndSetEvidenceForOracleProof(evidence, privateKey)
	s.Nil(err)

	pubKey, err := crypto.Ecrecover(hash.Bytes(), sig)
	s.Nil(err)
	isValid := crypto.VerifySignature(pubKey, hash.Bytes(), sig[:64])
	s.True(isValid)

	s.Equal(elliptic.Marshal(publicKeyECDSA, publicKeyECDSA.X, publicKeyECDSA.Y), pubKey)
}
