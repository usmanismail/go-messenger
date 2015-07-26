#!/bin/bash
set -e
godep restore
go build
echo "Build Successful"