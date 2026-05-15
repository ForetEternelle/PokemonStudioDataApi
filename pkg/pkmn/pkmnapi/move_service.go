package pkmnapi

import (
	"context"

	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/pkmn"
)

type MoveService struct {
	store               *pkmn.Store
	moveMapper          *MoveMapper
	accessPolicyFactory func(context.Context) *AccessPolicy
}

func NewMoveService(
	store *pkmn.Store,
	moveMapper *MoveMapper,
	accessPolicyFactory func(context.Context) *AccessPolicy,
) MovesAPIServicer {
	return &MoveService{
		store:               store,
		moveMapper:          moveMapper,
		accessPolicyFactory: accessPolicyFactory,
	}
}

func (s MoveService) GetMoveDetails(requestCtx context.Context, symbol string, lang string) (ImplResponse, error) {
	policy := s.accessPolicyFactory(requestCtx)
	m := s.store.FindMoveBySymbol(symbol, policy.MoveFilter)
	if m == nil {
		return ImplResponse{Code: 404, Body: nil}, nil
	}
	return ImplResponse{Code: 200, Body: s.moveMapper.ToMoveDetail(*m, lang, policy)}, nil
}

func (s MoveService) GetMove(requestCtx context.Context, symbol string, lang string) (ImplResponse, error) {
	policy := s.accessPolicyFactory(requestCtx)
	m := s.store.FindMoveBySymbol(symbol, policy.MoveFilter)
	if m == nil {
		return ImplResponse{Code: 404, Body: nil}, nil
	}
	return ImplResponse{Code: 200, Body: s.moveMapper.ToMovePartial(*m, lang, policy)}, nil
}
