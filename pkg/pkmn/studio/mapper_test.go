package studio

import (
	"testing"

	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/pkmn"
)

func TestNewPokemonMapper(t *testing.T) {
	store := pkmn.NewStore()
	mapper := NewPokemonMapper(store)

	if mapper == nil {
		t.Error("NewPokemonMapper should return non-nil mapper")
	}
	if mapper.store != store {
		t.Error("NewPokemonMapper should set store correctly")
	}
}

func TestNewTypeMapper(t *testing.T) {
	store := pkmn.NewStore()
	mapper := NewTypeMapper(store)

	if mapper == nil {
		t.Error("NewTypeMapper should return non-nil mapper")
	}
	if mapper.store != store {
		t.Error("NewTypeMapper should set store correctly")
	}
}

func TestNewAbilityMapper(t *testing.T) {
	store := pkmn.NewStore()
	mapper := NewAbilityMapper(store)

	if mapper == nil {
		t.Error("NewAbilityMapper should return non-nil mapper")
	}
	if mapper.store != store {
		t.Error("NewAbilityMapper should set store correctly")
	}
}

func TestNewMoveMapper(t *testing.T) {
	store := pkmn.NewStore()
	mapper := NewMoveMapper(store)

	if mapper == nil {
		t.Error("NewMoveMapper should return non-nil mapper")
	}
	if mapper.store != store {
		t.Error("NewMoveMapper should set store correctly")
	}
}
