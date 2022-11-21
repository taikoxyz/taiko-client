package prover

import (
	"context"
	"math/rand"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/taikoxyz/taiko-client/bindings"
)

func (s *ProverTestSuite) TestValidateAnchorTx() {
	wrongPrivKey, err := crypto.HexToECDSA("2bdd21761a483f71054e14f5b827213567971c676928d9a1808cbfa4b7501200")
	s.Nil(err)

	// 0x92954368afd3caa1f3ce3ead0069c1af414054aefe1ef9aeacc1bf426222ce38
	goldenTouchPriKey, err := crypto.HexToECDSA(bindings.GoldenTouchPrivKey[2:])
	s.Nil(err)

	// invalid To
	tx := types.NewTransaction(0, common.BytesToAddress(randBytes(1024)), common.Big0, 0, common.Big0, []byte{})
	s.ErrorContains(s.p.validateAnchorTx(context.Background(), tx), "invalid TaikoL2.anchor transaction to")

	// invalid sender
	dynamicFeeTxTx := &types.DynamicFeeTx{
		ChainID:    s.p.rpc.L2ChainID,
		Nonce:      0,
		GasTipCap:  common.Big1,
		GasFeeCap:  common.Big1,
		Gas:        1,
		To:         &s.p.cfg.TaikoL2Address,
		Value:      common.Big0,
		Data:       []byte{},
		AccessList: types.AccessList{},
	}

	signer := types.LatestSignerForChainID(s.p.rpc.L2ChainID)
	tx = types.MustSignNewTx(wrongPrivKey, signer, dynamicFeeTxTx)

	s.ErrorContains(
		s.p.validateAnchorTx(context.Background(), tx), "invalid TaikoL2.anchor transaction sender",
	)

	// invalid method selector
	tx = types.MustSignNewTx(goldenTouchPriKey, signer, dynamicFeeTxTx)
	s.ErrorContains(s.p.validateAnchorTx(context.Background(), tx), "invalid TaikoL2.anchor transaction selector")
}

func randBytes(l uint64) []byte {
	b := make([]byte, l)
	rand.Read(b)
	return b
}
