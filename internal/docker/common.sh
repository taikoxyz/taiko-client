#!/bin/bash

DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
COMPOSE="docker compose -f $DIR/docker-compose.yml"

# docker compose service list.
DOCKER_INIT_LIST=("create-beacon-chain-genesis" "geth-remove-db" "geth-genesis")
DOCKER_SERVICE_LIST=("beacon-chain" "geth" "validator" "l2_execution_engine")

check_command() {
  if ! command -v "$1" &> /dev/null; then
    print_error "$1 could not be found"
    exit
  fi
}

compose_down() {
  local services=("$@")
  echo
  echo "stopping services..."
  $COMPOSE down "${services[@]}" #--remove-orphans
  echo "done"
}

compose_up() {
  local services=("$@")
  echo
  echo "launching services..."
  $COMPOSE up --quiet-pull "${services[@]}" -d --wait
  echo "done"
}
