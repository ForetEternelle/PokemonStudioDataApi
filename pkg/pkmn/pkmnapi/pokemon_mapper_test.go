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
		Category:    pkmn.Translation{lang: "testCategory"},
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
	if detail.Form.Category != form.Category[lang] {
		t.Error("Mapper should map category")
	}
}

func TestFormToPokemonFormDetails_Category(t *testing.T) {
	store := pkmn.NewStore()
	mapper := NewPokemonMapper(NewTypeMapper(), NewAbilityMapper(), store)
	policy := NewPokemonFilterPolicy()
	form := pkmn.PokemonForm{
		Type1:    &pkmn.PokemonType{DbSymbol: "normal"},
		Category: pkmn.Translation{"en": "Mouse Pokémon", "fr": "Pokémon Souris"},
	}

	for _, test := range []struct {
		lang     string
		expected string
	}{
		{"en", "Mouse Pokémon"},
		{"fr", "Pokémon Souris"},
		{"unknown", ""},
	} {
		t.Run(test.lang, func(t *testing.T) {
			details := mapper.FormToPokemonFormDetails(form, test.lang, *policy)
			if details.Category != test.expected {
				t.Errorf("Expected category %q, got %q", test.expected, details.Category)
			}
		})
	}

	t.Run("missing category", func(t *testing.T) {
		form.Category = nil
		details := mapper.FormToPokemonFormDetails(form, "en", *policy)
		if details.Category != "" {
			t.Errorf("Expected empty category, got %q", details.Category)
		}
	})
}
