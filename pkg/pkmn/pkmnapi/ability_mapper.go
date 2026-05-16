package pkmnapi

import (
	"log/slog"

	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/pkmn"
	. "github.com/ForetEternelle/PokemonStudioDataApi/pkg/pkmn/pkmnapispec"
)

type AbilityMapper struct {
	store *pkmn.Store
}

// NewAbilityMapper Create a new ability mapper
// store the store for abilities
func NewAbilityMapper() *AbilityMapper {
	return &AbilityMapper{}
}

// ToAbilityDetail map an ability to an ability details transfer object
// ability the ability to map
// lang the language expected
func (m AbilityMapper) ToAbilityDetail(ability pkmn.Ability, lang string) AbilityDetails {
	slog.Debug("Mapping ability to details", "ability", ability.DbSymbol(), "lang", lang)
	return AbilityDetails{
		Symbol:      ability.DbSymbol(),
		Name:        ability.Name(lang),
		Description: ability.Description(lang),
	}
}

// ToAbilityPartial map an ability to an ability partial transfer object
// ability the ability to map
// lang the language expected
func (m AbilityMapper) ToAbilityPartial(ability pkmn.Ability, lang string) AbilityPartial {
	slog.Debug("Mapping ability to partial", "ability", ability.DbSymbol(), "lang", lang)
	return AbilityPartial{
		Symbol:      ability.DbSymbol(),
		Name:        ability.Name(lang),
		Description: ability.Description(lang),
	}
}
