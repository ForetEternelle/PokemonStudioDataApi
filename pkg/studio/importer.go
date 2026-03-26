package studio

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"iter"
	"log/slog"
	"os"
	"path"

	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/file"
)

const (
	pokemonNameTranslationFile            = "100000.csv"
	pokemonDescriptionTranslationFile     = "100002.csv"
	pokemonFormNameTranslationFile        = "100067.csv"
	pokemonFormDescriptionTranslationFile = "100068.csv"
	typeNameTranslationFile               = "100003.csv"
	abilityNameTranslationFile            = "100004.csv"
	abilityDescriptionTranslationFile     = "100005.csv"
	moveNameTranslationFile               = "100006.csv"
	moveDescriptionTranslationFile        = "100007.csv"
)

func ImportAbility(studioFolder, translationFolder string) (iter.Seq[*AbilityDescriptor], error) {
	abilityFolderPath := path.Join(studioFolder, "abilities/")
	slog.Info("Importing ability folder", "path", abilityFolderPath)
	abilityFileIterator, err := file.ImportFolder(abilityFolderPath)
	if err != nil {
		return nil, err
	}

	// Import translations
	abilityNametranslationsPath := path.Join(translationFolder, abilityNameTranslationFile)
	abilityTranslations, err := ImportTranslations(abilityNametranslationsPath)
	if err != nil {
		slog.Warn("Failed to import ability translations", "path", abilityNametranslationsPath, "error", err)
		abilityTranslations = []Translation{}
	}

	abilityDescriptiontranslationsPath := path.Join(translationFolder, abilityDescriptionTranslationFile)
	abilityDescTranslations, err := ImportTranslations(abilityDescriptiontranslationsPath)
	if err != nil {
		slog.Warn("Failed to import ability description translations", "path", abilityDescriptiontranslationsPath, "error", err)
		abilityDescTranslations = []Translation{}
	}

	return func(yield func(*AbilityDescriptor) bool) {
		for abilityFile := range abilityFileIterator {
			abilityDesc, err := UnmarshalAbilityDescriptor(abilityFile.Content)
			if err != nil {
				slog.Warn("Failed to unmarshal ability descriptor content", "file", abilityFile.Path, "error", err)
				continue
			}

			abilityDesc.Name = MapTranslation(abilityDesc.TextID, abilityTranslations)
			abilityDesc.Description = MapTranslation(abilityDesc.TextID, abilityDescTranslations)

			if abilityDesc != nil {
				if !yield(abilityDesc) {
					break
				}
			}
		}
	}, nil
}

// ImportPokemon import a pokemon folder to a store
// studioFolder pokemon studio folder
// translationFolder the translation folder
// store the store that import is sending data to
func ImportPokemon(studioFolder, translationFolder string) (iter.Seq[*PokemonDescriptor], error) {
	pokemonFolderPath := path.Join(studioFolder, "pokemon/")
	slog.Info("Importing pokemon folder", "path", pokemonFolderPath)
	pokemonFileIterator, err := file.ImportFolder(pokemonFolderPath)
	if err != nil {
		return nil, err
	}

	pokemonNameTranslationPath := path.Join(translationFolder, pokemonNameTranslationFile)
	pokemonNameTranslations := ImportTranslationsOrEmpty(pokemonNameTranslationPath)

	pokemonDescriptionTranslationPath := path.Join(translationFolder, pokemonDescriptionTranslationFile)
	pokemonDescriptionTranslations := ImportTranslationsOrEmpty(pokemonDescriptionTranslationPath)

	pokemonFormNameTranslationPath := path.Join(translationFolder, pokemonFormNameTranslationFile)
	pokemonFormNameTranslations := ImportTranslationsOrEmpty(pokemonFormNameTranslationPath)

	pokemonFormDescriptionTranslationPath := path.Join(translationFolder, pokemonFormDescriptionTranslationFile)
	pokemonFormDescriptionTranslations := ImportTranslationsOrEmpty(pokemonFormDescriptionTranslationPath)

	return func(yield func(*PokemonDescriptor) bool) {
		for pokemonFile := range pokemonFileIterator {
			pokemonDesc, err := UnmarshalPokemonDescriptor(pokemonFile.Content)
			if err != nil {
				slog.Warn("Failed to unmarshal pokemon descriptor content", "file", pokemonFile.Path, "error", err)
				continue
			}

			for i := range pokemonDesc.Forms {
				form := &pokemonDesc.Forms[i]
				formTextId := form.FormTextId

				if formTextId.Name == 0{
					form.Name = MapTranslation(int(pokemonDesc.ID), pokemonNameTranslations)
				}else{
					form.Name = MapTranslation(formTextId.Name, pokemonFormNameTranslations)
				}

				if formTextId.Description == 0 {
					form.Description = MapTranslation(int(pokemonDesc.ID), pokemonDescriptionTranslations)
				} else {
					form.Description = MapTranslation(formTextId.Description, pokemonFormDescriptionTranslations)
				}
			}

			if pokemonDesc != nil {
				if !yield(pokemonDesc) {
					break
				}
			}
		}
	}, nil
}

// ImportTypes import a type folder to a store
// studioFolder pokemon studio folder
// translationFolder the translation folder
// store the store that import is sending data to
func ImportTypes(studioFolder, translationFolder string) (iter.Seq[*PokemonTypeDescriptor], error) {
	typeFolderPath := path.Join(studioFolder, "types/")
	slog.Info("Importing type folder", "path", typeFolderPath)
	typeFileIterator, err := file.ImportFolder(typeFolderPath)
	if err != nil {
		return nil, err
	}

	// Import translations
	translationsPath := path.Join(translationFolder, typeNameTranslationFile)
	translations, err := ImportTranslations(translationsPath)
	if err != nil {
		slog.Warn("Failed to import type translations", "path", translationsPath, "error", err)
		translations = []Translation{}
	}

	return func(yield func(*PokemonTypeDescriptor) bool) {
		for typeFile := range typeFileIterator {
			typeDesc, err := UnmarshalTypeDescriptor(typeFile.Content)
			if err != nil {
				slog.Warn("Failed to unmarshal type descriptor content", "file", typeFile.Path, "error", err)
				continue
			}

			typeDesc.Name = MapTranslation(typeDesc.TextId, translations)

			if typeDesc != nil {
				if !yield(typeDesc) {
					break
				}
			}
		}
	}, nil
}

// ImportMoves import a moves folder to a store
// studioFolder pokemon studio folder
// translationFolder the translation folder
func ImportMoves(studioFolder, translationFolder string) (iter.Seq[*MoveDescriptor], error) {
	moveFolderPath := path.Join(studioFolder, "moves/")
	slog.Info("Importing move folder", "path", moveFolderPath)
	moveFileIterator, err := file.ImportFolder(moveFolderPath)
	if err != nil {
		return nil, err
	}

	// Import translations
	moveNameTranslationsPath := path.Join(translationFolder, moveNameTranslationFile)
	moveNameTranslations, err := ImportTranslations(moveNameTranslationsPath)
	if err != nil {
		slog.Warn("Failed to import move name translations", "path", moveNameTranslationsPath, "error", err)
		moveNameTranslations = []Translation{}
	}

	moveDescTranslationsPath := path.Join(translationFolder, moveDescriptionTranslationFile)
	moveDescTranslations, err := ImportTranslations(moveDescTranslationsPath)
	if err != nil {
		slog.Warn("Failed to import move description translations", "path", moveDescTranslationsPath, "error", err)
		moveDescTranslations = []Translation{}
	}

	return func(yield func(*MoveDescriptor) bool) {
		for moveFile := range moveFileIterator {
			moveDesc, err := UnmarshalMoveDescriptor(moveFile.Content)
			if err != nil {
				slog.Warn("Failed to unmarshal move descriptor content", "file", moveFile.Path, "error", err)
				continue
			}

			moveDesc.Name = MapTranslation(moveDesc.Id, moveNameTranslations)
			moveDesc.Description = MapTranslation(moveDesc.Id, moveDescTranslations)

			if moveDesc != nil {
				if !yield(moveDesc) {
					break
				}
			}
		}
	}, nil
}

// UnmarshalAbilityDescriptor unmarshal a json encoded ability to a descriptor
// abilityContent the encoded ability
func UnmarshalAbilityDescriptor(abilityContent []byte) (*AbilityDescriptor, error) {
	abilityDescriptor := &AbilityDescriptor{}
	if err := json.Unmarshal(abilityContent, abilityDescriptor); err != nil {
		return nil, err
	}

	return abilityDescriptor, nil
}

// UnmarshalPokemonDescriptor unmarshal a json encoded pokemon to a descriptor
// pokemonContent the encoded pokemon
func UnmarshalPokemonDescriptor(pokemonContent []byte) (*PokemonDescriptor, error) {
	pokemonDescriptor := &PokemonDescriptor{}
	if err := json.Unmarshal(pokemonContent, pokemonDescriptor); err != nil {
		return nil, err
	}

	// Handle __undef__ values by converting them to nil
	for i := range pokemonDescriptor.Forms {
		if pokemonDescriptor.Forms[i].Type2 != nil && *pokemonDescriptor.Forms[i].Type2 == UndefType {
			pokemonDescriptor.Forms[i].Type2 = nil
		}
	}

	return pokemonDescriptor, nil
}

// UnmarshalTypeDescriptor unmarshal a json encoded type to a descriptor
// typeContent the encoded type
func UnmarshalTypeDescriptor(typeContent []byte) (*PokemonTypeDescriptor, error) {
	typeDescriptor := &PokemonTypeDescriptor{}
	if err := json.Unmarshal(typeContent, typeDescriptor); err != nil {
		return nil, err
	}

	return typeDescriptor, nil
}

// UnmarshalMoveDescriptor unmarshal a json encoded move to a descriptor
// moveContent the encoded move
func UnmarshalMoveDescriptor(moveContent []byte) (*MoveDescriptor, error) {
	moveDescriptor := &MoveDescriptor{}
	if err := json.Unmarshal(moveContent, moveDescriptor); err != nil {
		return nil, err
	}

	return moveDescriptor, nil
}

// ImportTranslationsOrEmpty import translations from file, if it fails log the error and return an empty list
func ImportTranslationsOrEmpty(path string) []Translation {
	translations, err := ImportTranslations(path)
	if err != nil {
		slog.Error("Failed to import translations", "path", path, "error", err)
		return []Translation{}
	}
	return translations
}

// ImportTranslations import translations from file
// path the path of the file to import
func ImportTranslations(path string) ([]Translation, error) {
	curPath, _ := os.Getwd()
	slog.Info("Import translation file", "path", path, "currentPath", curPath)
	file, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			slog.Warn("Failed to close translation file", "path", path, "error", err)
		}
	}(file)

	reader := csv.NewReader(file)

	langs, err := reader.Read()
	if err != nil {
		return nil, err
	}

	results := make([]Translation, 0)
	for {
		records, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		translationMap := make(Translation)
		for index := range len(records) {
			translationMap[langs[index]] = records[index]
		}

		results = append(results, translationMap)
	}

	return results, nil
}

func MapTranslation(textId int, translations []Translation) Translation {
	if textId >= 0 && textId < len(translations) {
		translation := translations[textId]
		return translation
	}
	slog.Warn("Could not find translation", "TextID", textId, "available", len(translations))
	return nil
}
