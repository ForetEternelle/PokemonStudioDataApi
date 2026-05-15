package studio

import "github.com/ForetEternelle/PokemonStudioDataApi/pkg/pkmn"

// AbilityMapper maps Ability descriptors to Ability entities.
type AbilityMapper struct {
	store *pkmn.Store
}

// NewAbilityMapper creates a new AbilityMapper.
func NewAbilityMapper(store *pkmn.Store) *AbilityMapper {
	return &AbilityMapper{store: store}
}

// MapAbilityDescriptorToAbility maps an AbilityDescriptor to an Ability.
func (m *AbilityMapper) MapAbilityDescriptorToAbility(desc AbilityDescriptor) *pkmn.Ability {
	ability := pkmn.NewAbilityBuilder().
		DbSymbol(desc.DbSymbol).
		ID(desc.Id).
		TextId(desc.TextID).
		Name(desc.Name).
		Description(desc.Description).
		Build()

	return ability
}
