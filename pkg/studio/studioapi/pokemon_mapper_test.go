package studioapi_test

import (
	"testing"

	"github.com/ForetEternelle/ForetEternelleDataApi/pkg/studio"
	"github.com/ForetEternelle/ForetEternelleDataApi/pkg/studio/studioapi"
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
	store , _:= studio.NewStore([]studio.Pokemon{}, []studio.PokemonType{})
	pokemonMapper := studioapi.NewPokemonMapper(typeMapper, store)

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
