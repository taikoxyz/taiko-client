package basefee

import (
	"context"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common/math"
	"github.com/stretchr/testify/suite"
	"github.com/taikoxyz/taiko-client/pkg/sender"

	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/internal/testutils"
)

type BaseFeeSuite struct {
	testutils.ClientTestSuite

	baseFee *AuxBaseFee

	gasExcess uint64
	config    bindings.TaikoL2Config
}

func (s *BaseFeeSuite) testCalc1559BaseFee(numL1Blocks uint64, _parentGasUsed uint32) {
	_gasIssuance := numL1Blocks * uint64(s.config.GasTargetPerL1Block)
	res, err := s.baseFee.Calc1559BaseFee(
		nil,
		s.config.GasTargetPerL1Block,
		s.config.BasefeeAdjustmentQuotient,
		s.gasExcess,
		_gasIssuance,
		_parentGasUsed,
	)
	s.Nil(err)
	fmt.Printf(
		"numL1Blocks: %d\t\t, gasExcess: %d\t\t, parentGasUsed: %d\t\t, baseFee: %d\n",
		numL1Blocks,
		s.gasExcess,
		_parentGasUsed,
		res.Basefee.Uint64(),
	)
	s.gasExcess = res.GasExcess
}

type testNode struct {
	numL1Blocks    uint64
	gasUsed        uint32
	growthRate     uint32
	times          int
	resetGasExcess bool
}

func (s *BaseFeeSuite) TestDecreaseCalc1559BaseFee() {
	var testData = []*testNode{
		{
			numL1Blocks:    0,
			gasUsed:        70000000,
			growthRate:     110,
			times:          30,
			resetGasExcess: true,
		},
		{
			numL1Blocks:    0,
			gasUsed:        1110412930,
			growthRate:     90,
			times:          70,
			resetGasExcess: false,
		},
		{
			numL1Blocks:    1,
			gasUsed:        70000000,
			growthRate:     110,
			times:          30,
			resetGasExcess: true,
		},
		{
			numL1Blocks:    1,
			gasUsed:        1110412930,
			growthRate:     80,
			times:          90,
			resetGasExcess: false,
		},
	}
	for _, val := range testData {
		if val.resetGasExcess {
			s.gasExcess = 1
		}
		times := val.times
		for gasUsed := val.gasUsed; times > 0 && gasUsed < math.MaxUint32; gasUsed = gasUsed / 100 * val.growthRate {
			times--
			s.testCalc1559BaseFee(val.numL1Blocks, gasUsed)
		}
		fmt.Printf("\n\n")
	}
}

func (s *BaseFeeSuite) TestIncreaseCalc1559BaseFee() {
	testData := []*testNode{
		{
			numL1Blocks:    0,
			gasUsed:        70000000,
			growthRate:     110,
			times:          100,
			resetGasExcess: true,
		},
		{
			numL1Blocks:    1,
			gasUsed:        800000000,
			growthRate:     110,
			times:          40,
			resetGasExcess: true,
		},
		{
			numL1Blocks:    2,
			gasUsed:        800800000,
			growthRate:     110,
			times:          40,
			resetGasExcess: true,
		},
		{
			numL1Blocks:    3,
			gasUsed:        801200000,
			growthRate:     105,
			times:          50,
			resetGasExcess: true,
		},
		{
			numL1Blocks:    4,
			gasUsed:        801600000,
			growthRate:     105,
			times:          60,
			resetGasExcess: true,
		},
	}
	for _, val := range testData {
		if val.resetGasExcess {
			s.gasExcess = 1
		}
		times := val.times
		for gasUsed := val.gasUsed; times > 0 && gasUsed < math.MaxUint32; gasUsed = gasUsed / 100 * val.growthRate {
			times--
			s.testCalc1559BaseFee(val.numL1Blocks, gasUsed)
		}
		fmt.Printf("\n\n")
	}
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
}

func TestDriverTestSuite(t *testing.T) {
	suite.Run(t, new(BaseFeeSuite))
}
