package server

import (
	"crypto/rand"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
)

func (s *ProverServerTestSuite) TestGetStatusSuccess() {
	resp := s.sendReq("/status")
	s.Equal(http.StatusOK, resp.StatusCode)

	status := new(Status)

	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	s.Nil(err)
	s.Nil(json.Unmarshal(b, &status))

	s.Equal(s.ps.minProofFee.Uint64(), status.MinProofFee)
	s.Equal(uint64(s.ps.maxExpiry.Seconds()), status.MaxExpiry)
	s.Greater(status.CurrentCapacity, uint64(0))
}

func (s *ProverServerTestSuite) TestProposeBlockSuccess() {
	data, err := json.Marshal(CreateAssignmentRequestBody{
		FeeToken:   (common.Address{}),
		TierFees:   []*encoding.TierFee{{Tier: 0, Fee: common.Big256}},
		Expiry:     uint64(time.Now().Add(time.Minute).Unix()),
		TxListHash: common.BigToHash(common.Big1),
	})
	s.Nil(err)
	resp, err := http.Post(s.ws.URL+"/assignment", "application/json", strings.NewReader(string(data)))
	s.Nil(err)
	s.Equal(http.StatusOK, resp.StatusCode)
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	s.Nil(err)
	s.Contains(string(b), "signedPayload")
}

// randomHash generates a random blob of data and returns it as a hash.
func randomHash() common.Hash {
	var hash common.Hash
	if n, err := rand.Read(hash[:]); n != common.HashLength || err != nil {
		panic(err)
	}
	return hash
}
