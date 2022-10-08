package prover

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/taikochain/taiko-client/bindings"
)

const (
	proveBlockTxGasLimit = 1000000 // TODO: tune this value
)

// proofList represents a merkle proof to verify the inclusion of a key-value pair
type proofList [][]byte

// Put implements ethdb.KeyValueWriter interface.
func (n *proofList) Put(key []byte, value []byte) error {
	*n = append(*n, value)
	return nil
}

// Delete implements ethdb.KeyValueWriter interface.
func (n *proofList) Delete(key []byte) error {
	panic("proofList.Delete not supported")
}

// getBlockMetadataByID fetches the L2 block metadata with given block ID.
// TODO: add start height and end height in filter options.
func (p *Prover) getBlockMetadataByID(blockID *big.Int) (*bindings.LibDataBlockMetadata, error) {
	iter, err := p.taikoL1.FilterBlockProposed(nil, []*big.Int{blockID})
	if err != nil {
		return nil, err
	}

	for iter.Next() {
		return &iter.Event.Meta, nil
	}

	return nil, fmt.Errorf("block metadata not found, id: %d", blockID)
}

// generateTrieProof generates a merkle proof of the i'th item in a MPT of given
// elements.
func generateTrieProof[T types.DerivableList](list T, i uint64) (common.Hash, []byte, error) {
	trie := trie.NewEmpty(trie.NewDatabase(nil))

	types.DeriveSha(list, trie)

	var proof proofList
	if err := trie.Prove(rlp.AppendUint64([]byte{}, i), 0, &proof); err != nil {
		return common.Hash{}, nil, err
	}

	proofBytes, err := rlp.EncodeToBytes([][]byte(proof))
	if err != nil {
		return common.Hash{}, nil, err
	}

	return trie.Hash(), proofBytes, nil
}

// getProveBlockTxOpts creates a bind.TransactOpts instance with the sender's signatures.
func (p *Prover) getProveBlockTxOpts(ctx context.Context) (*bind.TransactOpts, error) {
	networkID, err := p.l1RPC.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	opts, err := bind.NewKeyedTransactorWithChainID(&p.cfg.L1ProverPrivKey, networkID)
	if err != nil {
		return nil, err
	}

	opts.GasLimit = proveBlockTxGasLimit

	return opts, nil
}

func (p *Prover) waitForL1Origin(ctx context.Context, blockID *big.Int) (*rawdb.L1Origin, error) {
	var (
		l1Origin *rawdb.L1Origin
		err      error
	)

	ticker := time.NewTicker(time.Second)
	timeout := time.After(time.Minute)
	defer ticker.Stop()

	log.Info("Start fetching L1Origin from L2 node", "blockID", blockID)

	for {
		select {
		case <-timeout:
			return nil, fmt.Errorf("fetch L1Origin timeout")
		case <-ticker.C:
			l1Origin, err = p.l2RPC.L1OriginByID(ctx, blockID)
			if err != nil {
				log.Warn("Failed to fetch L1Origin from L2 node", "blockID", blockID, "error", err)
				continue
			}

			if l1Origin == nil {
				continue
			}

			return l1Origin, nil
		}
	}
}
