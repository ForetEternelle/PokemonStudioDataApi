package studio

import (
	"testing"
)

func TestAddAndFindBySymbol(t *testing.T) {
	store := NewTypeStore()
	pokeType := PokemonType{
		DbSymbol: "test",
	}
	store.Add(pokeType)

	find := store.FindBySymbol(pokeType.DbSymbol)
	if find == nil {
		t.Error("Should find type with symbol", pokeType.DbSymbol)
	}
}

func TestFindAll(t *testing.T) {

	store := NewTypeStore()
	store.Add(PokemonType{
		DbSymbol: "1",
	})

	store.Add(PokemonType{
		DbSymbol: "2",
	})

	store.Add(PokemonType{
		DbSymbol: "3",
	})

	all := store.FindAll()
	allLen := len(all)
	if allLen != 3 {
		t.Error("Find all length should be 3, has", allLen)
	}
}