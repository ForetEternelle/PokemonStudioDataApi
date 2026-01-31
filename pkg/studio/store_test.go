package studio

import (
	"slices"
	"testing"
)

const DataFolder = "../../test/test_resources/valid-data"

func TestLoad(t *testing.T) {
	store, err := Load(DataFolder)
	if err != nil {
		t.Fatal("Import should succeed", err)
	}

	pokemonIter := store.FindAllPokemon()
	pokemonCount := len(slices.Collect(pokemonIter))

	if pokemonCount != 5 {
		t.Error("Import should have 5 pokemon, has", pokemonCount)
	}

	types := store.FindAllTypes()
	typeCount := len(types)
	if typeCount != 18 {
		t.Error("Import should have 18 types, has", typeCount)
	}
}

func TestFindTypeBySymbol(t *testing.T) {
	types := []PokemonType{{DbSymbol: "test"}}
	store := NewStore()

	for _, pokemonType := range types {
		store.AddType(pokemonType)
	}

	found := store.FindTypeBySymbol("test")
	if found == nil {
		t.Error("Should find type with symbol test")
	}
}

func TestFindAllTypes(t *testing.T) {
	types := []PokemonType{
		{DbSymbol: "1"},
		{DbSymbol: "2"},
		{DbSymbol: "3"},
	}
	store := NewStore()

	for _, pokemonType := range types {
		store.AddType(pokemonType)
	}

	all := store.FindAllTypes()
	if len(all) != 3 {
		t.Error("Find all length should be 3, has", len(all))
	}
}

func TestFindAllPokemon(t *testing.T) {
	pokemonList := []Pokemon{
		{Id: 1, DbSymbol: "1"},
		{Id: 2, DbSymbol: "2"},
		{Id: 4, DbSymbol: "4"},
	}

	store := NewStore()
	for _, pokemon := range pokemonList {
		store.AddPokemon(pokemon)
	}

	idLessThan3 := func(pkmn Pokemon) bool { return pkmn.Id < 3 }
	result := store.FindAllPokemon(idLessThan3)
	resultLen := len(slices.Collect(result))

	if resultLen != 2 {
		t.Error("Expected result to have length 2, has", resultLen)
	}
}

func TestFindPokemonBySymbol(t *testing.T) {
	pokemonList := []Pokemon{
		{Id: 1, DbSymbol: "1"},
		{Id: 2, DbSymbol: "2"},
		{Id: 4, DbSymbol: "4"},
	}
	store := NewStore()

	for _, pokemon := range pokemonList {
		store.AddPokemon(pokemon)
	}

	notFound := store.FindPokemonBySymbol("3")
	if notFound != nil {
		t.Error("Expect result to be null")
	}

	found := store.FindPokemonBySymbol("4")
	if found == nil {
		t.Error("Expect result not to be null")
	}
	if found.Id != 4 {
		t.Error("Expect result ID to be 4, is", found.Id)
	}
}
