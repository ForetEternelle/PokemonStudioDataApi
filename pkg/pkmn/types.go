package pkmn

import (
	"errors"
	"iter"
	"maps"
)

// PokemonTypeConfig configures a new PokemonType.
type PokemonTypeConfig struct {
	DbSymbol string
	Color    string
	TextId   int
	Name     Translation
	DamageTo map[string]float32
}

// NewPokemonType creates an immutable PokemonType from the given config.
func NewPokemonType(cfg PokemonTypeConfig) (*PokemonType, error) {
	if cfg.DbSymbol == "" {
		return nil, errors.New("pkmn: PokemonType dbSymbol is required")
	}

	return &PokemonType{
		dbSymbol: cfg.DbSymbol,
		color:    cfg.Color,
		textId:   cfg.TextId,
		name:     cfg.Name,
		damageTo: maps.Clone(cfg.DamageTo),
	}, nil
}

// PokemonType represents a Pokemon type (e.g., Fire, Water, Grass).
type PokemonType struct {
	dbSymbol string
	color    string
	textId   int
	name     Translation
	damageTo map[string]float32
}

// DbSymbol returns the database symbol of the PokemonType.
func (t *PokemonType) DbSymbol() string {
	return t.dbSymbol
}

// Color returns the color of the PokemonType.
func (t *PokemonType) Color() string {
	return t.color
}

// TextId returns the text ID of the PokemonType.
func (t *PokemonType) TextId() int {
	return t.textId
}

// Name returns the localized name of the PokemonType for the given language.
func (t *PokemonType) Name(lang string) string {
	return t.name[lang]
}

// DamageTo returns an iterator over the type damage relations.
func (t *PokemonType) DamageTo() iter.Seq2[string, float32] {
	return maps.All(t.damageTo)
}

// Damage returns the type damage relation for a defending type.
func (t *PokemonType) Damage(defType string) (float32, bool) {
	factor, ok := t.damageTo[defType]
	return factor, ok
}

// DamageToTypes calculates the overall damage factor when attacking a Pokemon with the given types.
func (t *PokemonType) DamageToTypes(type1, type2 string) float32 {
	var res float32 = 1
	type1Dmg, okType1 := t.Damage(type1)
	if okType1 {
		res *= type1Dmg
	}

	type2Dmg, okType2 := t.Damage(type2)
	if okType2 {
		res *= type2Dmg
	}
	return res
}


