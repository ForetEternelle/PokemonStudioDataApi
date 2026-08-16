---
# https://vitepress.dev/reference/default-theme-home-page
layout: home

hero:
  name: "PokemonStudioDataApi"
  text: "The official documentation of the Pokemon Studio data api"
  tagline: A Go-based REST API for accessing Pokémon Studio project data
  actions:
    - theme: brand
      text: Installation
      link: /installation
    - theme: alt
      text: Development Setup
      link: /dev/setup

features:
  - title: Standalone API
    details: Run it as a self-hosted REST API with Docker or a single Go binary.
  - title: Embeddable Library
    details: Use the router, store and services directly from your own Go application.
  - title: Pokémon & Data
    details: Access Pokémon, forms, types, abilities and moves with translations.
  - title: OpenAPI Spec
    details: Fully documented with an OpenAPI 3.0 specification and generated server code.
  - title: Multi-language
    details: Localized responses from the project's translation files.
---
