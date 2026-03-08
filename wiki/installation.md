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
docker run -p 8000:8000 -v /path/to/data:/data foreternelle/pokemon-studio-data-api
```

### Configuration

| Environment Variable | Default | Description |
|---------------------|---------|-------------|
| `PORT` | `8000` | Server port |
| `CORS` | `*` | CORS headers |
| `DATA_FOLDER` | `/data` | Data directory path |
| `LOG_LEVEL` | `INFO` | Logging level (DEBUG, INFO, WARN, ERROR) |

### Example with Custom Configuration

```bash
docker run -p 8080:8080 \
  -e PORT=8080 \
  -e LOG_LEVEL=DEBUG \
  -e CORS="https://example.com" \
  -v /my/pokemon/data:/data \
  foreternelle/pokemon-studio-data-api
```

### Docker Compose

```yaml
version: '3.8'
services:
  api:
    image: foreternelle/pokemon-studio-data-api
    ports:
      - "8000:8000"
    volumes:
      - ./data:/data
    environment:
      - PORT=8000
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

Retrieve data files:

1. **Texts**: From [foret-eternelle-texts](https://gitlab.com/Aerun/foret-eternelle-texts)
2. **Pokemon/Types**: From [foret-eternelle](https://gitlab.com/Aerun/foret-eternelle) repository

Place the data in the container's `/data` folder or your local data directory.
