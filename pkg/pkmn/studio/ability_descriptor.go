package studio

import "github.com/ForetEternelle/PokemonStudioDataApi/pkg/pkmn"

// AbilityDescriptor is the JSON descriptor for an Ability.
type AbilityDescriptor struct {
	DbSymbol    string `json:"dbSymbol"`
	Id          int    `json:"id"`
	TextID      int    `json:"textId"`
	Name        pkmn.Translation
	Description pkmn.Translation
}

