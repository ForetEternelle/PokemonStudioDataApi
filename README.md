# ForetEternelleDataApi - Pokémon Studio Data API

`ForetEternelleDataApi` is a Go-based API designed for accessing data related to a Pokémon studio project. This tool provides a standalone API server, allowing users to efficiently query and interact with various datasets. With `ForetEternelleDataApi`, you can seamlessly integrate data access into your applications or scripts.

## Features
- **Fast and lightweight:** Efficient data querying with minimal setup.
- **Flexible:** Customizable CORS settings, data directories, and log levels.
- **Containerized:** Run with Docker for easy deployment.
- **OpenAPI Documentation:** Comprehensive documentation for the API and its endpoints.

---
## Installation

### From source
  ```bash
  go install github.com/ForetEternelle/ForetEternelleDataApi
  ```

### Docker
*Coming in the next release*

## Usage
  ```bash
  ForetEternelleDataApi [flags]
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
./build/ForetEternelleDataApi -port=8080 -log-level=DEBUG
```

## Development
This project uses [Mise](https://mise.run/) to ensure a consistent development environment by managing tool versions automatically. We strongly recommend using it for local development.

### Recommended Setup (with mise)

1.  **Install `mise`**: Follow the instructions on the [official `mise` website](https://mise.run/).

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

./build/ForetEternelleDataApi [flags]
```


