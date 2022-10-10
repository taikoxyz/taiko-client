#!/bin/bash

set -eou pipefail

if [[ -z $NODE_URL ]]; then
    echo "Must pass NODE_URL"
    exit 1
fi

JSON_REQUEST_BODY='{"jsonrpc":"2.0","id":0,"method":"eth_chainId","params":[]}'

while ! curl \
    --fail \
    --silent \
    -X POST \
    -H "Content-Type: application/json" \
    -d "$JSON_REQUEST_BODY" \
    $NODE_URL > /dev/null
do
    sleep 1
done
