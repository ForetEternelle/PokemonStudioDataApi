package studioapi

import (
	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/iter2"
	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/studio"
)

type AccessPolicy struct {
	PokemonFilter iter2.FilterFunc[studio.Pokemon]
	FormFilter    iter2.FilterFunc[studio.PokemonForm]
	TypeFilter    iter2.FilterFunc[studio.PokemonType]
	AbilityFilter iter2.FilterFunc[studio.Ability]
	MoveFilter    iter2.FilterFunc[studio.Move]
}

type AccessPolicyOption func(*AccessPolicy)

var WithPokemonPolicy = func(filter iter2.FilterFunc[studio.Pokemon]) AccessPolicyOption {
	return func(p *AccessPolicy) {
		p.PokemonFilter = filter
	}
}

var WithFormPolicy = func(filters ...iter2.FilterFunc[studio.PokemonForm]) AccessPolicyOption {
	return func(p *AccessPolicy) {
		p.FormFilter = iter2.And(filters...)
	}
}

var WithTypePolicy = func(filters ...iter2.FilterFunc[studio.PokemonType]) AccessPolicyOption {
	return func(p *AccessPolicy) {
		p.TypeFilter = iter2.And(filters...)
	}
}

var WithAbilityPolicy = func(filters ...iter2.FilterFunc[studio.Ability]) AccessPolicyOption {
	return func(p *AccessPolicy) {
		p.AbilityFilter = iter2.And(filters...)
	}
}

var WithMovePolicy = func(filters ...iter2.FilterFunc[studio.Move]) AccessPolicyOption {
	return func(p *AccessPolicy) {
		p.MoveFilter = iter2.And(filters...)
	}
}

func NewAccessPolicy(opts ...AccessPolicyOption) *AccessPolicy {
	p := &AccessPolicy{
		PokemonFilter: iter2.True[studio.Pokemon],
		FormFilter:    iter2.True[studio.PokemonForm],
		TypeFilter:    iter2.True[studio.PokemonType],
		AbilityFilter: iter2.True[studio.Ability],
		MoveFilter:    iter2.True[studio.Move],
	}

	for _, opt := range opts {
		opt(p)
	}
	return p
}
