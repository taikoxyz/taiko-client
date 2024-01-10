#!/bin/bash

RED='\033[1;31m'
NC='\033[0m' # No Color

print_error() {
  local msg="$1"
  echo -e "${RED}$msg${NC}"
}

check_env() {
  local name="$1"
  local value="${!name}"

  if [ -z "$value" ]; then
    print_error "$name not set in env"
    exit 1
  fi
}

check_command() {
  if ! command -v "$1" &> /dev/null; then
    print_error "$1 could not be found"
    exit
  fi
}

