#!/bin/bash

source scripts/common.sh

DOCKER_SERVICE_LIST=("l2_execution_engine" "validator" "beacon-chain" "geth")

echo "stop docker compose service: ${DOCKER_SERVICE_LIST[*]}"

compose_down "${DOCKER_SERVICE_LIST[@]}"
