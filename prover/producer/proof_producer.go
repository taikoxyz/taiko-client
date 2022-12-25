package producer

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/taikoxyz/taiko-client/bindings"
)

// ProofRequestOptions contains all options that need to be passed to zkEVM rpcd service.
type ProofRequestOptions struct {
	Height     *big.Int                         // the block number
	Meta       *bindings.TaikoDataBlockMetadata // block meta data
	L2Endpoint string                           // a L2 execution engine's RPC endpoint
	Retry      bool                             // retry proof computation if error
	Param      string                           // parameter file to use
}

type ProofWithHeader struct {
	BlockID *big.Int
	Meta    *bindings.TaikoDataBlockMetadata
	Header  *types.Header
	ZkProof []byte
}

type ProofProducer interface {
	RequestProof(
		opts *ProofRequestOptions,
		blockID *big.Int,
		meta *bindings.TaikoDataBlockMetadata,
		header *types.Header,
		resultCh chan *ProofWithHeader,
	) error
}
