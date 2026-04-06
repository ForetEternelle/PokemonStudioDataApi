#!/bin/bash
set -e
echo "Building Go binary..."
go build -o build/bin main.go
echo "Done."
