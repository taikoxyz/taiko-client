#!/bin/bash

DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

export L1_NODE_HTTP_ENDPOINT=http://localhost:$(docker port l1_node | awk -F ':' '{print $2}')
export L1_NODE_WS_ENDPOINT=ws://localhost:$(docker port l1_node | awk -F ':' '{print $2}')

export L2_EXECUTION_ENGINE_HTTP_ENDPOINT=http://localhost:$(docker port l2_node | awk -F ':' 'NR==1 {print $2}')
export L2_EXECUTION_ENGINE_WS_ENDPOINT=ws://localhost:$(docker port l2_node | awk -F ':' 'NR==2 {print $2}')
export L2_EXECUTION_ENGINE_AUTH_ENDPOINT=http://localhost:$(docker port l2_node | awk -F ':' 'NR==3 {print $2}')
export JWT_SECRET=$DIR/nodes/jwt.hex

export DOCKER_SERVICE_LIST=("l1_node" "l2_execution_engine")

echo "L1_NODE_HTTP_ENDPOINT: $L1_NODE_HTTP_ENDPOINT"
echo "L1_NODE_WS_ENDPOINT: $L1_NODE_WS_ENDPOINT"
echo "L2_EXECUTION_ENGINE_HTTP_ENDPOINT: $L2_EXECUTION_ENGINE_HTTP_ENDPOINT"
echo "L2_EXECUTION_ENGINE_WS_ENDPOINT: $L2_EXECUTION_ENGINE_WS_ENDPOINT"
echo "L2_EXECUTION_ENGINE_AUTH_ENDPOINT: $L2_EXECUTION_ENGINE_AUTH_ENDPOINT"