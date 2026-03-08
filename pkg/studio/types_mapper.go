package studio


// TypeMapper maps Type descriptors to PokemonType entities.
type TypeMapper struct {
	store *Store
}

// NewTypeMapper creates a new TypeMapper.
func NewTypeMapper(store *Store) *TypeMapper {
	return &TypeMapper{store: store}
}

// MapPokemonTypeDescriptorToPokemonType maps a PokemonTypeDescriptor to a PokemonType.
func (m *TypeMapper) MapPokemonTypeDescriptorToPokemonType(desc PokemonTypeDescriptor) *PokemonType {
	pokemonType := NewTypeBuilder().
		DbSymbol(desc.DbSymbol).
		Color(desc.Color).
		TextId(desc.TextId).
		Name(desc.Name).
		DamageTo(m.MapTypeDamages(desc.DamageTo)).
		Build()

	return pokemonType
}

// MapTypeDamages maps type damage descriptors to type damage entities.
func (m *TypeMapper) MapTypeDamages(damages []TypeDamageDescriptor) []TypeDamage {
	if len(damages) == 0 {
		return nil
	}

	mapped := make([]TypeDamage, len(damages))
	for i, tdDesc := range damages {
		mapped[i] = TypeDamage{
			DefensiveType: tdDesc.DefensiveType,
			Factor:        tdDesc.Factor,
		}
	}
	return mapped
}
