package studio

import (
	"iter"
	"maps"
)

// PokemonType represents a Pokemon type (e.g., Fire, Water, Grass).
type PokemonType struct {
	dbSymbol string
	color    string
	textId   int
	name     Translation
	damageTo map[string]float32
}

// DbSymbol returns the database symbol of the PokemonType.
func (t PokemonType) DbSymbol() string {
	return t.dbSymbol
}

// Color returns the color of the PokemonType.
func (t PokemonType) Color() string {
	return t.color
}

// TextId returns the text ID of the PokemonType.
func (t PokemonType) TextId() int {
	return t.textId
}

// Name returns the localized name of the PokemonType for the given language.
func (t PokemonType) Name(lang string) string {
	return t.name[lang]
}

// DamageTo returns an iterator over the type damage relations.
func (t PokemonType) DamageTo() iter.Seq2[string, float32] {
	return maps.All(t.damageTo)
}

// Damage returns the type damage relation for a defending type.
func (t PokemonType) Damage(defType string) (float32, bool) {
	factor, ok := t.damageTo[defType]
	return factor, ok
}
