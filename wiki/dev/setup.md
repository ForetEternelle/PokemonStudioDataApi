# Development Setup

## Prerequisites

- **Go** 1.24 or higher
- **Java** 17 or higher (for OpenAPI Generator)
- **OpenAPI Generator CLI** 2.13.2 or higher

## Using Mise (Recommended)

This project uses [Mise](https://mise.jdx.dev/) for consistent development environment management.

### 1. Install Mise

Follow the instructions on the [official Mise website](https://mise.jdx.dev/).

### 2. Activate Environment

Simply `cd` into the project directory. Mise will automatically install the correct tool versions defined in `.mise.toml`.

### Available Commands

| Command | Description |
|---------|-------------|
| `mise run install` | Install Go dependencies |
| `mise run dev` | Start development server with live reload (Air) |
| `mise run start` | Run the application from source |
| `mise run build` | Build the production binary |
| `mise run generate` | Regenerate API client from OpenAPI specs |
| `mise run test` | Run the test suite |
| `mise run clean` | Remove build artifacts |

## Manual Setup

If you prefer not to use Mise, run the scripts directly:

```bash
# Install dependencies
./scripts/install.sh

# Build the binary
./scripts/build.sh

# Run tests
./scripts/test.sh

# Generate API code
./scripts/generate.sh
```

## Development Workflow

1. **Update OpenAPI Spec**: Edit definitions in `docs/api/`
2. **Generate Code**: Run `mise run generate`
3. **Implement Logic**: Add business logic to handlers
4. **Test**: Run `mise run test`

## Running the Server

```bash
# Development with live reload
mise run dev

# Or from source
mise run start

# Or run the built binary
./build/PokemonStudioDataApi -port=8080 -log-level=DEBUG
```
