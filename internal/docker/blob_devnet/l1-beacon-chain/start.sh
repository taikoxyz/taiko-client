#!/bin/bash

GENESIS_TIME=$(echo "$(date +%s) / 3600 * 3600" | bc)
echo "GENESIS_TIME=$GENESIS_TIME"

prysmctl \
  testnet \
  generate-genesis \
  --fork=deneb \
  --num-validators=64 \
  --genesis-time="$GENESIS_TIME" \
  --output-ssz=/genesis.ssz \
  --chain-config-file=/config.yml \
  --geth-genesis-json-in=/genesis.json \
  --geth-genesis-json-out=/genesis.json

cat /genesis.json

beacon-chain \
  --datadir=beacondata \
  --min-sync-peers=0 \
  --genesis-state=/genesis.ssz \
  --bootstrap-node= \
  --interop-eth1data-votes \
  --chain-config-file=/config.yml \
  --contract-deployment-block=0 \
  --chain-id=32382 \
  --rpc-host=0.0.0.0 \
  --grpc-gateway-host=0.0.0.0 \
  --execution-endpoint=http://geth:8551 \
  --accept-terms-of-use \
  --jwt-secret=/jwtsecret \
  --suggested-fee-recipient=0x123463a4b065722e99115d6c222f267d9cabb524 \
  --minimum-peers-per-subnet=0 \
  --enable-debug-rpc-endpoints
