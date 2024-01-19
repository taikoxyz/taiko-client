#!/bin/bash

# Generate go contract bindings.
# ref: https://geth.ethereum.org/docs/dapp/native-bindings

source scripts/common.sh

set -eou pipefail

DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null && pwd)"

# Check taiko-mono dir path environment.
check_env "TAIKO_MONO_DIR"
# Make sure abigen is available.
check_command "abigen"

echo ""
echo "TAIKO_MONO_DIR: ${TAIKO_MONO_DIR}"
echo ""

cd "${TAIKO_MONO_DIR}"/packages/protocol &&
  pnpm clean &&
  pnpm compile &&
  cd -

echo ""
echo "Start generating go contract bindings..."
echo ""

cat ${TAIKO_MONO_DIR}/packages/protocol/out/TaikoL1.sol/TaikoL1.json |
	jq .abi |
	abigen --abi - --type TaikoL1Client --pkg bindings --out $DIR/../bindings/gen_taiko_l1.go

cat ${TAIKO_MONO_DIR}/packages/protocol/out/TaikoL2.sol/TaikoL2.json |
	jq .abi |
	abigen --abi - --type TaikoL2Client --pkg bindings --out $DIR/../bindings/gen_taiko_l2.go

cat ${TAIKO_MONO_DIR}/packages/protocol/out/TaikoToken.sol/TaikoToken.json |
	jq .abi |
	abigen --abi - --type TaikoToken --pkg bindings --out $DIR/../bindings/gen_taiko_token.go

cat ${TAIKO_MONO_DIR}/packages/protocol/out/AddressManager.sol/AddressManager.json |
	jq .abi |
	abigen --abi - --type AddressManager --pkg bindings --out $DIR/../bindings/gen_address_manager.go

cat ${TAIKO_MONO_DIR}/packages/protocol/out/GuardianProver.sol/GuardianProver.json |
	jq .abi |
	abigen --abi - --type GuardianProver --pkg bindings --out $DIR/../bindings/gen_guardian_prover.go

cat ${TAIKO_MONO_DIR}/packages/protocol/out/AssignmentHook.sol/AssignmentHook.json |
	jq .abi |
	abigen --abi - --type AssignmentHook --pkg bindings --out $DIR/../bindings/gen_assignment_hook.go

cat ${TAIKO_MONO_DIR}/packages/protocol/out/TaikoTimelockController.sol/TaikoTimelockController.json |
	jq .abi |
	abigen --abi - --type TaikoTimelockController --pkg bindings --out $DIR/../bindings/gen_taiko_timelock_controller.go

git -C ${TAIKO_MONO_DIR} log --format="%H" -n 1 >./bindings/.githead

echo "🍻 Go contract bindings generated!"
