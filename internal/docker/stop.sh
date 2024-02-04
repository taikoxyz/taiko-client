#!/bin/bash

DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "$DIR"/common.sh

echo "stop docker compose service: ${DOCKER_INIT_LIST[*]}"
compose_down "${DOCKER_INIT_LIST[@]}"

echo "stop docker compose service: ${DOCKER_SERVICE_LIST[*]}"
compose_down "${DOCKER_SERVICE_LIST[@]}"

# Delete exited containers.
docker rm $(docker ps -aqf "status=exited") 2>/dev/null

rm -rf "$DIR"/consensus/beacondata "$DIR"/consensus/validatordata "$DIR"/consensus/genesis.ssz
rm -rf "$DIR"/execution/geth
rm -rf "$DIR"/taikogeth/taiko-geth
