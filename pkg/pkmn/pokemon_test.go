package pkmn

import (
	"testing"
)

func TestComparePokemonId(t *testing.T) {
	p1 := NewPokemonBuilder().ID(1).Build()
	p2 := NewPokemonBuilder().ID(2).Build()
	p3 := NewPokemonBuilder().ID(1).Build()

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

func TestMaxHp(t *testing.T) {
	expect := int32(386)
	result := MaxHp(91)
	if result != expect {
		t.Error("MaxHp with base stat 91 should return", expect, "but has", result)
	}
}

func TestMinHp(t *testing.T) {
	expect := int32(292)
	result := MinHp(91)
	if result != expect {
		t.Error("MinHp with base stat 91 should return", expect, "but has", result)
	}
}

func TestMaxStat(t *testing.T) {
	expect := int32(403)
	result := MaxStat(134)
	if result != expect {
		t.Error("MaxStat with base stat 91 should return", expect, "but has", result)
	}
}

func TestMinStat(t *testing.T) {
	expect := int32(245)
	result := MinStat(134)
	if result != expect {
		t.Error("MinStat with base stat 91 should return", expect, "but has", result)
	}
}
