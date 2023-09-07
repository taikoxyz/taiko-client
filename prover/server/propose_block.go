package server

import (
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
)

// ProposeBlockResponse represents the JSON response which will be returned by
// the ProposeBlock request handler.
type ProposeBlockResponse struct {
	SignedPayload []byte         `json:"signedPayload"`
	Prover        common.Address `json:"prover"`
}

// ProposeBlock handles a propose block request, decides if this prover wants to
// handle this block, and if so, returns a signed payload the proposer
// can submit onchain.
func (srv *ProverServer) ProposeBlock(c echo.Context) error {
	res := new(encoding.ProposeBlockData)
	if err := c.Bind(res); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	if res.Fee.Cmp(srv.minProofFee) < 0 {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "proof fee too low")
	}

	if srv.capacityManager.ReadCapacity() == 0 {
		log.Warn("Prover does not have capacity")
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "prover does not have capacity")
	}

	encoded, err := encoding.EncodeProposeBlockData(res)
	if err != nil {
		log.Error("Failed to encode proposeBlock data", "error", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err)
	}

	signed, err := crypto.Sign(crypto.Keccak256Hash(encoded).Bytes(), srv.proverPrivateKey)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	resp := &ProposeBlockResponse{
		SignedPayload: signed,
		Prover:        srv.proverAddress,
	}

	return c.JSON(http.StatusOK, resp)
}
