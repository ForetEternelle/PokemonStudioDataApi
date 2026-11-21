@echo off
setlocal

openapi-generator-cli generate -c ./openapitools-go-server.json || exit /b %errorlevel%
rem The go-server template indents nullable model constraints with spaces before tabs; gofmt keeps generated code clean.
gofmt -w .\pkg\pkmn\pkmnapispec || exit /b %errorlevel%
openapi-generator-cli generate -c ./openapitools-yaml.json || exit /b %errorlevel%
openapi-generator-cli generate -c ./openapitools-json.json || exit /b %errorlevel%

endlocal
