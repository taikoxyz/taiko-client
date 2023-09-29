package server

import (
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/labstack/echo/v4"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
)

// CreateAssignmentRequestBody represents a request body when handling assignment creation request.
type CreateAssignmentRequestBody struct {
	FeeToken   common.Address
	TierFees   []*encoding.TierFee
	Expiry     uint64
	TxListHash common.Hash
}

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
//	@Failure		422		{string} string	"invalid txList hash"
//	@Failure		422		{string} string	"only receive ETH"
//	@Failure		422		{string} string	"insufficient prover balance"
//	@Failure		422		{string} string	"proof fee too low"
//	@Failure		422		{string} string "expiry too long"
//	@Failure		422		{string} string "prover does not have capacity"
//	@Router			/assignment [post]
func (srv *ProverServer) CreateAssignment(c echo.Context) error {
	req := new(CreateAssignmentRequestBody)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	log.Info(
		"Propose block data",
		"feeToken", req.FeeToken,
		"expiry", req.Expiry,
		"tierFees", req.TierFees,
		"txListHash", req.TxListHash,
	)

	if req.TxListHash == (common.Hash{}) {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "invalid txList hash")
	}

	if req.FeeToken != (common.Address{}) {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "only receive ETH")
	}

	if !srv.isOracle {
		ok, err := rpc.CheckProverBalance(c.Request().Context(), srv.rpc, srv.proverAddress, srv.taikoL1Address, srv.bond)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		if !ok {
			log.Warn(
				"Insufficient prover balance, please get more tokens or wait for verification of the blocks you proved",
				"prover", srv.proverAddress,
			)
			return echo.NewHTTPError(http.StatusUnprocessableEntity, "insufficient prover balance")
		}
	}

	for _, tier := range req.TierFees {
		if tier.Fee.Cmp(srv.minProofFee) < 0 {
			log.Warn(
				"Proof fee too low",
				"tier", tier.Tier,
				"fee", tier.Fee,
				"srvMinProofFee", srv.minProofFee.String(),
				"proposerIP", c.RealIP(),
			)
			return echo.NewHTTPError(http.StatusUnprocessableEntity, "proof fee too low")
		}
	}

	if req.Expiry > uint64(time.Now().Add(srv.maxExpiry).Unix()) {
		log.Warn(
			"Expiry too long",
			"requestExpiry", req.Expiry,
			"srvMaxExpiry", srv.maxExpiry,
			"proposerIP", c.RealIP(),
		)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "expiry too long")
	}

	if _, ok := srv.capacityManager.TakeOneTempCapacity(); !ok {
		log.Warn("Prover unable to take a temporary capacity", "proposerIP", c.RealIP())
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "prover does not have capacity")
	}

	encoded, err := encoding.EncodeProverAssignmentPayload(req.TxListHash, req.FeeToken, req.Expiry, req.TierFees)
	if err != nil {
		log.Error("Failed to encode proverAssignment payload data", "error", err)
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
