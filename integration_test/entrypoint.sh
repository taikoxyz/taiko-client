#!/bin/bash

set -eou pipefail

DIR=$(
    cd $(dirname ${BASH_SOURCE[0]})
    pwd
)

if ! command -v docker &>/dev/null 2>&1; then
    echo "ERROR: docker command not found"
    exit 1
fi

if ! docker info >/dev/null 2>&1; then
    echo "ERROR: docker daemon isn't running"
    exit 1
fi

TESTNET_CONFIG=$DIR/testnet/docker-compose.yml
COMPILE_PROTOCOL=${COMPILE_PROTOCOL:-false}

TESTNET_CONFIG=$TESTNET_CONFIG \
COMPILE_PROTOCOL=$COMPILE_PROTOCOL \
TAIKO_MONO_DIR=$TAIKO_MONO_DIR \
    $DIR/testnet/init.sh

DEPLOYMENT_JSON=$(cat $TAIKO_MONO_DIR/packages/protocol/deployments/l1_test_L1.json)
L2_GENESIS_ALLOC=$(cat $DIR/testnet/deployments/mainnet.json)

TAIKO_L1_CONTRACT_ADDRESS=$(echo $DEPLOYMENT_JSON | jq .contracts.TaikoL1 | sed 's/\"//g')
TAIKO_L2_CONTRACT_ADDRESS=$(echo $L2_GENESIS_ALLOC | jq 'to_entries[] | select(.value.contractName=="V1TaikoL2") | .key' | sed 's/\"//g')

trap "docker compose -f $TESTNET_CONFIG down" EXIT INT KILL ERR

RUN_TESTS=${RUN_TESTS:-false}

if [ "$RUN_TESTS" == "true" ]; then
    L1_NODE_ENDPOINT=ws://localhost:18546 \
    L2_NODE_ENDPOINT=ws://localhost:28546 \
    L2_NODE_ENGINE_ENDPOINT=http://localhost:28551 \
    TAIKO_L1_ADDRESS=$TAIKO_L1_CONTRACT_ADDRESS \
    TAIKO_L2_ADDRESS=$TAIKO_L2_CONTRACT_ADDRESS \
    L1_PROPOSER_PRIVATE_KEY=2bdd21761a483f71054e14f5b827213567971c676928d9a1808cbfa4b7501200 \
    L2_SUGGESTED_FEE_RECIPIENT=0xDf08F82De32B8d460adbE8D72043E3a7e25A3B39 \
    L1_PROVER_PRIVATE_KEY=2bdd21761a483f71054e14f5b827213567971c676928d9a1808cbfa4b7501200 \
    THROWAWAY_BLOCKS_BUILDER_PRIV_KEY=2bdd21761a483f71054e14f5b827213567971c676928d9a1808cbfa4b7501200 \
    JWT_SECRET=$DIR/testnet/jwt.hex \
        go test -v -p=1 ./... -coverprofile=coverage.out -covermode=atomic -timeout=300s
else
    echo "💻 Local dev net started"
    docker-compose -f $TESTNET_CONFIG logs -f l2_node
fi
