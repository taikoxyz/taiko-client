package producer

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/taikoxyz/taiko-client/bindings"
)

const (
	CircuitsDegree10Txs = 19
	CircuitsDegree80Txs = 21

	CircuitsIdx10Txs = 0
	CircuitsIdx80Txs = 1
)

// ProofRequestOptions contains all options that need to be passed to zkEVM rpcd service.
type ProofRequestOptions struct {
	Height             *big.Int // the block number
	ProverAddress      common.Address
	ProposeBlockTxHash common.Hash
	L1SignalService    common.Address
	L2SignalService    common.Address
	TaikoL2            common.Address
	MetaHash           common.Hash
	BlockHash          common.Hash
	ParentHash         common.Hash
	SignalRoot         common.Hash
	EventL1Hash        common.Hash
	Graffiti           string
	GasUsed            uint64
	ParentGasUsed      uint64
}

type ProofWithHeader struct {
	BlockID *big.Int
	Meta    *bindings.TaikoDataBlockMetadata
	Header  *types.Header
	ZkProof []byte
	Degree  uint64
	Opts    *ProofRequestOptions
}

type ProofProducer interface {
	RequestProof(
		ctx context.Context,
		opts *ProofRequestOptions,
		blockID *big.Int,
		meta *bindings.TaikoDataBlockMetadata,
		header *types.Header,
		resultCh chan *ProofWithHeader,
	) error
	Cancel(ctx context.Context, blockID *big.Int) error
}

func DegreeToCircuitsIdx(degree uint64) (uint16, error) {
	switch degree {
	case CircuitsDegree10Txs:
		return CircuitsIdx10Txs, nil
	case CircuitsDegree80Txs:
		return CircuitsIdx80Txs, nil
	default:
		return 0, fmt.Errorf("invalid degree: %d", degree)
	}
}
