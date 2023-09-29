#!/bin/bash
# TAIKO_MONO_DIR=../taiko-mono COMPILE_PROTOCOL=false PACKAGE=prover make test
# TAIKO_MONO_DIR=../taiko-mono COMPILE_PROTOCOL=false PACKAGE=proposer make test
# TAIKO_MONO_DIR=../taiko-mono COMPILE_PROTOCOL=false PACKAGE=driver make test
# TAIKO_MONO_DIR=../taiko-mono COMPILE_PROTOCOL=false PACKAGE=cmd make test
TAIKO_MONO_DIR=../taiko-mono COMPILE_PROTOCOL=false make dev_net
# go test github.com/taikoxyz/taiko-client/driver

INFO [09-29 | 21:43:16.028]
DEBUG[09-29|21:46:05.625] Genesis hash