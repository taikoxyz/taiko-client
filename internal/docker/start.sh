#!/bin/bash

source scripts/common.sh
source internal/docker/docker_env.sh

# start docker compose service list
echo "start docker compose service: ${DOCKER_SERVICE_LIST[*]}"

compose_up "${DOCKER_SERVICE_LIST[@]}"

# check until L1 chain is ready
until cast chain-id --rpc-url "$L2_EXECUTION_ENGINE_HTTP_ENDPOINT"; do
    sleep 1
done

# show all the running containers
echo
docker ps --format "table {{.ID}}\t{{.Names}}\t{{.Ports}}\t{{.Status}}"
