package basefee

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common/math"
	"github.com/stretchr/testify/suite"

	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/internal/sender"
	"github.com/taikoxyz/taiko-client/internal/testutils"
	"github.com/taikoxyz/taiko-client/internal/utils"
)

type BaseFeeSuite struct {
	testutils.ClientTestSuite

	baseFee *AuxBaseFee

	lastSyncedBlock uint64
	gasExcess       uint64
	config          bindings.TaikoL2Config
}

func (s *BaseFeeSuite) calc1559BaseFee(_l1BlockId, _parentGasUsed uint64) uint64 {
	var (
		basefee_ *big.Int
		err      error
	)
	if s.gasExcess > 0 {
		excess := s.gasExcess + _parentGasUsed
		var numL1Blocks uint64
		if s.lastSyncedBlock > 0 && _l1BlockId > s.lastSyncedBlock {
			numL1Blocks = _l1BlockId - s.lastSyncedBlock
		}
		if numL1Blocks > 0 {
			issuance := numL1Blocks * uint64(s.config.GasTargetPerL1Block)
			if excess > issuance {
				excess -= issuance
			} else {
				excess = 1
			}
		}
		gasExcess_ := utils.Min(excess, math.MaxUint64)
		basefee_, err = s.baseFee.BaseFee(nil, new(big.Int).SetUint64(gasExcess_), s.config.BasefeeAdjustmentQuotient, s.config.GasTargetPerL1Block)
		s.Nil(err)
	}

	if basefee_ == nil || basefee_.Uint64() == 0 {
		return 1
	}
	return basefee_.Uint64()
}

func (s *BaseFeeSuite) TestVerifyCalc1559BaseFee() {
	l1CLi := s.RPCClient.L1
	l2CLi := s.RPCClient.L2
	taikoL2 := s.RPCClient.TaikoL2

	l1Header, err := l1CLi.HeaderByNumber(context.Background(), nil)
	s.Nil(err)
	_l1BlockId := l1Header.Number.Uint64()

	l2Header, err := l2CLi.HeaderByNumber(context.Background(), nil)
	s.Nil(err)

	l2BaseFee_, err := taikoL2.GetBasefee(nil, _l1BlockId, uint32(l2Header.GasUsed))
	s.Nil(err)

	mockBaseFee := s.calc1559BaseFee(_l1BlockId, l2Header.GasUsed)

	s.Equal(l2BaseFee_.Uint64(), mockBaseFee)
}

func (s *BaseFeeSuite) TestBaseFee() {}

func (s *BaseFeeSuite) SetupTest() {
	s.ClientTestSuite.SetupTest()

	send, err := sender.NewSender(context.Background(), nil, s.RPCClient.L1, s.TestAddrPrivKey)
	s.Nil(err)

	_, tx, baseFee, err := DeployAuxBaseFee(send.Opts, s.RPCClient.L1)
	s.Nil(err)
	s.baseFee = baseFee
	id, err := send.SendTransaction(tx)
	s.Nil(err)
	confirm := <-send.TxToConfirmChannel(id)
	s.Nil(confirm.Err)

	taikoL2 := s.RPCClient.TaikoL2
	s.gasExcess, err = taikoL2.GasExcess(nil)
	s.Nil(err)

	s.config, err = taikoL2.GetConfig(nil)
	s.Nil(err)

	s.lastSyncedBlock, err = taikoL2.LastSyncedBlock(nil)
	s.Nil(err)
}

func TestDriverTestSuite(t *testing.T) {
	suite.Run(t, new(BaseFeeSuite))
}
