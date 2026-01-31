package studio

import (
	"os"
	"testing"
)

const (
	TypeInvalid = "../../test/test_resources/bug-invalid.json"
	TypeValid   = "../../test/test_resources/valid-data/Studio/types/bug.json"

	PokemonInvalid = "../../test/test_resources/abra-invalid.json"
	PokemonValid   = "../../test/test_resources/valid-data/Studio/pokemon/abra.json"

	AbilityInvalid = "../../test/test_resources/ability-invalid.json"
	AbilityValid   = "../../test/test_resources/valid-data/Studio/abilities/adaptability.json"

	InvalidPath        = "invalid/path"
	TranslationInvalid = "../../test/test_resources/100003-invalid.csv"
	TranslationValid   = "../../test/test_resources/valid-data/Text/Dialogs/100003.csv"
)

func TestUnmarshalTypeDescriptor_Error(t *testing.T) {
	content, err := os.ReadFile(TypeInvalid)
	if err != nil {
		t.Fatal("Error reading test file", "file", TypeInvalid)
	}
	_, err = UnmarshalTypeDescriptor(content)
	if err == nil {
		t.Error("Unmarshal invalid type descriptor should return error")
	}
}

func TestUnmarshalTypeDescriptor_Ok(t *testing.T) {
	content, err := os.ReadFile(TypeValid)
	if err != nil {
		t.Fatal("Error reading test file", "file", TypeValid)
	}
	_, err = UnmarshalTypeDescriptor(content)
	if err != nil {
		t.Error("Unmarshal valid type descriptor should not return error", err)
	}
}

func TestMapPokemonTypeDescriptorToPokemonType_Ok(t *testing.T) {
	desc := &PokemonTypeDescriptor{
		DbSymbol: "test",
		TextId:   0,
		Color:    "#FF0000",
	}

	store := NewStore()
	mapper := NewTypeMapper(store)
	pokemonType := mapper.MapPokemonTypeDescriptorToPokemonType(*desc)

	if pokemonType.DbSymbol != "test" {
		t.Error("DbSymbol should be 'test'")
	}
	if pokemonType.Color != "#FF0000" {
		t.Error("Color should be '#FF0000'")
	}
	if pokemonType.Name != nil {
		t.Error("Translation for type name should be nil when FormTextId removed")
	}
}

func TestUnmarshalPokemonDescriptor_Error(t *testing.T) {
	content, err := os.ReadFile(PokemonInvalid)
	if err != nil {
		t.Fatal("Error reading test file", "file", PokemonInvalid)
	}
	_, err = UnmarshalPokemonDescriptor(content)
	if err == nil {
		t.Error("Unmarshal invalid pokemon descriptor should return error")
	}
}

func TestUnmarshalPokemonDescriptor_Ok(t *testing.T) {
	content, err := os.ReadFile(PokemonValid)
	if err != nil {
		t.Fatal("Error reading test file", "file", PokemonValid)
	}

	descriptor, err := UnmarshalPokemonDescriptor(content)
	if err != nil {
		t.Error("Unmarshal valid pokemon descriptor should not return error")
	}

	form := descriptor.Forms[0]
	if form.Type2 != nil {
		t.Error("Unmarshal undefined type2 should set type to nil, has", *form.Type2)
	}
}

func TestUnmarshalAbilityDescriptor_Ok(t *testing.T) {
	content, err := os.ReadFile(AbilityValid)
	if err != nil {
		t.Fatal("Error reading test file", "file", AbilityValid)
	}

	_, err = UnmarshalAbilityDescriptor(content)
	if err != nil {
		t.Error("Unmarshal valid ability descriptor should not return error")
	}
}

func TestUnmarshalAbilityDescriptor_Error(t *testing.T) {
	content, err := os.ReadFile(AbilityInvalid)
	if err != nil {
		t.Fatal("Error reading test file", "file", AbilityInvalid)
	}

	_, err = UnmarshalAbilityDescriptor(content)
	if err == nil {
		t.Error("Unmarshal invalid ability descriptor should return error")
	}
}

func TestMapTranslation_Valid(t *testing.T) {
	translations := []Translation{
		{"en": "first"},
		{"en": "second"},
	}

	result := MapTranslation(1, translations)
	if result == nil || result["en"] != "second" {
		t.Error("MapTranslation with valid ID should return correct translation")
	}
}

func TestNewPokemonMapper(t *testing.T) {
	store := NewStore()
	mapper := NewPokemonMapper(store)

	if mapper == nil {
		t.Error("NewPokemonMapper should return non-nil mapper")
	}
	if mapper.store != store {
		t.Error("NewPokemonMapper should set store correctly")
	}
}

func TestNewTypeMapper(t *testing.T) {
	store := NewStore()
	mapper := NewTypeMapper(store)

	if mapper == nil {
		t.Error("NewTypeMapper should return non-nil mapper")
	}
	if mapper.store != store {
		t.Error("NewTypeMapper should set store correctly")
	}
}

func TestNewAbilityMapper(t *testing.T) {
	store := NewStore()
	mapper := NewAbilityMapper(store)

	if mapper == nil {
		t.Error("NewAbilityMapper should return non-nil mapper")
	}
	if mapper.store != store {
		t.Error("NewAbilityMapper should set store correctly")
	}
}

func TestImportTranslations_FileNotFound(t *testing.T) {
	_, err := ImportTranslations(InvalidPath)
	if err == nil {
		t.Error("Import translation at invalid path should return error")
	}
}

func TestImportTranslations_InvalidFormat(t *testing.T) {
	_, err := ImportTranslations(TranslationInvalid)
	if err == nil {
		t.Error("Import translation with invalid format should return error")
	}
}

func TestImportTranslations_Ok(t *testing.T) {
	_, err := ImportTranslations(TranslationValid)
	if err != nil {
		t.Error("Import translation with valid file should not return error")
	}
}
