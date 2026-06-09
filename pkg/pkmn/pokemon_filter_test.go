package pkmn

import (
	"testing"
)

func TestNewPokemonQueryFilter(t *testing.T) {
	pwf := PokemonWithForm{
		Pokemon: NewPokemonBuilder().ID(25).DbSymbol("pikachu").Forms([]*PokemonForm{
			NewPokemonFormBuilder().Form(0).Name(Translation{"en": "Pikachu", "fr": "Pikachu", "de": "Pikachu"}).Build(),
		}).Build(),
		FormId: 0,
	}

	t.Run("dbSymbol match", func(t *testing.T) {
		if !NewPokemonQueryFilter("pika")(pwf) {
			t.Error("Expected match 'pika' in dbSymbol 'pikachu'")
		}
	})

	t.Run("dbSymbol case-insensitive", func(t *testing.T) {
		if !NewPokemonQueryFilter("PIKA")(pwf) {
			t.Error("Expected match 'PIKA' in dbSymbol 'pikachu'")
		}
	})

	t.Run("ID match", func(t *testing.T) {
		if !NewPokemonQueryFilter("25")(pwf) {
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
	fire := &PokemonType{dbSymbol: "fire"}
	water := &PokemonType{dbSymbol: "water"}
	grass := &PokemonType{dbSymbol: "grass"}
	flying := &PokemonType{dbSymbol: "flying"}

	charizard := NewPokemonFormBuilder().Form(0).Type1(fire).Type2(flying).Build()
	squirtle := NewPokemonFormBuilder().Form(0).Type1(water).Build()
	bulbasaur := NewPokemonFormBuilder().Form(0).Type1(grass).Build()

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
		if NewPokemonFormTypesFilter([]string{})(charizard) {
			t.Error("Expected no match with empty types list")
		}
	})
}
