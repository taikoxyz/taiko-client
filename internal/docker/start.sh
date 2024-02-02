#!/bin/bash
source scripts/common.sh

DOCKER_INIT_LIST=("create-beacon-chain-genesis" "geth-remove-db" "geth-genesis")
DOCKER_SERVICE_LIST=("beacon-chain" "geth" "validator" "l2_execution_engine")

# start docker compose service list
echo "start docker compose service: ${DOCKER_SERVICE_LIST[*]}"

# Init docker
compose_up "${DOCKER_INIT_LIST[@]}"

# Start docker containers.
compose_up "${DOCKER_SERVICE_LIST[@]}"

# show all the running containers
echo
docker ps --format "table {{.ID}}\t{{.Names}}\t{{.Ports}}\t{{.Status}}"
