package pkmn

import (
	"slices"
	"testing"
)

const DataFolder = "../../../test/test_resources/valid-data"

func TestFindTypeBySymbol(t *testing.T) {
	store := NewStore()
	store.AddType(	newTestType(t, nil))

	found := store.FindTypeBySymbol("test")
	if found == nil {
		t.Error("Should find type with symbol test")
	}
}

func TestFindAllTypes(t *testing.T) {
	store := NewStore()
	store.AddType(	newTestType(t, nil))
	store.AddType(	newTestType(t, nil))
	store.AddType(	newTestType(t, nil))

	allIter := store.FindAllTypes()
	allSlice := slices.Collect(allIter)
	if len(allSlice) != 3 {
		t.Error("Find all length should be 3, has", len(allSlice))
	}
}

func TestFindAllPokemon(t *testing.T) {
	store := NewStore()
	store.AddPokemon(	newTestPokemon(t, 1, "1"))
	store.AddPokemon(	newTestPokemon(t, 2, "2"))
	store.AddPokemon(	newTestPokemon(t, 4, "4"))

	idLessThan3 := func(pkmn *Pokemon) bool { return pkmn.ID() < 3 }
	result := store.FindAllPokemon(idLessThan3)
	resultLen := len(slices.Collect(result))

	if resultLen != 2 {
		t.Error("Expected result to have length 2, has", resultLen)
	}
}

func TestFindPokemonBySymbol(t *testing.T) {
	store := NewStore()
	store.AddPokemon(	newTestPokemon(t, 1, "1"))
	store.AddPokemon(	newTestPokemon(t, 2, "2"))
	store.AddPokemon(	newTestPokemon(t, 4, "4"))

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
	store := NewStore()
	store.AddPokemon(	newTestPokemon(t, 1, "pikachu"))
	store.AddPokemon(	newTestPokemon(t, 2, "bulbasaur"))
	store.AddPokemon(	newTestPokemon(t, 3, "charmander"))

	idGreaterThan1 := func(p *Pokemon) bool { return p.ID() > 1 }
	result := store.FindAllPokemon(idGreaterThan1)
	resultSlice := slices.Collect(result)

	if len(resultSlice) != 2 {
		t.Error("Expected 2 pokemon after filter, got", len(resultSlice))
	}
}

func TestFindPokemonByName_RealData(t *testing.T) {
	store := NewStore()

	form, err := NewPokemonForm(PokemonFormConfig{Form: 0, Name: Translation{"en": "Abomasnow", "fr": "Blizzaroi"}})
	if err != nil {
		t.Fatal(err)
	}
	pokemon, err := NewPokemon(PokemonConfig{
		ID:       460,
		DbSymbol: "abomasnow",
		Forms:    []*PokemonForm{form},
	})
	if err != nil {
		t.Fatal(err)
	}
	store.AddPokemon(pokemon)

	p1 := store.FindPokemonByName("abomasnow")
	if p1 == nil {
		t.Error("Should find abomasnow by name (symbol)")
	}

	p2 := store.FindPokemonByName("Blizzaroi")
	if p2 == nil {
		t.Error("Should find abomasnow by name (French CSV fallback)")
	} else if p2.DbSymbol() != "abomasnow" {
		t.Errorf("Expected abomasnow, got %s", p2.DbSymbol())
	}

	p3 := store.FindPokemonByName("blizzaroi")
	if p3 == nil {
		t.Error("Should find abomasnow by name (lowercase French CSV fallback)")
	}

	p4 := store.FindPokemonByName("  Blizzaroi  ")
	if p4 == nil {
		t.Error("Should find abomasnow by name (French CSV fallback with spaces)")
	}
}
