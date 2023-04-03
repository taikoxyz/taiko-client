package encoding

import (
	"bytes"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/taikoxyz/taiko-client/bindings"
)

// ABI arguments marshaling components.
var (
	blockMetadataInputComponents = []abi.ArgumentMarshaling{
		{
			Name: "txListHash",
			Type: "bytes32",
		},
		{
			Name: "beneficiary",
			Type: "address",
		},
		{
			Name: "gasLimit",
			Type: "uint32",
		},
		{
			Name: "txListByteStart",
			Type: "uint24",
		},
		{
			Name: "txListByteEnd",
			Type: "uint24",
		},
		{
			Name: "cacheTxListInfo",
			Type: "uint8",
		},
	}
	blockMetadataComponents = []abi.ArgumentMarshaling{
		{
			Name: "id",
			Type: "uint64",
		},
		{
			Name: "timestamp",
			Type: "uint64",
		},
		{
			Name: "l1Height",
			Type: "uint64",
		},
		{
			Name: "gasLimit",
			Type: "uint32",
		},
		{
			Name: "l1Hash",
			Type: "bytes32",
		},
		{
			Name: "mixHash",
			Type: "bytes32",
		},
		{
			Name: "txListHash",
			Type: "bytes32",
		},
		{
			Name: "txListByteStart",
			Type: "uint24",
		},
		{
			Name: "txListByteEnd",
			Type: "uint24",
		},
		{
			Name: "beneficiary",
			Type: "address",
		},
	}
	zkProofComponents = []abi.ArgumentMarshaling{
		{
			Name: "data",
			Type: "bytes",
		},
		{
			Name: "verifierId",
			Type: "uint16",
		},
	}
	evidenceComponents = []abi.ArgumentMarshaling{
		{
			Name:       "meta",
			Type:       "tuple",
			Components: blockMetadataComponents,
		},
		{
			Name:       "zkproof",
			Type:       "tuple",
			Components: zkProofComponents,
		},
		{
			Name: "parentHash",
			Type: "bytes32",
		},
		{
			Name: "blockHash",
			Type: "bytes32",
		},
		{
			Name: "signalRoot",
			Type: "bytes32",
		},
		{
			Name: "graffiti",
			Type: "bytes32",
		},
		{
			Name: "prover",
			Type: "address",
		},
	}
)

var (
	// BlockMetadataInput
	blockMetadataInputType, _ = abi.NewType("tuple", "TaikoData.BlockMetadataInput", blockMetadataInputComponents)
	blockMetadataInputArgs    = abi.Arguments{{Name: "BlockMetadataInput", Type: blockMetadataInputType}}
	// BlockMetadata
	blockMetadataType, _ = abi.NewType("tuple", "LibData.BlockMetadata", blockMetadataComponents)
	blockMetadataArgs    = abi.Arguments{{Name: "BlockMetadata", Type: blockMetadataType}}
	// Evidence
	EvidenceType, _ = abi.NewType("tuple", "TaikoData.BlockEvidence", evidenceComponents)
	EvidenceArgs    = abi.Arguments{{Name: "Evidence", Type: EvidenceType}}
)

// Contract ABIs.
var (
	TaikoL1ABI *abi.ABI
	TaikoL2ABI *abi.ABI
)

func init() {
	var err error

	if TaikoL1ABI, err = bindings.TaikoL1ClientMetaData.GetAbi(); err != nil {
		log.Crit("Get TaikoL1 ABI error", "error", err)
	}

	if TaikoL2ABI, err = bindings.TaikoL2ClientMetaData.GetAbi(); err != nil {
		log.Crit("Get TaikoL2 ABI error", "error", err)
	}
}

// EncodeBlockMetadataInput performs the solidity `abi.encode` for the given blockMetadataInput.
func EncodeBlockMetadataInput(meta *TaikoL1BlockMetadataInput) ([]byte, error) {
	b, err := blockMetadataInputArgs.Pack(meta)
	if err != nil {
		return nil, fmt.Errorf("failed to abi.encode block metadata input, %w", err)
	}
	return b, nil
}

// EncodeBlockMetadata performs the solidity `abi.encode` for the given blockMetadata.
func EncodeBlockMetadata(meta *bindings.TaikoDataBlockMetadata) ([]byte, error) {
	b, err := blockMetadataArgs.Pack(meta)
	if err != nil {
		return nil, fmt.Errorf("failed to abi.encode block metadata, %w", err)
	}
	return b, nil
}

// EncodeEvidence performs the solidity `abi.encode` for the given evidence.
func EncodeEvidence(e *TaikoL1Evidence) ([]byte, error) {
	b, err := EvidenceArgs.Pack(e)
	if err != nil {
		return nil, fmt.Errorf("failed to abi.encode evidence, %w", err)
	}
	return b, nil
}

// EncodeCommitHash performs the solidity `abi.encodePacked` for the given
// commitHash components.
func EncodeCommitHash(beneficiary common.Address, txListHash [32]byte) []byte {
	// keccak256(abi.encodePacked(beneficiary, txListHash));
	return crypto.Keccak256(
		bytes.Join([][]byte{beneficiary.Bytes(), txListHash[:]}, nil),
	)
}

// EncodeProposeBlockInput encodes the input params for TaikoL1.proposeBlock.
func EncodeProposeBlockInput(metadataInput *TaikoL1BlockMetadataInput) ([]byte, error) {
	metaBytes, err := EncodeBlockMetadataInput(metadataInput)
	if err != nil {
		return nil, err
	}
	return metaBytes, nil
}

// EncodeProveBlockInput encodes the input params for TaikoL1.proveBlock.
func EncodeProveBlockInput(
	evidence *TaikoL1Evidence,
	anchorTx *types.Transaction,
	anchorReceipt *types.Receipt,
) ([]byte, error) {
	evidenceBytes, err := EncodeEvidence(evidence)
	if err != nil {
		return nil, err
	}

	return evidenceBytes, nil
}

// EncodeProveBlockInvalidInput encodes the input params for TaikoL1.proveBlockInvalid.
func EncodeProveBlockInvalidInput(
	evidence *TaikoL1Evidence,
	target *bindings.TaikoDataBlockMetadata,
	receipt *types.Receipt,
) ([][]byte, error) {
	evidenceBytes, err := EncodeEvidence(evidence)
	if err != nil {
		return nil, err
	}

	metaBytes, err := EncodeBlockMetadata(target)
	if err != nil {
		return nil, err
	}

	receiptBytes, err := rlp.EncodeToBytes(receipt)
	if err != nil {
		return nil, err
	}

	return [][]byte{evidenceBytes, metaBytes, receiptBytes}, nil
}

// UnpackTxListBytes unpacks the input data of a TaikoL1.proposeBlock transaction, and returns the txList bytes.
func UnpackTxListBytes(txData []byte) ([]byte, error) {
	method, err := TaikoL1ABI.MethodById(txData)
	if err != nil {
		return nil, err
	}

	// Only check for safety.
	if method.Name != "proposeBlock" {
		return nil, fmt.Errorf("invalid method name: %s", method.Name)
	}

	args := map[string]interface{}{}

	if err := method.Inputs.UnpackIntoMap(args, txData[4:]); err != nil {
		return nil, err
	}

	inputs, ok := args["txList"].([]byte)

	if !ok {
		return nil, errors.New("failed to get txList bytes")
	}

	return inputs, nil
}

// UnpackEvidenceHeader unpacks the evidence data of a TaikoL1.proveBlock transaction, and returns
// the block header inside.
func UnpackEvidenceHeader(txData []byte) (*BlockHeader, error) {
	method, err := TaikoL1ABI.MethodById(txData)
	if err != nil {
		return nil, err
	}

	// Only check for safety.
	if method.Name != "proveBlock" {
		return nil, fmt.Errorf("invalid method name: %s", method.Name)
	}

	args := map[string]interface{}{}

	if err := method.Inputs.UnpackIntoMap(args, txData[4:]); err != nil {
		return nil, err
	}

	inputs, ok := args["inputs"].([][]byte)

	if !ok || len(inputs) < 3 {
		return nil, fmt.Errorf("invalid transaction inputs map length, get: %d", len(inputs))
	}

	return decodeEvidenceHeader(inputs[0])
}

// decodeEvidenceHeader decodes the encoded evidence bytes, and then returns its inner header.
func decodeEvidenceHeader(evidenceBytes []byte) (*BlockHeader, error) {
	unpacked, err := EvidenceArgs.Unpack(evidenceBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to decode evidence meta")
	}

	evidence := new(TaikoL1Evidence)
	if err := EvidenceArgs.Copy(&evidence, unpacked); err != nil {
		return nil, err
	}

	return &BlockHeader{
		ParentHash:  evidence.ParentHash,
		OmmersHash:  types.EmptyUncleHash,
		Beneficiary: evidence.Meta.Beneficiary,
		// StateRoot: , // ?
		// TransactionsRoot: , // ?
		// ReceiptsRoot: , // ?
		// LogsBloom:  , // ?
		Difficulty: common.Big0,
		Height:     new(big.Int).SetUint64(evidence.Meta.Id),
		GasLimit:   uint64(evidence.Meta.GasLimit),
		// GasUsed: , // ?
		Timestamp:     evidence.Meta.Timestamp,
		ExtraData:     []byte{},
		MixHash:       evidence.Meta.MixHash,
		Nonce:         0,
		BaseFeePerGas: nil,
	}, nil
}
