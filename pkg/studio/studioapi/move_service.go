package studioapi

import (
	"context"

	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/studio"
)

type MoveService struct {
	store      *studio.Store
	moveMapper *MoveMapper
}

func NewMoveService(store *studio.Store, moveMapper *MoveMapper) MovesAPIServicer {
	return &MoveService{
		store,
		moveMapper,
	}
}

func (s MoveService) GetMoveDetails(requestCtx context.Context, symbol string, lang string) (ImplResponse, error) {
	m := s.store.FindMoveBySymbol(symbol)
	if m == nil {
		return ImplResponse{Code: 404, Body: nil}, nil
	}
	return ImplResponse{Code: 200, Body: s.moveMapper.ToMoveDetail(*m, lang)}, nil
}

func (s MoveService) GetMove(requestCtx context.Context, symbol string, lang string) (ImplResponse, error) {
	m := s.store.FindMoveBySymbol(symbol)
	if m == nil {
		return ImplResponse{Code: 404, Body: nil}, nil
	}
	return ImplResponse{Code: 200, Body: s.moveMapper.ToMovePartial(*m, lang)}, nil
}
