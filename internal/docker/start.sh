#!/bin/bash

source scripts/common.sh

# start docker compose service list
DOCKER_SERVICE_LIST=("l1_node" "l2_execution_engine")
echo "start docker compose service: ${DOCKER_SERVICE_LIST[*]}"
compose_up "${DOCKER_SERVICE_LIST[@]}"

# start blob devnet service list
DOCKER_BLOB_DEVNET_LIST=("create-beacon-chain-genesis" "geth-remove-db" "geth-genesis" "beacon-chain" "geth" "validator")
echo "start blob devnet service: ${DOCKER_BLOB_DEVNET_LIST[*]}"
compose_up "${DOCKER_BLOB_DEVNET_LIST[@]}"

# show all the running containers
echo
docker ps --format "table {{.ID}}\t{{.Names}}\t{{.Ports}}\t{{.Status}}"
