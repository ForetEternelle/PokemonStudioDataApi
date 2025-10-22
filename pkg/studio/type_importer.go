package studio

import (
	"encoding/json"
	"log/slog"
	"path"

	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/file"
)

const (
	TypeFolder             = "types/"
	TypeTranslationFilName = "100003.csv"
)

// ImportTypes import a type folder to a store
// studioFolder the pokemon studio folder
// translationFolder the translation folder
// store the store the import is sending data to
func ImportTypes(studioFolder, translationFolder string) ([]PokemonType, error) {
	slog.Info("Importing type name translation")
	typeNameFilePath := path.Join(translationFolder, TypeTranslationFilName)
	typeNameTranslations, err := ImportTranslations(typeNameFilePath)
	if err != nil {
		return nil, err
	}

	typeFolderPath := path.Join(studioFolder, TypeFolder)
	slog.Info("Importing type folder", "path", typeFolderPath)
	typeFileIterator, err := file.ImportFolder(typeFolderPath)
	if err != nil {
		return nil, err
	}

	types := make([]PokemonType, 0)
	for typeFile := range typeFileIterator {
		pokemonType, err := UnmarshalType(typeFile.Content)
		if err != nil {
			slog.Warn("Failed to unmarshal type content", "file", typeFile.Path)
			continue
		}
		TranslateType(pokemonType, typeNameTranslations)
		types = append(types, *pokemonType)
	}
	return types, nil
}

// UnmarshalType unmarshal a json encoded type to an object
// typeContent the encoded type
func UnmarshalType(typeContent []byte) (*PokemonType, error) {
	pokemonType := &PokemonType{}
	if err := json.Unmarshal(typeContent, pokemonType); err != nil {
		return nil, err
	}
	return pokemonType, nil
}

// TranslateType add a translation to a type name
// pokemonType the type to add translation to
// typeNameTranslations the datastructure containing all type names translations
func TranslateType(pokemonType *PokemonType, typeNameTranslations []Translation) {
	if pokemonType.TextId < len(typeNameTranslations) {
		pokemonType.Name = typeNameTranslations[pokemonType.TextId]
	} else {
		slog.Warn("Could not find translation for type ", "symbol", pokemonType.DbSymbol, "TextID", pokemonType.TextId)
	}
}
