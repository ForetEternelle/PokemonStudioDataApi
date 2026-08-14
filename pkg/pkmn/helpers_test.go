package pkmn

import "testing"

func newTestType(t *testing.T, damageTo map[string]float32) *PokemonType {
	t.Helper()
	pokemonType, err := NewPokemonType(PokemonTypeConfig{DbSymbol: "test", DamageTo: damageTo})
	if err != nil {
		t.Fatal(err)
	}
	return pokemonType
}

func newTestForm(t *testing.T, form int32) *PokemonForm {
	t.Helper()
	f, err := NewPokemonForm(PokemonFormConfig{Form: form})
	if err != nil {
		t.Fatal(err)
	}
	return f
}

func newTestPokemon(t *testing.T, id int32, dbSymbol string, forms ...*PokemonForm) *Pokemon {
	t.Helper()
	pokemon, err := NewPokemon(PokemonConfig{ID: id, DbSymbol: dbSymbol, Forms: forms})
	if err != nil {
		t.Fatal(err)
	}
	return pokemon
}
