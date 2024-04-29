package studio

import (
	"os"
	"testing"
)

const (
	PokemonInvalid = "../../test/test_resources/abra-invalid.json"
	PokemonValid   = "../../test/test_resources/valid-data/Studio/pokemon/abra.json"
)

func TestUnmarshalPokemon_Error(t *testing.T) {
	content, err := os.ReadFile(PokemonInvalid)
	if err != nil {
		t.Error("Error when reading test file", "file", PokemonInvalid)
	}
	_, err = UnmarshalPokemon(content)
	if err == nil {
		t.Error("Unmarshal invalid pokemon should return error")
	}
}

func TestUnmarshalPokemon_Ok(t *testing.T) {
	content, err := os.ReadFile(PokemonValid)
	if err != nil {
		t.Error("Error when reading test file", "file", PokemonValid)
	}

	pokemon, err := UnmarshalPokemon(content)
	if err != nil {
		t.Error("Unmarshal valid pokemon should not return error")
	}

	form := pokemon.Forms[0]
	if form.Type2 != nil {
		t.Error("Unmarshal undefined type2 should set type to nil, has", *form.Type2)
	}
}

func TestTranslatePokemon_NameOob(t *testing.T) {
	pokemon := Pokemon{
		Forms: []PokemonForm{
			{
				FormTextId: FormTextId{
					Name:        1000,
					Description: 1000,
				},
			},
		},
	}

	form := pokemon.Forms[0]
	translation := []Translation{
		{"en": "test"},
	}

	TranslatePokemon(&pokemon, translation, translation)
	if form.Name["en"] != "" {
		t.Error("Translation for pokemon name should be empty")
	}
	if form.Description["en"] != "" {
		t.Error("Translation for pokemon description should be empty")
	}

}

func TestTranslatePokemon_Ok(t *testing.T) {
	pokemon := Pokemon{
		Forms: []PokemonForm{
			{
				FormTextId: FormTextId{
					Name:        0,
					Description: 0,
				},
			},
		},
	}

	translation := []Translation{
		{"en": "test"},
	}

	TranslatePokemon(&pokemon, translation, translation)
	form := &pokemon.Forms[0]
	if form.Name["en"] != "test" {
		t.Error("Translation for pokemon name should not be empty")
	}
	if form.Description["test"] != "" {
		t.Error("Translation for pokemon description should not be empty")
	}

}
