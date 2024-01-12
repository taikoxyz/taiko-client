#!/bin/bash

source docker/docker_env.sh
source scripts/common.sh

# make sure environment variables are set.
check_env "TAIKO_MONO_DIR"

# get deployed contract address.
DEPLOYMENT_JSON=$(cat "$TAIKO_MONO_DIR"/packages/protocol/deployments/deploy_l1.json)
export TAIKO_L1_ADDRESS=$(echo "$DEPLOYMENT_JSON" | jq '.taiko' | sed 's/\"//g')
export TAIKO_L2_ADDRESS=0x1670010000000000000000000000000000010001
export TAIKO_TOKEN_ADDRESS=$(echo "$DEPLOYMENT_JSON" | jq '.taiko_token' | sed 's/\"//g')
export ASSIGNMENT_HOOK_ADDRESS=$(echo "$DEPLOYMENT_JSON" | jq '.assignment_hook' | sed 's/\"//g')
export TIMELOCK_CONTROLLER=$(echo "$DEPLOYMENT_JSON" | jq '.timelock_controller' | sed 's/\"//g')
export ROLLUP_ADDRESS_MANAGER_CONTRACT_ADDRESS=$(echo "$DEPLOYMENT_JSON" | jq '.rollup_address_manager' | sed 's/\"//g')
export GUARDIAN_PROVER_CONTRACT_ADDRESS=$(echo "$DEPLOYMENT_JSON" | jq '.guardian_prover' | sed 's/\"//g')
export L1_SIGNAL_SERVICE_CONTRACT_ADDRESS=$(echo "$DEPLOYMENT_JSON" | jq '.signal_service' | sed 's/\"//g')
export L1_CONTRACT_OWNER_PRIVATE_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
export L1_SECURITY_COUNCIL_PRIVATE_KEY=0xdbda1821b80551c9d65939329250298aa3472ba22feea921c0cf5d620ea67b97
export L1_PROPOSER_PRIVATE_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
export L1_PROVER_PRIVATE_KEY=0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d
export TREASURY=0x1670010000000000000000000000000000010001
