#!/bin/bash

DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# check until L1 chain is ready
L1_PROBE_URL=http://localhost:$(docker port l1_node | grep '0.0.0.0' | awk -F ':' '{print $2}')
until cast chain-id --rpc-url "$L1_PROBE_URL" &>/dev/null 2>&1; do
  sleep 1
done
L1_NODE_PORT=$(docker port l1_node | grep '0.0.0.0' | awk -F ':' '{print $2}')
export L1_NODE_HTTP_ENDPOINT=http://localhost:$L1_NODE_PORT
export L1_NODE_WS_ENDPOINT=ws://localhost:$L1_NODE_PORT

# check until L2 chain is ready
L2_PROBE_URL=http://localhost:$(docker port l2_node | grep "0.0.0.0" | awk -F ':' 'NR==1 {print $2}')
until cast chain-id --rpc-url "$L2_PROBE_URL" &>/dev/null 2>&1; do
  sleep 1
done
export L2_EXECUTION_ENGINE_HTTP_ENDPOINT=http://localhost:$(docker port l2_node | grep "0.0.0.0" | awk -F ':' 'NR==1 {print $2}')
export L2_EXECUTION_ENGINE_WS_ENDPOINT=ws://localhost:$(docker port l2_node | grep "0.0.0.0" | awk -F ':' 'NR==2 {print $2}')
export L2_EXECUTION_ENGINE_AUTH_ENDPOINT=http://localhost:$(docker port l2_node | grep "0.0.0.0" | awk -F ':' 'NR==3 {print $2}')
export JWT_SECRET=$DIR/nodes/jwt.hex

# check until BLOB node is ready
BLOB_GETH_NODE_ENDPOINT=ws://localhost:$(docker port blob_node | grep '0.0.0.0' | awk -F ':' '{print $2}')
until cast block-number --rpc-url "$BLOB_GETH_NODE_ENDPOINT" &>/dev/null 2>&1; do
  sleep 1
done
export BLOB_BEACON_NODE_ENDPOINT=http://localhost:$(docker port beacon-chain | grep "0.0.0.0" | awk -F ':' 'NR==1 {print $2}')

echo -e "L1_NODE PORTS: \n$(docker port l1_node)"
echo -e "L2_NODE PORTS: \n$(docker port l2_node)"
echo -e "BLOB_NODE PORTS: \n$(docker port blob_node)"

echo "L1_NODE_HTTP_ENDPOINT: $L1_NODE_HTTP_ENDPOINT"
echo "L1_NODE_WS_ENDPOINT: $L1_NODE_WS_ENDPOINT"
echo "L2_EXECUTION_ENGINE_HTTP_ENDPOINT: $L2_EXECUTION_ENGINE_HTTP_ENDPOINT"
echo "L2_EXECUTION_ENGINE_WS_ENDPOINT: $L2_EXECUTION_ENGINE_WS_ENDPOINT"
echo "L2_EXECUTION_ENGINE_AUTH_ENDPOINT: $L2_EXECUTION_ENGINE_AUTH_ENDPOINT"
echo "BLOB_GETH_NODE_ENDPOINT: $BLOB_GETH_NODE_ENDPOINT"
echo "BLOB_BEACON_NODE_ENDPOINT: $BLOB_BEACON_NODE_ENDPOINT"
