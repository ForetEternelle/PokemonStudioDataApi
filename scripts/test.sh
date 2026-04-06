#!/bin/bash
set -e
echo "Running tests..."
./scripts/build.sh
go test ./...
echo "Tests finished."
