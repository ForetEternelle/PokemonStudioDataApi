#!/bin/sh
set -e
openapi-generator-cli generate -c ./openapitools-go-server.json
openapi-generator-cli generate -c ./openapitools-yaml.json
