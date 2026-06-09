package pkmnapi

import (
	"context"
	"log/slog"
	"slices"

	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/iter2"
	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/pkmn"
	. "github.com/ForetEternelle/PokemonStudioDataApi/pkg/pkmn/pkmnapispec"
	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/scroll"
)

type PokemonService struct {
	store                      *pkmn.Store
	pokemonMapper              *PokemonMapper
	pokemonAccessPolicyFactory func(context.Context) *PokemonFilterPolicy
}

func NewPokemonService(
	store *pkmn.Store,
	pokemonMapper *PokemonMapper,
	pokemonAccessPolicyFactory func(context.Context) *PokemonFilterPolicy,
) PokemonAPIServicer {
	return &PokemonService{
		store:                      store,
		pokemonMapper:              pokemonMapper,
		pokemonAccessPolicyFactory: pokemonAccessPolicyFactory,
	}
}

func (s PokemonService) GetPokemonDetails(requestCtx context.Context, symbol string, lang string) (ImplResponse, error) {
	policy := s.pokemonAccessPolicyFactory(requestCtx)
	pkmn := s.store.FindPokemonBySymbol(symbol, policy.PokemonFilter)

	if pkmn == nil {
		return ImplResponse{Code: 404, Body: nil}, nil
	}

	return ImplResponse{Code: 200, Body: s.pokemonMapper.PokemonToDetail(*pkmn, lang, *policy)}, nil
}

func (s PokemonService) GetPokemonDetailsByName(requestCtx context.Context, name string, lang string) (ImplResponse, error) {
	policy := s.pokemonAccessPolicyFactory(requestCtx)
	pkmn := s.store.FindPokemonByName(name, policy.PokemonFilter)

	if pkmn == nil {
		return ImplResponse{Code: 404, Body: nil}, nil
	}

	return ImplResponse{Code: 200, Body: s.pokemonMapper.PokemonToDetail(*pkmn, lang, *policy)}, nil
}

func (s PokemonService) GetPokemon(requestCtx context.Context, lang string, size int32, lastId *int32, lastForm *int32, mainFormsOnly bool, query *string, types []string) (ImplResponse, error) {
	slog.Debug("GetPokemon called with parameters", "lang", lang, "size", size, "lastId", lastId, "lastForm", lastForm, "mainFormsOnly", mainFormsOnly, "query", query, "types", types)
	policy := s.pokemonAccessPolicyFactory(requestCtx)
	formFilters := iter2.And(
		policy.FormFilter,
		pkmn.NewPokemonFormTypesFilter(types),
	)
	pokemonWithFormOptions := []pkmn.FindAllPokemonWithFormOption{
		pkmn.WithPokemonFilter(policy.PokemonFilter),
		pkmn.WithFormFilter(formFilters),
	}

	if mainFormsOnly {
		slog.Debug("Main forms only option enabled for GetPokemon")
		pokemonWithFormOptions = append(pokemonWithFormOptions, pkmn.WithMainFormsOnly())
	}

	if lastId != nil {
		slog.Debug("Last ID provided for GetPokemon", "lastId", *lastId)
		pokemonWithFormOptions = append(pokemonWithFormOptions, pkmn.WithLastId(*lastId))
	}

	if lastForm != nil {
		slog.Debug("Last form ID provided for GetPokemon", "lastFormId", *lastForm)
		pokemonWithFormOptions = append(pokemonWithFormOptions, pkmn.WithLastForm(*lastForm))
	}

	pwfIt := s.store.FindAllPokemonWithForm(pokemonWithFormOptions...)

	if query != nil {
		slog.Debug("Query filter provided for GetPokemon", "query", *query)
		pwfIt = iter2.Filter(pwfIt, pkmn.NewPokemonQueryFilter(*query, lang))
	}

	thumbnailsIt := iter2.Map(pwfIt, func(pwf pkmn.PokemonWithForm) *PokemonThumbnail {
		return s.pokemonMapper.PokemonToThumbnail(*pwf.Pokemon, pwf.FormId, lang, *policy)
	})

	return ImplResponse{Code: 200, Body: scroll.Of(thumbnailsIt, int(size))}, nil
}

func (s PokemonService) GetFormsByPokemon(requestCtx context.Context, symbol string, lang string) (ImplResponse, error) {
	policy := s.pokemonAccessPolicyFactory(requestCtx)
	poke := s.store.FindPokemonBySymbol(symbol, policy.PokemonFilter)

	if poke == nil {
		return ImplResponse{Code: 404, Body: nil}, nil
	}

	formsIter := poke.Forms()
	formPartialsIter := iter2.Map(formsIter, func(form *pkmn.PokemonForm) *FormPartial {
		return s.pokemonMapper.FormToPokemonFormPartial(*form, lang, *policy)
	})

	return ImplResponse{Code: 200, Body: slices.Collect(formPartialsIter)}, nil
}

func (s PokemonService) GetPokemonForm(requestCtx context.Context, symbol string, form int32, lang string) (ImplResponse, error) {
	policy := s.pokemonAccessPolicyFactory(requestCtx)
	pkmn := s.store.FindPokemonBySymbol(symbol, policy.PokemonFilter)

	if pkmn == nil {
		return ImplResponse{Code: 404, Body: nil}, nil
	}

	pkmnForm, ok := pkmn.Form(form)
	if !ok || !policy.FormFilter(pkmnForm) {
		return ImplResponse{Code: 404, Body: nil}, nil
	}

	return ImplResponse{Code: 200, Body: s.pokemonMapper.FormToPokemonFormDetails(*pkmnForm, lang, *policy)}, nil
}
