package studioapi

import (
	"testing"

	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/studio"
)

func TestToTypeDetail(t *testing.T) {
	lang := "test"
	pokemonType := studio.NewTypeBuilder().
		DbSymbol("testDbSymbol").
		Color("testColor").
		Name(studio.Translation{lang: "testName"}).
		DamageTo(map[string]float32{
			"defType2": 0.5,
		}).
		Build()

	typeMapper := NewTypeMapper()
	policy := NewAccessPolicy()
	typeDetail := typeMapper.ToTypeDetail(*pokemonType, lang, policy)

	if typeDetail.Name != pokemonType.Name(lang) {
		t.Error("Mapper should map name, expected", pokemonType.Name(lang), ", has", typeDetail.Name)
	}

	if typeDetail.Color != pokemonType.Color() {
		t.Error("Mapper should map color, expected", pokemonType.Color(), ", has", typeDetail.Color)
	}

	if typeDetail.Symbol != pokemonType.DbSymbol() {
		t.Error("Mapper should map db symbol, expected", pokemonType.DbSymbol(), ", has", typeDetail.Symbol)
	}

	def, ok := typeDetail.TypeDamage["defType2"]
	if !ok {
		t.Error("Mapper should map damage to, expected key 'defType2' to be present")
	} else if def != 0.5 {
		t.Error("Mapper should map damage to, expected value 0.5 for key 'defType2', has", def)
	}
}

func TestToTypePartial(t *testing.T) {
	lang := "test"
	pokemonType := studio.NewTypeBuilder().
		DbSymbol("testDbSymbol").
		Color("testColor").
		Name(studio.Translation{lang: "testName"}).
		Build()

	typeMapper := NewTypeMapper()
	policy := NewAccessPolicy()
	typePartial := typeMapper.ToTypePartial(*pokemonType, lang, policy)

	if typePartial.Name != pokemonType.Name(lang) {
		t.Error("Mapper should map name, expected", pokemonType.Name(lang), ", has", typePartial.Name)
	}

	if typePartial.Color != pokemonType.Color() {
		t.Error("Mapper should map color, expected", pokemonType.Color(), ", has", typePartial.Color)
	}

	if typePartial.Symbol != pokemonType.DbSymbol() {
		t.Error("Mapper should map db symbol, expected", pokemonType.DbSymbol(), ", has", typePartial.Symbol)
	}
}
