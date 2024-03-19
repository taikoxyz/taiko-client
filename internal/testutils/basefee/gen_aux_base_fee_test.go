package basefee

import (
	"bytes"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/taikoxyz/taiko-client/internal/utils"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"github.com/taikoxyz/taiko-client/bindings"
)

type BaseFeeSuite struct {
	suite.Suite
	//testutils.ClientTestSuite

	buffer bytes.Buffer

	baseFee *AuxBaseFee

	gasExcess uint64
	config    bindings.TaikoL2Config
}

func (s *BaseFeeSuite) testCalc1559BaseFee(numL1Blocks, gasExcess uint64, _parentGasUsed uint32) (uint64, uint64) {
	_gasIssuance := numL1Blocks * uint64(s.config.GasTargetPerL1Block)
	res, err := s.baseFee.Calc1559BaseFee(
		nil,
		s.config.GasTargetPerL1Block,
		s.config.BasefeeAdjustmentQuotient,
		gasExcess,
		_gasIssuance,
		_parentGasUsed,
	)
	s.Nil(err)
	/*fmt.Printf(
		"numL1Blocks: %d\t, gasExcess: %d\t\t\t, parentGasUsed: %d\t\t, baseFee: %d\n",
		numL1Blocks,
		s.gasExcess,
		_parentGasUsed,
		res.Basefee.Uint64(),
	)*/
	return res.Basefee.Uint64(), res.GasExcess
}

type testNode struct {
	numL1Blocks    uint64
	gasUsed        uint32
	growthRate     uint32
	times          int
	resetGasExcess bool
}

func (s *BaseFeeSuite) TestDecreaseCalc1559BaseFee() {
	for _, numL1Blocks := range []int{1, 2, 4} {
		s.buffer.Reset()
		s.buffer.Write([]byte("blockTime,gasExcess,gasUsed,baseFee\n"))
		var (
			gasExcess uint64 = 1
			baseFee   uint64 = 1
		)
		for gasUsed := 0; gasUsed < (30_000_000 * 8); gasUsed += 100_000 {
			baseFee, gasExcess = s.testCalc1559BaseFee(uint64(numL1Blocks), gasExcess, uint32(gasUsed))
			s.buffer.Write([]byte(fmt.Sprintf("%d,%d,%d,%d\n", numL1Blocks*12, gasExcess, gasUsed, baseFee)))
		}
		for gasUsed := 30_000_000 * 8; gasUsed >= 0; gasUsed -= 100_000 {
			baseFee, gasExcess = s.testCalc1559BaseFee(uint64(numL1Blocks), gasExcess, uint32(gasUsed))
			s.buffer.Write([]byte(fmt.Sprintf("%d,%d,%d,%d\n", numL1Blocks*12, gasExcess, gasUsed, baseFee)))
		}
		s.Nil(os.WriteFile(fmt.Sprintf("/Users/huan/Documents/taiko/%d_basefee.csv", numL1Blocks*12), s.buffer.Bytes(), 0644))
	}
}

func (s *BaseFeeSuite) SetupTest() {
	utils.LoadEnv()
	l1Client, err := rpc.NewEthClient(context.Background(), os.Getenv("L1_NODE_WS_ENDPOINT"), time.Second*30)
	s.Nil(err)

	//priv, err := crypto.ToECDSA(common.FromHex(os.Getenv("L1_PROPOSER_PRIVATE_KEY")))
	//s.Nil(err)
	//send, err := sender.NewSender(context.Background(), nil, l1Client, priv)
	//s.Nil(err)
	//opts := send.GetOpts()
	//addr, tx, baseFee, err := DeployAuxBaseFee(opts, l1Client)
	//s.Nil(err)
	//s.baseFee = baseFee
	//id, err := send.SendTransaction(tx)
	//s.Nil(err)
	//confirm := <-send.TxToConfirmChannel(id)
	//s.Nil(confirm.Err)
	//fmt.Println("contract address: ", addr.String())

	s.baseFee, err = NewAuxBaseFee(common.HexToAddress("0x4C2F7092C2aE51D986bEFEe378e50BD4dB99C901"), l1Client)

	s.gasExcess = 1
	s.config = bindings.TaikoL2Config{
		BasefeeAdjustmentQuotient: 4,
		GasTargetPerL1Block:       60000000,
	}

	//taikoL2 := s.RPCClient.TaikoL2
	//s.gasExcess, err = taikoL2.GasExcess(nil)
	//s.Nil(err)
	//fmt.Println("gasExcess: ", s.gasExcess)
	//
	//s.config, err = taikoL2.GetConfig(nil)
	//s.Nil(err)
	//fmt.Println("config: ", s.config.BasefeeAdjustmentQuotient, s.config.GasTargetPerL1Block)
}

func TestDriverTestSuite(t *testing.T) {
	suite.Run(t, new(BaseFeeSuite))
}
