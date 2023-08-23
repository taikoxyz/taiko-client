package http

import (
	"context"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
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
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	if r.Fee.Cmp(srv.minProofFee) < 0 {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "proof fee too low")
	}

	srv.requestCurrentCapacityCh <- struct{}{}

	ctx, cancel := context.WithTimeout(c.Request().Context(), 8*time.Second)
	defer cancel()

	for {
		select {
		case capacity := <-srv.receiveCurrentCapacityCh:
			if capacity == 0 {
				return echo.NewHTTPError(http.StatusUnprocessableEntity, "prover does not have capacity")
			}

			encoded, err := encoding.EncodeProposeBlockData(r)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnprocessableEntity, err)
			}

			hashed := crypto.Keccak256Hash(encoded)

			signed, err := crypto.Sign(hashed.Bytes(), srv.proverPrivateKey)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err)
			}

			resp := &proposeBlockResp{
				SignedPayload: signed,
				Prover:        srv.proverAddress,
			}

			return c.JSON(http.StatusOK, resp)
		case <-ctx.Done():
			log.Info("timed out trying to get capacity")
			return echo.NewHTTPError(http.StatusUnprocessableEntity, "timed out trying to get capacity")
		}
	}
}
