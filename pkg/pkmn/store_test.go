package pkmn

import (
	"slices"
	"testing"
)

const DataFolder = "../../../test/test_resources/valid-data"

func TestFindTypeBySymbol(t *testing.T) {
	store := NewStore()
	store.AddType(newTestTypeDto(nil))

	found := store.FindTypeBySymbol("test")
	if found == nil {
		t.Error("Should find type with symbol test")
	}
}

func TestFindAllTypes(t *testing.T) {
	store := NewStore()
	store.AddType(newTestTypeDto(nil))
	store.AddType(newTestTypeDto(nil))
	store.AddType(newTestTypeDto(nil))

	allIter := store.FindAllTypes()
	allSlice := slices.Collect(allIter)
	if len(allSlice) != 3 {
		t.Error("Find all length should be 3, has", len(allSlice))
	}
}

func TestFindAllPokemon(t *testing.T) {
	store := NewStore()
	store.AddPokemon(newTestPokemonDto(1, "1"))
	store.AddPokemon(newTestPokemonDto(2, "2"))
	store.AddPokemon(newTestPokemonDto(4, "4"))

	idLessThan3 := func(pkmn *Pokemon) bool { return pkmn.ID < 3 }
	result := store.FindAllPokemon(idLessThan3)
	resultLen := len(slices.Collect(result))

	if resultLen != 2 {
		t.Error("Expected result to have length 2, has", resultLen)
	}
}

func TestFindPokemonBySymbol(t *testing.T) {
	store := NewStore()
	store.AddPokemon(newTestPokemonDto(1, "1"))
	store.AddPokemon(newTestPokemonDto(2, "2"))
	store.AddPokemon(newTestPokemonDto(4, "4"))

	notFound := store.FindPokemonBySymbol("3")
	if notFound != nil {
		t.Error("Expect result to be null")
	}

	found := store.FindPokemonBySymbol("4")
	if found == nil {
		t.Error("Expect result not to be null")
	}
	if found.ID != 4 {
		t.Error("Expect result ID to be 4, is", found.ID)
	}
}

func TestFindAllPokemonWithFilters(t *testing.T) {
	store := NewStore()
	store.AddPokemon(newTestPokemonDto(1, "pikachu"))
	store.AddPokemon(newTestPokemonDto(2, "bulbasaur"))
	store.AddPokemon(newTestPokemonDto(3, "charmander"))

	idGreaterThan1 := func(p *Pokemon) bool { return p.ID > 1 }
	result := store.FindAllPokemon(idGreaterThan1)
	resultSlice := slices.Collect(result)

	if len(resultSlice) != 2 {
		t.Error("Expected 2 pokemon after filter, got", len(resultSlice))
	}
}

func TestFindPokemonByName_RealData(t *testing.T) {
	store := NewStore()

	store.AddPokemon(AddPokemonDto{
		ID:       460,
		DbSymbol: "abomasnow",
		Forms: []AddPokemonFormDto{
			{Form: 0, Name: Translation{"en": "Abomasnow", "fr": "Blizzaroi"}},
		},
	})

	p1 := store.FindPokemonByName("abomasnow", "fr")
	if p1 == nil {
		t.Error("Should find abomasnow by name (symbol)")
	}

	p2 := store.FindPokemonByName("Blizzaroi", "fr")
	if p2 == nil {
		t.Error("Should find abomasnow by name (French CSV fallback)")
	} else if p2.DbSymbol != "abomasnow" {
		t.Errorf("Expected abomasnow, got %s", p2.DbSymbol)
	}

	p3 := store.FindPokemonByName("blizzaroi", "fr")
	if p3 == nil {
		t.Error("Should find abomasnow by name (lowercase French CSV fallback)")
	}

	p4 := store.FindPokemonByName("  Blizzaroi  ", "fr")
	if p4 == nil {
		t.Error("Should find abomasnow by name (French CSV fallback with spaces)")
	}
}

// Pokemon are imported in file (alphabetical) order, not dex order, so name
// lookups must not depend on insertion order.
func TestFindPokemonByName_InsertionOrderMismatch(t *testing.T) {
	store := NewStore()
	store.AddPokemon(AddPokemonDto{ID: 63, DbSymbol: "abra", Forms: []AddPokemonFormDto{
		{Form: 0, Name: Translation{"en": "Abra", "fr": "Abra"}},
	}})
	store.AddPokemon(AddPokemonDto{ID: 59, DbSymbol: "arcanine", Forms: []AddPokemonFormDto{
		{Form: 0, Name: Translation{"en": "Arcanine", "fr": "Arcanin"}},
	}})
	store.AddPokemon(AddPokemonDto{ID: 3, DbSymbol: "venusaur", Forms: []AddPokemonFormDto{
		{Form: 0, Name: Translation{"en": "Venusaur", "fr": "Florizarre"}},
	}})

	for name, expected := range map[string]string{"Abra": "abra", "arcanin": "arcanine", "Florizarre": "venusaur"} {
		found := store.FindPokemonByName(name, "fr")
		if found == nil || found.DbSymbol != expected {
			t.Errorf("FindPokemonByName(%q, fr): expected %s, got %v", name, expected, found)
		}
	}
}

// A name matching in the requested language wins over a name matching in another language.
func TestFindPokemonByName_PrefersRequestedLang(t *testing.T) {
	store := NewStore()
	store.AddPokemon(AddPokemonDto{ID: 1, DbSymbol: "first", Forms: []AddPokemonFormDto{
		{Form: 0, Name: Translation{"en": "Same", "fr": "Premier"}},
	}})
	store.AddPokemon(AddPokemonDto{ID: 2, DbSymbol: "second", Forms: []AddPokemonFormDto{
		{Form: 0, Name: Translation{"en": "Second", "fr": "Same"}},
	}})

	if found := store.FindPokemonByName("same", "fr"); found == nil || found.DbSymbol != "second" {
		t.Errorf("lang fr: expected second, got %v", found)
	}
	if found := store.FindPokemonByName("same", "en"); found == nil || found.DbSymbol != "first" {
		t.Errorf("lang en: expected first, got %v", found)
	}
	// Falls back to other languages when nothing matches in the requested one
	if found := store.FindPokemonByName("premier", "en"); found == nil || found.DbSymbol != "first" {
		t.Errorf("fallback: expected first, got %v", found)
	}
	if found := store.FindPokemonByName("unknown", "fr"); found != nil {
		t.Errorf("expected nil, got %v", found.DbSymbol)
	}
}

// A main form name in the requested language wins over another pokemon's alternate form name.
func TestFindPokemonByName_PrefersMainForm(t *testing.T) {
	store := NewStore()
	store.AddPokemon(AddPokemonDto{ID: 1, DbSymbol: "base", Forms: []AddPokemonFormDto{
		{Form: 0, Name: Translation{"fr": "Base"}},
		{Form: 1, Name: Translation{"fr": "Cible"}},
	}})
	store.AddPokemon(AddPokemonDto{ID: 2, DbSymbol: "target", Forms: []AddPokemonFormDto{
		{Form: 0, Name: Translation{"fr": "Cible"}},
	}})

	if found := store.FindPokemonByName("cible", "fr"); found == nil || found.DbSymbol != "target" {
		t.Errorf("expected target, got %v", found)
	}
}
