#!/bin/bash


source scripts/common.sh
source docker/docker_env.sh

echo "stop docker compose service: ${DOCKER_SERVICE_LIST[*]}"

compose_down "${DOCKER_SERVICE_LIST[@]}"
