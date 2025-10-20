package studio

import (
	"encoding/json"
	"log/slog"
	"path"

	"github.com/ForetEternelle/ForetEternelleDataApi/pkg/file"
)

const (
	PokemonFolder                         = "pokemon/"
	PokemonTranslationFileName            = "100067.csv"
	PokemonDescriptionTranslationFileName = "100068.csv"
)

// ImportPokemon import a pokemon folder to a store
// studioFolder the pokemon studio folder
// translationFolder the translation folder
// store the store the import is sending data to
func ImportPokemon(studioFolder, translationFolder string) ([]Pokemon, error) {
	slog.Info("Importing pokemon name translation")
	pokemonNameFilePath := path.Join(translationFolder, PokemonTranslationFileName)
	pokemonNameTranslations, err := ImportTranslations(pokemonNameFilePath)
	if err != nil {
		return nil, err
	}

	slog.Info("Importing pokemon description translation")
	pokemonDescriptionFilePath := path.Join(translationFolder, PokemonDescriptionTranslationFileName)
	pokemonDescriptionTranslations, err := ImportTranslations(pokemonDescriptionFilePath)
	if err != nil {
		return nil, err
	}

	pokemonFolderPath := path.Join(studioFolder, PokemonFolder)
	slog.Info("Importing pokemon folder", "path", pokemonFolderPath)
	pokemonFileIterator, err := file.ImportFolder(pokemonFolderPath)
	if err != nil {
		return nil, err
	}

	pokemonList := make([]Pokemon, 0)
	for pokemonFile := range pokemonFileIterator {
		p, err := UnmarshalPokemon(pokemonFile.Content)
		if err != nil {
			slog.Warn("Failed to unmarshal pokemon content", "file", pokemonFile.Path)
			continue
		}
		TranslatePokemon(p, pokemonNameTranslations, pokemonDescriptionTranslations)
		pokemonList = append(pokemonList, *p)
	}
	
	return pokemonList, nil
}

// UnmarshalPokemon unmarshal a json encoded pokemon to an object
// pokemonContent the encoded pokemon
func UnmarshalPokemon(pokemonContent []byte) (*Pokemon, error) {
	pokemon := &Pokemon{}
	if err := json.Unmarshal(pokemonContent, pokemon); err != nil {
		return nil, err
	}
	for i := range pokemon.Forms {
		form := &pokemon.Forms[i]
		if form.Type2 != nil && *form.Type2 == UndefType {
			form.Type2 = nil
		}
	}
	return pokemon, nil
}

// TranslatePokemon add a translation to a pokemon name and description
// pokemon the pokemon to add translation to
// pokemonNameTranslations the datastucture containing all pokemon names translations
// pokemonDescriptionTranslations the datastructure containing all pokemon description translations
func TranslatePokemon(pokemon *Pokemon, pokemonNameTranslations, pokemonDescriptionTranslations []Translation) {
	nameTranslationSize := len(pokemonNameTranslations)
	descriptionTranslationSize := len(pokemonDescriptionTranslations)

	for i := range pokemon.Forms {
		form := &pokemon.Forms[i]
		if form.FormTextId.Name < nameTranslationSize {
			form.Name = pokemonNameTranslations[form.FormTextId.Name]
		} else {
			slog.Warn("Could not find translation for pokemon ", "symbol", pokemon.DbSymbol, "form", form.Form, "TextID", form.FormTextId.Name)
		}

		if form.FormTextId.Description < descriptionTranslationSize {
			form.Description = pokemonDescriptionTranslations[form.FormTextId.Description]
		} else {
			slog.Warn("Could not find translation for pokemon description ", "symbol", pokemon.DbSymbol, "form", form.Form, "TextID", form.FormTextId.Description)
		}
	}

}
