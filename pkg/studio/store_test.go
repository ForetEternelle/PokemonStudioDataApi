package studio

import (
	"slices"
	"testing"
)

const (
	DataFolder = "../../test/test_resources/valid-data"
)

func TestLoad(t *testing.T) {
	store, err := Load(DataFolder)
	if err != nil {
		t.Error("Import should suceed", err)
	}

	pokemonIter := store.FindAllPokemon()
	nbPokemon := len(slices.Collect(pokemonIter))

	if nbPokemon != 5 {
		t.Error("Import should have 5 pokemon", "has", nbPokemon)
	}

	types := store.FindAllTypes()

	nbTypes := len(types)
	if nbTypes != 18 {
		t.Error("Import should have 18 types", "has", nbTypes)
	}
}

func TestFindTypeBySymbol(t *testing.T) {
	symbol := "test"
	types := []PokemonType{
		{
			DbSymbol: "test",
		},
	}
	abilities := []Ability{}
	store, _ := NewStore([]Pokemon{}, types, abilities)

	find := store.FindTypeBySymbol(symbol)
	if find == nil {
		t.Error("Should find type with symbol", symbol)
	}
}

func TestFindAllTypes(t *testing.T) {
	types := []PokemonType{
		{
			DbSymbol: "1",
		},
		{
			DbSymbol: "2",
		},
		{
			DbSymbol: "3",
		},
	}
	abilities := []Ability{}
	store, _ := NewStore([]Pokemon{}, types, abilities)
	all := store.FindAllTypes()
	allLen := len(all)
	if allLen != 3 {
		t.Error("Find all length should be 3, has", allLen)
	}
}
func TestFindAllPokemon(t *testing.T) {
	pokemonList := []Pokemon{{
		Id:       1,
		DbSymbol: "1",
	}, {
		Id:       2,
		DbSymbol: "2",
	}, {
		Id:       4,
		DbSymbol: "4",
	}}

	store, _ := NewStore(pokemonList, []PokemonType{}, []Ability{})

	idLessThan3 := func(pkmn Pokemon) bool {
		return pkmn.Id < 3
	}
	result := store.FindAllPokemon(idLessThan3)

	expectLen := 2
	resultLen := len(slices.Collect(result))
	if expectLen != resultLen {
		t.Error("Expected result to have length", expectLen, ", has", resultLen)
	}
}

func TestFindPokemonBySymbol(t *testing.T) {
	pokemonList := []Pokemon{
		{
			Id:       1,
			DbSymbol: "1",
		}, {
			Id:       2,
			DbSymbol: "2",
		}, {
			Id:       4,
			DbSymbol: "4",
		},
	}
	store, _ := NewStore(pokemonList, []PokemonType{}, []Ability{})
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
