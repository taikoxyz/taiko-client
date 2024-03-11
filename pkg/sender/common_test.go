package sender_test

import (
	"context"
	"math"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/taikoxyz/taiko-client/pkg/sender"
)

func (s *SenderTestSuite) TestSetConfigWithDefaultValues() {
	priv, err := crypto.ToECDSA(common.FromHex(os.Getenv("L1_PROPOSER_PRIVATE_KEY")))
	s.Nil(err)

	sender, err := sender.NewSender(
		context.Background(),
		&sender.Config{MaxRetrys: 1, MaxBlobFee: 1024},
		s.RPCClient.L1,
		priv,
	)
	s.Nil(err)
	s.Equal(uint64(50), sender.Config.GasGrowthRate)
	s.Equal(uint64(1), sender.Config.MaxRetrys)
	s.Equal(5*time.Minute, sender.Config.MaxWaitingTime)
	s.Equal(uint64(math.MaxUint64), sender.Config.MaxGasFee)
	s.Equal(uint64(1024), sender.Config.MaxBlobFee)
}
