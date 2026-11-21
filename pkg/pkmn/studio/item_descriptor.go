package studio

import "github.com/ForetEternelle/PokemonStudioDataApi/pkg/pkmn"

// techItemClass is the Studio klass of a technical machine item.
const techItemClass = "TechItem"

// TechItemDescriptor is the JSON descriptor for a technical machine item.
type TechItemDescriptor struct {
	ID       int    `json:"id"`
	DbSymbol string `json:"dbSymbol"`
	Klass    string `json:"klass"`
	IsHm     bool   `json:"isHm"`
	Move     string `json:"move"`
	Name     pkmn.Translation
}
