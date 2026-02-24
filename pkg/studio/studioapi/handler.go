package studioapi

import (
	"errors"

	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/studio"
	"github.com/go-chi/chi/v5"
)

type RouterBuilderOption func(*RouterBuilderConfig)
var WithStore = func(store *studio.Store) RouterBuilderOption {
	return func(config *RouterBuilderConfig) {
		config.store = store
	}
}

type RouterBuilderConfig struct {
	store *studio.Store

}

func GetRouter(opts ...RouterBuilderOption) (chi.Router, error) {
	config := &RouterBuilderConfig{}

	for _, opt := range opts {
		opt(config)
	}

	if config.store == nil {
		return nil, errors.New("store is required")
	}

	abilityMapper := NewAbilityMapper()
	typeMapper := NewTypeMapper()
	moveMapper := NewMoveMapper(typeMapper)
	pokemonMapper := NewPokemonMapper(typeMapper, abilityMapper, config.store)

	abilityService := NewAbilityService(config.store, abilityMapper)
	typeService := NewTypeService(config.store, typeMapper)
	moveService := NewMoveService(config.store, moveMapper)
	pokemonService := NewPokemonService(config.store, pokemonMapper)

	abilityController := NewAbilitiesAPIController(abilityService)
	typeController := NewTypesAPIController(typeService)
	moveController := NewMovesAPIController(moveService)
	pokemonController := NewPokemonAPIController(pokemonService)
	return NewRouter(pokemonController, typeController, abilityController, moveController), nil

}
