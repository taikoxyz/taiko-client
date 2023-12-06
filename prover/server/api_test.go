package server

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
	"github.com/taikoxyz/taiko-client/prover/db"
)

func (s *ProverServerTestSuite) TestGetStatusSuccess() {
	res := s.sendReq("/status")
	s.Equal(http.StatusOK, res.StatusCode)

	status := new(Status)

	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	s.Nil(err)
	s.Nil(json.Unmarshal(b, &status))

	s.Equal(s.s.minOptimisticTierFee.Uint64(), status.MinOptimisticTierFee)
	s.Equal(s.s.minSgxTierFee.Uint64(), status.MinSgxTierFee)
	s.Equal(s.s.minSgxAndPseZkevmTierFee.Uint64(), status.MinSgxTierFee)
	s.Equal(uint64(s.s.maxExpiry.Seconds()), status.MaxExpiry)
	s.Greater(status.CurrentCapacity, uint64(0))
	s.NotEmpty(status.HeartBeatSignature)
	pubKey, err := crypto.SigToPub(crypto.Keccak256Hash([]byte("HEART_BEAT")).Bytes(), status.HeartBeatSignature)
	s.Nil(err)
	s.NotEmpty(status.Prover)
	s.Equal(status.Prover, crypto.PubkeyToAddress(*pubKey).Hex())
}

func (s *ProverServerTestSuite) TestProposeBlockSuccess() {
	data, err := json.Marshal(CreateAssignmentRequestBody{
		FeeToken: (common.Address{}),
		TierFees: []encoding.TierFee{
			{Tier: encoding.TierOptimisticID, Fee: common.Big256},
			{Tier: encoding.TierSgxID, Fee: common.Big256},
			{Tier: encoding.TierSgxAndPseZkevmID, Fee: common.Big256},
		},
		Expiry:     uint64(time.Now().Add(time.Minute).Unix()),
		TxListHash: common.BigToHash(common.Big1),
	})
	s.Nil(err)
	res, err := http.Post(s.testServer.URL+"/assignment", "application/json", strings.NewReader(string(data)))
	s.Nil(err)
	s.Equal(http.StatusOK, res.StatusCode)
	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	s.Nil(err)
	s.Contains(string(b), "signedPayload")
}

func (s *ProverServerTestSuite) TestGetSignedBlocks() {
	latest, err := s.s.rpc.L2.BlockByNumber(context.Background(), nil)
	s.Nil(err)

	signed, err := crypto.Sign(latest.Hash().Bytes(), s.s.proverPrivateKey)
	s.Nil(err)

	s.Nil(s.s.db.Put(db.BuildBlockKey(latest.Time()), signed))
	res := s.sendReq("/signedBlocks")
	s.Equal(http.StatusOK, res.StatusCode)

	signedBlocks := make([]SignedBlock, 0)

	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	s.Nil(err)
	s.Nil(json.Unmarshal(b, &signedBlocks))

	s.Equal(1, len(signedBlocks))
	s.Equal(latest.Hash().Hex(), signedBlocks[0].BlockHash)
	s.Equal(latest.Number().Uint64(), signedBlocks[0].BlockID)
	s.Equal(common.Bytes2Hex(signed), signedBlocks[0].Signature)
}
