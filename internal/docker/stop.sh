#!/bin/bash

source scripts/common.sh

DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

docker compose -f "$DIR"/nodes/docker-compose.yml down

docker compose -f "$DIR"/blob_devnet/docker-compose.yml down
