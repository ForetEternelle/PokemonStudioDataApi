# Pokémon Studio Data API

A Go-based REST API for accessing Pokémon Studio project data. Provides efficient querying of Pokemon, types, abilities, and moves datasets.

## Key Features

- Fast and lightweight data querying
- Flexible configuration (CORS, data directories, logging)
- Containerized with Docker
- OpenAPI 3.0 documentation
- Multi-language support

## Quick Start

```bash
# Run with Docker
docker run -p 8000:8000 -v /path/to/data:/data foreternelle/pokemon-studio-data-api
```

## Documentation

For detailed documentation, see the [Wiki](wiki):

- [Installation](wiki/installation) - Docker and source installation
- [Development Setup](wiki/dev/setup) - Local development environment
- [Project Structure](wiki/dev/structure) - Code organization
- [API Reference](docs/api/openapi.yml) - OpenAPI specification

## Available Data

- **Pokemon**: List and detailed information
- **Types**: Type relationships and details
- **Abilities**: Ability information
- **Moves**: Move data (coming soon)

## Tech Stack

- Go 1.24
- OpenAPI 3.0
- Docker
- Mise (development tools)
