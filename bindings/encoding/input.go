package encoding

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/taikoxyz/taiko-client/bindings"
)

// ABI arguments marshaling components.
var (
	blockMetadataComponents = []abi.ArgumentMarshaling{
		{
			Name: "l1Hash",
			Type: "bytes32",
		},
		{
			Name: "difficulty",
			Type: "bytes32",
		},
		{
			Name: "blobHash",
			Type: "bytes32",
		},
		{
			Name: "extraData",
			Type: "bytes32",
		},
		{
			Name: "depositsHash",
			Type: "bytes32",
		},
		{
			Name: "coinbase",
			Type: "address",
		},
		{
			Name: "id",
			Type: "uint64",
		},
		{
			Name: "gasLimit",
			Type: "uint32",
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
			Name: "txListByteOffset",
			Type: "uint24",
		},
		{
			Name: "txListByteSize",
			Type: "uint24",
		},
		{
			Name: "minTier",
			Type: "uint16",
		},
		{
			Name: "blobUsed",
			Type: "bool",
		},
		{
			Name: "parentMetaHash",
			Type: "bytes32",
		},
	}
	transitionComponents = []abi.ArgumentMarshaling{
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
	}
	tierProofComponents = []abi.ArgumentMarshaling{
		{
			Name: "tier",
			Type: "uint16",
		},
		{
			Name: "data",
			Type: "bytes",
		},
	}
	blockParamsComponents = []abi.ArgumentMarshaling{
		{
			Name:       "assignment",
			Type:       "tuple",
			Components: proverAssignmentComponents,
		},
		{
			Name: "extraData",
			Type: "bytes32",
		},
		{
			Name: "blobHash",
			Type: "bytes32",
		},
		{
			Name: "txListByteOffset",
			Type: "uint24",
		},
		{
			Name: "txListByteSize",
			Type: "uint24",
		},
		{
			Name: "cacheBlobForReuse",
			Type: "bool",
		},
		{
			Name: "parentMetaHash",
			Type: "bytes32",
		},
	}
	proverAssignmentComponents = []abi.ArgumentMarshaling{
		{
			Name: "prover",
			Type: "address",
		},
		{
			Name: "feeToken",
			Type: "address",
		},
		{
			Name: "tierFees",
			Type: "tuple[]",
			Components: []abi.ArgumentMarshaling{
				{
					Name: "tier",
					Type: "uint16",
				},
				{
					Name: "fee",
					Type: "uint128",
				},
			},
		},
		{
			Name: "expiry",
			Type: "uint64",
		},
		{
			Name: "maxBlockId",
			Type: "uint64",
		},
		{
			Name: "maxProposedIn",
			Type: "uint64",
		},
		{
			Name: "metaHash",
			Type: "bytes32",
		},
		{
			Name: "signature",
			Type: "bytes",
		},
	}
)

var (
	// BlockParams
	blockParamsComponentsType, _ = abi.NewType("tuple", "TaikoData.BlockParams", blockParamsComponents)
	blockParamsComponentsArgs    = abi.Arguments{{Name: "TaikoData.BlockParams", Type: blockParamsComponentsType}}
	// ProverAssignmentPayload
	stringType, _   = abi.NewType("string", "", nil)
	bytes32Type, _  = abi.NewType("bytes32", "", nil)
	addressType, _  = abi.NewType("address", "", nil)
	uint64Type, _   = abi.NewType("uint64", "", nil)
	tierFeesType, _ = abi.NewType(
		"tuple[]",
		"",
		[]abi.ArgumentMarshaling{
			{
				Name: "tier",
				Type: "uint16",
			},
			{
				Name: "fee",
				Type: "uint128",
			},
		},
	)
	proverAssignmentPayloadArgs = abi.Arguments{
		{Name: "PROVER_ASSIGNMENT", Type: stringType},
		{Name: "taikoAddress", Type: addressType},
		{Name: "blobHash", Type: bytes32Type},
		{Name: "assignment.feeToken", Type: addressType},
		{Name: "assignment.expiry", Type: uint64Type},
		{Name: "assignment.maxBlockId", Type: uint64Type},
		{Name: "assignment.maxProposedIn", Type: uint64Type},
		{Name: "assignment.tierFees", Type: tierFeesType},
	}
	blockMetadataComponentsType, _ = abi.NewType("tuple", "TaikoData.BlockMetadata", blockMetadataComponents)
	transitionComponentsType, _    = abi.NewType("tuple", "TaikoData.Transition", transitionComponents)
	tierProofComponentsType, _     = abi.NewType("tuple", "TaikoData.TierProof", tierProofComponents)
	proveBlockInputArgs            = abi.Arguments{
		{Name: "TaikoData.BlockMetadata", Type: blockMetadataComponentsType},
		{Name: "TaikoData.Transition", Type: transitionComponentsType},
		{Name: "TaikoData.TierProof", Type: tierProofComponentsType},
	}
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

// EncodeBlockParams performs the solidity `abi.encode` for the given blockParams.
func EncodeBlockParams(params *BlockParams) ([]byte, error) {
	b, err := blockParamsComponentsArgs.Pack(params)
	if err != nil {
		return nil, fmt.Errorf("failed to abi.encode block params, %w", err)
	}
	return b, nil
}

// EncodeProverAssignmentPayload performs the solidity `abi.encode` for the given proverAssignment payload.
func EncodeProverAssignmentPayload(
	taikoAddress common.Address,
	txListHash common.Hash,
	feeToken common.Address,
	expiry uint64,
	maxBlockID uint64,
	maxProposedIn uint64,
	tierFees []TierFee,
) ([]byte, error) {
	b, err := proverAssignmentPayloadArgs.Pack(
		"PROVER_ASSIGNMENT",
		taikoAddress,
		txListHash,
		feeToken,
		expiry,
		maxBlockID,
		maxProposedIn,
		tierFees,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to abi.encode prover assignment hash payload, %w", err)
	}
	return b, nil
}

// EncodeProveBlockInput performs the solidity `abi.encode` for the given TaikoL1.proveBlock input.
func EncodeProveBlockInput(
	meta *bindings.TaikoDataBlockMetadata,
	transition *bindings.TaikoDataTransition,
	tierProof *bindings.TaikoDataTierProof,
) ([]byte, error) {
	b, err := proveBlockInputArgs.Pack(meta, transition, tierProof)
	if err != nil {
		return nil, fmt.Errorf("failed to abi.encode TakoL1.proveBlock input, %w", err)
	}
	return b, nil
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
