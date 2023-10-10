package evidence

import (
	"context"
	"encoding/binary"
	"fmt"

	"github.com/ethereum/go-ethereum/log"
	"github.com/taikoxyz/taiko-client/bindings/encoding"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	anchorTxValidator "github.com/taikoxyz/taiko-client/prover/anchor_tx_validator"
	proofProducer "github.com/taikoxyz/taiko-client/prover/proof_producer"
)

// Assembler is responsible for assembling evidence for the given L2 block proof.
type Assembler struct {
	rpc               *rpc.Client
	anchorTxValidator *anchorTxValidator.AnchorTxValidator
	graffiti          [32]byte
}

// NewAssembler creates a new EvidenceAssembler instance.
func NewAssembler(cli *rpc.Client, anchorTxValidator *anchorTxValidator.AnchorTxValidator, graffiti string) *Assembler {
	return &Assembler{
		rpc:               cli,
		anchorTxValidator: anchorTxValidator,
		graffiti:          rpc.StringToBytes32(graffiti),
	}
}

// assembleEvidence assembles the evidence for the given L2 block proof.
func (a *Assembler) AssembleEvidence(
	ctx context.Context,
	proofWithHeader *proofProducer.ProofWithHeader,
) (*encoding.BlockEvidence, error) {
	var (
		blockID = proofWithHeader.BlockID
		header  = proofWithHeader.Header
		proof   = proofWithHeader.Proof
	)

	log.Info(
		"Assemble new evidence",
		"blockID", blockID,
		"parentHash", proofWithHeader.Header.ParentHash,
		"hash", proofWithHeader.Header.Hash(),
		"signalRoot", proofWithHeader.Opts.SignalRoot,
		"tier", proofWithHeader.Tier,
	)

	// Get the corresponding L2 block.
	block, err := a.rpc.L2.BlockByHash(ctx, header.Hash())
	if err != nil {
		return nil, fmt.Errorf("failed to get L2 block with given hash %s: %w", header.Hash(), err)
	}

	if block.Transactions().Len() == 0 {
		return nil, fmt.Errorf("invalid block without anchor transaction, blockID %s", blockID)
	}

	// Validate TaikoL2.anchor transaction inside the L2 block.
	anchorTx := block.Transactions()[0]
	if err := a.anchorTxValidator.ValidateAnchorTx(ctx, anchorTx); err != nil {
		return nil, fmt.Errorf("invalid anchor transaction: %w", err)
	}

	// Get and validate this anchor transaction's receipt.
	if _, err = a.anchorTxValidator.GetAndValidateAnchorTxReceipt(ctx, anchorTx); err != nil {
		return nil, fmt.Errorf("failed to fetch anchor transaction receipt: %w", err)
	}

	evidence := &encoding.BlockEvidence{
		MetaHash:   proofWithHeader.Opts.MetaHash,
		ParentHash: proofWithHeader.Opts.ParentHash,
		BlockHash:  proofWithHeader.Opts.BlockHash,
		SignalRoot: proofWithHeader.Opts.SignalRoot,
		Graffiti:   a.graffiti,
		Tier:       proofWithHeader.Tier,
		Proof:      proof,
	}

	if proofWithHeader.Tier == encoding.TierPseZkevmID {
		circuitsIdx, err := proofProducer.DegreeToCircuitsIdx(proofWithHeader.Degree)
		if err != nil {
			return nil, err
		}
		evidence.Proof = append(uint16ToBytes(circuitsIdx), evidence.Proof...)
	}

	return evidence, nil
}

// uint16ToBytes converts an uint16 to bytes.
func uint16ToBytes(i uint16) []byte {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, i)
	return b
}
