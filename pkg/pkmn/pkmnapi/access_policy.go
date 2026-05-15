package pkmnapi

import (
	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/iter2"
	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/pkmn"
)

type AccessPolicy struct {
	PokemonFilter iter2.FilterFunc[pkmn.Pokemon]
	FormFilter    iter2.FilterFunc[pkmn.PokemonForm]
	TypeFilter    iter2.FilterFunc[pkmn.PokemonType]
	AbilityFilter iter2.FilterFunc[pkmn.Ability]
	MoveFilter    iter2.FilterFunc[pkmn.Move]
}

type AccessPolicyOption func(*AccessPolicy)

var WithPokemonPolicy = func(filter iter2.FilterFunc[pkmn.Pokemon]) AccessPolicyOption {
	return func(p *AccessPolicy) {
		p.PokemonFilter = filter
	}
}

var WithFormPolicy = func(filters ...iter2.FilterFunc[pkmn.PokemonForm]) AccessPolicyOption {
	return func(p *AccessPolicy) {
		p.FormFilter = iter2.And(filters...)
	}
}

var WithTypePolicy = func(filters ...iter2.FilterFunc[pkmn.PokemonType]) AccessPolicyOption {
	return func(p *AccessPolicy) {
		p.TypeFilter = iter2.And(filters...)
	}
}

var WithAbilityPolicy = func(filters ...iter2.FilterFunc[pkmn.Ability]) AccessPolicyOption {
	return func(p *AccessPolicy) {
		p.AbilityFilter = iter2.And(filters...)
	}
}

var WithMovePolicy = func(filters ...iter2.FilterFunc[pkmn.Move]) AccessPolicyOption {
	return func(p *AccessPolicy) {
		p.MoveFilter = iter2.And(filters...)
	}
}

func NewAccessPolicy(opts ...AccessPolicyOption) *AccessPolicy {
	p := &AccessPolicy{
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
