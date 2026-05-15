package studio

// AbilityMapper maps Ability descriptors to Ability entities.
type AbilityMapper struct {
	store *Store
}

// NewAbilityMapper creates a new AbilityMapper.
func NewAbilityMapper(store *Store) *AbilityMapper {
	return &AbilityMapper{store: store}
}

// MapAbilityDescriptorToAbility maps an AbilityDescriptor to an Ability.
func (m *AbilityMapper) MapAbilityDescriptorToAbility(desc AbilityDescriptor) *Ability {
	ability := NewAbilityBuilder().
		DbSymbol(desc.DbSymbol).
		ID(desc.Id).
		TextId(desc.TextID).
		Name(desc.Name).
		Description(desc.Description).
		Build()

	return ability
}
