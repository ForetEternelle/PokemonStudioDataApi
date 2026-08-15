package pkmn

import "testing"

func newTestType(t *testing.T, damageTo map[string]float32) *PokemonType {
	t.Helper()
	return &PokemonType{DbSymbol: "test", DamageTo: damageTo}
}

func newTestTypeDto(damageTo map[string]float32) AddTypeDto {
	return AddTypeDto{DbSymbol: "test", DamageTo: damageTo}
}

func newTestForm(t *testing.T, form int32) *PokemonForm {
	t.Helper()
	return &PokemonForm{Form: form}
}

func newTestFormDto(form int32) AddPokemonFormDto {
	return AddPokemonFormDto{Form: form}
}

func newTestPokemon(t *testing.T, id int32, dbSymbol string, forms ...*PokemonForm) *Pokemon {
	t.Helper()
	return &Pokemon{ID: id, DbSymbol: dbSymbol, Forms: forms}
}

func newTestPokemonDto(id int32, dbSymbol string, forms ...AddPokemonFormDto) AddPokemonDto {
	return AddPokemonDto{ID: id, DbSymbol: dbSymbol, Forms: forms}
}
