package selector

import (
	"context"
	"math/big"
	"net/url"

	"github.com/taikoxyz/taiko-client/bindings/encoding"
)

type ProverSelector interface {
	AssignProver(
		ctx context.Context,
		meta *encoding.TaikoL1BlockMetadataInput,
	) (signedPayload []byte, fee *big.Int, err error)
	ProverEndpoints() []*url.URL
}
