package studioapi

import (
	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/iter2"
	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/studio"
)

type AccessPolicy struct {
	PokemonFilters []iter2.FilterFunc[studio.Pokemon]
	FormFilters    []iter2.FilterFunc[studio.PokemonForm]
	TypeFilters    []iter2.FilterFunc[studio.PokemonType]
	AbilityFilters []iter2.FilterFunc[studio.Ability]
	MoveFilters    []iter2.FilterFunc[studio.Move]
}

type AccessPolicyOption func(*AccessPolicy)

var WithPokemonPolicy = func(filters ...iter2.FilterFunc[studio.Pokemon]) AccessPolicyOption {
	return func(p *AccessPolicy) {
		p.PokemonFilters = filters
	}
}

var WithFormPolicy = func(filters ...iter2.FilterFunc[studio.PokemonForm]) AccessPolicyOption {
	return func(p *AccessPolicy) {
		p.FormFilters = filters
	}
}

var WithTypePolicy = func(filters ...iter2.FilterFunc[studio.PokemonType]) AccessPolicyOption {
	return func(p *AccessPolicy) {
		p.TypeFilters = filters
	}
}

var WithAbilityPolicy = func(filters ...iter2.FilterFunc[studio.Ability]) AccessPolicyOption {
	return func(p *AccessPolicy) {
		p.AbilityFilters = filters
	}
}

var WithMovePolicy = func(filters ...iter2.FilterFunc[studio.Move]) AccessPolicyOption {
	return func(p *AccessPolicy) {
		p.MoveFilters = filters
	}
}

func NewAccessPolicy(opts ...AccessPolicyOption) *AccessPolicy {
	p := &AccessPolicy{}
	for _, opt := range opts {
		opt(p)
	}
	return p
}
