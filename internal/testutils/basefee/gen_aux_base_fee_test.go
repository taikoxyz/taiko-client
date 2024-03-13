package basefee

import (
	"context"
	"fmt"
	"github.com/taikoxyz/taiko-client/pkg/sender"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common/math"
	"github.com/stretchr/testify/suite"

	"github.com/taikoxyz/taiko-client/bindings"
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

func (s *BaseFeeSuite) calc1559BaseFee(_l1BlockID, _parentGasUsed uint64) uint64 {
	var (
		baseFee *big.Int
		err     error
	)
	if s.gasExcess > 0 {
		excess := s.gasExcess + _parentGasUsed
		var numL1Blocks uint64
		if s.lastSyncedBlock > 0 && _l1BlockID > s.lastSyncedBlock {
			numL1Blocks = _l1BlockID - s.lastSyncedBlock
		}
		if numL1Blocks > 0 {
			issuance := numL1Blocks * uint64(s.config.GasTargetPerL1Block)
			if excess > issuance {
				excess -= issuance
			} else {
				excess = 1
			}
		}
		gasExcess := utils.Min(excess, math.MaxUint64)
		baseFee, err = s.baseFee.BaseFee(nil,
			new(big.Int).SetUint64(gasExcess),
			s.config.BasefeeAdjustmentQuotient,
			s.config.GasTargetPerL1Block,
		)
		s.Nil(err)
	}

	if baseFee == nil || baseFee.Uint64() == 0 {
		return 1
	}
	return baseFee.Uint64()
}

func (s *BaseFeeSuite) testCalc1559BaseFee(numL1Blocks uint64, startGasUsed uint64, times int) {
	s.lastSyncedBlock = 1
	for ; times > 0; times-- {
		fmt.Println("GasUsed: ", startGasUsed)
		startGasUsed = s.calc1559BaseFee(s.lastSyncedBlock+numL1Blocks, startGasUsed)
	}
}

func (s *BaseFeeSuite) TestCalc1559BaseFee() {
	var numL1Blocks uint64 = 1
	times := 10
	for i := 7740000000; times > 0; i += 400000000 {
		times--
		s.testCalc1559BaseFee(numL1Blocks, numL1Blocks*uint64(s.config.GasTargetPerL1Block)+uint64(i), 3)
	}
}

func (s *BaseFeeSuite) TestBaseFee() {
	times := 1000
	for i := 9550000000; times > 0; i += 50000000 {
		times--
		gasExcess := new(big.Int).SetUint64(uint64(i))
		baseFee, err := s.baseFee.BaseFee(nil,
			gasExcess,
			s.config.BasefeeAdjustmentQuotient,
			s.config.GasTargetPerL1Block,
		)
		s.Nil(err)
		//if baseFee.Cmp(gasExcess) >= 0 {
		fmt.Println(i, ", baseFee: ", baseFee)
		//}
	}
}

func (s *BaseFeeSuite) TestVerifyCalc1559BaseFee() {
	l1CLi := s.RPCClient.L1
	l2CLi := s.RPCClient.L2

	l1Header, err := l1CLi.HeaderByNumber(context.Background(), nil)
	s.Nil(err)
	l1BlockID := l1Header.Number.Uint64()

	l2Header, err := l2CLi.HeaderByNumber(context.Background(), nil)
	s.Nil(err)

	l2BaseFee, err := s.RPCClient.TaikoL2.GetBasefee(nil, l1BlockID, uint32(l2Header.GasUsed))
	s.Nil(err)

	mockBaseFee := s.calc1559BaseFee(l1BlockID, l2Header.GasUsed)

	s.Equal(l2BaseFee.Uint64(), mockBaseFee)
}

func (s *BaseFeeSuite) SetupTest() {
	s.ClientTestSuite.SetupTest()

	send, err := sender.NewSender(context.Background(), nil, s.RPCClient.L1, s.TestAddrPrivKey)
	s.Nil(err)

	opts := send.GetOpts()
	_, tx, baseFee, err := DeployAuxBaseFee(opts, s.RPCClient.L1)
	s.Nil(err)
	s.baseFee = baseFee
	id, err := send.SendTransaction(tx)
	s.Nil(err)
	confirm := <-send.TxToConfirmChannel(id)
	s.Nil(confirm.Err)

	taikoL2 := s.RPCClient.TaikoL2
	s.gasExcess, err = taikoL2.GasExcess(nil)
	s.Nil(err)
	fmt.Println("gasExcess: ", s.gasExcess)

	s.config, err = taikoL2.GetConfig(nil)
	s.Nil(err)
	fmt.Println("config: ", s.config.BasefeeAdjustmentQuotient, s.config.GasTargetPerL1Block)

	s.lastSyncedBlock, err = taikoL2.LastSyncedBlock(nil)
	s.Nil(err)

	fmt.Println("lastSyncedBlock: ", s.lastSyncedBlock)
}

func TestDriverTestSuite(t *testing.T) {
	suite.Run(t, new(BaseFeeSuite))
}
