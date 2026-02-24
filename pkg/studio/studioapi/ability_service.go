package studioapi

import (
	"context"
	"slices"

	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/iter2"
	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/studio"
)

type AbilityService struct {
	store               *studio.Store
	abilityMapper       *AbilityMapper
	accessPolicyFactory func(context.Context) *AccessPolicy
}

func NewAbilityService(
	store *studio.Store,
	abilityMapper *AbilityMapper,
	accessPolicyFactory func(context.Context) *AccessPolicy,
) AbilitiesAPIServicer {
	return &AbilityService{
		store:               store,
		abilityMapper:       abilityMapper,
		accessPolicyFactory: accessPolicyFactory,
	}
}

func (s AbilityService) GetAbilities(requestCtx context.Context, lang string) (ImplResponse, error) {
	policy := s.accessPolicyFactory(requestCtx)
	abilitiesIter := s.store.FindAllAbilities(policy.AbilityFilters...)
	mappedIter := iter2.Map(func(a studio.Ability) AbilityPartial {
		return s.abilityMapper.ToAbilityPartial(a, lang)
	}, abilitiesIter)
	return ImplResponse{Code: 200, Body: slices.Collect(mappedIter)}, nil
}

func (s AbilityService) GetAbilityDetails(requestCtx context.Context, symbol string, lang string) (ImplResponse, error) {
	a := s.store.FindAbilityBySymbol(symbol)
	if a == nil {
		return ImplResponse{Code: 200, Body: nil}, nil
	}
	return ImplResponse{Code: 200, Body: s.abilityMapper.ToAbilityDetail(*a, lang)}, nil
}
