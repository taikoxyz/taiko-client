package server

import (
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/labstack/echo/v4"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
)

// Status represents the current prover server status.
type Status struct {
	MinProofFee     uint64 `json:"minProofFee"`
	MaxExpiry       uint64 `json:"maxExpiry"`
	CurrentCapacity uint64 `json:"currentCapacity"`
}

// GetStatus handles a query to the current prover server status.
//
//	@Summary		Get current prover server status
//	@ID			   	get-status
//	@Accept			json
//	@Produce		json
//	@Success		200	{object} Status
//	@Router			/status [get]
func (srv *ProverServer) GetStatus(c echo.Context) error {
	return c.JSON(http.StatusOK, &Status{
		MinProofFee:     srv.minProofFee.Uint64(),
		MaxExpiry:       uint64(srv.maxExpiry.Seconds()),
		CurrentCapacity: srv.capacityManager.ReadCapacity(),
	})
}

// ProposeBlockResponse represents the JSON response which will be returned by
// the ProposeBlock request handler.
type ProposeBlockResponse struct {
	SignedPayload []byte         `json:"signedPayload"`
	Prover        common.Address `json:"prover"`
}

// CreateAssignment handles a block proof assignment request, decides if this prover wants to
// handle this block, and if so, returns a signed payload the proposer
// can submit onchain.
//
//	@Summary		Try to accept a block proof assignment
//	@ID			   	create-assignment
//	@Accept			json
//	@Produce		json
//	@Success		200		{object} ProposeBlockResponse
//	@Failure		422		{string} string	"proof fee too low"
//	@Failure		422		{string} string "expiry too long"
//	@Failure		422		{string} string "prover does not have capacity"
//	@Router			/assignment [post]
func (srv *ProverServer) CreateAssignment(c echo.Context) error {
	req := new(encoding.ProposeBlockData)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	log.Info("Propose block data", "fee", req.Fee, "expiry", req.Expiry)

	if req.Fee.Cmp(srv.minProofFee) < 0 {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "proof fee too low")
	}

	if req.Expiry > uint64(time.Now().Add(srv.maxExpiry).Unix()) {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "expiry too long")
	}

	if srv.capacityManager.ReadCapacity() == 0 {
		log.Warn("Prover does not have capacity")
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "prover does not have capacity")
	}

	encoded, err := encoding.EncodeProposeBlockData(req)
	if err != nil {
		log.Error("Failed to encode proposeBlock data", "error", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err)
	}

	signed, err := crypto.Sign(crypto.Keccak256Hash(encoded).Bytes(), srv.proverPrivateKey)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, &ProposeBlockResponse{
		SignedPayload: signed,
		Prover:        srv.proverAddress,
	})
}
