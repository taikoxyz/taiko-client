package encoding

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/taikochain/taiko-client/bindings"
)

type BlockHeader struct {
	ParentHash       [32]byte
	OmmersHash       [32]byte
	Beneficiary      common.Address
	StateRoot        [32]byte
	TransactionsRoot [32]byte
	ReceiptsRoot     [32]byte
	LogsBloom        [8][32]byte
	Difficulty       *big.Int
	Height           *big.Int
	GasLimit         uint64
	GasUsed          uint64
	Timestamp        uint64
	ExtraData        []byte
	MixHash          [32]byte
	Nonce            uint64
	BaseFeePerGas    *big.Int
}

type TaikoL1Evidence struct {
	Meta   bindings.LibDataBlockMetadata
	Header BlockHeader
	Prover common.Address
	Proofs [][]byte
}

// FromGethHeader converts geth *types.Header to *BlockHeader
func FromGethHeader(header *types.Header) *BlockHeader {
	baseFeePerGas := header.BaseFee
	if baseFeePerGas == nil {
		baseFeePerGas = common.Big0
	}
	return &BlockHeader{
		ParentHash:       header.ParentHash,
		OmmersHash:       header.UncleHash,
		Beneficiary:      header.Coinbase,
		StateRoot:        header.Root,
		TransactionsRoot: header.TxHash,
		ReceiptsRoot:     header.ReceiptHash,
		LogsBloom:        BloomToBytes(header.Bloom),
		Difficulty:       header.Difficulty,
		Height:           header.Number,
		GasLimit:         header.GasLimit,
		GasUsed:          header.GasUsed,
		Timestamp:        header.Time,
		ExtraData:        header.Extra,
		MixHash:          header.MixDigest,
		Nonce:            header.Nonce.Uint64(),
		BaseFeePerGas:    baseFeePerGas,
	}
}

// BloomToBytes converts a types.Bloom to [8][32]byte.
func BloomToBytes(bloom types.Bloom) [8][32]byte {
	b := [8][32]byte{}

	for i := 0; i < 8; i++ {
		copy(b[i][:], bloom[i*32:(i+1)*32])
	}

	return b
}
