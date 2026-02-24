package studioapi

import (
	"context"

	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/pagination"
	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/studio"
)

type PokemonService struct {
	store *studio.Store

	pokemonMapper        *PokemonMapper
	pokemonContextFilter ContextFilter[studio.Pokemon]
	formContextFilter    ContextFilter[studio.PokemonForm]
}

func NewPokemonService(
	store *studio.Store,
	pokemonMapper *PokemonMapper,
	pokemonContextFilter ContextFilter[studio.Pokemon],
	formContextFilter ContextFilter[studio.PokemonForm],
) PokemonAPIServicer {
	return &PokemonService{
		store:                store,
		pokemonMapper:        pokemonMapper,
		pokemonContextFilter: pokemonContextFilter,
		formContextFilter:    formContextFilter,
	}
}

func (s PokemonService) GetPokemonDetails(requestCtx context.Context, symbol string, lang string) (ImplResponse, error) {
	pkmn := s.store.FindPokemonBySymbol(symbol)

	if pkmn == nil {
		return ImplResponse{Code: 404, Body: nil}, nil
	}

	pokemonFilter := s.pokemonContextFilter(requestCtx)
	if !pokemonFilter(*pkmn) {
		return ImplResponse{Code: 404, Body: nil}, nil
	}

	return ImplResponse{Code: 200, Body: s.pokemonMapper.PokemonToDetail(*pkmn, lang)}, nil
}

func (s PokemonService) GetPokemon(requestCtx context.Context, page int32, pageSize int32, lang string) (ImplResponse, error) {
	p := int(page)
	size := int(pageSize)
	pr := pagination.NewPageRequest(p, size)

	pokemonFilter := s.pokemonContextFilter(requestCtx)
	pkmnIter := s.store.FindAllPokemon(pokemonFilter)
	pkmnPage := pagination.Collect(pkmnIter, pr)
	thumbnails := make([]*PokemonThumbnail, len(pkmnPage.Content))

	for i, pkmn := range pkmnPage.Content {
		thumbnails[i] = s.pokemonMapper.PokemonToThumbnail(pkmn, lang)
	}

	return ImplResponse{Code: 200, Body: pagination.NewPage(pr.Page, pr.Size, thumbnails, pkmnPage.Total)}, nil
}

func (s PokemonService) GetPokemonForm(requestCtx context.Context, symbol string, form int32, lang string) (ImplResponse, error) {
	pkmn := s.store.FindPokemonBySymbol(symbol)

	if pkmn == nil {
		return ImplResponse{Code: 404, Body: nil}, nil
	}

	pokemonFilter := s.pokemonContextFilter(requestCtx)
	if !pokemonFilter(*pkmn) {
		return ImplResponse{Code: 404, Body: nil}, nil
	}

	pkmnForm, ok := pkmn.Forms[form]
	if !ok {
		return ImplResponse{Code: 404, Body: nil}, nil
	}

	formFilter := s.formContextFilter(requestCtx)
	if !formFilter(pkmnForm) {
		return ImplResponse{Code: 404, Body: nil}, nil
	}

	return ImplResponse{Code: 200, Body: s.pokemonMapper.FormToPokemonFormDetails(pkmnForm, lang)}, nil
}
