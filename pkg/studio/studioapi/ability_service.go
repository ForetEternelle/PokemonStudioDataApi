package studioapi

import (
	"context"

	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/studio"
)

type AbilityService struct {
	store         *studio.Store
	abilityMapper *AbilityMapper
}

func NewAbilityService(store *studio.Store, abilityMapper *AbilityMapper) AbilitiesAPIServicer {
	return &AbilityService{
		store,
		abilityMapper,
	}
}

func (s AbilityService) GetAbilities(requestCtx context.Context, lang string) (ImplResponse, error) {
	abilities := s.store.FindAllAbilities()
	res := make([]AbilityPartial, len(abilities))

	for i, a := range abilities {
		res[i] = s.abilityMapper.ToAbilityPartial(a, lang)
	}
	return ImplResponse{Code: 200, Body: res}, nil
}

func (s AbilityService) GetAbilityDetails(requestCtx context.Context, symbol string, lang string) (ImplResponse, error) {
	a := s.store.FindAbilityBySymbol(symbol)
	if a == nil {
		return ImplResponse{Code: 200, Body: nil}, nil
	}
	return ImplResponse{Code: 200, Body: s.abilityMapper.ToAbilityDetail(*a, lang)}, nil
}
