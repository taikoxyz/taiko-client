# taiko-client

[![CI](https://github.com/taikochain/taiko-client/actions/workflows/test.yml/badge.svg)](https://github.com/taikochain/taiko-client/actions/workflows/test.yml)

Taiko protocol's client software implementation in Go.

## Project structure

| Path                | Description                                                                                                                              |
| ------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| `bindings/`         | [Go contract bindings](https://geth.ethereum.org/docs/dapp/native-bindings) for Taiko smart contracts, and few related utility functions |
| `cmd/`              | Main executable for this project                                                                                                         |
| `docs/`             | Documentation                                                                                                                            |
| `driver/`           | Driver sub-command                                                                                                                       |
| `integration_test/` | Scripts to do the integration testing of all client softwares                                                                            |
| `metrics/`          | Metrics related                                                                                                                          |
| `pkg/`              | Library code which used by all sub-commands                                                                                              |
| `proposer/`         | Proposer sub-command                                                                                                                     |
| `prover/`           | Prover sub-command                                                                                                                       |
| `scripts/`          | Helpful scripts                                                                                                                          |
| `version/`          | Version information                                                                                                                      |

## Build the source

Building the `taiko-client` binary requires a Go compiler. Once installed, run:

```sh
make build
```

## Usage

Review all available sub-commands:

```sh
bin/taiko-client --help
```

Review each sub-command's command line flags:

```sh
bin/taiko-client <sub-command> --help
```

## Testing

> NOTE: the `taiko-mono` repository has not been open-sourced yet.

Ensure you have Docker running, and Yarn installed.

Then, run the integration tests:

1. Start Docker locally
2. Perform a `yarn install` in `taiko-mono/packages/protocol`
3. Replace `<PATH_TO_TAIKO_MONO_REPO>` and execute:

   ```bash
   TAIKO_MONO_DIR=<PATH_TO_TAIKO_MONO_REPO> \
   COMPILE_PROTOCOL=true \
     make test
   ```
