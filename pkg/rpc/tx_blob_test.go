package rpc

import (
	"context"
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/stretchr/testify/assert"

	"github.com/taikoxyz/taiko-client/internal/utils"
)

func TestSendingBlobTx(t *testing.T) {
	t.SkipNow()
	// Load environment variables.
	utils.LoadEnv()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	url := os.Getenv("L1_NODE_WS_ENDPOINT")
	l1Client, err := NewEthClient(ctx, url, time.Second*20)
	assert.NoError(t, err)

	priv := os.Getenv("L1_PROPOSER_PRIVATE_KEY")
	sk, err := crypto.ToECDSA(common.FromHex(priv))
	assert.NoError(t, err)

	opts, err := bind.NewKeyedTransactorWithChainID(sk, l1Client.ChainID)
	assert.NoError(t, err)
	opts.Context = ctx
	//opts.NoSend = true

	balance, err := l1Client.BalanceAt(ctx, opts.From, nil)
	assert.NoError(t, err)
	t.Logf("address: %s, balance: %s", opts.From.String(), balance.String())

	data, dErr := os.ReadFile("./tx_blob.go")
	assert.NoError(t, dErr)
	//data := []byte{'s'}
	sideCar, sErr := MakeSidecar(data)
	assert.NoError(t, sErr)

	tx, err := l1Client.TransactBlobTx(opts, common.Address{}, nil, sideCar)
	assert.NoError(t, err)

	receipt, err := bind.WaitMined(ctx, l1Client, tx)
	assert.NoError(t, err)
	assert.Equal(t, true, receipt.Status == types.ReceiptStatusSuccessful)

	t.Log("blob hash: ", tx.BlobHashes()[0].String())
	t.Log("block number: ", receipt.BlockNumber.Uint64())
	t.Log("tx hash: ", receipt.TxHash.String())
}

func TestMakeSideCar(t *testing.T) {
	origin, err := os.ReadFile("./tx_blob.go")
	assert.NoError(t, err)

	sideCar, mErr := MakeSidecar(origin)
	assert.NoError(t, mErr)

	origin1, dErr := DecodeBlobs(sideCar.Blobs)
	assert.NoError(t, dErr)
	assert.Equal(t, origin, origin1)
}

func TestSpecialEndWith0(t *testing.T) {
	// nolint: lll
	var txsData = `
[{"type":"0x2","chainId":"0x28c59","nonce":"0x1cca","to":"0x0167001000000000000000000000000000010099","gas":"0x86b3","gasPrice":null,"maxPriorityFeePerGas":"0x59682f00","maxFeePerGas":"0x59682f02","value":"0x0","input":"0xa9059cbb00000000000000000000000001670010000000000000000000000000000100990000000000000000000000000000000000000000000000000000000000000001","accessList":[],"v":"0x0","r":"0x2d554e149d15575030f271403a3b359cd9d5df8acb47ae7df5845aadc54b1ee2","s":"0x39b7ce8e803c443d8fd33679948fbd0a485d88b6a55812a53d9a03a922142100","yParity":"0x0","hash":"0x27aa02a44ea343a72131fc67734c67d410ab6f65429637fbb17a08f781e77f7e"}]
`

	var txs types.Transactions
	err := json.Unmarshal([]byte(txsData), &txs)
	assert.NoError(t, err)

	origin, err := rlp.EncodeToBytes(txs)
	assert.NoError(t, err)

	blobs := EncodeBlobs(origin)

	data, dErr := DecodeBlobs(blobs)
	assert.NoError(t, dErr)

	assert.Equal(t, crypto.Keccak256Hash(origin), crypto.Keccak256Hash(data))
}
