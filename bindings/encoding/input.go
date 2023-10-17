package encoding

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	taikol1 "github.com/taikoxyz/taiko-client/bindings/taikol1"
	taikol2 "github.com/taikoxyz/taiko-client/bindings/taikol2"
)

// ABI arguments marshaling components.
var (
	evidenceComponents = []abi.ArgumentMarshaling{
		{
			Name: "metaHash",
			Type: "bytes32",
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
			Name: "tier",
			Type: "uint16",
		},
		{
			Name: "proof",
			Type: "bytes",
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
					Type: "uint256",
				},
			},
		},
		{
			Name: "expiry",
			Type: "uint64",
		},
		{
			Name: "signature",
			Type: "bytes",
		},
	}
)

var (
	// Evidence
	evidenceType, _ = abi.NewType("tuple", "TaikoData.BlockEvidence", evidenceComponents)
	evidenceArgs    = abi.Arguments{{Name: "Evidence", Type: evidenceType}}
	// ProverAssignment
	proverAssignmentType, _ = abi.NewType("tuple", "ProverAssignment", proverAssignmentComponents)
	proverAssignmentArgs    = abi.Arguments{{Name: "ProverAssignment", Type: proverAssignmentType}}
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
				Type: "uint256",
			},
		},
	)
	proverAssignmentPayloadArgs = abi.Arguments{
		{Name: "PROVER_ASSIGNMENT", Type: stringType},
		{Name: "txListHash", Type: bytes32Type},
		{Name: "assignment.feeToken", Type: addressType},
		{Name: "assignment.expiry", Type: uint64Type},
		{Name: "assignment.tierFees", Type: tierFeesType},
	}
)

// Contract ABIs.
var (
	TaikoL1ABI *abi.ABI
	TaikoL2ABI *abi.ABI
)

func init() {
	var err error

	if TaikoL1ABI, err = taikol1.TaikoL1ClientMetaData.GetAbi(); err != nil {
		log.Crit("Get TaikoL1 ABI error", "error", err)
	}

	if TaikoL2ABI, err = taikol2.TaikoL2ClientMetaData.GetAbi(); err != nil {
		log.Crit("Get TaikoL2 ABI error", "error", err)
	}
}

// EncodeProverAssignment performs the solidity `abi.encode` for the given proverAssignment.
func EncodeProverAssignment(assignment *ProverAssignment) ([]byte, error) {
	b, err := proverAssignmentArgs.Pack(assignment)
	if err != nil {
		return nil, fmt.Errorf("failed to abi.encode prover assignment, %w", err)
	}
	return b, nil
}

// EncodeEvidence performs the solidity `abi.encode` for the given evidence.
func EncodeEvidence(e *BlockEvidence) ([]byte, error) {
	b, err := evidenceArgs.Pack(e)
	if err != nil {
		return nil, fmt.Errorf("failed to abi.encode evidence, %w", err)
	}
	return b, nil
}

// EncodeProverAssignmentPayload performs the solidity `abi.encode` for the given proverAssignment payload.
func EncodeProverAssignmentPayload(
	txListHash common.Hash,
	feeToken common.Address,
	expiry uint64,
	tierFees []TierFee,
) ([]byte, error) {
	b, err := proverAssignmentPayloadArgs.Pack("PROVER_ASSIGNMENT", txListHash, feeToken, expiry, tierFees)
	if err != nil {
		return nil, fmt.Errorf("failed to abi.encode prover assignment hash payload, %w", err)
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
