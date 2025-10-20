package studioapi

import (
	"context"

	"github.com/ForetEternelle/ForetEternelleDataApi/pkg/studio"
	"github.com/ForetEternelle/ForetEternelleDataApi/pkg/pagination"
)

type pokemonService struct {
	store         *studio.Store

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

func (s pokemonService) GetPokemon(requestCtx context.Context, page int32, pageSize int32, lang string) (ImplResponse, error) {
	p := int(page)
	size := int(pageSize)
	pr := pagination.NewPageRequest(p, size)

	pkmnIter := s.store.FindAllPokemon()
	pkmnPage  := pagination.Collect(pkmnIter, pr)
	thumbnails := make([]*PokemonThumbnail, len(pkmnPage.Content))

	for i, pkmn := range pkmnPage.Content {
		thumbnails[i] = s.pokemonMapper.PokemonToThumbnail(pkmn, lang)
	}

	return ImplResponse{Code: 200, Body: pagination.NewPage(pr.Page, pr.Size, thumbnails, pkmnPage.Total)}, nil
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
