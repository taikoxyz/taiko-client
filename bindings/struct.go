package bindings

import (
	"math/big"
)

// ProtocolConstants contains some constants used by Taiko protocol, defined in protocol's LibConstants.
// NOTE: this struct *MUST* match the return values of TaikoL1.getConstants method.
// ref: https://github.com/taikoxyz/taiko-mono/blob/main/packages/protocol/contracts/libs/LibConstants.sol
type ProtocolConstants struct {
	ZKProofsPerBlock         *big.Int // uint256 K_ZKPROOFS_PER_BLOCK
	ChainID                  *big.Int // uint256 K_CHAIN_ID
	MaxNumBlocks             *big.Int // uint256 K_MAX_NUM_BLOCKS
	MaxVerificationsPerTx    *big.Int // uint256 K_MAX_VERIFICATIONS_PER_TX
	CommitDelayConfirmations *big.Int // uint256 K_COMMIT_DELAY_CONFIRMS
	MaxProofsPerForkChoice   *big.Int // uint256 K_MAX_PROOFS_PER_FORK_CHOICE
	BlockMaxGasLimit         *big.Int // uint256 K_BLOCK_MAX_GAS_LIMIT
	BlockMaxTxs              *big.Int // uint256 K_BLOCK_MAX_TXS
	TxListMaxBytes           *big.Int // uint256 K_TXLIST_MAX_BYTES
	TxMinGasLimit            *big.Int // uint256 K_TX_MIN_GAS_LIMIT
	AnchorTxGasLimit         *big.Int // uint256 K_ANCHOR_TX_GAS_LIMIT
}

// ProtocolStateVariables contains some state variables used by Taiko protocol, defined in protocol's LibData.
// NOTE: this struct *MUST* match the return values of TaikoL1.getStateVariables method.
// ref: https://github.com/taikoxyz/taiko-mono/blob/main/packages/protocol/contracts/L1/LibData.sol
type ProtocolStateVariables struct {
	GenesisHeight        uint64
	LatestVerifiedHeight uint64
	LatestVerifiedID     uint64
	NextBlockID          uint64
}
