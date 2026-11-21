#!/bin/sh
set -e
openapi-generator-cli generate -c ./openapitools-go-server.json
# The go-server template indents nullable model constraints with spaces before tabs; gofmt keeps generated code clean.
gofmt -w ./pkg/pkmn/pkmnapispec
openapi-generator-cli generate -c ./openapitools-yaml.json
openapi-generator-cli generate -c ./openapitools-json.json
