#!/bin/bash

set -eou pipefail

DIR=$(
    cd $(dirname ${BASH_SOURCE[0]})
    pwd
)

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

# Premint ETHs for the test account.
curl \
    --silent \
    -X POST \
    -H "Content-Type: application/json" \
    -d '{"jsonrpc":"2.0","id":0,"method":"hardhat_setBalance","params":["0xDf08F82De32B8d460adbE8D72043E3a7e25A3B39", "0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"]}' \
    localhost:18545

echo ""
echo "Premint ETHs to the contracts deployer"

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

docker cp \
    $(docker-compose -f $TESTNET_CONFIG ps -q l2_execution_engine):/deployments/mainnet.json \
    $DIR/deployments/mainnet.json

L2_GENESIS_ALLOC=$(cat $DIR/deployments/mainnet.json)

TAIKO_L2_CONTRACT_ADDRESS=$(echo $L2_GENESIS_ALLOC | jq 'to_entries[] | select(.value.contractName=="TaikoL2") | .key' | sed 's/\"//g')

# Deploy Taiko protocol.
cd $TAIKO_MONO_DIR/packages/protocol &&
    LOG_LEVEL=debug \
    PRIVATE_KEY=0x2bdd21761a483f71054e14f5b827213567971c676928d9a1808cbfa4b7501200 \
    npx hardhat deploy_L1 \
    --network l1_test \
    --dao-vault 0xdf08f82de32b8d460adbe8d72043e3a7e25a3b39 \
    --team-vault 0xdf08f82de32b8d460adbe8d72043e3a7e25a3b39 \
    --oracle-prover 0xdf08f82de32b8d460adbe8d72043e3a7e25a3b39 \
    --l2-genesis-block-hash $L2_GENESIS_HASH \
    --taiko-l2 $TAIKO_L2_CONTRACT_ADDRESS \
    --confirmations 1
