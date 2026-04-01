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

	if pokemonCount != 8 {
		t.Error("Import should have 8 pokemon, has", pokemonCount)
	}

	typesIter := store.FindAllTypes()
	typesCount := len(slices.Collect(typesIter))
	if typesCount != 18 {
		t.Error("Import should have 18 types, has", typesCount)
	}
}

func TestFindTypeBySymbol(t *testing.T) {
	types := []PokemonType{*NewTypeBuilder().DbSymbol("test").Build()}
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
		*NewTypeBuilder().DbSymbol("1").Build(),
		*NewTypeBuilder().DbSymbol("2").Build(),
		*NewTypeBuilder().DbSymbol("3").Build(),
	}
	store := NewStore()

	for _, pokemonType := range types {
		store.AddType(pokemonType)
	}

	allIter := store.FindAllTypes()
	allSlice := slices.Collect(allIter)
	if len(allSlice) != 3 {
		t.Error("Find all length should be 3, has", len(allSlice))
	}
}

func TestFindAllPokemon(t *testing.T) {
	pokemonList := []Pokemon{
		*NewPokemonBuilder().ID(1).DbSymbol("1").Build(),
		*NewPokemonBuilder().ID(2).DbSymbol("2").Build(),
		*NewPokemonBuilder().ID(4).DbSymbol("4").Build(),
	}

	store := NewStore()
	for _, pokemon := range pokemonList {
		store.AddPokemon(pokemon)
	}

	idLessThan3 := func(pkmn Pokemon) bool { return pkmn.ID() < 3 }
	result := store.FindAllPokemon(idLessThan3)
	resultLen := len(slices.Collect(result))

	if resultLen != 2 {
		t.Error("Expected result to have length 2, has", resultLen)
	}
}

func TestFindPokemonBySymbol(t *testing.T) {
	pokemonList := []Pokemon{
		*NewPokemonBuilder().ID(1).DbSymbol("1").Build(),
		*NewPokemonBuilder().ID(2).DbSymbol("2").Build(),
		*NewPokemonBuilder().ID(4).DbSymbol("4").Build(),
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
	if found.ID() != 4 {
		t.Error("Expect result ID to be 4, is", found.ID())
	}
}

func TestFindAllPokemonWithFilters(t *testing.T) {
	pokemonList := []Pokemon{
		*NewPokemonBuilder().ID(1).DbSymbol("pikachu").Build(),
		*NewPokemonBuilder().ID(2).DbSymbol("bulbasaur").Build(),
		*NewPokemonBuilder().ID(3).DbSymbol("charmander").Build(),
	}
	store := NewStore()

	for _, pokemon := range pokemonList {
		store.AddPokemon(pokemon)
	}

	idGreaterThan1 := func(p Pokemon) bool { return p.ID() > 1 }
	result := store.FindAllPokemon(idGreaterThan1)
	resultSlice := slices.Collect(result)

	if len(resultSlice) != 2 {
		t.Error("Expected 2 pokemon after filter, got", len(resultSlice))
	}
}

func TestFindPokemonByName_RealData(t *testing.T) {
	store, err := Load(DataFolder)
	if err != nil {
		t.Fatal("Import should succeed", err)
	}

	// Test direct match (DbSymbol or mapped name)
	p1 := store.FindPokemonByName("abomasnow")
	if p1 == nil {
		t.Error("Should find abomasnow by name (symbol)")
	}

	// Test fallback via CSV (Blizzaroi is French name for Abomasnow in CSV)
	p2 := store.FindPokemonByName("Blizzaroi")
	if p2 == nil {
		t.Error("Should find abomasnow by name (French CSV fallback)")
	} else if p2.DbSymbol() != "abomasnow" {
		t.Errorf("Expected abomasnow, got %s", p2.DbSymbol())
	}

	// Test case insensitivity
	p3 := store.FindPokemonByName("blizzaroi")
	if p3 == nil {
		t.Error("Should find abomasnow by name (lowercase French CSV fallback)")
	}

	// Test with spaces
	p4 := store.FindPokemonByName("  Blizzaroi  ")
	if p4 == nil {
		t.Error("Should find abomasnow by name (French CSV fallback with spaces)")
	}
}
