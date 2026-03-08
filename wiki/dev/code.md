# Code Documentation

This page explains how the codebase is organized and how to add new features.

## Architecture Overview

The API follows a standard Go web application pattern:

```
Request → Router → Handler → Service → Data
         ↓
      Middleware
```

## Key Components

### Handlers

Handlers are generated from the OpenAPI specification in `docs/api/`. After modifying the spec, run code generation:

```bash
mise run generate
```

Generated handlers are located in the `internal/` or `pkg/` directory (check `.mise.toml` for output location).

### Services

Business logic lives in service packages. They handle:
- Data loading and caching
- Data transformation
- Business rules

### Middleware

Located in `pkg/middleware/`, provides cross-cutting functionality:
- CORS handling
- Response caching
- Logging

## Adding a New Endpoint

### 1. Define the OpenAPI Spec

Add your endpoint in `docs/api/paths/`:

```yaml
# docs/api/paths/myresource.yml
get:
  summary: Get my resource
  parameters:
    - $ref: '../parameters/lang.yml'
  responses:
    '200':
      description: Successful response
      content:
        application/json:
          schema:
            $ref: '../schemas/MyResource.yml'
```

### 2. Create the Schema

Define your data model in `docs/api/schemas/`:

```yaml
# docs/api/schemas/MyResource.yml
type: object
properties:
  id:
    type: integer
  name:
    type: string
```

### 3. Generate Code

```bash
mise run generate
```

### 4. Implement the Handler

Edit the generated handler file to add business logic.

## Code Patterns

### Builder Pattern

Used for constructing complex domain entities. Each entity has a corresponding builder in `pkg/studio/`:

```go
func NewPokemonBuilder() *PokemonBuilder

// Fluent API for building entities
pokemon := studio.NewPokemonBuilder().
    ID(25).
    DbSymbol("pikachu").
    Name(studio.Translation{"en": "Pikachu", "fr": "Pikachu"}).
    Build()
```

### Service Layer

Services contain business logic and coordinate between handlers and the data store:

```go
type PokemonService struct {
    store               *studio.Store
    pokemonMapper       *PokemonMapper
    accessPolicyFactory func(context.Context) *AccessPolicy
}
```

Services are injected with dependencies (store, mappers, factories) for testability.

### Mapper Pattern

Mappers transform domain entities to API response models. Located in `pkg/studio/studioapi/`:

```go
func (m PokemonMapper) PokemonToThumbnail(p studio.Pokemon, lang string, policy *AccessPolicy) *PokemonThumbnail
```

### Store Pattern

The `Store` is the central data repository that loads and indexes all data:

```go
type Store struct {
    pokemonList     []Pokemon
    types           []PokemonType
    abilities       []Ability
    moves           []Move

    pokemonBySymbol      map[string]*Pokemon
    pokemonTypesBySymbol map[string]*PokemonType
    abilitiesBySymbol    map[string]*Ability
    movesBySymbol        map[string]*Move
}
```

Data is loaded via `Load(folder)` which imports from JSON files and creates in-memory indexes for fast lookups.

### Iterator Pattern

Go 1.23+ iterators are used throughout for lazy evaluation of collections:

```go
// In studio/store.go
func (s *Store) FindAllPokemon(filters ...PokemonFilter) iter.Seq[Pokemon]

// Usage with pagination
pkmnIter := s.store.FindAllPokemon(policy.PokemonFilters...)
pkmnPage := pagination.Collect(pkmnIter, pr)
```

### Access Policy Pattern

Access policies determine what data a request can access. Factories create policies from request context:

```go
accessPolicyFactory func(context.Context) *AccessPolicy

// In handler
policy := s.accessPolicyFactory(requestCtx)
```

### Middleware Pattern

Middleware provides cross-cutting concerns. See `pkg/middleware/` for examples like caching:

```go
// Cache middleware example
func Cache(next http.Handler) http.Handler
```

## Code Style

- Follow standard Go conventions
- Use meaningful variable names
- Add comments for exported functions
- Write tests for new functionality
