package selector

import (
	"context"
	"math/big"
	"net/url"

	"github.com/ethereum/go-ethereum/common"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
)

type ProverSelector interface {
	AssignProver(
		ctx context.Context,
		tierFees []encoding.TierFee,
		txListHash common.Hash,
	) (signedPayload []byte, prover common.Address, fee *big.Int, err error)
	ProverEndpoints() []*url.URL
}
