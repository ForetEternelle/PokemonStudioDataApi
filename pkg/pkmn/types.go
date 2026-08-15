package pkmn

// PokemonType represents a Pokemon type (e.g., Fire, Water, Grass).
//
// Entities are immutable by convention only: fields are exported but must not
// be mutated after the entity has been registered in a store.
type PokemonType struct {
	// DbSymbol is the database symbol of the PokemonType.
	// Immutable: used to index PokemonTypes in the store.
	DbSymbol string
	Color    string
	TextId   int
	Name     Translation
	DamageTo map[string]float32
}

// Damage returns the type damage relation for a defending type.
func (t *PokemonType) Damage(defType string) (float32, bool) {
	factor, ok := t.DamageTo[defType]
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
