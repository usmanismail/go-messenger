#!/bin/bash
set -e
# Set directory to where we expect code to be
cd /go/src/${SOURCE_PATH}
echo "Downloading dependencies"
godep restore
echo "Building source"
go test ./... 
go build
echo "Build Successful"
