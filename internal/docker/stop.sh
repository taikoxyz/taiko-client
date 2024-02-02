#!/bin/bash

source scripts/common.sh

DOCKER_INIT_LIST=("create-beacon-chain-genesis" "geth-remove-db" "geth-genesis")
DOCKER_SERVICE_LIST=("beacon-chain" "geth" "validator" "l2_execution_engine")

echo "stop docker compose service: ${DOCKER_INIT_LIST[*]}"
compose_down "${DOCKER_INIT_LIST[@]}"

echo "stop docker compose service: ${DOCKER_SERVICE_LIST[*]}"
compose_down "${DOCKER_SERVICE_LIST[@]}"

# Delete exited containers.
docker rm $(docker ps -aqf "status=exited") 2>/dev/null

rm -rf ./consensus/beacondata ./consensus/validatordata ./consensus/genesis.ssz
rm -rf ./execution/geth
rm -rf taikogeth/taiko-geth
