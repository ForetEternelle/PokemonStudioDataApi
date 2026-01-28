package studioapi

import (
	"log/slog"

	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/studio"
)

type AbilityMapper struct {
	store *studio.Store
}

// NewAbilityMapper Create a new ability mapper
// store the store for abilities
func NewAbilityMapper() *AbilityMapper {
	return &AbilityMapper{}
}

// ToAbilityDetail map an ability to an ability details transfer object
// ability the ability to map
// lang the language expected
func (m AbilityMapper) ToAbilityDetail(ability studio.Ability, lang string) AbilityDetails {
	slog.Debug("Mapping ability to details", "ability", ability, "lang", lang)
	return AbilityDetails{
		Symbol:      ability.DbSymbol,
		Name:        ability.Name[lang],
		Description: ability.Description[lang],
	}
}

// ToAbilityPartial map an ability to an ability partial transfer object
// ability the ability to map
// lang the language expected
func (m AbilityMapper) ToAbilityPartial(ability studio.Ability, lang string) AbilityPartial {
	slog.Debug("Mapping ability to partial", "ability", ability, "lang", lang)
	return AbilityPartial{
		Symbol:      ability.DbSymbol,
		Name:        ability.Name[lang],
		Description: ability.Description[lang],
	}
}
