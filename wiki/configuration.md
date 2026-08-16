# Configuration

The API is configured through command-line flags when running from source, and through environment variables when running the Docker image.

## Command-line Flags

When running the binary or `go run` from source:

| Flag | Default | Description |
|------|---------|-------------|
| `-port` | `8000` | Server port |
| `-cors` | `*` | CORS headers |
| `-data` | `data` | Data directory path |
| `-log-level` | `DEBUG` | Logging level (`DEBUG`, `INFO`, `WARN`, `ERROR`) |

Example:

```bash
./build/bin -port=8080 -cors="https://example.com" -data=/path/to/data -log-level=DEBUG
```

## Environment Variables (Docker)

When running the [Docker image](installation), the following environment variables are supported:

| Environment Variable | Default | Description |
|---------------------|---------|-------------|
| `DATA` | `/app/data/` | Data directory path |
| `CORS` | `*` | CORS headers |
| `LOG_LEVEL` | `INFO` | Logging level (`DEBUG`, `INFO`, `WARN`, `ERROR`) |

The container listens on port `8000`; map it to a host port with `-p`.

## Mise

The `mise run start` task wires these values through [Mise](https://mise.jdx.dev/) variables defined in `.mise.toml`:

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `8080` | Server port |
| `CORS` | `*` | CORS headers |
| `DATA_FOLDER` | `./test/test_resources/valid-data` | Data directory path |
| `LOG_LEVEL` | `DEBUG` | Logging level |
