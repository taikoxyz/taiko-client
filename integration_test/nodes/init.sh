#!/bin/bash

set -eou pipefail

DIR=$(
    cd $(dirname ${BASH_SOURCE[0]})
    pwd
)

# Download solc for PlonkVerifier
protocol_dir=$TAIKO_MONO_DIR/packages/protocol
solc_bin=${protocol_dir}/bin/solc

if [ -f "${solc_bin}" ]; then
  echo "solc already exists, skipping download."
else
  mkdir -p "$(dirname ${solc_bin})"
  VERSION=v0.8.18

  if [ "$(uname)" = 'Darwin' ]; then
    SOLC_FILE_NAME=solc-macos
  elif [ "$(uname)" = 'Linux' ]; then
    SOLC_FILE_NAME=solc-static-linux
  else
    echo "unsupported platform $(uname)"
    exit 1
  fi

  wget -O "${solc_bin}" https://github.com/ethereum/solidity/releases/download/$VERSION/$SOLC_FILE_NAME
  chmod +x "${solc_bin}"
fi

echo "Starting testnet..."

docker compose -f $TESTNET_CONFIG down -v --remove-orphans &>/dev/null
docker compose -f $TESTNET_CONFIG up -d

if [ "$COMPILE_PROTOCOL" == "true" ]; then
    cd $TAIKO_MONO_DIR/packages/protocol && yarn run clean && yarn run compile
    cd -
fi

echo "Waiting till testnet nodes fully started..."

NODE_URL=localhost:18545 $DIR/../util/wait_for_node.sh
NODE_URL=localhost:28545 $DIR/../util/wait_for_node.sh
rm -rf $DIR/deployments/mainnet.json

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
    PRIVATE_KEY=ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
    ORACLE_PROVER=0x70997970C51812dc3A010C7d01b50e0d17dc79C8 \
    SOLO_PROPOSER=0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
    OWNER=0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC \
    TREASURY=0xdf09A0afD09a63fb04ab3573922437e1e637dE8b \
    TAIKO_L2_ADDRESS=0x1000777700000000000000000000000000000001 \
    L2_SIGNAL_SERVICE=0x1000777700000000000000000000000000000007 \
    SHARED_SIGNAL_SERVICE=0x0000000000000000000000000000000000000000 \
    TAIKO_TOKEN_PREMINT_RECIPIENTS=0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266,0x70997970C51812dc3A010C7d01b50e0d17dc79C8 \
    TAIKO_TOKEN_PREMINT_AMOUNTS=$PREMINT_TOKEN_AMOUNT,$PREMINT_TOKEN_AMOUNT \
    L2_GENESIS_HASH=$L2_GENESIS_HASH \
    L2_CHAIN_ID=167001 \
    forge script script/DeployOnL1.s.sol:DeployOnL1 \
        --fork-url http://localhost:18545 \
        --broadcast \
        --ffi \
        -vvvv \
        --block-gas-limit 100000000
