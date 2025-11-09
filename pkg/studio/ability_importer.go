package studio

import (
	"encoding/json"
	"log/slog"
	"path"

	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/file"
)

const (
	AbilityFolder                         = "abilities/"
	AbilityTranslationFileName            = "100004.csv"
	AbilityDescriptionTranslationFileName = "100005.csv"
)

func ImportAbility(studioFolder, translationFolder string) ([]Ability, error) {
	slog.Info("Importing ability name translation")
	abilityNameFilePath := path.Join(translationFolder, AbilityTranslationFileName)
	abilityNameTranslations, err := ImportTranslations(abilityNameFilePath)
	if err != nil {
		return nil, err
	}

	slog.Info("Importing ability description translation")
	abilityDescriptionFilePath := path.Join(translationFolder, AbilityDescriptionTranslationFileName)
	abilityDescriptionTranslations, err := ImportTranslations(abilityDescriptionFilePath)
	if err != nil {
		return nil, err
	}

	abilityFolderPath := path.Join(studioFolder, AbilityFolder)
	slog.Info("Importing ability folder", "path", abilityFolderPath)
	abilityFileIterator, err := file.ImportFolder(abilityFolderPath)
	if err != nil {
		return nil, err
	}

	abilities := make([]Ability, 0)
	for abilityFile := range abilityFileIterator {
		ability, err := UnmarshalAbility(abilityFile.Content)
		if err != nil {
			slog.Warn("Failed to unmarshal ability content", "file", abilityFile.Path)
			continue
		}
		TranslateAbility(ability, abilityNameTranslations, abilityDescriptionTranslations)
		abilities = append(abilities, *ability)
	}
	return abilities, nil
}

// UnmarshalAbility unmarshal a json encoded ability to an object
// abilityContent the encoded ability
func UnmarshalAbility(abilityContent []byte) (*Ability, error) {
	ability := &Ability{}
	if err := json.Unmarshal(abilityContent, ability); err != nil {
		return nil, err
	}

	return ability, nil
}

// TranslateAbility add a translation to an ability name and description
// ability the ability to add translation to
// abilityNameTranslations the datastructure containing all ability names translations
// abilityDescriptionTranslations the datastructure containing all ability descriptions translations
func TranslateAbility(ability *Ability, abilityNameTranslations, abilityDescriptionTranslations []Translation) {
	// prevent out of range access like other Translate* functions
	if ability.TextID < len(abilityNameTranslations) {
		ability.Name = abilityNameTranslations[ability.TextID]
	} else {
		slog.Warn("Could not find translation for ability", "symbol", ability.DbSymbol, "TextID", ability.TextID)
	}

	if ability.TextID < len(abilityDescriptionTranslations) {
		ability.Description = abilityDescriptionTranslations[ability.TextID]
	} else {
		slog.Warn("Could not find description translation for ability", "symbol", ability.DbSymbol, "TextID", ability.TextID)
	}
}
