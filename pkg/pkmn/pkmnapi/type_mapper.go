package pkmnapi

import (
	"log/slog"
	"maps"

	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/pkmn"
	. "github.com/ForetEternelle/PokemonStudioDataApi/pkg/pkmn/pkmnapispec"
)

type TypeMapper struct {
}

func NewTypeMapper() *TypeMapper {
	return &TypeMapper{}
}

func (t TypeMapper) ToTypeDetail(pokemonType pkmn.PokemonType, lang string, policy *AccessPolicy) *TypeDetails {
	slog.Debug("Mapping type to details", "type", pokemonType.DbSymbol(), "lang", lang)

	return &TypeDetails{
		Symbol:     pokemonType.DbSymbol(),
		Name:       pokemonType.Name(lang),
		Color:      pokemonType.Color(),
		TypeDamage: maps.Collect(pokemonType.DamageTo()),
	}
}

func (t TypeMapper) ToTypePartial(pokemonType pkmn.PokemonType, lang string, policy *AccessPolicy) *TypePartial {
	slog.Debug("Mapping type to partial", "type", pokemonType.DbSymbol(), "lang", lang)
	return &TypePartial{
		Symbol: pokemonType.DbSymbol(),
		Name:   pokemonType.Name(lang),
		Color:  pokemonType.Color(),
	}
}
