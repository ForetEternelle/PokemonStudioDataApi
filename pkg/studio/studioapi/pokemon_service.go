package studioapi

import (
	"context"
	"slices"

	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/iter2"
	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/pagination"
	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/studio"
)

type PokemonService struct {
	store               *studio.Store
	pokemonMapper       *PokemonMapper
	accessPolicyFactory func(context.Context) *AccessPolicy
}

func NewPokemonService(
	store *studio.Store,
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

func (s PokemonService) GetPokemon(requestCtx context.Context, page int32, pageSize int32, lang string) (ImplResponse, error) {
	policy := s.accessPolicyFactory(requestCtx)

	p := int(page)
	size := int(pageSize)
	pr := pagination.NewPageRequest(p, size)

	pkmnIter := s.store.FindAllPokemon(policy.PokemonFilter)
	pkmnPage := pagination.Collect(pkmnIter, pr)

	thumbnailsIter := iter2.Map(func(pkmn studio.Pokemon) *PokemonThumbnail {
		return s.pokemonMapper.PokemonToThumbnail(pkmn, lang, policy)
	}, slices.Values(pkmnPage.Content))

	return ImplResponse{Code: 200, Body: pagination.NewPage(pr.Page, pr.Size, slices.Collect(thumbnailsIter), pkmnPage.Total)}, nil
}

func (s PokemonService) GetFormsByPokemon(requestCtx context.Context, symbol string, lang string) (ImplResponse, error) {
	policy := s.accessPolicyFactory(requestCtx)
	pkmn := s.store.FindPokemonBySymbol(symbol, policy.PokemonFilter)

	if pkmn == nil {
		return ImplResponse{Code: 404, Body: nil}, nil
	}

	formsIter := iter2.Map2(func(_ int32, form studio.PokemonForm) *FormPartial {
		return s.pokemonMapper.FormToPokemonFormPartial(form, lang, policy)
	}, pkmn.Forms())

	return ImplResponse{Code: 200, Body: slices.Collect(formsIter)}, nil
}

func (s PokemonService) GetPokemonForm(requestCtx context.Context, symbol string, form int32, lang string) (ImplResponse, error) {
	policy := s.accessPolicyFactory(requestCtx)
	pkmn := s.store.FindPokemonBySymbol(symbol, policy.PokemonFilter)

	if pkmn == nil {
		return ImplResponse{Code: 404, Body: nil}, nil
	}

	var pkmnForm studio.PokemonForm
	hasForm := false
	formFilter := iter2.And(policy.FormFilter)
	for formId, f := range pkmn.Forms() {
		if formId == form && formFilter(f) {
			pkmnForm = f
			hasForm = true
			break
		}
	}

	if !hasForm {
		return ImplResponse{Code: 404, Body: nil}, nil
	}

	return ImplResponse{Code: 200, Body: s.pokemonMapper.FormToPokemonFormDetails(pkmnForm, lang, policy)}, nil
}
