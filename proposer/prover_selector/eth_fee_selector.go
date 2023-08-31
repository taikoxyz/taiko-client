package selector

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"net/url"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
)

var (
	httpScheme              = "http"
	errEmptyProverEndpoints = errors.New("empty prover endpoints")
	errUnableToFindProver   = errors.New("unable to find prover")
)

type assignProverResponse struct {
	SignedPayload []byte         `json:"signedPayload"`
	Prover        common.Address `json:"prover"`
}

type ETHFeeSelector struct {
	protocolConfigs       *bindings.TaikoDataConfig
	rpc                   *rpc.Client
	taikoL1Address        common.Address
	feeBase               *big.Int
	feeIncreasePercentage *big.Int
	proverEndpoints       []*url.URL
	proposalFeeIterations uint64
	proposalExpiry        time.Duration
	requestTimeout        time.Duration
}

func NewETHFeeSelector(
	protocolConfigs *bindings.TaikoDataConfig,
	rpc *rpc.Client,
	taikoL1Address common.Address,
	feeBase *big.Int,
	feeIncreasePercentage *big.Int,
	proverEndpoints []*url.URL,
	proposalFeeIterations uint64,
	proposalExpiry time.Duration,
	requestTimeout time.Duration,
) (*ETHFeeSelector, error) {
	if len(proverEndpoints) == 0 {
		return nil, errEmptyProverEndpoints
	}

	for _, endpoint := range proverEndpoints {
		if endpoint.Scheme != httpScheme {
			return nil, fmt.Errorf("invalid prover endpoint %s", endpoint)
		}
	}

	return &ETHFeeSelector{
		protocolConfigs,
		rpc,
		taikoL1Address,
		feeBase,
		feeIncreasePercentage,
		proverEndpoints,
		proposalFeeIterations,
		proposalExpiry,
		requestTimeout,
	}, nil
}

func (s *ETHFeeSelector) ProverEndpoints() []*url.URL { return s.proverEndpoints }

func (s *ETHFeeSelector) AssignProver(
	ctx context.Context,
	meta *encoding.TaikoL1BlockMetadataInput,
) ([]byte, *big.Int, error) {
	// Iterate over each configured endpoint, and see if someone wants to accept this block.
	// If it is denied, we continue on to the next endpoint.
	// If we do not find a prover, we can increase the fee up to a point, or give up.
	for i := 0; i < int(s.proposalFeeIterations); i++ {
		var (
			fee    = s.feeBase
			expiry = uint64(time.Now().Add(s.proposalExpiry).Unix())
		)

		// Increase fee on each failed loop
		if i > 0 {
			cumulativePercent := new(big.Int).Mul(s.feeIncreasePercentage, big.NewInt(int64(i)))
			increase := new(big.Int).Mul(fee, cumulativePercent)
			increase.Div(increase, big.NewInt(100))
			fee.Add(fee, increase)
		}

		proposeBlockReq := &encoding.ProposeBlockData{
			Expiry: expiry,
			Input:  *meta,
			Fee:    fee,
		}

		jsonBody, err := json.Marshal(proposeBlockReq)
		if err != nil {
			return nil, nil, err
		}

		r := bytes.NewReader(jsonBody)

		for _, endpoint := range s.proverEndpoints {
			log.Info(
				"Attempting to assign prover",
				"endpoint", endpoint,
				"fee", fee.String(),
				"expiry", expiry,
			)
			client := &http.Client{Timeout: s.requestTimeout}

			req, err := http.NewRequestWithContext(
				ctx,
				"POST",
				fmt.Sprintf("%v/%v", endpoint, "proposeBlock"),
				r,
			)
			if err != nil {
				log.Error("Init http request error", "endpoint", endpoint, "err", err)
				continue
			}
			req.Header.Add("Content-Type", "application/json")

			res, err := client.Do(req)
			if err != nil {
				log.Error("Request prover server error", "endpoint", endpoint, "err", err)
				continue
			}

			if res.StatusCode != http.StatusOK {
				log.Info(
					"Prover rejected request to assign block",
					"endpoint", endpoint,
					"fee", fee.String(),
					"expiry", expiry,
				)
				continue
			}

			resBody, err := io.ReadAll(res.Body)
			if err != nil {
				log.Error("Read response body error", "endpoint", endpoint, "err", err)
				continue
			}

			resp := &assignProverResponse{}
			if err := json.Unmarshal(resBody, resp); err != nil {
				log.Error("Unmarshal response body error", "endpoint", endpoint, "err", err)
				continue
			}

			// ensure prover in response is the same as the one recovered
			// from the signature
			encodedBlockData, err := encoding.EncodeProposeBlockData(proposeBlockReq)
			if err != nil {
				log.Error("Encode block data error", "endpoint", endpoint, "error", err)
				continue
			}

			pubKey, err := crypto.SigToPub(crypto.Keccak256Hash(encodedBlockData).Bytes(), resp.SignedPayload)
			if err != nil {
				log.Error("Failed to get public key from signature", "endpoint", endpoint, "error", err)
				continue
			}

			if crypto.PubkeyToAddress(*pubKey).Hex() != resp.Prover.Hex() {
				log.Info(
					"Assigned prover signature did not recover to provided prover address",
					"endpoint", endpoint,
					"recoveredAddress", crypto.PubkeyToAddress(*pubKey).Hex(),
					"providedProver", resp.Prover.Hex(),
				)
				continue
			}

			// make sure the prover has the necessary balance either in TaikoL1 token balances
			// or, if not, check allowance, as contract will attempt to burn directly after
			// if it doesnt have the available tokenbalance in-contract.
			taikoTokenBalance, err := s.rpc.TaikoL1.GetTaikoTokenBalance(&bind.CallOpts{Context: ctx}, resp.Prover)
			if err != nil {
				log.Error(
					"Get taiko token balance error",
					"endpoint", endpoint,
					"providedProver", resp.Prover.Hex(),
					"error", err,
				)
				continue
			}

			if s.protocolConfigs.ProofBond.Cmp(taikoTokenBalance) > 0 {
				// check allowance on taikotoken contract
				allowance, err := s.rpc.TaikoToken.Allowance(&bind.CallOpts{Context: ctx}, resp.Prover, s.taikoL1Address)
				if err != nil {
					log.Error(
						"Get taiko token allowance error",
						"endpoint", endpoint,
						"providedProver", resp.Prover.Hex(),
						"error", err,
					)
					continue
				}

				if s.protocolConfigs.ProofBond.Cmp(allowance) > 0 {
					log.Info(
						"Assigned prover does not have required on-chain token balance or allowance",
						"endpoint", endpoint,
						"providedProver", resp.Prover.Hex(),
						"taikoTokenBalance", taikoTokenBalance.String(),
						"allowance", allowance.String(),
						"proofBond", s.protocolConfigs.ProofBond,
						"requiredFee", fee.String(),
					)
					continue
				}
			}

			// convert signature to one solidity can recover by adding 27 to 65th byte
			resp.SignedPayload[64] = uint8(uint(resp.SignedPayload[64])) + 27

			encoded, err := encoding.EncodeProverAssignment(&encoding.ProverAssignment{
				Prover: resp.Prover,
				Expiry: proposeBlockReq.Expiry,
				Data:   resp.SignedPayload,
			})
			if err != nil {
				return nil, nil, err
			}

			log.Info(
				"Prover assigned for block",
				"prover", resp.Prover.Hex(),
				"signedPayload", common.Bytes2Hex(resp.SignedPayload),
			)

			return encoded, fee, nil
		}
	}

	return nil, nil, errUnableToFindProver
}
