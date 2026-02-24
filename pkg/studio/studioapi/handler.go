package studioapi

import (
	"context"
	"errors"

	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/iter2"
	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/studio"
	"github.com/go-chi/chi/v5"
)

type RouterBuilderOption func(*RouterBuilderConfig)
type ContextFilter[T any] func(context.Context) iter2.FilterFunc[T]

var WithStore = func(store *studio.Store) RouterBuilderOption {
	return func(config *RouterBuilderConfig) {
		config.store = store
	}
}

var WithPokemonContextFilter = func(contextFilter ContextFilter[studio.Pokemon]) RouterBuilderOption {
	return func(config *RouterBuilderConfig) {
		config.pokemonContextFilter = contextFilter
	}
}

var WithFormContextFilter = func(contextFilter ContextFilter[studio.PokemonForm]) RouterBuilderOption {
	return func(config *RouterBuilderConfig) {
		config.formContextFilter = contextFilter
	}
}

var DefaultPokemonContextFilter = func(context.Context) iter2.FilterFunc[studio.Pokemon] {
	return func(studio.Pokemon) bool {
		return true
	}
}

var DefaultFormContextFilter = func(context.Context) iter2.FilterFunc[studio.PokemonForm] {
	return func(studio.PokemonForm) bool {
		return true
	}
}

type RouterBuilderConfig struct {
	store                 *studio.Store
	pokemonContextFilter ContextFilter[studio.Pokemon]
	formContextFilter    ContextFilter[studio.PokemonForm]
}

func GetRouter(opts ...RouterBuilderOption) (chi.Router, error) {
	config := &RouterBuilderConfig{}

	for _, opt := range opts {
		opt(config)
	}

	if config.store == nil {
		return nil, errors.New("store is required")
	}

	if config.pokemonContextFilter == nil {
		config.pokemonContextFilter = DefaultPokemonContextFilter
	}

	if config.formContextFilter == nil {
		config.formContextFilter = DefaultFormContextFilter
	}

	abilityMapper := NewAbilityMapper()
	typeMapper := NewTypeMapper()
	moveMapper := NewMoveMapper(typeMapper)
	pokemonMapper := NewPokemonMapper(typeMapper, abilityMapper, config.store)

	abilityService := NewAbilityService(config.store, abilityMapper)
	typeService := NewTypeService(config.store, typeMapper)
	moveService := NewMoveService(config.store, moveMapper)
	pokemonService := NewPokemonService(config.store, pokemonMapper, config.pokemonContextFilter, config.formContextFilter)

	abilityController := NewAbilitiesAPIController(abilityService)
	typeController := NewTypesAPIController(typeService)
	moveController := NewMovesAPIController(moveService)
	pokemonController := NewPokemonAPIController(pokemonService)
	return NewRouter(pokemonController, typeController, abilityController, moveController), nil
}
