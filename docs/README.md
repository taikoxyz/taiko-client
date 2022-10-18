# Overview

The compiled binary `bin/taiko-client` is the main entrypoint which including three sub-commands:

- `proposer`: propose new transactions from the L2 node's transaction pool to the `TaikoL1` protocol contract.
- `driver`: keep the L2 node's chain in sync with `TaikoL1` protocol contact by directing the L2 node's [execution engine](https://github.com/ethereum/execution-apis/tree/main/src/engine).
- `prover`: request ZK proofs from zkEVM and then send transactions to prove the proposed blocks valid or invalid.

## Proposer

### Proposing Strategy

Since tokenomics has not been fully implemented in Taiko protocol, the current proposing strategy is simply based on time interval (which is a required command line flag).

### Proposing Operation

The proposing operation contains several steps:

1. Fetching the pending transactions from L2 node through `txpool_content` RPC method.
2. Since the Taiko protocol restricts the max size of each proposed txList, if there are too many pending transactions in L2 node, split them into several smaller txLists.
3. Commit hashes of the splitted txLists by sending `TaikoL1.commitBlock` transactions to L1.
4. Wait `LibConstants.TAIKO_COMMIT_DELAY_CONFIRMATIONS` (currently `4`) L1 blocks confirmations.
5. Propose all txLists by sending `TaikoL1.proposeBlock` transactions.

## Driver

### Engine API

Driver directs the L2 node's execution engine to insert new blocks or reorg local chain through Engine API.

> Reference link, Engine API Specification: <https://github.com/ethereum/execution-apis/blob/main/src/engine/specification.md>

### Chain Syncrhonization Process

> NOTE: Taiko protocol allows a block's timestamp equals to its parent block's timestamp, which is different from the original Ethereum protocol. So it's fine that there are two `TaikoL1.proposeBlock` transactions included in one L1 block.

Driver subscribes `TaikoL1.BlockProposed` events, and when a new proposed block comes:

1. Get the corresponding `TaikoL1.proposeBlock` L1 transaction.
2. Decode the txList and block metadata from the transaction's calldata.
3. Check whether the txList is valid based on the rules defined in Taiko protocol.

If the txList is valid:

4. Assemble a deterministic `V1TaikoL2.anchor` transaction based on the rules defined in protocol, put it as the first transaction in the proposed txList.
5. Use this txList and the decoded block metadata to assemble a deterministic L2 block.
6. Direct L2 nodes' execution engine to insert this assembled block and set it as current chain's head via Engine API.

If the txList is invalid:

4. Create a `V1TaikoL2.invalidateBlock` transaction and then assemble a L2 block only including this transaction.
5. Direct L2 nodes' execution engine to insert this block but not set it as the chain's head via Engine API.

> NOTE: For more detailed information about: block metadata, please see white paper's `5.2.2 Block Metadata`.

> NOTE: For more detailed information about txList validation rules, please see white paper's `5.3.1 Validation`.

> NOTE: For more detailed information about the `V1TaikoL2.anchor` transaction and proposed block's determination, please see white paper's `5.4.1 Construction of Anchor Transactions`

## Prover

### Proving Strategy

Since tokenomics has not been fully implemented in Taiko protocol, the prover software currently proves all proposed blocks.

### Proving Process

When a new proposed block comes:

1. Get the `TaikoL1.proposeBlock` L1 transaction calldata, decode it, and validate the txList, just like what the `driver` software does.
2. Wait util the corresponding block inserted by the L2 node's `driver` software.
3. Generate zk proof for that block asynchronously.

If the proven block has a valid txList:

4. Generate the merkel proof of the block's `V1TaikoL2.anchor` transaction to prove its existence in the `block.txRoot`'s [MPT](https://ethereum.org/en/developers/docs/data-structures-and-encoding/patricia-merkle-trie/) and also this transaction receipt's merkel proof in the `block.receiptRoot`'s MPT from L2 node.
5. Submit the `V1TaikoL2.anchor` transaction's RLP encoded bytes, its receipt's RLP encoded bytes, generated merkel proofs and zk proof to prove this block valid, by sending a `TaikoL1.proveBlock` transaction.

If the proven block has an invalid txList:

4. Generate the merkel proof of the block's `V1TaikoL2.invalidateBlock` transaction receipt to prove its existence in the `block.receiptRoot`'s MPT from L2 node.
5. Submit the `V1TaikoL2.invalidateBlock` transaction receipt's RLP encoded bytes, generated merkel proof and zk proof to prove this block invalid, by sending a `TaikoL1.proveBlockInvalid` transaction.

> NOTE: For more information about why we need these merkel proofs when proving, please see white paper's `5.5 Proving Blocks`
