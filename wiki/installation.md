# Installation

## Docker (Recommended)

### Prerequisites

- [Docker](https://www.docker.com/get-started) installed on your system

### Pull the Image

```bash
docker pull foreternelle/pokemon-studio-data-api
```

### Run the Container

```bash
docker run -p 8000:8000 -v /path/to/data:/app/data/ foreternelle/pokemon-studio-data-api
```

### Configuration

| Environment Variable | Default | Description |
|---------------------|---------|-------------|
| `DATA` | `/app/data/` | Data directory path |
| `CORS` | `*` | CORS headers |
| `LOG_LEVEL` | `INFO` | Logging level (DEBUG, INFO, WARN, ERROR) |

The server port is fixed to `8000` inside the container, so map it to a host port of your choice with `-p`.

### Example with Custom Configuration

```bash
docker run -p 8080:8000 \
  -e LOG_LEVEL=DEBUG \
  -e CORS="https://example.com" \
  -v /my/pokemon/data:/app/data/ \
  foreternelle/pokemon-studio-data-api
```

### Docker Compose

```yaml
services:
  api:
    image: foreternelle/pokemon-studio-data-api
    ports:
      - "8000:8000"
    volumes:
      - ./data:/app/data/
    environment:
      - LOG_LEVEL=DEBUG
      - CORS=*
```

Run with:
```bash
docker-compose up
```

## From Source

See [Development Setup](dev/setup) for building from source.

## Data Setup

The API requires a data folder containing the Pokémon Studio project files.

Place the data in the container's `/app/data/` folder or your local data directory.
