package studio

import (
	"iter"
)

// PokemonType represents a Pokemon type (e.g., Fire, Water, Grass).
type PokemonType struct {
	dbSymbol string
	color    string
	textId   int
	name     Translation
	damageTo []TypeDamage
}

// TypeDamage represents a type damage relation.
type TypeDamage struct {
	DefensiveType string
	Factor        float32
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
func (t PokemonType) DamageTo() iter.Seq2[int, TypeDamage] {
	return func(yield func(int, TypeDamage) bool) {
		for i, d := range t.damageTo {
			if !yield(i, d) {
				return
			}
		}
	}
}

// Damage returns the type damage relation for a defending type.
func (t PokemonType) Damage(defType string) (TypeDamage, bool) {
	for _, d := range t.damageTo {
		if d.DefensiveType == defType {
			return d, true
		}
	}
	return TypeDamage{}, false
}




