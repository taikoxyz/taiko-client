package utils

import (
	"errors"
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/params"
)

func GWeiToWei(gwei float64) (*big.Int, error) {
	if math.IsNaN(gwei) || math.IsInf(gwei, 0) {
		return nil, fmt.Errorf("invalid gwei value: %v", gwei)
	}

	// convert float GWei value into integer Wei value
	wei, _ := new(big.Float).Mul(
		big.NewFloat(gwei),
		big.NewFloat(params.GWei)).
		Int(nil)

	if wei.Cmp(abi.MaxUint256) == 1 {
		return nil, errors.New("gwei value larger than max uint256")
	}

	return wei, nil
}

func EtherToWei(ether float64) (*big.Int, error) {
	if math.IsNaN(ether) || math.IsInf(ether, 0) {
		return nil, fmt.Errorf("invalid ether value: %v", ether)
	}

	// convert float GWei value into integer Wei value
	wei, _ := new(big.Float).Mul(
		big.NewFloat(ether),
		big.NewFloat(params.GWei)).
		Int(nil)

	if wei.Cmp(abi.MaxUint256) == 1 {
		return nil, errors.New("gwei value larger than max uint256")
	}

	return wei, nil
}

func WeiToEther(wei *big.Int) *big.Float {
	return new(big.Float).Quo(new(big.Float).SetInt(wei), new(big.Float).SetInt(big.NewInt(params.Ether)))
}
