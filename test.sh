#!/bin/bash
# TAIKO_MONO_DIR=../taiko-mono COMPILE_PROTOCOL=false PACKAGE=prover make test
# TAIKO_MONO_DIR=../taiko-mono COMPILE_PROTOCOL=false PACKAGE=proposer make test
# TAIKO_MONO_DIR=../taiko-mono COMPILE_PROTOCOL=false PACKAGE=driver make test
# TAIKO_MONO_DIR=../taiko-mono COMPILE_PROTOCOL=false PACKAGE=cmd make test
TAIKO_MONO_DIR=../taiko-mono COMPILE_PROTOCOL=false make test
