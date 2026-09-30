# Pokémon Studio Data API

A Go-based REST API for accessing Pokémon Studio project data.

## Features

- **Pokemon**: List and get details, including forms
- **Types**: Query type relationships
- **Abilities**: Access ability information
- **Moves**: List and get move details

## Quick Start

```bash
# Run with Docker
docker run -p 8080:8080 -v /path/to/data:/app/data ghcr.io/foreteternelle/pokemonstudiodataapi:latest
```

## Documentation

All documentation is available at [https://foreteternelle.github.io/PokemonStudioDataApi/](https://foreteternelle.github.io/PokemonStudioDataApi/):

| Page | Description |
|------|-------------|
| [Installation](https://foreteternelle.github.io/PokemonStudioDataApi/installation.html) | Docker and source installation |
| [Configuration](https://foreteternelle.github.io/PokemonStudioDataApi/configuration.html) | Flags, environment variables |
| [Development Setup](https://foreteternelle.github.io/PokemonStudioDataApi/dev/setup.html) | Local development environment |
| [Project Structure](https://foreteternelle.github.io/PokemonStudioDataApi/dev/structure.html) | Code organization |
| [Code Documentation](https://foreteternelle.github.io/PokemonStudioDataApi/dev/code.html) | Architecture and patterns |

## API Reference

The API is documented with OpenAPI 3.0. See [docs/api/openapi.yml](docs/api/openapi.yml) for the full specification, or the [rendered Swagger UI](https://foreteternelle.github.io/PokemonStudioDataApi/api.html).

## Configuration

| Flag | Env | Default | Description |
|------|-----|---------|-------------|
| `-port` | `PORT` | `8080` | Server port |
| `-cors` | `CORS` | `*` | CORS headers |
| `-data` | `DATA` | `data` (binary) / `/app/data` (Docker) | Data directory |
| `-log-level` | `LOG_LEVEL` | `INFO` | Logging level |

Flags take precedence over environment variables.

## Development

See [Development Setup](https://foreteternelle.github.io/PokemonStudioDataApi/dev/setup.html) for:
- Mise-based development environment
- Available commands
- Code generation workflow
