package anchorTxConstructor

import (
	"context"
	"math/rand"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/suite"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/testutils"
)

type AnchorTxConstructorTestSuite struct {
	testutils.ClientTestSuite
	c *AnchorTxConstructor
}

func (s *AnchorTxConstructorTestSuite) SetupTest() {
	s.ClientTestSuite.SetupTest()
	protocolConfigs, err := s.RpcClient.TaikoL1.GetConfig(nil)
	s.Nil(err)
	c, err := New(
		s.RpcClient,
		protocolConfigs.AnchorTxGasLimit.Uint64(),
		bindings.GoldenTouchAddress,
		bindings.GoldenTouchPrivKey,
	)
	s.Nil(err)
	s.c = c
}

func (s *AnchorTxConstructorTestSuite) TestNewAnchorTransactor() {
	gasLimit := rand.Uint64()
	c, err := New(
		s.RpcClient,
		gasLimit, bindings.GoldenTouchAddress,
		bindings.GoldenTouchPrivKey,
	)
	s.Nil(err)

	opts, err := c.transactOpts(context.Background(), common.Big0)
	s.Nil(err)
	s.Equal(true, opts.NoSend)
	s.Equal(gasLimit, opts.GasLimit)
	s.Equal(common.Big0, opts.GasPrice)
	s.Equal(common.Big0, opts.Nonce)
	s.Equal(bindings.GoldenTouchAddress, opts.From)
}

func (s *AnchorTxConstructorTestSuite) TestSign() {
	// Payload 1
	hash := hexutil.MustDecode("0x44943399d1507f3ce7525e9be2f987c3db9136dc759cb7f92f742154196868b9")
	signatureBytes := testutils.SignatureFromRSV(
		"0x79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798",
		"0x782a1e70872ecc1a9f740dd445664543f8b7598c94582720bca9a8c48d6a4766",
		1,
	)
	pubKey, err := crypto.Ecrecover(hash, signatureBytes)
	s.Nil(err)
	isValid := crypto.VerifySignature(pubKey, hash, signatureBytes[:64])
	s.True(isValid)
	signed, err := s.c.signTxPayload(hash)
	s.Nil(err)
	s.Equal(signatureBytes, signed)

	// Payload 2
	hash = hexutil.MustDecode("0x663d210fa6dba171546498489de1ba024b89db49e21662f91bf83cdffe788820")
	signatureBytes = testutils.SignatureFromRSV(
		"0x79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798",
		"0x568130fab1a3a9e63261d4278a7e130588beb51f27de7c20d0258d38a85a27ff",
		1,
	)
	pubKey, err = crypto.Ecrecover(hash, signatureBytes)
	s.Nil(err)
	isValid = crypto.VerifySignature(pubKey, hash, signatureBytes[:64])
	s.True(isValid)
	signed, err = s.c.signTxPayload(hash)
	s.Nil(err)
	s.Equal(signatureBytes, signed)
}

func TestAnchorTxConstructorTestSuite(t *testing.T) {
	suite.Run(t, new(AnchorTxConstructorTestSuite))
}
