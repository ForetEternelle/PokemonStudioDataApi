# Pokémon Studio Data API

`PokemonStudioDataApi` is a Go-based API designed for accessing data related to a Pokémon studio project. This tool provides a standalone API server, allowing users to efficiently query and interact with various datasets. With `PokemonStudioDataApi`, you can seamlessly integrate data access into your applications or scripts.

- **Fast and lightweight:** Efficient data querying with minimal setup.
- **Flexible:** Customizable CORS settings, data directories, and log levels.
- **Containerized:** Run with Docker for easy deployment.
- **OpenAPI Documentation:** Comprehensive documentation for the API and its endpoints.

## Features
- List pokemon
- Get pokemon details

- List types

- List abilities
- Get ability details

### Incoming
- List moves
- Get move details

- List trainer
- Get trainer details

- Textual search (pokemon name, move, ...)
- Parameterized search (types, ...)

---
## Installation

### From source
  ```bash
  go install github.com/ForetEternelle/PokemonStudioDataApi
  ```

### Docker
*Coming in the next release*

## Usage
  ```bash
  PokemonStudioDataApi [flags]
  ```

| Flag | Description | Default |
|---|---|---|
| `-cors=<string>` | Set the CORS headers for the API | `*` |
| `-data=<string>` | Specify the directory containing the data files | `data` |
| `-log-level=<string>` | Set the logging level (e.g., DEBUG, INFO, WARN) | `INFO` |
| `-port=<int>` | Specify the port to run the API on | `8000` |

### Example

Start the API server on port 8080 with DEBUG log level:

```bash
./build/PokemonStudioDataApi -port=8080 -log-level=DEBUG
```

## Development
This project uses [Mise](https://mise.jdx.dev/) to ensure a consistent development environment by managing tool versions automatically. We strongly recommend using it for local development.

### Recommended Setup (with mise)

1.  **Install `mise`**: Follow the instructions on the [official `mise` website](https://mise.jdx.dev/).

2.  **Activate the environment**: Simply `cd` into the project directory. `mise` will automatically install the correct versions of Go, Java, and other tools defined in the `.mise.toml` file the first time you run a task.

Once `mise` is set up, you can use the following commands:

| Command | Description |
|---|---|
| `mise run install` | Install Go dependencies. |
| `mise run dev` | Start the development server with live reload. |
| `mise run start` | Run the application from source. |
| `mise run build` | Build the production binary. |
| `mise run generate`| (Re)generate the API client from OpenAPI specs. |
| `mise run test` | Run the test suite. |
| `mise run clean` | Remove build artifacts. |

### Manual Setup (without mise)

If you prefer not to use `mise`, you can set up the project manually.

**1. Prerequisites**

You must have the following tools installed on your system:

- **Go** (version 1.24 or higher)
- **OpenAPI Generator CLI** (version 2.13.2 or higher)
- **Java** (version 17 or higher, required by OpenAPI Generator)

**2. Using the Scripts**

The project includes scripts for common tasks. All `.sh` scripts have a `.bat` equivalent for Windows.

| Script | Description |
|---|---|
| `./scripts/install.sh` | Install Go dependencies. |
| `./scripts/build.sh` | Build the production binary. |
| `./scripts/generate.sh`| (Re)generate the API client from OpenAPI specs. |
| `./scripts/test.sh` | Run the test suite. |
| `./scripts/clean.sh` | Remove build artifacts. |

To run the server, build the binary and execute it:

```bash
./scripts/build.sh

./build/PokemonStudioDataApi [flags]
```

### Workflow

The general development workflow is as follows:

1.  **Write/Update OpenAPI Specification:** Define or modify the API endpoints, models, and parameters in the `docs/api/` directory using OpenAPI (YAML) format.
2.  **Generate Go Files:** After updating the OpenAPI spec, run the generation script to update the Go server stubs and models. This ensures your Go code reflects the latest API definition.
    ```bash
    mise run generate
    # or manually
    ./scripts/generate.sh
    ```


--------

## Installation

Retrieve the Texts via the [foret-eternelle-texts](https://gitlab.com/Aerun/foret-eternelle-texts) repository, and the pokemon (Data/Studio/pokemon) and the types (Data/Studio/types) folders via the [foret-eternelle](https://gitlab.com/Aerun/foret-eternelle) repository.
Place them in the "data" folder located at the root of the project.

## Running

Launch the main.go file located at the root of the project.

### Access a pokemon

You can access a pokemon, once the server is running, via the following URL: `http://localhost:8000/api/pokemon/{symbol}` where `{symbol}` is the pokemon symbol. For instance: `http://localhost:8000/api/pokemon/pikachu`

The url `/api/pokemon/{symbol}/{formId}` will display a specific form of a pokemon. For instance: `http://localhost:8000/api/pokemon/charizard/2`

### Language

Default language will be English (en), you can specify any supported language (depending of the project's data) by adding a query parameter `lang` : `http://localhost:8000/api/pokemon?lang=fr`
