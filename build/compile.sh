#!/bin/bash
set -e
# Set directory to where we expect code to be
cd /go/src/${SOURCE_PATH}
godep restore
go build
echo "Build Successful"