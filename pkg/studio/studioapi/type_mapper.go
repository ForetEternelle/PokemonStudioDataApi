package studioapi

import (
	"log/slog"

	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/studio"
)

type TypeMapper struct {
}

// NewTypeMapper create a new pokemon type mapper
func NewTypeMapper() *TypeMapper {
	return &TypeMapper{}
}

// ToTypeDetail map a type to a type details transfer object
// pokemonType the pokemon type to map
// lang the language expected
func (t TypeMapper) ToTypeDetail(pokemonType studio.PokemonType, lang string) *TypeDetails {
	slog.Debug("Mapping type to details", "type", pokemonType, "lang", lang)
	typeDamage := make([]TypeDamage, len(pokemonType.DamageTo))
	for i, damage := range pokemonType.DamageTo {
		typeDamage[i] = TypeDamage{
			DefensiveType: damage.DefensiveType,
			Factor:        &damage.Factor,
		}
		typeDamage[i].Factor = &damage.Factor
	}
	return &TypeDetails{
		Symbol:     pokemonType.DbSymbol,
		Name:       pokemonType.Name[lang],
		Color:      pokemonType.Color,
		TypeDamage: typeDamage,
	}
}

// ToTypePartial map a type to a type partial transfer object
// pokemonType the pokemon type to map
// lang the language expected
func (t TypeMapper) ToTypePartial(pokemonType studio.PokemonType, lang string) *TypePartial {
	slog.Debug("Mapping type to partial", "type", pokemonType, "lang", lang)
	return &TypePartial{
		Symbol: pokemonType.DbSymbol,
		Name:   pokemonType.Name[lang],
		Color:  pokemonType.Color,

	}
}
