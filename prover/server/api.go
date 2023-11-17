package server

import (
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/labstack/echo/v4"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	"github.com/taikoxyz/taiko-client/prover/db"
)

// @title Taiko Prover Server API
// @version 1.0
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url https://community.taiko.xyz/
// @contact.email info@taiko.xyz

// @license.name MIT
// @license.url https://github.com/taikoxyz/taiko-client/blob/main/LICENSE.md

// CreateAssignmentRequestBody represents a request body when handling assignment creation request.
type CreateAssignmentRequestBody struct {
	FeeToken   common.Address
	TierFees   []encoding.TierFee
	Expiry     uint64
	TxListHash common.Hash
}

// Status represents the current prover server status.
type Status struct {
	MinOptimisticTierFee uint64 `json:"minOptimisticTierFee"`
	MinSgxTierFee        uint64 `json:"minSgxTierFee"`
	MinPseZkevmTierFee   uint64 `json:"minPseZkevmTierFee"`
	MaxExpiry            uint64 `json:"maxExpiry"`
	CurrentCapacity      uint64 `json:"currentCapacity"`
	Prover               string `json:"prover"`
	HeartBeatSignature   []byte `json:"heartBeatSignature"`
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
	sig, err := crypto.Sign(crypto.Keccak256Hash([]byte("HEART_BEAT")).Bytes(), srv.proverPrivateKey)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, &Status{
		MinOptimisticTierFee: srv.minOptimisticTierFee.Uint64(),
		MinSgxTierFee:        srv.minSgxTierFee.Uint64(),
		MinPseZkevmTierFee:   srv.minPseZkevmTierFee.Uint64(),
		MaxExpiry:            uint64(srv.maxExpiry.Seconds()),
		CurrentCapacity:      srv.capacityManager.ReadCapacity(),
		Prover:               srv.proverAddress.Hex(),
		HeartBeatSignature:   sig,
	})
}

// ProposeBlockResponse represents the JSON response which will be returned by
// the ProposeBlock request handler.
type ProposeBlockResponse struct {
	SignedPayload []byte         `json:"signedPayload"`
	Prover        common.Address `json:"prover"`
	MaxBlockID    uint64         `json:"maxBlockID"`
	MaxProposedIn uint64         `json:"maxProposedIn"`
}

// CreateAssignment handles a block proof assignment request, decides if this prover wants to
// handle this block, and if so, returns a signed payload the proposer
// can submit onchain.
//
//	@Summary		Try to accept a block proof assignment
//	@Param          body        body    CreateAssignmentRequestBody   true    "assignment request body"
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
		"Proof assignment request body",
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

	if !srv.isGuardian {
		ok, err := rpc.CheckProverBalance(
			c.Request().Context(),
			srv.rpc,
			srv.proverAddress,
			srv.taikoL1Address,
			srv.livenessBond,
		)
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
		if tier.Tier == encoding.TierGuardianID {
			continue
		}

		var minTierFee *big.Int
		switch tier.Tier {
		case encoding.TierOptimisticID:
			minTierFee = srv.minOptimisticTierFee
		case encoding.TierSgxID:
			minTierFee = srv.minSgxTierFee
		case encoding.TierPseZkevmID:
			minTierFee = srv.minPseZkevmTierFee
		case encoding.TierSgxAndPseZkevmID:
			minTierFee = srv.minSgxAndPseZkevmTierFee
		default:
			log.Warn("Unknown tier", "tier", tier.Tier, "fee", tier.Fee, "proposerIP", c.RealIP())
		}

		if tier.Fee.Cmp(minTierFee) < 0 {
			log.Warn(
				"Proof fee too low",
				"tier", tier.Tier,
				"fee", tier.Fee,
				"minTierFee", minTierFee,
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

	if ok := srv.capacityManager.HoldOneCapacity(time.Duration(req.Expiry) * time.Second); !ok {
		log.Warn("Prover unable to hold a capacity", "proposerIP", c.RealIP())
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "prover does not have capacity")
	}

	l1Head, err := srv.rpc.L1.BlockNumber(c.Request().Context())
	if err != nil {
		log.Error("Failed to get L1 block head", "error", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err)
	}

	encoded, err := encoding.EncodeProverAssignmentPayload(
		srv.taikoL1Address,
		req.TxListHash,
		req.FeeToken,
		req.Expiry,
		l1Head+srv.maxSlippage,
		srv.maxProposedIn,
		req.TierFees,
	)
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
		MaxBlockID:    l1Head + srv.maxSlippage,
		MaxProposedIn: srv.maxProposedIn,
	})
}

type SignedBlock struct {
	BlockID   uint64         `json:"blockID"`
	BlockHash string         `json:"blockHash"`
	Signature string         `json:"signature"`
	Prover    common.Address `json:"proverAddress"`
}

// GetSignedBlocks handles a query to retrieve the most recent signed blocks from the database.
//
//	@Summary		Get signed blocks
//	@ID			   	get-signed-blocks
//	@Accept			json
//	@Produce		json
//	@Success		200	{object} []SignedBlocks
//	@Router			/signedBlocks [get]
func (srv *ProverServer) GetSignedBlocks(c echo.Context) error {
	latestBlock, err := srv.rpc.L2.BlockByNumber(c.Request().Context(), nil)
	if err != nil {
		if err != nil {
			log.Error("Failed to get latest L2 block", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
	}

	var signedBlocks []SignedBlock

	// start iterator at 0
	start := big.NewInt(0)

	// if latestBlock is greater than the number of blocks to return, we only want to return
	// the most recent N blocks signed by this guardian prover.
	if latestBlock.NumberU64() > numBlocksToReturn.Uint64() {
		start = new(big.Int).Sub(latestBlock.Number(), numBlocksToReturn)
	}

	iter := srv.db.NewIterator([]byte(db.BlockKeyPrefix), start.Bytes())

	defer iter.Release()

	for iter.Next() {
		k := strings.Split(string(iter.Key()), "-")

		blockID, err := strconv.Atoi(k[1])

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		signedBlocks = append(signedBlocks, SignedBlock{
			BlockID:   uint64(blockID),
			BlockHash: latestBlock.Hash().Hex(),
			Signature: common.Bytes2Hex(iter.Value()),
			Prover:    srv.proverAddress,
		})
	}

	return c.JSON(http.StatusOK, signedBlocks)
}
