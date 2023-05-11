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
	errProtocolAddressMismatch = errors.New("special prover private key does not match protocol setting")
)

// SpecialProofProducer is responsible for generating a fake "zkproof" consisting
// of a signature of the evidence.
type SpecialProofProducer struct {
	rpc               *rpc.Client
	proverPrivKey     *ecdsa.PrivateKey
	anchorTxValidator *anchorTxValidator.AnchorTxValidator
	proofTimeTarget   time.Duration
	graffiti          [32]byte
	isSystemProver    bool
}

// NewSpecialProofProducer creates a new NewSpecialProofProducer instance, which can be either
// an oracle proof producer, or a system proofproducer.
func NewSpecialProofProducer(
	rpc *rpc.Client,
	proverPrivKey *ecdsa.PrivateKey,
	taikoL2Address common.Address,
	proofTimeTarget time.Duration,
	protocolSpecialProverAddress common.Address,
	graffiti string,
	isSystemProver bool,
) (*SpecialProofProducer, error) {
	proverAddress := crypto.PubkeyToAddress(proverPrivKey.PublicKey)
	if proverAddress != protocolSpecialProverAddress {
		return nil, errProtocolAddressMismatch
	}

	anchorValidator, err := anchorTxValidator.New(taikoL2Address, rpc.L2ChainID, rpc)
	if err != nil {
		return nil, err
	}

	var graffitiBytes [32]byte
	copy(graffitiBytes[:], []byte(graffiti))

	return &SpecialProofProducer{rpc, proverPrivKey, anchorValidator, proofTimeTarget, graffitiBytes, isSystemProver}, nil
}

// RequestProof implements the ProofProducer interface.
func (p *SpecialProofProducer) RequestProof(
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

	block, err := p.rpc.L2.BlockByHash(ctx, header.Hash())
	if err != nil {
		return fmt.Errorf("failed to get L2 block with given hash %s: %w", header.Hash(), err)
	}

	anchorTx := block.Transactions()[0]
	if err := p.anchorTxValidator.ValidateAnchorTx(ctx, anchorTx); err != nil {
		return fmt.Errorf("invalid anchor transaction: %w", err)
	}

	signalRoot, err := p.rpc.GetStorageRoot(ctx, p.rpc.L2GethClient, opts.L2SignalService, block.Number())
	if err != nil {
		return fmt.Errorf("error getting storageroot: %w", err)
	}

	if err := p.anchorTxValidator.ValidateAnchorTx(ctx, anchorTx); err != nil {
		return fmt.Errorf("invalid anchor transaction: %w", err)
	}

	parent, err := p.rpc.L2.BlockByHash(ctx, block.ParentHash())
	if err != nil {
		return err
	}

	blockInfo, err := p.rpc.TaikoL1.GetBlock(nil, blockID)
	if err != nil {
		return err
	}

	// the only difference from a client perspective when generating a special proof,
	// either an oracle proof or a system proof, is the prover address which should be set to 1
	// if system prover, and 0 if oracle prover, and the protocol will use that to decide
	// whether a proof can be overwritten or not.
	var prover common.Address
	if p.isSystemProver {
		prover = common.HexToAddress("0x0000000000000000000000000000000000000001")
	} else {
		prover = common.HexToAddress("0x0000000000000000000000000000000000000000")
	}
	// signature should be done with proof set to nil, verifierID set to 0,
	// and prover set to 0 address.
	evidence := &encoding.TaikoL1Evidence{
		MetaHash:      blockInfo.MetaHash,
		ParentHash:    block.ParentHash(),
		BlockHash:     block.Hash(),
		SignalRoot:    signalRoot,
		Graffiti:      p.graffiti,
		Prover:        prover,
		ParentGasUsed: uint32(parent.GasUsed()),
		GasUsed:       uint32(block.GasUsed()),
		VerifierId:    0,
		Proof:         []byte{},
	}

	proof, err := hashAndSignForSpecialProof(evidence, p.proverPrivKey)
	if err != nil {
		return fmt.Errorf("failed to sign evidence: %w", err)
	}

	var (
		delay     time.Duration = 0
		now                     = time.Now()
		blockTime               = time.Unix(int64(block.Time()), 0)
	)
	if now.Before(blockTime.Add(p.proofTimeTarget)) {
		delay = blockTime.Add(p.proofTimeTarget).Sub(now)
	}

	log.Info("Oracle proof submission delay", "delay", delay)

	time.AfterFunc(delay, func() {
		resultCh <- &ProofWithHeader{
			BlockID: blockID,
			Header:  header,
			Meta:    meta,
			ZkProof: proof,
		}
	})

	return nil
}

// HashSignAndSetEvidenceForSpecialProof hashes and signs the TaikoL1Evidence according to the
// protocol spec to generate a special proof via the signature and v value.
func hashAndSignForSpecialProof(
	evidence *encoding.TaikoL1Evidence,
	privateKey *ecdsa.PrivateKey,
) ([]byte, error) {
	inputToSign, err := encoding.EncodeProveBlockInput(evidence)
	if err != nil {
		return nil, fmt.Errorf("failed to encode TaikoL1.proveBlock inputs: %w", err)
	}

	hashed := crypto.Keccak256Hash(inputToSign)

	sig, err := crypto.Sign(hashed.Bytes(), privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to sign TaikoL1Evidence: %w", err)
	}

	// add 27 to be able to be ecrecover in solidity
	sig[64] = uint8(int(sig[64])) + 27

	return sig, nil
}

// Cancel cancels an existing proof generation.
// Since oracle and system proofs are not "real" proofs, there is nothing to cancel.
func (d *SpecialProofProducer) Cancel(ctx context.Context, blockID *big.Int) error {
	return nil
}
