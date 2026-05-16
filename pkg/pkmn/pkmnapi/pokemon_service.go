package pkmnapi

import (
	"context"
	"slices"

	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/iter2"
	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/pagination"
	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/pkmn"
	. "github.com/ForetEternelle/PokemonStudioDataApi/pkg/pkmn/pkmnapispec"
)

type PokemonService struct {
	store               *pkmn.Store
	pokemonMapper       *PokemonMapper
	accessPolicyFactory func(context.Context) *AccessPolicy
}

func NewPokemonService(
	store *pkmn.Store,
	pokemonMapper *PokemonMapper,
	accessPolicyFactory func(context.Context) *AccessPolicy,
) PokemonAPIServicer {
	return &PokemonService{
		store:               store,
		pokemonMapper:       pokemonMapper,
		accessPolicyFactory: accessPolicyFactory,
	}
}

func (s PokemonService) GetPokemonDetails(requestCtx context.Context, symbol string, lang string) (ImplResponse, error) {
	policy := s.accessPolicyFactory(requestCtx)
	pkmn := s.store.FindPokemonBySymbol(symbol, policy.PokemonFilter)

	if pkmn == nil {
		return ImplResponse{Code: 404, Body: nil}, nil
	}

	return ImplResponse{Code: 200, Body: s.pokemonMapper.PokemonToDetail(*pkmn, lang, policy)}, nil
}

func (s PokemonService) GetPokemonDetailsByName(requestCtx context.Context, name string, lang string) (ImplResponse, error) {
	policy := s.accessPolicyFactory(requestCtx)
	pkmn := s.store.FindPokemonByName(name, policy.PokemonFilter)

	if pkmn == nil {
		return ImplResponse{Code: 404, Body: nil}, nil
	}

	return ImplResponse{Code: 200, Body: s.pokemonMapper.PokemonToDetail(*pkmn, lang, policy)}, nil
}

func (s PokemonService) GetPokemon(requestCtx context.Context, page int32, pageSize int32, lang string) (ImplResponse, error) {
	policy := s.accessPolicyFactory(requestCtx)

	p := int(page)
	size := int(pageSize)
	pr := pagination.NewPageRequest(p, size)

	pkmnIter := s.store.FindAllPokemon(policy.PokemonFilter)
	pkmnPage := pagination.Collect(pkmnIter, pr)

	thumbnailsIter := iter2.Map(slices.Values(pkmnPage.Content), func(poke pkmn.Pokemon) *PokemonThumbnail {
		return s.pokemonMapper.PokemonToThumbnail(poke, lang, policy)
	})

	return ImplResponse{Code: 200, Body: pagination.NewPage(pr.Page, pr.Size, slices.Collect(thumbnailsIter), pkmnPage.Total)}, nil
}

func (s PokemonService) GetFormsByPokemon(requestCtx context.Context, symbol string, lang string) (ImplResponse, error) {
	policy := s.accessPolicyFactory(requestCtx)
	poke := s.store.FindPokemonBySymbol(symbol, policy.PokemonFilter)

	if poke == nil {
		return ImplResponse{Code: 404, Body: nil}, nil
	}

	formsIter := iter2.Values(poke.Forms())
	formPartialsIter := iter2.Map(formsIter, func(form pkmn.PokemonForm) *FormPartial {
		return s.pokemonMapper.FormToPokemonFormPartial(form, lang, policy)
	})

	return ImplResponse{Code: 200, Body: slices.Collect(formPartialsIter)}, nil
}

func (s PokemonService) GetPokemonForm(requestCtx context.Context, symbol string, form int32, lang string) (ImplResponse, error) {
	policy := s.accessPolicyFactory(requestCtx)
	pkmn := s.store.FindPokemonBySymbol(symbol, policy.PokemonFilter)

	if pkmn == nil {
		return ImplResponse{Code: 404, Body: nil}, nil
	}

	pkmnForm, ok := pkmn.Form(form)
	if !ok {
		return ImplResponse{Code: 404, Body: nil}, nil
	}

	return ImplResponse{Code: 200, Body: s.pokemonMapper.FormToPokemonFormDetails(pkmnForm, lang, policy)}, nil
}
