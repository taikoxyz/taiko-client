package submitter

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie"
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
	log.Crit("proofList.Delete not supported")
	return nil
}

// generateTrieProof generates a merkle proof of the i'th item in a MPT of given
// elements.
func generateTrieProof(list types.DerivableList, i uint64) (common.Hash, []byte, error) {
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
