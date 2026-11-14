# Code Documentation

This page explains how the codebase is organized and how to add new features.

## Contributing

### Choose and coordinate work

Find useful work in the [public GitHub Project associated with this repository](https://github.com/ForetEternelle/PokemonStudioDataApi/projects). Use the existing `area:*` labels to explore a domain, `good first issue` for a first contribution, and `help wanted` for work seeking help.

Before starting, read the issue's outcome, acceptance criteria, dependencies, assignees, recent comments and linked pull requests. An unassigned issue may already have a contributor's claim in its comments. A `needs refinement` label means essential decisions are still open; help resolve those in the issue before implementing the unsettled contract.

1. **Take the initiative.** Choose available, ready work, assign yourself where GitHub permissions allow, and comment that you are taking it on. If self-assignment is unavailable, your comment signals the claim; you do not need to wait for a maintainer to assign you.
2. **Make progress visible at meaningful changes.** Link your PR to the issue. An early draft PR can share an approach or ask a focused question. Update the Project status when you have permission; otherwise the issue and PR provide the visible coordination trail.
3. **Pause or hand over openly.** If you stop working on an issue, leave a short note and any useful branch or PR link, then unassign yourself if possible. If someone has already claimed a topic, coordinate in that issue before starting overlapping work.

Contributions are voluntary. Choose work that fits your interests and availability; there are no required estimates, individual delivery deadlines or recurring progress reports.

### Keep one source for each kind of information

| Information | Source |
|---|---|
| Available work and progress | Public GitHub Project, using issue labels and assignees |
| Scope, contract decisions and delivery dependencies | The issue and its linked PR |
| API contract | OpenAPI source in `docs/api/`, rendered in the [API reference](../api.md) |
| Contribution rules and development patterns | This guide and the [development setup](./setup.md) |

Record contract decisions in the issue and carry the agreed result into OpenAPI and documentation in the same PR. Describe unresolved questions explicitly rather than assuming a fixed official dataset or a consumer-specific behavior.

For an API change, include the relevant importer/domain/service wiring, source OpenAPI and regenerated artifacts, representative validation, and any affected documentation. The endpoint workflow below explains the implementation steps.

Use `Closes #<number>` in a PR only when it completes the issue; use `Refs #<number>` for a partial contribution. Current availability and progress belong in GitHub, so this guide does not maintain a ticket list or coverage-status table.

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

Generated models and handler interfaces are located in `pkg/pkmn/pkmnapispec/` (see `openapitools-go-server.json` for the output location). Handlers are wired together with services in `pkg/pkmn/pkmnapi/handler.go`.

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

Implement business logic in the service layer under `pkg/pkmn/pkmnapi/` and wire it in `handler.go`. Do not edit generated files in `pkg/pkmn/pkmnapispec/`; they are overwritten by `mise run generate`.

## Code Patterns

### Entity Construction

Domain entities (`Pokemon`, `PokemonForm`, `Move`, `Ability`, `PokemonType`) have exported fields and are immutable by convention only: fields must not be mutated once an entity has been registered in a store. The fields used to index/order entities in the store (`Pokemon.ID`, `Pokemon.DbSymbol`, `Move.DbSymbol`, `Ability.DbSymbol`, `PokemonType.DbSymbol`) are documented as immutable. Entities are built by the store `Add*` methods, which take a DTO and return the constructed entity:


The `Add*` methods resolve symbol references (types, abilities), default `CustomProperties`/`Tags`, sort forms, and register the entity in the store indexes.


Services are injected with dependencies (store, mappers, factories) for testability.

### Mapper Pattern

Mappers transform domain entities to API response models. Located in `pkg/pkmn/pkmnapi/`:

```go
func (m PokemonMapper) PokemonToThumbnail(p pkmn.Pokemon, formId int32, lang string, policy PokemonFilterPolicy) *PokemonThumbnail
```

### Store Pattern

The `Store` is the central data repository that loads and indexes all data:

Data is loaded via `Load(folder)` which imports from JSON files and creates in-memory indexes for fast lookups.

### Iterator Pattern

Go 1.23+ iterators are used throughout for lazy evaluation of collections:

```go
// In pkg/pkmn/store.go
func (s *Store) FindAllPokemon(filters ...iter2.FilterFunc[*Pokemon]) iter.Seq[*Pokemon]

// Usage with pagination
pkmnIter := s.store.FindAllPokemon(filters...)
pkmnPage := scroll.Of(pkmnIter, size)
```

### Access Policy Pattern

Access policies determine what data a request can access. Factories create policies from request context:

```go
accessPolicyFactory func(context.Context) *PokemonFilterPolicy

// In handler
policy := s.accessPolicyFactory(requestCtx)
```

## Code Style

- Follow standard Go conventions
- Use meaningful variable names
- Add comments for exported functions
- Write tests for new functionality
