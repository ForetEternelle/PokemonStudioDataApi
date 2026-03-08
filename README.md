# Pokémon Studio Data API

A Go-based REST API for accessing Pokémon Studio project data.

## Features

- **Pokemon**: List and get details
- **Types**: Query type relationships
- **Abilities**: Access ability information
- **Moves**: Coming soon

## Quick Start

```bash
# Run with Docker
docker run -p 8000:8000 -v /path/to/data:/data foreternelle/pokemon-studio-data-api
```

## Documentation

All documentation is available in the [Wiki](wiki):

| Page | Description |
|------|-------------|
| [Installation](wiki/installation) | Docker and source installation |
| [Development Setup](wiki/dev/setup) | Local development environment |
| [Project Structure](wiki/dev/structure) | Code organization |

## API Reference

The API is documented with OpenAPI 3.0. See [docs/api/openapi.yml](docs/api/openapi.yml) for the full specification.

## Configuration

| Flag | Default | Description |
|------|---------|-------------|
| `-port` | `8000` | Server port |
| `-cors` | `*` | CORS headers |
| `-data` | `data` | Data directory |
| `-log-level` | `INFO` | Logging level |

## Development

See [Development Setup](wiki/dev/setup) for:
- Mise-based development environment
- Available commands
- Code generation workflow
