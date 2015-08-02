#!/bin/bash
set -e
# Set directory to where we expect code to be
cd /go/src/${SOURCE_PATH}
echo "Downloading dependencies"
godep restore
echo "Fix formatting"
go fmt ./...
echo "Running Tests"
go test ./... 
echo "Building source"
go build
echo "Build Successful"
