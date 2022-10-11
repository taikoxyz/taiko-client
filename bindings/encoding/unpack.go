package encoding

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/log"
	"github.com/taikochain/taiko-client/bindings"
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

	if TaikoL2ABI, err = bindings.V1TaikoL2ClientMetaData.GetAbi(); err != nil {
		log.Crit("Get TaikoL2 ABI error", "error", err)
	}
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

	inputs, ok := args["inputs"].([][]byte)

	if !ok || len(inputs) < 2 {
		return nil, fmt.Errorf("invalid transaction inputs map length, get: %d", len(inputs))
	}

	return inputs[1], nil
}
