package studioapi

import (
	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/studio"
	"github.com/go-chi/chi/v5"
)

func MakeDefaultRouter(store *studio.Store) chi.Router {
	abilityMapper := NewAbilityMapper()
	abilityService := NewAbilityService(store, abilityMapper)
	abilityController := NewAbilitiesAPIController(abilityService)

	typeMapper := NewTypeMapper()
	typeService := NewTypeService(store, typeMapper)
	typeController := NewTypesAPIController(typeService)

	moveMapper := NewMoveMapper(typeMapper)
	moveService := NewMoveService(store, moveMapper)
	moveController := NewMovesAPIController(moveService)

	pokemonMapper := NewPokemonMapper(typeMapper, abilityMapper, store)
	pokemonService := NewPokemonService(store, pokemonMapper)
	pokemonController := NewPokemonAPIController(pokemonService)

	return NewRouter(pokemonController, typeController, abilityController, moveController)
}
