package studio

import "github.com/ForetEternelle/PokemonStudioDataApi/pkg/pkmn"

// TypeMapper maps Type descriptors to PokemonType entities.
type TypeMapper struct {
	store *pkmn.Store
}

// NewTypeMapper creates a new TypeMapper.
func NewTypeMapper(store *pkmn.Store) *TypeMapper {
	return &TypeMapper{store: store}
}

// MapPokemonTypeDescriptorToPokemonType maps a PokemonTypeDescriptor to a PokemonType.
func (m *TypeMapper) MapPokemonTypeDescriptorToPokemonType(desc PokemonTypeDescriptor) *pkmn.PokemonType {
	pokemonType := pkmn.NewTypeBuilder().
		DbSymbol(desc.DbSymbol).
		Color(desc.Color).
		TextId(desc.TextId).
		Name(desc.Name).
		DamageTo(m.MapTypeDamages(desc.DamageTo)).
		Build()

	return pokemonType
}

// MapTypeDamages maps type damage descriptors to type damage entities.
func (m *TypeMapper) MapTypeDamages(damages []TypeDamageDescriptor) map[string]float32 {
	mapped := make(map[string]float32, len(damages))
	for _, tdDesc := range damages {
		mapped[tdDesc.DefensiveType] = tdDesc.Factor
	}
	return mapped
}
