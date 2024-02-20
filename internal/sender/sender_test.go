package sender_test

import (
	"context"
	"math/big"
	"os"
	"runtime"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/errgroup"

	"github.com/taikoxyz/taiko-client/internal/sender"
	"github.com/taikoxyz/taiko-client/internal/utils"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
)

func TestSender(t *testing.T) {
	utils.LoadEnv()

	ctx := context.Background()

	client, err := rpc.NewEthClient(ctx, os.Getenv("L1_NODE_WS_ENDPOINT"), time.Second*10)
	assert.NoError(t, err)

	priv, err := crypto.ToECDSA(common.FromHex(os.Getenv("L1_PROPOSER_PRIVATE_KEY")))
	assert.NoError(t, err)

	send, err := sender.NewSender(ctx, &sender.Config{
		Confirmations: 1,
		MaxGasPrice:   1000000000000,
		GasRate:       10,
		MaxPendTxs:    10,
		RetryTimes:    3,
	}, client, priv)
	assert.NoError(t, err)
	defer send.Stop()

	var (
		batchSize  = 10
		eg         errgroup.Group
		confirmsCh = make([]<-chan *sender.TxConfirm, 0, batchSize)
	)
	eg.SetLimit(runtime.NumCPU())
	for i := 0; i < batchSize; i++ {
		i := i
		eg.Go(func() error {
			addr := common.BigToAddress(big.NewInt(int64(i)))
			txID, err := send.SendRaw(&addr, big.NewInt(1), nil)
			if err == nil {
				confirmCh, _ := send.WaitTxConfirm(txID)
				confirmsCh = append(confirmsCh, confirmCh)
			}
			return err
		})
	}
	err = eg.Wait()
	assert.NoError(t, err)

	for len(confirmsCh) > 0 {
		confirmCh := confirmsCh[0]
		select {
		case confirm := <-confirmCh:
			assert.NoError(t, confirm.Error)
			confirmsCh = confirmsCh[1:]
		default:
		}
	}
}
