#!/bin/bash

/validator \
  --beacon-rpc-provider=beacon-chain:4000 \
  --datadir=data \
  --accept-terms-of-use \
  --interop-num-validators=64 \
  --interop-start-index=0 \
  --chain-config-file=/config.yml \
  --force-clear-db
