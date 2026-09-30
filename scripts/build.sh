#!/bin/bash
set -e
echo "Building Go binary..."
CGO_ENABLED=0 go build -trimpath -o build/bin main.go
echo "Done."
