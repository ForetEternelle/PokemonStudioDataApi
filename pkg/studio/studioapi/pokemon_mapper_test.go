package studioapi

import (
	"testing"

	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/studio"
)

func TestPokemonToThumbnail(t *testing.T) {
	lang := "test"
	normalType := studio.NewTypeBuilder().DbSymbol("normal").Build()
	form := studio.NewPokemonFormBuilder().
		Form(0).
		Type1(normalType).
		Name(studio.Translation{lang: "testName"}).
		Build()
	pokemon := studio.NewPokemonBuilder().
		ID(1).
		DbSymbol("test").
		Forms(map[int32]studio.PokemonForm{0: *form}).
		Build()

	typeMapper := NewTypeMapper()
	abilityMapper := NewAbilityMapper()
	store := studio.NewStore()
	pokemonMapper := NewPokemonMapper(typeMapper, abilityMapper, store)

	policy := NewAccessPolicy()
	thumbnail := pokemonMapper.PokemonToThumbnail(*pokemon, lang, policy)

	if thumbnail.Image != pokemon.DbSymbol() {
		t.Error("Mapper should map image, expected", pokemon.DbSymbol(), ", has", thumbnail.Image)
	}
	if thumbnail.Symbol != pokemon.DbSymbol() {
		t.Error("Mapper should map db symbol, expected", pokemon.DbSymbol(), ", has", thumbnail.Symbol)
	}
	if thumbnail.Number != pokemon.ID() {
		t.Error("Mapper should map Id, expected", pokemon.ID(), ", has", thumbnail.Number)
	}
	if thumbnail.Name != form.Name(lang) {
		t.Error("Mapper should map name, expected", form.Name(lang), ", has", thumbnail.Name)
	}
}

func TestPokemonToDetail(t *testing.T) {
	lang := "test"
	normalType := studio.NewTypeBuilder().DbSymbol("normal").Build()

	form := studio.NewPokemonFormBuilder().
		Form(0).
		Type1(normalType).
		BaseHp(100).
		BaseAtk(50).
		Name(studio.Translation{lang: "testName"}).
		Description(studio.Translation{lang: "testDesc"}).
		Build()
	pokemon := studio.NewPokemonBuilder().
		ID(1).
		DbSymbol("test").
		Forms(map[int32]studio.PokemonForm{0: *form}).
		Build()

	typeMapper := NewTypeMapper()
	abilityMapper := NewAbilityMapper()
	store := studio.NewStore()
	store.AddType(*normalType)
	pokemonMapper := NewPokemonMapper(typeMapper, abilityMapper, store)

	policy := NewAccessPolicy()
	detail := pokemonMapper.PokemonToDetail(*pokemon, lang, policy)

	if detail.Symbol != pokemon.DbSymbol() {
		t.Error("Mapper should map symbol")
	}
	if detail.Number != pokemon.ID() {
		t.Error("Mapper should map number")
	}
	if detail.MainForm.Form == nil {
		t.Error("Mapper should map main form")
	}
	if detail.MainForm.Name != form.Name(lang) {
		t.Error("Mapper should map name")
	}
	if detail.MainForm.Description != form.Description(lang) {
		t.Error("Mapper should map description")
	}
}
