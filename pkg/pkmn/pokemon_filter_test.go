package pkmn

import (
	"testing"
)

func TestNewPokemonQueryFilter(t *testing.T) {
	form := &PokemonForm{Form: 0, Name: Translation{"en": "Pikachu", "fr": "Pikachu", "de": "Pikachu"}}
	pokemon := &Pokemon{
		ID:       25,
		DbSymbol: "pikachu",
		Forms:    []*PokemonForm{form},
	}

	pwf := PokemonWithForm{
		Pokemon: pokemon,
		FormId:  0,
	}

	t.Run("ID match", func(t *testing.T) {
		if !NewPokemonQueryFilter("25")(pwf) {
			t.Error("Expected match ID '25'")
		}
	})

	t.Run("ID contained match", func(t *testing.T) {
		if !NewPokemonQueryFilter("2")(pwf) {
			t.Error("Expected match ID '25'")
		}
	})

	t.Run("ID no match", func(t *testing.T) {
		if NewPokemonQueryFilter("26")(pwf) {
			t.Error("Expected no match ID '26'")
		}
	})

	t.Run("name match", func(t *testing.T) {
		if !NewPokemonQueryFilter("Pika", "en")(pwf) {
			t.Error("Expected match name 'Pika' in 'en'")
		}
	})

	t.Run("name specific language", func(t *testing.T) {
		if !NewPokemonQueryFilter("Pika", "fr")(pwf) {
			t.Error("Expected match name 'Pika' in 'fr'")
		}
	})

	t.Run("name case-insensitive", func(t *testing.T) {
		if !NewPokemonQueryFilter("PIKACHU", "en")(pwf) {
			t.Error("Expected match name case-insensitively")
		}
	})

	t.Run("no match", func(t *testing.T) {
		if NewPokemonQueryFilter("bulbasaur", "en")(pwf) {
			t.Error("Expected no match 'bulbasaur'")
		}
	})

	t.Run("empty query", func(t *testing.T) {
		if !NewPokemonQueryFilter("")(pwf) {
			t.Error("Expected empty query to match everything")
		}
	})

	t.Run("no languages", func(t *testing.T) {
		if NewPokemonQueryFilter("raichu")(pwf) {
			t.Error("Expected no match without langs for name-only query")
		}
	})

	t.Run("multiple languages", func(t *testing.T) {
		if !NewPokemonQueryFilter("chu", "en", "fr")(pwf) {
			t.Error("Expected match 'chu' in multiple languages")
		}
	})
}

func TestNewPokemonFormTypesFilter(t *testing.T) {
	fire := &PokemonType{DbSymbol: "fire"}
	water := &PokemonType{DbSymbol: "water"}
	grass := &PokemonType{DbSymbol: "grass"}
	flying := &PokemonType{DbSymbol: "flying"}

	charizard := &PokemonForm{Form: 0, Type1: fire, Type2: flying}
	squirtle := &PokemonForm{Form: 0, Type1: water}
	bulbasaur := &PokemonForm{Form: 0, Type1: grass}

	t.Run("match type1", func(t *testing.T) {
		if !NewPokemonFormTypesFilter([]string{"fire"})(charizard) {
			t.Error("Expected match 'fire' on charizard (type1)")
		}
	})

	t.Run("match type2", func(t *testing.T) {
		if !NewPokemonFormTypesFilter([]string{"flying"})(charizard) {
			t.Error("Expected match 'flying' on charizard (type2)")
		}
	})

	t.Run("match any of multiple types", func(t *testing.T) {
		if !NewPokemonFormTypesFilter([]string{"electric", "water"})(squirtle) {
			t.Error("Expected match 'water' from multiple types")
		}
	})

	t.Run("single type no type2", func(t *testing.T) {
		if !NewPokemonFormTypesFilter([]string{"grass"})(bulbasaur) {
			t.Error("Expected match 'grass' on bulbasaur")
		}
	})

	t.Run("no match", func(t *testing.T) {
		if NewPokemonFormTypesFilter([]string{"electric"})(charizard) {
			t.Error("Expected no match 'electric' on charizard")
		}
	})

	t.Run("empty types list", func(t *testing.T) {
		if !NewPokemonFormTypesFilter([]string{})(charizard) {
			t.Error("Expected match with empty types list")
		}
	})

	t.Run("nil types list", func(t *testing.T) {
		if !NewPokemonFormTypesFilter(nil)(charizard) {
			t.Error("Expected match with nil types list")
		}
	})
}

func TestNewPokemonWithTagsFilter(t *testing.T) {
	form := &PokemonForm{Form: 0, Name: Translation{"en": "Raichu"}}
	form.Tags = []string{"alolan", "evolved"}
	pokemon := &Pokemon{
		ID:       26,
		DbSymbol: "raichu",
		Forms:    []*PokemonForm{form},
	}
	pokemon.Tags = []string{"special"}

	pwf := PokemonWithForm{
		Pokemon: pokemon,
		FormId:  0,
	}

	t.Run("no tags matches all", func(t *testing.T) {
		if !NewPokemonWithTagsFilter(nil)(pwf) {
			t.Error("Expected match with no tags")
		}
	})

	t.Run("empty tags matches all", func(t *testing.T) {
		if !NewPokemonWithTagsFilter([]string{})(pwf) {
			t.Error("Expected match with empty tags")
		}
	})

	t.Run("match pokemon tag", func(t *testing.T) {
		if !NewPokemonWithTagsFilter([]string{"special"})(pwf) {
			t.Error("Expected match pokemon tag 'special'")
		}
	})

	t.Run("match form tag", func(t *testing.T) {
		if !NewPokemonWithTagsFilter([]string{"alolan"})(pwf) {
			t.Error("Expected match form tag 'alolan'")
		}
	})

	t.Run("match any of multiple tags", func(t *testing.T) {
		if !NewPokemonWithTagsFilter([]string{"other", "evolved"})(pwf) {
			t.Error("Expected match 'evolved' from multiple tags")
		}
	})

	t.Run("no match", func(t *testing.T) {
		if NewPokemonWithTagsFilter([]string{"legendary"})(pwf) {
			t.Error("Expected no match 'legendary'")
		}
	})
}

func TestNewPokemonWithoutTagsFilter(t *testing.T) {
	form := &PokemonForm{Form: 0, Name: Translation{"en": "Raichu"}}
	form.Tags = []string{"alolan", "evolved"}
	pokemon := &Pokemon{
		ID:       26,
		DbSymbol: "raichu",
		Forms:    []*PokemonForm{form},
	}
	pokemon.Tags = []string{"special"}

	pwf := PokemonWithForm{
		Pokemon: pokemon,
		FormId:  0,
	}

	t.Run("no tags matches all", func(t *testing.T) {
		if !NewPokemonWithoutTagsFilter(nil)(pwf) {
			t.Error("Expected match with no tags")
		}
	})

	t.Run("empty tags matches all", func(t *testing.T) {
		if !NewPokemonWithoutTagsFilter([]string{})(pwf) {
			t.Error("Expected match with empty tags")
		}
	})

	t.Run("exclude pokemon tag", func(t *testing.T) {
		if NewPokemonWithoutTagsFilter([]string{"special"})(pwf) {
			t.Error("Expected no match with excluded pokemon tag 'special'")
		}
	})

	t.Run("exclude form tag", func(t *testing.T) {
		if NewPokemonWithoutTagsFilter([]string{"evolved"})(pwf) {
			t.Error("Expected no match with excluded form tag 'evolved'")
		}
	})

	t.Run("exclude unrelated tag keeps result", func(t *testing.T) {
		if !NewPokemonWithoutTagsFilter([]string{"legendary"})(pwf) {
			t.Error("Expected match when only unrelated tag is excluded")
		}
	})
}
