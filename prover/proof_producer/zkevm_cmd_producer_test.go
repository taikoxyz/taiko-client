package producer

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

func TestOutputToCalldata(t *testing.T) {
	output, err := os.ReadFile("../../testutils/testdata/block-5_proof.json")
	require.Nil(t, err)

	var (
		testCalldataHexHash = common.HexToHash("0x6a37a238e75278c1dd49b84a730dfa43a85ad01edad3f0f97b5a1c7e47f5123a")
		proverCmdOutput     ProverCmdOutput
	)
	require.Nil(t, json.Unmarshal(output, &proverCmdOutput))

	calldata := new(ZkevmCmdProducer).outputToCalldata(&proverCmdOutput)

	require.Equal(t, testCalldataHexHash, crypto.Keccak256Hash(calldata))
}
