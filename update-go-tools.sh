#!/usr/bin/env bash
#
# Updates the go.tool.mod
#

set -e

# prepare dir
cd $(dirname "$0")

mkdir -p tmp
cd tmp
echo "module github.com/uponusolutions/go-smtp" > go.mod

# install tool
go get -tool github.com/golangci/golangci-lint/v2/cmd/golangci-lint
go get -tool mvdan.cc/gofumpt

# replace go.tool.mod and go.tool.sum
cd ..
rm go.tool.mod go.tool.sum
mv tmp/go.mod go.tool.mod
mv tmp/go.sum go.tool.sum
rm -r -f tmp
