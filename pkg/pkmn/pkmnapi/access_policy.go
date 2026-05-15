package pkmnapi

import (
	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/iter2"
	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/pkmn"
)

type PokemonFilterPolicy struct {
	PokemonFilter iter2.FilterFunc[pkmn.Pokemon]
	FormFilter    iter2.FilterFunc[pkmn.PokemonForm]
	TypeFilter    iter2.FilterFunc[pkmn.PokemonType]
	AbilityFilter iter2.FilterFunc[pkmn.Ability]
	MoveFilter    iter2.FilterFunc[pkmn.Move]
}

type PokemonFilterPolicyOption func(*PokemonFilterPolicy)

var WithPokemonFilter = func(filter iter2.FilterFunc[pkmn.Pokemon]) PokemonFilterPolicyOption {
	return func(p *PokemonFilterPolicy) {
		p.PokemonFilter = filter
	}
}

var WithFormFilter = func(filters ...iter2.FilterFunc[pkmn.PokemonForm]) PokemonFilterPolicyOption {
	return func(p *PokemonFilterPolicy) {
		p.FormFilter = iter2.And(filters...)
	}
}

var WithTypeFilter = func(filters ...iter2.FilterFunc[pkmn.PokemonType]) PokemonFilterPolicyOption {
	return func(p *PokemonFilterPolicy) {
		p.TypeFilter = iter2.And(filters...)
	}
}

var WithAbilityFilter = func(filters ...iter2.FilterFunc[pkmn.Ability]) PokemonFilterPolicyOption {
	return func(p *PokemonFilterPolicy) {
		p.AbilityFilter = iter2.And(filters...)
	}
}

var WithMoveFilter = func(filters ...iter2.FilterFunc[pkmn.Move]) PokemonFilterPolicyOption {
	return func(p *PokemonFilterPolicy) {
		p.MoveFilter = iter2.And(filters...)
	}
}

func NewPokemonFilterPolicy(opts ...PokemonFilterPolicyOption) *PokemonFilterPolicy {
	p := &PokemonFilterPolicy{
		PokemonFilter: iter2.True[pkmn.Pokemon],
		FormFilter:    iter2.True[pkmn.PokemonForm],
		TypeFilter:    iter2.True[pkmn.PokemonType],
		AbilityFilter: iter2.True[pkmn.Ability],
		MoveFilter:    iter2.True[pkmn.Move],
	}

	for _, opt := range opts {
		opt(p)
	}
	return p
}

func FromPokemonFilterPolicy(policy PokemonFilterPolicy, opts ...PokemonFilterPolicyOption) *PokemonFilterPolicy {
	p := &PokemonFilterPolicy{
		PokemonFilter: policy.PokemonFilter,
		FormFilter:    policy.FormFilter,
		TypeFilter:    policy.TypeFilter,
		AbilityFilter: policy.AbilityFilter,
		MoveFilter:    policy.MoveFilter,
	}

	for _, opt := range opts {
		opt(p)
	}
	return p
}
