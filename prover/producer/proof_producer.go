package producer

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/taikoxyz/taiko-client/bindings"
)

// ProofRequestOptions contains all options that need to be passed to zkEVM rpcd service.
type ProofRequestOptions struct {
	Height         *big.Int                       // the block number
	Meta           *bindings.LibDataBlockMetadata // block meta data
	L2NodeEndpoint string                         // the L2 node rpc endpoint url
	Retry          bool                           // retry proof computation if error
	Param          string                         // parameter file to use
}

type ProofWithHeader struct {
	BlockID *big.Int
	Meta    *bindings.LibDataBlockMetadata
	Header  *types.Header
	ZkProof []byte
}

type ProofProducer interface {
	RequestProof(
		opts *ProofRequestOptions,
		blockID *big.Int,
		meta *bindings.LibDataBlockMetadata,
		header *types.Header,
		resultCh chan *ProofWithHeader,
	) error
}
