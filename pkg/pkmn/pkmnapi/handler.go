package pkmnapi

import (
	"context"

	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/pkmn"
	"github.com/go-chi/chi/v5"
	. "github.com/ForetEternelle/PokemonStudioDataApi/pkg/pkmn/pkmnapispec"
)

type GetRouterOption func(*GetRouterConfig)

type GetRouterConfig struct {
	store               *pkmn.Store
	accessPolicyFactory func(context.Context) *AccessPolicy
}

var WithStore = func(store *pkmn.Store) GetRouterOption {
	return func(config *GetRouterConfig) {
		config.store = store
	}
}

var WithAccessPolicyFactory = func(factory func(context.Context) *AccessPolicy) GetRouterOption {
	return func(config *GetRouterConfig) {
		config.accessPolicyFactory = factory
	}
}

func GetRouter(opts ...GetRouterOption) (chi.Router, error) {
	config := &GetRouterConfig{}

	for _, opt := range opts {
		opt(config)
	}

	if config.store == nil {
		config.store = pkmn.NewStore()
	}

	if config.accessPolicyFactory == nil {
		config.accessPolicyFactory = func(ctx context.Context) *AccessPolicy {
			return NewAccessPolicy()
		}
	}

	abilityMapper := NewAbilityMapper()
	typeMapper := NewTypeMapper()
	moveMapper := NewMoveMapper(typeMapper)
	pokemonMapper := NewPokemonMapper(typeMapper, abilityMapper, config.store)

	abilityService := NewAbilityService(config.store, abilityMapper, config.accessPolicyFactory)
	typeService := NewTypeService(config.store, typeMapper, config.accessPolicyFactory)
	moveService := NewMoveService(config.store, moveMapper, config.accessPolicyFactory)
	pokemonService := NewPokemonService(config.store, pokemonMapper, config.accessPolicyFactory)

	abilityController := NewAbilitiesAPIController(abilityService)
	typeController := NewTypesAPIController(typeService)
	moveController := NewMovesAPIController(moveService)
	pokemonController := NewPokemonAPIController(pokemonService)
	return NewRouter(pokemonController, typeController, abilityController, moveController), nil
}
