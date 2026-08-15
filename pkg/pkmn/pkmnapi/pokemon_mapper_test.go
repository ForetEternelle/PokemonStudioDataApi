package pkmnapi

import (
	"testing"

	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/pkmn"
)

func TestPokemonToThumbnail(t *testing.T) {
	lang := "test"
	normalType := &pkmn.PokemonType{DbSymbol: "normal"}
	form := &pkmn.PokemonForm{
		Form:  0,
		Type1: normalType,
		Name:  pkmn.Translation{lang: "testName"},
	}
	pokemon := &pkmn.Pokemon{
		ID:       1,
		DbSymbol: "test",
		Forms:    []*pkmn.PokemonForm{form},
	}

	typeMapper := NewTypeMapper()
	abilityMapper := NewAbilityMapper()
	store := pkmn.NewStore()
	pokemonMapper := NewPokemonMapper(typeMapper, abilityMapper, store)

	policy := NewPokemonFilterPolicy()
	thumbnail := pokemonMapper.PokemonToThumbnail(*pokemon, 0, lang,
		*policy)

	if thumbnail.Image != pokemon.DbSymbol {
		t.Error("Mapper should map image, expected", pokemon.DbSymbol, ", has", thumbnail.Image)
	}
	if thumbnail.Symbol != pokemon.DbSymbol {
		t.Error("Mapper should map db symbol, expected", pokemon.DbSymbol, ", has", thumbnail.Symbol)
	}
	if thumbnail.Number != pokemon.ID {
		t.Error("Mapper should map Id, expected", pokemon.ID, ", has", thumbnail.Number)
	}
	if thumbnail.Name != form.Name[lang] {
		t.Error("Mapper should map name, expected", form.Name[lang], ", has", thumbnail.Name)
	}
}

func TestPokemonToDetail(t *testing.T) {
	lang := "test"
	normalType := &pkmn.PokemonType{DbSymbol: "normal"}

	form := &pkmn.PokemonForm{
		Form:        1,
		Type1:       normalType,
		BaseHp:      100,
		BaseAtk:     50,
		Name:        pkmn.Translation{lang: "testName"},
		Description: pkmn.Translation{lang: "testDesc"},
	}
	pokemon := &pkmn.Pokemon{
		ID:       1,
		DbSymbol: "test",
		Forms:    []*pkmn.PokemonForm{form},
	}

	typeMapper := NewTypeMapper()
	abilityMapper := NewAbilityMapper()
	store := pkmn.NewStore()
	store.AddType(pkmn.AddTypeDto{DbSymbol: "normal"})
	pokemonMapper := NewPokemonMapper(typeMapper, abilityMapper, store)

	policy := NewPokemonFilterPolicy()
	detail := pokemonMapper.PokemonToDetail(*pokemon, 1, lang, *policy)

	if detail.Symbol != pokemon.DbSymbol {
		t.Error("Mapper should map symbol")
	}
	if detail.Number != pokemon.ID {
		t.Error("Mapper should map number")
	}
	if detail.Form.Number != 1 {
		t.Error("Mapper should map given form")
	}
	if detail.Form.Name != form.Name[lang] {
		t.Error("Mapper should map name")
	}
	if detail.Form.Description != form.Description[lang] {
		t.Error("Mapper should map description")
	}
}
