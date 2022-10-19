package prover

import (
	"math/big"
	"math/rand"
	"testing"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/stretchr/testify/require"
)

var (
	maxBlocksGasLimit = uint64(50)
	maxBlockNumTxs    = uint64(11)
	maxTxlistBytes    = uint64(10000)
	minTxGasLimit     = uint64(1)
	chainID           = genesis.Config.ChainID
)

func newTestTxListValidator(t *testing.T) *TxListValidator {
	return &TxListValidator{
		maxBlocksGasLimit: maxBlocksGasLimit,
		maxBlockNumTxs:    maxBlockNumTxs,
		maxTxlistBytes:    maxTxlistBytes,
		minTxGasLimit:     minTxGasLimit,
		chainID:           chainID,
	}
}

func rlpEncodedTransactionBytes(l int, signed bool) []byte {
	txs := make(types.Transactions, 0)
	for i := 0; i < l; i++ {
		var tx *types.Transaction
		if signed {
			txData := &types.LegacyTx{
				Nonce:    1,
				To:       &testAddr,
				GasPrice: big.NewInt(100),
				Value:    big.NewInt(1),
				Gas:      10,
			}

			tx = types.MustSignNewTx(testKey, types.LatestSigner(genesis.Config), txData)
		} else {
			tx = types.NewTransaction(1, testAddr, big.NewInt(1), 10, big.NewInt(100), nil)
		}
		txs = append(
			txs,
			tx,
		)
	}
	b, _ := rlp.EncodeToBytes(txs)
	return b
}

func randBytes(l uint64) []byte {
	b := make([]byte, l)
	rand.Read(b)
	return b
}

func Test_ValidateTxList(t *testing.T) {
	v := newTestTxListValidator(t)
	tests := []struct {
		name                string
		blockID             *big.Int
		proposeBlockTxInput []byte
		wantReason          InvalidTxListReason
		wantTxIdx           int
		wantErr             bool
	}{
		{
			"binary not decodable",
			chainID,
			randBytes(5),
			HintBinaryNotDecodable,
			0,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reason, txIdx, err := v.ValidateTxList(tt.blockID, tt.proposeBlockTxInput)
			require.Equal(t, tt.wantReason, reason)
			require.Equal(t, tt.wantTxIdx, txIdx)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}
func Test_isTxListValid(t *testing.T) {
	v := newTestTxListValidator(t)
	tests := []struct {
		name        string
		blockID     *big.Int
		txListBytes []byte
		wantReason  InvalidTxListReason
		wantTxIdx   int
	}{
		{
			"txListBytes binary too large",
			chainID,
			randBytes(maxTxlistBytes + 1),
			HintBinaryTooLarge,
			0,
		},
		{
			"txListBytes not decodable to rlp",
			chainID,
			randBytes(1),
			HintBinaryNotDecodable,
			0,
		},
		{
			"txListBytes too many transactions",
			chainID,
			rlpEncodedTransactionBytes(int(maxBlockNumTxs)+1, true),
			HintBlockTooManyTxs,
			0,
		},
		{
			"txListBytes gas limit too large",
			chainID,
			rlpEncodedTransactionBytes(6, true),
			HintBlockGasLimitTooLarge,
			0,
		},
		{
			"invalid signature",
			chainID,
			rlpEncodedTransactionBytes(1, false),
			HintTxInvalidSig,
			0,
		},
		{
			"success empty tx list",
			chainID,
			rlpEncodedTransactionBytes(0, true),
			HintOK,
			0,
		},
		{
			"success non-empty tx list",
			chainID,
			rlpEncodedTransactionBytes(1, true),
			HintOK,
			0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reason, txIdx := v.isTxListValid(tt.blockID, tt.txListBytes)
			require.Equal(t, tt.wantReason, reason)
			require.Equal(t, tt.wantTxIdx, txIdx)
		})
	}
}
