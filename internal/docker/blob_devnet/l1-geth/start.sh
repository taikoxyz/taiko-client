#!/bin/bash

GENESIS_TIME=$(echo "$(date +%s) / 100 * 100" | bc)
echo "CURRENT_TIME=$(date +%s)"
echo "GENESIS_TIME=$GENESIS_TIME"

prysmctl \
  testnet \
  generate-genesis \
  --fork=deneb \
  --num-validators=64 \
  --genesis-time="$GENESIS_TIME" \
  --genesis-time-delay=100 \
  --output-ssz=genesis.ssz \
  --chain-config-file=config.yml \
  --geth-genesis-json-in=genesis.json \
  --geth-genesis-json-out=genesis.json

cat genesis.json

# Init geth.
/usr/local/bin/geth init --datadir=data genesis.json

# Move keystore file into data/keystore.
mv keyfile.json data/keystore

# Run geth service.
/usr/local/bin/geth \
  --http --http.api=eth,net,web3 --http.addr=0.0.0.0  --http.corsdomain=* \
  --ws --ws.api=eth,net,web3 --ws.addr=0.0.0.0 --ws.origins=* \
  --authrpc.vhosts=* --authrpc.addr=0.0.0.0 --authrpc.jwtsecret=jwtsecret \
  --datadir=data \
  --allow-insecure-unlock \
  --unlock=0x123463a4b065722e99115d6c222f267d9cabb524 \
  --password=geth_password.txt \
  --nodiscover \
  --gcmode=archive \
  --syncmode=full