package studioapi

import (
	"context"

	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/studio"
)

type pokemonService struct {
	store *studio.Store

	pokemonMapper *PokemonMapper
}

func NewPokemonService(store *studio.Store, pokemonMapper *PokemonMapper) PokemonAPIServicer {
	return &pokemonService{
		store:         store,
		pokemonMapper: pokemonMapper,
	}
}

func (s pokemonService) GetPokemonDetails(requestCtx context.Context, symbol string, lang string) (ImplResponse, error) {
	pkmn := s.store.FindPokemonBySymbol(symbol)

	if pkmn == nil {
		return ImplResponse{Code: 200, Body: nil}, nil
	}
	return ImplResponse{Code: 200, Body: s.pokemonMapper.PokemonToDetail(*pkmn, lang)}, nil
}

func (s pokemonService) GetPokemon(requestCtx context.Context, lang string) (ImplResponse, error) {
	pkmnIter := s.store.FindAllPokemon()
	thumbnails := make([]*PokemonThumbnail, 0, 1000)

	for pkmn := range pkmnIter {
		thumbnails = append(thumbnails, s.pokemonMapper.PokemonToThumbnail(pkmn, lang))
	}

	return ImplResponse{Code: 200, Body: thumbnails}, nil
}

func (s pokemonService) GetPokemonForm(requestCtx context.Context, symbol string, form int32, lang string) (ImplResponse, error) {
	f := int(form)
	pkmn := s.store.FindPokemonBySymbol(symbol)

	if pkmn == nil {
		return ImplResponse{Code: 404, Body: nil}, nil
	}

	if f > len(pkmn.Forms)-1 {
		return ImplResponse{Code: 404, Body: nil}, nil
	}

	pkmnForm := pkmn.Forms[f]
	return ImplResponse{Code: 200, Body: s.pokemonMapper.FormToPokemonFormDetails(pkmnForm, lang)}, nil
}
