package studioapi

import (
	"context"

	"github.com/ForetEternelle/ForetEternelleDataApi/pkg/studio"
)

type TypeService struct {
	store      *studio.Store
	typeMapper *TypeMapper
}

func NewTypeService(store *studio.Store, typeMapper *TypeMapper) TypesAPIServicer {
	return &TypeService{
		store,
		typeMapper,
	}
}

func (s TypeService) GetTypes(requestCtx context.Context, lang string) (ImplResponse, error) {
	types := s.store.FindAllTypes()
	res := make([]TypePartial, len(types))

	for i, t := range types {
		res[i] = s.typeMapper.ToTypePartial(t, lang)
	}
	return ImplResponse{Code: 200, Body: res}, nil
}

func (s TypeService) GetTypeDetails(requestCtx context.Context, symbol string, lang string) (ImplResponse, error) {
	t := s.store.FindTypeBySymbol(symbol)
	if t == nil {
		return ImplResponse{Code: 200, Body: nil}, nil
	}
	return ImplResponse{Code: 200, Body: s.typeMapper.ToTypeDetail(*t, lang)}, nil
}
