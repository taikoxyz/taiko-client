#!/bin/bash

DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

source "$DIR"/common.sh

# start docker compose service list
echo "start docker compose service: ${DOCKER_SERVICE_LIST[*]}"

# Init docker
compose_up "${DOCKER_INIT_LIST[@]}"

# Start docker containers.
compose_up "${DOCKER_SERVICE_LIST[@]}"

# show all the running containers
echo
docker ps --format "table {{.ID}}\t{{.Names}}\t{{.Ports}}\t{{.Status}}"
