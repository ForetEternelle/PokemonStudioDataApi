package psapi

import (
	"github.com/go-chi/chi/v5"

	"github.com/rcharre/psapi/pkg/studio"
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
