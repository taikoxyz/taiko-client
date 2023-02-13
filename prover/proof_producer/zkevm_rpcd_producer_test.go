package producer

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
	"github.com/taikoxyz/taiko-client/bindings"
)

func TestNewZkevmRpcdProducer(t *testing.T) {
	_, err := NewZkevmRpcdProducer("http://localhost:28551", "", "", false)
	require.EqualError(t, err, errRpcdUnhealthy.Error())

	dummpyZkevmRpcdProducer, err := NewZkevmRpcdProducer("http://localhost:18545", "", "", false)
	require.Nil(t, err)

	dummpyZkevmRpcdProducer.CustomProofHook = func() ([]byte, error) {
		return []byte{0}, nil
	}

	resCh := make(chan *ProofWithHeader, 1)

	blockID := common.Big32
	header := &types.Header{
		ParentHash:  randHash(),
		UncleHash:   randHash(),
		Coinbase:    bindings.GoldenTouchAddress,
		Root:        randHash(),
		TxHash:      randHash(),
		ReceiptHash: randHash(),
		Difficulty:  common.Big0,
		Number:      common.Big256,
		GasLimit:    1024,
		GasUsed:     1024,
		Time:        uint64(time.Now().Unix()),
		Extra:       randHash().Bytes(),
		MixDigest:   randHash(),
		Nonce:       types.BlockNonce{},
	}
	require.Nil(t, dummpyZkevmRpcdProducer.RequestProof(
		&ProofRequestOptions{},
		blockID,
		&bindings.TaikoDataBlockMetadata{},
		header,
		resCh,
	))

	res := <-resCh
	require.Equal(t, res.BlockID, blockID)
	require.Equal(t, res.Header, header)
	require.NotEmpty(t, res.ZkProof)
}

var testCalldataHexHash = "0xf50afda3076f7102e4d7d20fc82856b47d8f357d5007ccd1b541fa4b42ba7cba"

func TestZkevmRpcdProducerOutputToCalldata(t *testing.T) {
	output, err := os.ReadFile("../../testutils/testdata/zkchain_proof.json")
	require.Nil(t, err)

	var zkevmRpcdOutput RequestProofBodyResponse
	require.Nil(t, json.Unmarshal(output, &zkevmRpcdOutput))

	calldata := new(ZkevmRpcdProducer).outputToCalldata(zkevmRpcdOutput.Result)

	require.Equal(t, common.HexToHash(testCalldataHexHash), crypto.Keccak256Hash(calldata))
}
