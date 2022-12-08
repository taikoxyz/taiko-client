#!/bin/bash

# Generate go contract bindings.
# ref: https://geth.ethereum.org/docs/dapp/native-bindings

set -eou pipefail

DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null && pwd)"

echo ""
echo "TAIKO_MONO_DIR: ${TAIKO_MONO_DIR}"
echo "TAIKO_GETH_DIR: ${TAIKO_GETH_DIR}"
echo ""

cd ${TAIKO_GETH_DIR} &&
  make all &&
  cd -

cd ${TAIKO_MONO_DIR}/packages/protocol &&
  pnpm clean &&
  pnpm compile &&
  cd -

ABIGEN_BIN=$TAIKO_GETH_DIR/build/bin/abigen

echo ""
echo "Start generating go contract bindings..."
echo ""

cat ${TAIKO_MONO_DIR}/packages/protocol/artifacts/contracts/L1/TaikoL1.sol/TaikoL1.json |
	jq .abi |
	${ABIGEN_BIN} --abi - --type TaikoL1Client --pkg bindings --out $DIR/../bindings/gen_taiko_l1.go

cat ${TAIKO_MONO_DIR}/packages/protocol/artifacts/contracts/L2/TaikoL2.sol/TaikoL2.json |
	jq .abi |
	${ABIGEN_BIN} --abi - --type TaikoL2Client --pkg bindings --out $DIR/../bindings/gen_taiko_l2.go

git -C ${TAIKO_MONO_DIR} log --format="%H" -n 1 >./bindings/.githead

echo "🍻 Go contract bindings generated!"
