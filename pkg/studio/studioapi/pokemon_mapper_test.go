package studioapi_test

import (
	"testing"

	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/studio"
	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/studio/studioapi"
)

func TestPokemonToThumbnail(t *testing.T) {
	lang := "test"
	pokemon := studio.Pokemon{
		Id:       1,
		DbSymbol: "test",
		Forms: []studio.PokemonForm{
			{
				Name: studio.Translation{lang: "testName"},
			},
		},
	}

	typeMapper := studioapi.NewTypeMapper()
	abilityMapper := studioapi.NewAbilityMapper()
	store := studio.NewStore()
	pokemonMapper := studioapi.NewPokemonMapper(typeMapper, abilityMapper, store)

	thumbnail := pokemonMapper.PokemonToThumbnail(pokemon, lang)

	if thumbnail.Image != pokemon.DbSymbol {
		t.Error("Mapper should map image, expected", pokemon.DbSymbol, ", has", thumbnail.Image)
	}

	if thumbnail.Name != pokemon.Forms[0].Name[lang] {
		t.Error("Mapper should map name, expected", pokemon.Forms[0].Name[lang], ", has", thumbnail.Name)
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
		Id:       1,
		DbSymbol: "test",
		Forms: []studio.PokemonForm{
			{
				Form:        0,
				Name:        studio.Translation{lang: "testName"},
				Description: studio.Translation{lang: "testDesc"},
				Type1:       &normalType,
				BaseHp:      100,
				BaseAtk:     50,
			},
		},
	}

	typeMapper := studioapi.NewTypeMapper()
	abilityMapper := studioapi.NewAbilityMapper()
	store := studio.NewStore()
	store.AddType(normalType)
	pokemonMapper := studioapi.NewPokemonMapper(typeMapper, abilityMapper, store)

	detail := pokemonMapper.PokemonToDetail(pokemon, lang)

	if detail.Symbol != pokemon.DbSymbol {
		t.Error("Mapper should map symbol")
	}
	if detail.Number != pokemon.Id {
		t.Error("Mapper should map number")
	}
	if detail.MainForm.Form == nil {
		t.Error("Mapper should map main form")
	}
}
