package bindings

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

var (
	// Account address and private key of golden touch account, defined in protocol's LibAnchorSignature.
	// ref: https://github.com/taikoxyz/taiko-mono/blob/main/packages/protocol/contracts/libs/LibAnchorSignature.sol
	GoldenTouchAddress = common.HexToAddress("0x0000777735367b36bC9B61C50022d9D0700dB4Ec")
	GoldenTouchPrivKey = "0x92954368afd3caa1f3ce3ead0069c1af414054aefe1ef9aeacc1bf426222ce38"
)

// ProtocolConstants contains some constants used by Taiko protocol, defined in protocol's LibConstants.
// NOTE: this struct *MUST* match the return values of TaikoL1.getConstants method.
// ref: https://github.com/taikoxyz/taiko-mono/blob/main/packages/protocol/contracts/libs/LibConstants.sol
type ProtocolConstants struct {
	ZKProofsPerBlock         *big.Int    // uint256 K_ZKPROOFS_PER_BLOCK
	ChainID                  *big.Int    // uint256 TAIKO_CHAIN_ID
	MaxProposedBlocks        *big.Int    // uint256 TAIKO_MAX_PROPOSED_BLOCKS
	MaxVerificationsPerTx    *big.Int    // uint256 TAIKO_MAX_VERIFICATIONS_PER_TX
	CommitDelayConfirmations *big.Int    // uint256 K_COMMIT_DELAY_CONFIRMATIONS
	MaxProofsPerForkChoice   *big.Int    // uint256 TAIKO_MAX_PROOFS_PER_FORK_CHOICE
	BlockMaxGasLimit         *big.Int    // uint256 TAIKO_BLOCK_MAX_GAS_LIMIT
	BlockMaxTxs              *big.Int    // uint256 TAIKO_BLOCK_MAX_TXS
	BlockDeadendHash         common.Hash // bytes32 TAIKO_BLOCK_DEADEND_HASH
	TxListMaxBytes           *big.Int    // uint256 TAIKO_TXLIST_MAX_BYTES
	TxMinGasLimit            *big.Int    // uint256 TAIKO_TX_MIN_GAS_LIMIT
	AnchorTxGasLimit         *big.Int    // uint256 V1_ANCHOR_TX_GAS_LIMIT
	AnchorTxSelector         [4]byte     // bytes4  V1_ANCHOR_TX_SELECTOR
	InvalidateBlockLogTopic  [32]byte    // bytes32 V1_INVALIDATE_BLOCK_LOG_TOPIC
}
