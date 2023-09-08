package server

import (
	"math/big"
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
	MinProofFee     *big.Int      `json:"minProofFee"`
	MaxExpiry       time.Duration `json:"maxExpiry"`
	CurrentCapacity uint64        `json:"currentCapacity"`
}

// GetStatus handles a query to the current prover server status.
func (srv *ProverServer) GetStatus(c echo.Context) error {
	return c.JSON(http.StatusOK, &Status{
		MinProofFee:     srv.minProofFee,
		MaxExpiry:       srv.maxExpiry,
		CurrentCapacity: srv.capacityManager.ReadCapacity(),
	})
}

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
