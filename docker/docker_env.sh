#!/bin/bash

DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

export L1_NODE_HTTP_ENDPOINT=http://localhost:18545
export L1_NODE_WS_ENDPOINT=ws://localhost:18546
export L1_EXECUTION_ENGINE_AUTH_ENDPOINT=http://localhost:18551

export L2_EXECUTION_ENGINE_HTTP_ENDPOINT=http://localhost:28545
export L2_EXECUTION_ENGINE_WS_ENDPOINT=ws://localhost:28546
export L2_EXECUTION_ENGINE_AUTH_ENDPOINT=http://localhost:28551
export JWT_SECRET=$DIR/nodes/jwt.hex

export DOCKER_SERVICE_LIST=("l1_node" "l2_execution_engine")
