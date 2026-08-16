# Project Structure

The project follows a standard Go application layout with additional directories for API documentation and scripts.

## Core Directories

### `main.go`

Application entry point. Initializes and starts the HTTP server.

### `pkg/`

Core packages containing business logic and utilities:

- **`pkg/iter2/`** - Iterator utilities for collection processing
- **`pkg/scroll/`** - Scroll-based pagination for list endpoints
- **`pkg/file/`** - File system utilities
- **`pkg/middleware/`** - HTTP middleware (caching, logging, etc.)

### `pkg/pkmn/`

The Pokémon Studio domain model:

- **`pkg/pkmn/pokemon.go`**, **`move.go`**, **`ability.go`**, **`types.go`** - Domain entities
- **`pkg/pkmn/store.go`** - In-memory data store and indexes
- **`pkg/pkmn/pokemon_filter.go`** - Filtering logic for list queries

### `pkg/pkmn/studio/`

Data importers that load the Pokémon Studio JSON/CSV data files into the store.

### `pkg/pkmn/pkmnapi/`

API layer: services, mappers and the router wiring (`handler.go`).

### `pkg/pkmn/pkmnapispec/`

Generated OpenAPI server code (models and handlers), regenerated from `docs/api/`.

### `docs/api/`

OpenAPI 3.0 specification files:

- **`docs/api/openapi.yml`** - Main API definition
- **`docs/api/paths/`** - API endpoint definitions
- **`docs/api/schemas/`** - Data models and schemas
- **`docs/api/responses/`** - Response definitions
- **`docs/api/parameters/`** - Reusable parameters

### `scripts/`

Shell scripts for common tasks (install, build, test, generate, clean). Windows batch equivalents also available.

### `test/`

Test resources including:
- Valid test data in `test/test_resources/valid-data/`
- Invalid test data files for validation testing

## Configuration Files

- **`.mise.toml`** - Mise tool and task definitions
- **`go.mod` / `go.sum`** - Go dependencies
- **`package.json` / `package-lock.json`** - Node dependencies (VitePress wiki)
- **`.air.toml`** - Live reload configuration for development (Air)
- **`docker-compose.yml`** - Docker composition for production
- **`docker-compose-dev.yml`** - Docker composition for development
- **`Dockerfile`** - Container image definition
- **`openapitools*.json`** - OpenAPI Generator configuration
- **`.github/workflows/`** - CI pipeline (build, release) and Dependabot configuration
