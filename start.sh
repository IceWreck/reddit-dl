#!/usr/bin/env bash

set -e
go test
go build
./reddit-img-dl \
