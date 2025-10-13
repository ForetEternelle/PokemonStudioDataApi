@echo off
setlocal

openapi-generator-cli generate -c ./openapitools-go-server.json || exit /b %errorlevel%
openapi-generator-cli generate -c ./openapitools-yaml.json || exit /b %errorlevel%

endlocal
