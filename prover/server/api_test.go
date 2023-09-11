package server

import (
	"crypto/rand"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/cyberhorsey/webutils/testutils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo/v4"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
)

func (s *ProverServerTestSuite) TestGetStatusSuccess() {
	rec := s.sendReq("/status")
	s.Equal(http.StatusOK, rec.Code)

	status := new(Status)
	s.Nil(json.Unmarshal(rec.Body.Bytes(), &status))

	s.Equal(s.srv.minProofFee.Uint64(), status.MinProofFee)
	s.Equal(uint64(s.srv.maxExpiry.Seconds()), status.MaxExpiry)
	s.Greater(status.CurrentCapacity, uint64(0))
}

func (s *ProverServerTestSuite) TestProposeBlockSuccess() {
	rec := httptest.NewRecorder()

	s.srv.ServeHTTP(rec, testutils.NewUnauthenticatedRequest(
		echo.POST,
		"/assignment",
		&encoding.ProposeBlockData{
			Fee:    common.Big256,
			Expiry: uint64(time.Now().Add(time.Minute).Unix()),
			Input: encoding.TaikoL1BlockMetadataInput{
				Proposer:        common.BytesToAddress(randomHash().Bytes()),
				TxListHash:      randomHash(),
				TxListByteStart: common.Big0,
				TxListByteEnd:   common.Big0,
				CacheTxListInfo: false,
			},
		},
	))

	testutils.AssertStatusAndBody(s.T(), rec, http.StatusOK, []string{"signedPayload"})
}

// randomHash generates a random blob of data and returns it as a hash.
func randomHash() common.Hash {
	var hash common.Hash
	if n, err := rand.Read(hash[:]); n != common.HashLength || err != nil {
		panic(err)
	}
	return hash
}
