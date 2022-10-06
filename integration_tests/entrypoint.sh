#!/bin/bash

set -eou pipefail

DIR=$(
    cd $(dirname ${BASH_SOURCE[0]})
    pwd
)

TESTNET_CONFIG=$DIR/testnet/docker-compose.yml
COMPILE_PROTOCOL=${COMPILE_PROTOCOL:-false}

TESTNET_CONFIG=$TESTNET_CONFIG \
COMPILE_PROTOCOL=$COMPILE_PROTOCOL \
TAIKO_MONO_DIR=$TAIKO_MONO_DIR \
    $DIR/testnet/init.sh

DEPLOYMENT_JSON=$(cat $TAIKO_MONO_DIR/packages/protocol/deployments/l1_test_L1.json)
L2_GENESIS_ALLOC=$(cat $DIR/testnet/deployments/genesis_alloc.json)

TAIKO_L1_CONTRACT_ADDRESS=$(echo $DEPLOYMENT_JSON | jq .contracts.TaikoL1 | sed 's/\"//g')
TAIKO_L2_CONTRACT_ADDRESS=$(echo $L2_GENESIS_ALLOC | jq 'to_entries[] | select(.value.contractName=="V1TaikoL2") | .key' | sed 's/\"//g')

trap "docker compose -f $TESTNET_CONFIG down" EXIT INT KILL ERR

L1_NODE_ENDPOINT=ws://localhost:18546 \
L2_NODE_ENDPOINT=ws://localhost:28546 \
TAIKO_L1_ADDRESS=$TAIKO_L1_CONTRACT_ADDRESS \
TAIKO_L2_ADDRESS=$TAIKO_L2_CONTRACT_ADDRESS \
L1_TRANSACTOR_PRIVATE_KEY=2bdd21761a483f71054e14f5b827213567971c676928d9a1808cbfa4b7501200 \
L2_PROPOSER_PRIVATE_KEY=6bff9a8ffd7f94f43f4f5f642be8a3f32a94c1f316d90862884b2e276293b6ee \
    go test -v ./...
