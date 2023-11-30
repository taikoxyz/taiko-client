#!/bin/bash

set -eou pipefail

DIR=$(
    cd $(dirname ${BASH_SOURCE[0]})
    pwd
)

# Download solc for PlonkVerifier
$TAIKO_MONO_DIR/packages/protocol/script/download_solc.sh

echo "Starting testnet..."

docker compose -f $TESTNET_CONFIG down -v --remove-orphans &>/dev/null
docker compose -f $TESTNET_CONFIG up -d

echo "Waiting till testnet nodes fully started..."

NODE_URL=localhost:18545 $DIR/../util/wait_for_node.sh
NODE_URL=localhost:28545 $DIR/../util/wait_for_node.sh

# Get the hash of L2 genesis.
L2_GENESIS_HASH=$(
    curl \
        --silent \
        -X POST \
        -H "Content-Type: application/json" \
        -d '{"jsonrpc":"2.0","id":0,"method":"eth_getBlockByNumber","params":["0x0", false]}' \
        localhost:28545 | jq .result.hash | sed 's/\"//g'
)

# Deploy Taiko protocol.
cd $TAIKO_MONO_DIR/packages/protocol &&
    PRIVATE_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
    GUARDIAN_PROVERS=0x70997970C51812dc3A010C7d01b50e0d17dc79C8,0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266,0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC,0x90F79bf6EB2c4f870365E785982E1f101E93b906,0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65 \
    OWNER=0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC \
    TAIKO_L2_ADDRESS=0x1670010000000000000000000000000000010001 \
    L2_SIGNAL_SERVICE=0x1670010000000000000000000000000000010005 \
    SECURITY_COUNCIL=0x23618e81E3f5cdF7f54C3d65f7FBc0aBf5B21E8f \
    SHARED_ADDRESS_MANAGER=0x0000000000000000000000000000000000000000 \
    PROPOSER=0x0000000000000000000000000000000000000000 \
    PROPOSER_ONE=0x0000000000000000000000000000000000000000 \
    TAIKO_TOKEN_PREMINT_RECIPIENT=0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
    TAIKO_TOKEN_NAME="Taiko Token Test" \
    TAIKO_TOKEN_SYMBOL=TTKOt \
    L2_GENESIS_HASH=$L2_GENESIS_HASH \
    MIN_GUARDIANS=5 \
    forge script script/DeployOnL1.s.sol:DeployOnL1 \
        --fork-url http://localhost:18545 \
        --broadcast \
        --ffi \
        -vvvvv \
        --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
        --block-gas-limit 100000000
