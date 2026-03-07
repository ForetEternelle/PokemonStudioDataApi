package studio

import (
	"testing"
)

func TestComparePokemonId(t *testing.T) {
	p1 := NewPokemon(WithID(1))
	p2 := NewPokemon(WithID(2))
	p3 := NewPokemon(WithID(1))

	if ComparePokemonId(p1, p2) != -1 {
		t.Error("ComparePokemonId with p1:", p1.ID(), "and p2:", p2.ID(), "should return -1")
	}
	if ComparePokemonId(p2, p1) != 1 {
		t.Error("ComparePokemonId with p2:", p2.ID(), "and p1:", p1.ID(), "should return 1")
	}
	if ComparePokemonId(p1, p3) != 1 {
		t.Error("ComparePokemonId with p1:", p1.ID(), "and p3:", p3.ID(), "should return 1")
	}
}

func TestNewPokemonMapper(t *testing.T) {
	store := NewStore()
	mapper := NewPokemonMapper(store)

	if mapper == nil {
		t.Error("NewPokemonMapper should return non-nil mapper")
	}
	if mapper.store != store {
		t.Error("NewPokemonMapper should set store correctly")
	}
}

func TestNewTypeMapper(t *testing.T) {
	store := NewStore()
	mapper := NewTypeMapper(store)

	if mapper == nil {
		t.Error("NewTypeMapper should return non-nil mapper")
	}
	if mapper.store != store {
		t.Error("NewTypeMapper should set store correctly")
	}
}

func TestNewAbilityMapper(t *testing.T) {
	store := NewStore()
	mapper := NewAbilityMapper(store)

	if mapper == nil {
		t.Error("NewAbilityMapper should return non-nil mapper")
	}
	if mapper.store != store {
		t.Error("NewAbilityMapper should set store correctly")
	}
}

func TestNewMoveMapper(t *testing.T) {
	store := NewStore()
	mapper := NewMoveMapper(store)

	if mapper == nil {
		t.Error("NewMoveMapper should return non-nil mapper")
	}
	if mapper.store != store {
		t.Error("NewMoveMapper should set store correctly")
	}
}
