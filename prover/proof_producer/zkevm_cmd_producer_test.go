package producer

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

func TestZkevmCmdProducerOutputToCalldata(t *testing.T) {
	output, err := os.ReadFile("../../testutils/testdata/block-5_proof.json")
	require.Nil(t, err)

	var (
		testCalldataHexHash = common.HexToHash("0xfbc74eec1aa02cadd59cf2fdfb8c311b199a4f83d0046fd20a1a53081bb0de22")
		proverCmdOutput     ProverCmdOutput
	)
	require.Nil(t, json.Unmarshal(output, &proverCmdOutput))

	calldata := new(ZkevmCmdProducer).outputToCalldata(&proverCmdOutput)

	require.Equal(t, testCalldataHexHash, crypto.Keccak256Hash(calldata))
}
