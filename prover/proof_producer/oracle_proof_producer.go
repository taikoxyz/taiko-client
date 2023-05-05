package producer

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/taikoxyz/taiko-client/bindings"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	anchorTxValidator "github.com/taikoxyz/taiko-client/prover/anchor_tx_validator"
)

var (
	errProtocolAddressMismatch = errors.New("oracle prover private key does not match protocol setting")
)

// OracleProducer is responsible for generating a fake "zkproof" consisting
// of a signature of the evidence.
type OracleProducer struct {
	rpc               *rpc.Client
	proverPrivKey     *ecdsa.PrivateKey
	anchorTxValidator *anchorTxValidator.AnchorTxValidator
	proofTimeTarget   time.Duration
}

// NewOracleProducer creates a new NewOracleProducer instance.
func NewOracleProducer(
	rpc *rpc.Client,
	proverPrivKey *ecdsa.PrivateKey,
	taikoL2Address common.Address,
	proofTimeTarget time.Duration,
	protocolOracleProverAddress common.Address,
) (*OracleProducer, error) {
	proverAddress := crypto.PubkeyToAddress(proverPrivKey.PublicKey)
	if proverAddress != protocolOracleProverAddress {
		return nil, errProtocolAddressMismatch
	}

	anchorValidator, err := anchorTxValidator.New(taikoL2Address, rpc.L2ChainID, rpc)
	if err != nil {
		return nil, err
	}

	return &OracleProducer{rpc, proverPrivKey, anchorValidator, proofTimeTarget}, nil
}

// RequestProof implements the ProofProducer interface.
func (d *OracleProducer) RequestProof(
	ctx context.Context,
	opts *ProofRequestOptions,
	blockID *big.Int,
	meta *bindings.TaikoDataBlockMetadata,
	header *types.Header,
	resultCh chan *ProofWithHeader,
) error {
	log.Info(
		"Request oracle proof",
		"blockID", blockID,
		"beneficiary", meta.Beneficiary,
		"height", header.Number,
		"hash", header.Hash(),
	)

	block, err := d.rpc.L2.BlockByHash(ctx, header.Hash())
	if err != nil {
		return fmt.Errorf("failed to get L2 block with given hash %s: %w", header.Hash(), err)
	}

	anchorTx := block.Transactions()[0]
	if err := d.anchorTxValidator.ValidateAnchorTx(ctx, anchorTx); err != nil {
		return fmt.Errorf("invalid anchor transaction: %w", err)
	}

	signalRoot, err := d.anchorTxValidator.GetAnchoredSignalRoot(ctx, anchorTx)
	if err != nil {
		return err
	}

	parent, err := d.rpc.L2.BlockByHash(ctx, block.ParentHash())
	if err != nil {
		return err
	}

	blockInfo, err := d.rpc.TaikoL1.GetBlock(nil, blockID)
	if err != nil {
		return err
	}

	// signature should be done with proof set to nil, verifierID set to 0,
	// and prover set to 0 address.
	evidence := &encoding.TaikoL1Evidence{
		MetaHash:      blockInfo.MetaHash,
		ParentHash:    block.ParentHash(),
		BlockHash:     block.Hash(),
		SignalRoot:    signalRoot,
		Graffiti:      [32]byte{},
		Prover:        common.HexToAddress("0x0000000000000000000000000000000000000000"),
		ParentGasUsed: uint32(parent.GasUsed()),
		GasUsed:       uint32(block.GasUsed()),
		VerifierId:    0,
		Proof:         nil,
	}

	proof, _, err := hashAndSignForOracleProof(evidence, d.proverPrivKey)
	if err != nil {
		return fmt.Errorf("failed to sign evidence: %w", err)
	}

	time.AfterFunc(d.proofTimeTarget, func() {
		resultCh <- &ProofWithHeader{
			BlockID: blockID,
			Header:  header,
			Meta:    meta,
			ZkProof: proof,
		}
	})

	return nil
}

// HashSignAndSetEvidenceForOracleProof hashes and signs the TaikoL1Evidence according to the
// protocol spec to generate an "oracle proof" via the signature and v value.
func hashAndSignForOracleProof(
	evidence *encoding.TaikoL1Evidence,
	privateKey *ecdsa.PrivateKey,
) ([]byte, uint8, error) {
	inputToSign, err := encoding.EncodeProveBlockInput(evidence)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to encode TaikoL1.proveBlock inputs: %w", err)
	}

	hashed := crypto.Keccak256Hash(inputToSign)

	sig, err := crypto.Sign(hashed.Bytes(), privateKey)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to sign TaikoL1Evidence: %w", err)
	}

	// add 27 to be able to be ecrecover in solidity
	v := uint8(int(sig[64])) + 27

	return sig, v, nil
}

// Cancel cancels an existing proof generation.
// Since Oracle proofs are not "real" proofs, there is nothing to cancel.
func (d *OracleProducer) Cancel(ctx context.Context, blockID *big.Int) error {
	return nil
}
