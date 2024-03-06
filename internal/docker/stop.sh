#!/bin/bash

source scripts/common.sh

DOCKER_SERVICE_LIST=("l1_node" "l2_execution_engine")
echo "stop docker compose service: ${DOCKER_SERVICE_LIST[*]}"
compose_down "${DOCKER_SERVICE_LIST[@]}"

DOCKER_BLOB_DEVNET_LIST=("create-beacon-chain-genesis" "geth-remove-db" "geth-genesis" "beacon-chain" "geth" "validator")
echo "stop docker blob devnet: ${DOCKER_BLOB_DEVNET_LIST[*]}"
compose_down "${DOCKER_BLOB_DEVNET_LIST[@]}"
