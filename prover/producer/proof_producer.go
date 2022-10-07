package producer

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
)

// ProofRequestOptions contains all options that need to be passed to zkEVM rpcd service.
type ProofRequestOptions struct {
	Height         *big.Int // the block number
	L2NodeEndpoint string   // the L2 node rpc endpoint url
	Retry          bool     // retry proof computation if error
	Param          string   // parameter file to use
}

type ProofWithHeader struct {
	BlockID *big.Int
	Header  *types.Header
	ZkProof []byte
}

type ProofProducer interface {
	RequestProof(opts *ProofRequestOptions, blockID *big.Int, header *types.Header, resultCh chan *ProofWithHeader) error
}
