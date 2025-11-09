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
				Name: studio.Translation{
					lang: "testName",
				},
				Resources: studio.Resources{
					Front: "testFrontImage",
				},
			},
		},
	}

	typeMapper := studioapi.NewTypeMapper()
	abilityMapper := studioapi.NewAbilityMapper()
	store, _ := studio.NewStore([]studio.Pokemon{}, []studio.PokemonType{}, []studio.Ability{})
	pokemonMapper := studioapi.NewPokemonMapper(typeMapper, abilityMapper, store)

	thumbnail := pokemonMapper.PokemonToThumbnail(pokemon, lang)

	if thumbnail.Image != pokemon.Forms[0].Resources.Front {
		t.Error("Mapper should map image, expected", pokemon.Forms[0].Resources.Front, ", has", thumbnail.Image)
	}

	if thumbnail.Name != pokemon.Forms[0].Name[lang] {
		t.Error("Mapper should map name, expected", pokemon.Forms[0].Name, ", has", thumbnail.Name)
	}

	if thumbnail.Symbol != pokemon.DbSymbol {
		t.Error("Mapper should map db symbol, expected", pokemon.DbSymbol, ", has", thumbnail.Symbol)
	}
	if thumbnail.Number != pokemon.Id {
		t.Error("Mapper should map Id, expected", pokemon.Id, ", has", thumbnail.Number)
	}
}
