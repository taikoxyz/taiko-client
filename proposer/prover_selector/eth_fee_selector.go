package selector

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"net/url"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/go-resty/resty/v2"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	"github.com/taikoxyz/taiko-client/prover/http"
)

var (
	httpScheme              = "http"
	errEmptyProverEndpoints = errors.New("empty prover endpoints")
	errUnableToFindProver   = errors.New("unable to find prover")
)

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
		for _, endpoint := range s.proverEndpoints {
			encodedAssignment, proverAddress, err := assignProver(ctx, meta, endpoint, fee, expiry, s.requestTimeout)
			if err != nil {
				log.Warn("Failed to assign prover", "endpoint", endpoint, "error", err)
				continue
			}

			ok, err := s.checkProverBalance(ctx, proverAddress)
			if err != nil {
				log.Warn("Failed to check prover balance", "endpoint", endpoint, "error", err)
				continue
			}
			if !ok {
				continue
			}

			return encodedAssignment, fee, nil
		}
	}

	return nil, nil, errUnableToFindProver
}

// checkProverBalance checks if the prover has the necessary balance either in TaikoL1 token balances
// or, if not, then check allowance, as contract will attempt to burn directly after
// if it doesnt have the available token balance in-contract.
func (s *ETHFeeSelector) checkProverBalance(ctx context.Context, prover common.Address) (bool, error) {
	taikoTokenBalance, err := s.rpc.TaikoL1.GetTaikoTokenBalance(&bind.CallOpts{Context: ctx}, prover)
	if err != nil {
		return false, err
	}

	if s.protocolConfigs.ProofBond.Cmp(taikoTokenBalance) > 0 {
		// Check allowance on taiko token contract
		allowance, err := s.rpc.TaikoToken.Allowance(&bind.CallOpts{Context: ctx}, prover, s.taikoL1Address)
		if err != nil {
			return false, err
		}

		if s.protocolConfigs.ProofBond.Cmp(allowance) > 0 {
			log.Info(
				"Assigned prover does not have required on-chain token balance or allowance",
				"providedProver", prover.Hex(),
				"taikoTokenBalance", taikoTokenBalance.String(),
				"allowance", allowance.String(),
				"proofBond", s.protocolConfigs.ProofBond,
			)
			return false, nil
		}
	}

	return true, nil
}

func assignProver(
	ctx context.Context,
	meta *encoding.TaikoL1BlockMetadataInput,
	endpoint *url.URL,
	fee *big.Int,
	expiry uint64,
	timeout time.Duration,
) ([]byte, common.Address, error) {
	log.Info(
		"Attempting to assign prover",
		"endpoint", endpoint,
		"fee", fee.String(),
		"expiry", expiry,
	)

	// Send the HTTP request
	var (
		client  = resty.New()
		reqBody = &encoding.ProposeBlockData{Expiry: expiry, Input: *meta, Fee: fee}
		result  = http.ProposeBlockResponse{}
	)
	requestUrl, err := url.JoinPath(endpoint.String(), "/proposeBlock")
	if err != nil {
		return nil, common.Address{}, err
	}

	ctxTimeout, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	resp, err := client.R().
		SetContext(ctxTimeout).
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(reqBody).
		SetResult(&result).
		Post(requestUrl)
	if err != nil {
		return nil, common.Address{}, err
	}
	if !resp.IsSuccess() {
		return nil, common.Address{}, fmt.Errorf("unsuccessful response %d", resp.StatusCode())
	}

	// Ensure prover in response is the same as the one recovered
	// from the signature
	encodedBlockData, err := encoding.EncodeProposeBlockData(reqBody)
	if err != nil {
		return nil, common.Address{}, err
	}

	pubKey, err := crypto.SigToPub(crypto.Keccak256Hash(encodedBlockData).Bytes(), result.SignedPayload)
	if err != nil {
		return nil, common.Address{}, err
	}

	if crypto.PubkeyToAddress(*pubKey).Hex() != result.Prover.Hex() {
		return nil, common.Address{}, fmt.Errorf(
			"assigned prover signature did not recover to provided prover address %s != %s",
			crypto.PubkeyToAddress(*pubKey).Hex(),
			result.Prover.Hex(),
		)
	}

	// Convert signature to one solidity can recover by adding 27 to 65th byte
	result.SignedPayload[64] = uint8(uint(result.SignedPayload[64])) + 27

	encoded, err := encoding.EncodeProverAssignment(&encoding.ProverAssignment{
		Prover: result.Prover,
		Expiry: reqBody.Expiry,
		Data:   result.SignedPayload,
	})
	if err != nil {
		return nil, common.Address{}, err
	}

	return encoded, result.Prover, nil
}
