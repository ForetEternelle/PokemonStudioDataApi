#!/bin/bash
echo "Running tests..."
./scripts/build.sh
go test ./...
echo "Tests finished."
