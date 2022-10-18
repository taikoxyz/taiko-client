# taiko-client

[![CI](https://github.com/taikochain/taiko-client/actions/workflows/test.yml/badge.svg)](https://github.com/taikochain/taiko-client/actions/workflows/test.yml)

Taiko protocol's client softwares implementation in Golang.

## Building the source

Building the `taiko-client` binary requires a Go compiler, once installed, run:

```shell
make build
```

## Project Structure

| Path                | Description                                                                                                                              |
| ------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| `bindings/`         | [Go contract bindings](https://geth.ethereum.org/docs/dapp/native-bindings) for Taiko smart contracts, and few related utility functions |
| `cmd/`              | Main executable for this project                                                                                                         |
| `integration_test/` | Scripts to do the integration testing of all client softwares                                                                            |
| `scripts/`          | Helpful scripts                                                                                                                          |
| `pkg/`              | Library code which used by all sub-commands                                                                                              |
| `proposer/`         | Proposer sub-command                                                                                                                     |
| `driver/`           | Driver sub-command                                                                                                                       |
| `prover/`           | Prover sub-command                                                                                                                       |
| `docs/`             | Documentations                                                                                                                           |
| `version/`          | Version infomation                                                                                                                       |

## Testing

> NOTE: the `taiko-mono` repository has not been open-sourced yet.

Run the integration tests:

```bash
TAIKO_MONO_DIR=<PATH_TO_TAIKO_MONO_REPO> \
COMPILE_PROTOCOL=true \
  make test
```

## Running

All available sub-commands can be reviewed with:

```bash
bin/taiko-client --help
```

And each sub-command's command line flags can be reviewed with:

```bash
bin/taiko-client <sub-command> --help
```
