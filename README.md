# taiko-client

[![CI](https://github.com/taikochain/taiko-client/actions/workflows/test.yml/badge.svg)](https://github.com/taikochain/taiko-client/actions/workflows/test.yml)

Taiko protocol's client softwares implementation in Golang.

## Building

Compile a binary:

```shell
make build
```

## Testing

Run the integration tests:

```bash
TAIKO_MONO_DIR=PATH_TO_TAIKO_MONO_REPO \
COMPILE_PROTOCOL=true \
  make test
```

## Running

All available sub-commands can be reviewed with:

```bash
bin/taiko-client --help
```
