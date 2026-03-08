# Project Structure

The project follows a standard Go application layout with additional directories for API documentation and scripts.

## Core Directories

### `main.go`

Application entry point. Initializes and starts the HTTP server.

### `pkg/`

Core packages containing business logic and utilities:

- **`pkg/iter2/`** - Iterator utilities for collection processing
- **`pkg/pagination/`** - Pagination logic for list endpoints
- **`pkg/file/`** - File system utilities
- **`pkg/middleware/`** - HTTP middleware (caching, logging, etc.)

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
- **`docker-compose.yml`** - Docker composition for production
- **`docker-compose-dev.yml`** - Docker composition for development
- **`Dockerfile`** - Container image definition
- **`openapitools*.json`** - OpenAPI Generator configuration
