package psapi

import (
	"github.com/ForetEternelle/ForetEternelleDataApi/pkg/studio"
	"github.com/go-chi/chi/v5"
)

func MakeDefaultRouter(store *studio.Store) chi.Router {
	typeMapper := NewTypeMapper()
	typeService := NewTypeService(store, typeMapper)
	typeController := NewTypesAPIController(typeService)

	pokemonMapper := NewPokemonMapper(typeMapper, store)
	pokemonService := NewPokemonService(store, pokemonMapper)
	pokemonController := NewPokemonAPIController(pokemonService)

	return NewRouter(pokemonController, typeController)
}
