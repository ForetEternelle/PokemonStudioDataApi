package studioapi

import (
	"context"
	"slices"

	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/iter2"
	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/studio"
)

type TypeService struct {
	store               *studio.Store
	typeMapper          *TypeMapper
	accessPolicyFactory func(context.Context) *AccessPolicy
}

func NewTypeService(
	store *studio.Store,
	typeMapper *TypeMapper,
	accessPolicyFactory func(context.Context) *AccessPolicy,
) TypesAPIServicer {
	return &TypeService{
		store:               store,
		typeMapper:          typeMapper,
		accessPolicyFactory: accessPolicyFactory,
	}
}

func (s TypeService) GetTypes(requestCtx context.Context, lang string) (ImplResponse, error) {
	policy := s.accessPolicyFactory(requestCtx)
	typesIter := s.store.FindAllTypes(policy.TypeFilters...)
	mappedIter := iter2.Map(func(t studio.PokemonType) TypePartial {
		return *s.typeMapper.ToTypePartial(t, lang, policy)
	}, typesIter)
	return ImplResponse{Code: 200, Body: slices.Collect(mappedIter)}, nil
}

func (s TypeService) GetTypeDetails(requestCtx context.Context, symbol string, lang string) (ImplResponse, error) {
	policy := s.accessPolicyFactory(requestCtx)
	t := s.store.FindTypeBySymbol(symbol, policy.TypeFilters...)
	if t == nil {
		return ImplResponse{Code: 200, Body: nil}, nil
	}
	return ImplResponse{Code: 200, Body: s.typeMapper.ToTypeDetail(*t, lang, policy)}, nil
}
