package studio

import (
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"
)

type rawAbility struct {
	Klass    string `json:"klass"`
	DbSymbol string `json:"dbSymbol"`
	ID       int    `json:"id"`
	TextID   int    `json:"textId"`
}

func LoadAbilitiesFromDir(dir string, store *Store) error {
	return filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		if filepath.Ext(path) != ".json" {
			return nil
		}
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()
		var r rawAbility
		if err := json.NewDecoder(f).Decode(&r); err != nil {
			return err
		}
		// only accept abilities
		if r.Klass != "Ability" {
			return nil
		}
		a := Ability{
			Symbol: r.DbSymbol,
			ID:     r.ID,
			TextID: r.TextID,
		}
		store.Abilities = append(store.Abilities, a)
		return nil
	})
}
