#!/bin/bash

set -e

go install github.com/incu6us/goimports-reviser/v3@latest

PROJECT_NAME=github.com/taikoxyz/taiko-client

find . -name '*.go' -print0 | while IFS= read -r -d '' file; do
  goimports-reviser -project-name "$PROJECT_NAME" "$file"
done
