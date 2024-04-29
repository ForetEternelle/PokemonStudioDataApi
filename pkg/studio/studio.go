package studio

import (
	"path"
)

type Translation map[string]string

const (
	StudioFolder   = "Studio"
	LanguageFolder = "Text/Dialogs"

	UndefType = "__undef__"
)

// Import import a pokemon studio folder into a store
// folder the studio project folder
// store the store to import data to
func Import(folder string, store *Store) error {
	translationFolder := path.Join(folder, LanguageFolder)
	studioFolder := path.Join(folder, StudioFolder)

	if err := ImportPokemon(studioFolder, translationFolder, store); err != nil {
		return err
	}

	if err := ImportTypes(studioFolder, translationFolder, store); err != nil {
		return err
	}

	return nil
}
