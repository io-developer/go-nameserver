#!/bin/bash

export GOBIN="$(pwd)/bin"
export CGO_ENABLED=0

go build -o $GOBIN -tags netgo -a
#go install
