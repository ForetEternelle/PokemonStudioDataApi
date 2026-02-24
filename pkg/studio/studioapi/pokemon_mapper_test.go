package studioapi_test

import (
	"testing"

	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/studio"
	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/studio/studioapi"
)

func TestPokemonToThumbnail(t *testing.T) {
	lang := "test"
	normalType := studio.PokemonType{DbSymbol: "normal"}
	pokemon := studio.Pokemon{
		Id:       1,
		DbSymbol: "test",
		Name:     studio.Translation{lang: "testName"},
		Forms: map[int32]studio.PokemonForm{
			0: {
				Type1: &normalType,
			},
		},
	}

	typeMapper := studioapi.NewTypeMapper()
	abilityMapper := studioapi.NewAbilityMapper()
	store := studio.NewStore()
	pokemonMapper := studioapi.NewPokemonMapper(typeMapper, abilityMapper, store)

	policy := studioapi.NewAccessPolicy()
	thumbnail := pokemonMapper.PokemonToThumbnail(pokemon, lang, policy)

	if thumbnail.Image != pokemon.DbSymbol {
		t.Error("Mapper should map image, expected", pokemon.DbSymbol, ", has", thumbnail.Image)
	}

	if thumbnail.Name != pokemon.Name[lang] {
		t.Error("Mapper should map name, expected", pokemon.Name[lang], ", has", thumbnail.Name)
	}

	if thumbnail.Symbol != pokemon.DbSymbol {
		t.Error("Mapper should map db symbol, expected", pokemon.DbSymbol, ", has", thumbnail.Symbol)
	}
	if thumbnail.Number != pokemon.Id {
		t.Error("Mapper should map Id, expected", pokemon.Id, ", has", thumbnail.Number)
	}
}

func TestPokemonToDetail(t *testing.T) {
	lang := "test"
	normalType := studio.PokemonType{DbSymbol: "normal"}

	pokemon := studio.Pokemon{
		Id:          1,
		DbSymbol:    "test",
		Name:        studio.Translation{lang: "testName"},
		Description: studio.Translation{lang: "testDesc"},
		Forms: map[int32]studio.PokemonForm{
			0: {
				Form:    0,
				Type1:   &normalType,
				BaseHp:  100,
				BaseAtk: 50,
			},
		},
	}

	typeMapper := studioapi.NewTypeMapper()
	abilityMapper := studioapi.NewAbilityMapper()
	store := studio.NewStore()
	store.AddType(normalType)
	pokemonMapper := studioapi.NewPokemonMapper(typeMapper, abilityMapper, store)

	policy := studioapi.NewAccessPolicy()
	detail := pokemonMapper.PokemonToDetail(pokemon, lang, policy)

	if detail.Symbol != pokemon.DbSymbol {
		t.Error("Mapper should map symbol")
	}
	if detail.Number != pokemon.Id {
		t.Error("Mapper should map number")
	}
	if detail.Name != pokemon.Name[lang] {
		t.Error("Mapper should map name")
	}
	if detail.Description != pokemon.Description[lang] {
		t.Error("Mapper should map description")
	}
	if detail.MainForm.Form == nil {
		t.Error("Mapper should map main form")
	}
}
