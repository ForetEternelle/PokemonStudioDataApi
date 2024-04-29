package studio

import (
	"testing"

	"github.com/rcharre/psapi/pkg/utils/pagination"
)

const (
	DataFolder = "../../test/test_resources/valid-data"
)

func TestImport(t *testing.T) {
	store := NewStore()
	Import(DataFolder, store)

	pokemonStore := store.PokemonStore
	page := pokemonStore.FindAll(pagination.NewPageRequest(0, 100))

	if page.Total != 5 {
		t.Error("Import should have 5 pokemon", "has", page.Total)
	}

	typesStore := store.TypeStore
	types := typesStore.FindAll()

	nbTypes := len(types)
	if nbTypes != 18 {
		t.Error("Import should have 18 types", "has", nbTypes)
	}
}