package studio

import (
	"os"
	"testing"
)

const (
	TypeInvalid = "../../test/test_resources/bug-invalid.json"
	TypeValid   = "../../test/test_resources/valid-data/Studio/types/bug.json"
)

func TestUnmarshalType_Error(t *testing.T) {
	content, err := os.ReadFile(TypeInvalid)
	if err != nil {
		t.Error("Error when reading test file", "file", TypeInvalid)
	}
	_, err = UnmarshalType(content)
	if err == nil {
		t.Error("Unmarshal invalid type should return error")
	}
}

func TestUnmarshalType_Ok(t *testing.T) {
	content, err := os.ReadFile(TypeValid)
	if err != nil {
		t.Error("Error when reading test file", "file", TypeValid)
	}

	_, err = UnmarshalType(content)
	if err != nil {
		t.Error("Unmarshal valid type should not return error", err)
	}
}

func TestTranslateType_Oob(t *testing.T) {
	pokemonType := PokemonType{
		TextId: 5,
	}

	translations := []Translation{
		{"en": "test"},
	}

	TranslateType(&pokemonType, translations)

	if pokemonType.Name["en"] != "" {
		t.Error("Translation for pokemon name should be empty")
	}
}

func TestTranslateType_Ok(t *testing.T) {
	pokemonType := PokemonType{
		TextId: 0,
	}

	translations := []Translation{
		{"en": "test"},
	}

	TranslateType(&pokemonType, translations)

	if pokemonType.Name["en"] != "test" {
		t.Error("Translation for pokemon name should not be empty")
	}

}
