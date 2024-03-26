#!/bin/bash

source scripts/common.sh

DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# start docker compose service list
docker compose -f "$DIR"/nodes/docker-compose.yml up -d

# start blob devnet service list
rm -rf /tmp/consensus /tmp/execution
cp -R "$DIR"/blob_devnet/consensus /tmp
cp -R "$DIR"/blob_devnet/execution /tmp
docker compose -f "$DIR"/blob_devnet/docker-compose.yml up -d

# show all the running containers
echo
docker ps --format "table {{.ID}}\t{{.Names}}\t{{.Ports}}\t{{.Status}}"
