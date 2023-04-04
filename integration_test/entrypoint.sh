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

TESTNET_CONFIG=$DIR/nodes/docker-compose.yml
COMPILE_PROTOCOL=${COMPILE_PROTOCOL:-false}

TESTNET_CONFIG=$TESTNET_CONFIG \
COMPILE_PROTOCOL=$COMPILE_PROTOCOL \
TAIKO_MONO_DIR=$TAIKO_MONO_DIR \
    $DIR/nodes/init.sh

DEPLOYMENT_JSON=$(cat $TAIKO_MONO_DIR/packages/protocol/broadcast/DeployOnL1.s.sol/31337/run-latest.json)
# TAIKO_L1_CONTRACT_ADDRESS=$(echo $DEPLOYMENT_JSON | jq .transactions | jq 'to_entries[] | select(.value.contractName=="TaikoL1") | .contractAddress' | sed 's/\"//g')
TAIKO_L1_CONTRACT_ADDRESS="0xA51c1fc2f0D1a1b8494Ed1FE312d7C3a78Ed91C0" # TODO: use jq to get TaikoL1 address
L1_SIGNAL_SERVICE_CONTRACT_ADDRESS="0x4ed7c70F96B99c776995fB64377f0d4aB3B0e1C1" # TODO: use jq to get L1 signal service address
TAIKO_L2_CONTRACT_ADDRESS=0x0000777700000000000000000000000000000001

trap "docker compose -f $TESTNET_CONFIG down -v" EXIT INT KILL ERR

RUN_TESTS=${RUN_TESTS:-false}

if [ "$RUN_TESTS" == "true" ]; then
    L1_NODE_HTTP_ENDPOINT=http://localhost:18545 \
    L1_NODE_WS_ENDPOINT=ws://localhost:18546 \
    L2_EXECUTION_ENGINE_HTTP_ENDPOINT=http://localhost:28545 \
    L2_EXECUTION_ENGINE_WS_ENDPOINT=ws://localhost:28546 \
    L2_EXECUTION_ENGINE_AUTH_ENDPOINT=http://localhost:28551 \
    TAIKO_L1_ADDRESS=$TAIKO_L1_CONTRACT_ADDRESS \
    TAIKO_L2_ADDRESS=$TAIKO_L2_CONTRACT_ADDRESS \
    L1_SIGNAL_SERVICE_CONTRACT_ADDRESS=$L1_SIGNAL_SERVICE_CONTRACT_ADDRESS \
    L1_CONTRACT_OWNER_PRIVATE_KEY=ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
    L1_PROPOSER_PRIVATE_KEY=ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
    L2_SUGGESTED_FEE_RECIPIENT=0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
    L1_PROVER_PRIVATE_KEY=59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d \
    JWT_SECRET=$DIR/nodes/jwt.hex \
        go test -v -p=1 ./... -coverprofile=coverage.out -covermode=atomic -timeout=300s
else
    echo "💻 Local dev net started"
    docker compose -f $TESTNET_CONFIG logs -f l2_execution_engine
fi
