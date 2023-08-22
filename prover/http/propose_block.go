package http

import (
	"net/http"

	"github.com/cyberhorsey/webutils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/labstack/echo/v4"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
)

type proposeBlockResp struct {
	SignedPayload []byte         `json:"signedPayload"`
	Prover        common.Address `json:"prover"`
}

// ProposeBlock handles a propose block request, decides if this prover wants to
// handle this block, and if so, returns a signed payload the proposer
// can submit onchain.
func (srv *Server) ProposeBlock(c echo.Context) error {
	r := &encoding.ProposeBlockData{}
	if err := c.Bind(r); err != nil {
		return webutils.LogAndRenderErrors(c, http.StatusUnprocessableEntity, err)
	}

	// TODO(jeff): logic to determine is prover wants this block.
	// check fee, check expiry, determine if its feasible/profitable.

	// TODO(jeff): check capacity

	encoded, err := encoding.EncodeProposeBlockData(r)
	if err != nil {
		return webutils.LogAndRenderErrors(c, http.StatusUnprocessableEntity, err)
	}

	hashed := crypto.Keccak256Hash(encoded)

	signed, err := crypto.Sign(hashed.Bytes(), srv.proverPrivateKey)
	if err != nil {
		return webutils.LogAndRenderErrors(c, http.StatusUnprocessableEntity, err)
	}

	resp := &proposeBlockResp{
		SignedPayload: signed,
		Prover:        srv.proverAddress,
	}

	return c.JSON(http.StatusOK, resp)
}
