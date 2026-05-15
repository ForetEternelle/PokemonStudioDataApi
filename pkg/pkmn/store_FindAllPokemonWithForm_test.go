package pkmn

import (
	"slices"
	"testing"
)

func TestFindAllPokemonWithForm_Basic(t *testing.T) {
	pokemonList := []Pokemon{
		*NewPokemonBuilder().ID(1).DbSymbol("bulbasaur").Forms([]PokemonForm{
			*NewPokemonFormBuilder().Form(0).Build(),
			*NewPokemonFormBuilder().Form(1).Build(),
		}).Build(),
		*NewPokemonBuilder().ID(2).DbSymbol("ivysaur").Forms([]PokemonForm{
			*NewPokemonFormBuilder().Form(0).Build(),
		}).Build(),
	}

	store := NewStore()
	for _, p := range pokemonList {
		store.AddPokemon(p)
	}

	result := store.FindAllPokemonWithForm()
	all := slices.Collect(result)

	if len(all) != 3 {
		t.Errorf("Expected 3 PokemonWithForm entries, got %d", len(all))
	}

	if all[0].Pokemon.ID() != 1 || all[0].FormId != 0 {
		t.Errorf("Expected first to be bulbasaur form 0, got ID %d form %d", all[0].Pokemon.ID(), all[0].FormId)
	}
	if all[1].Pokemon.ID() != 1 || all[1].FormId != 1 {
		t.Errorf("Expected second to be bulbasaur form 1, got ID %d form %d", all[1].Pokemon.ID(), all[1].FormId)
	}
	if all[2].Pokemon.ID() != 2 || all[2].FormId != 0 {
		t.Errorf("Expected third to be ivysaur form 0, got ID %d form %d", all[2].Pokemon.ID(), all[2].FormId)
	}
}

func TestFindAllPokemonWithForm_PokemonFilter(t *testing.T) {
	pokemonList := []Pokemon{
		*NewPokemonBuilder().ID(1).DbSymbol("bulbasaur").Forms([]PokemonForm{
			*NewPokemonFormBuilder().Form(0).Build(),
		}).Build(),
		*NewPokemonBuilder().ID(2).DbSymbol("ivysaur").Forms([]PokemonForm{
			*NewPokemonFormBuilder().Form(0).Build(),
		}).Build(),
	}

	store := NewStore()
	for _, p := range pokemonList {
		store.AddPokemon(p)
	}

	filter := func(p Pokemon) bool { return p.ID() == 1 }
	result := store.FindAllPokemonWithForm(WithPokemonFilter(filter))
	all := slices.Collect(result)

	if len(all) != 1 {
		t.Errorf("Expected 1 PokemonWithForm, got %d", len(all))
	}
	if len(all) > 0 && all[0].Pokemon.ID() != 1 {
		t.Errorf("Expected bulbasaur (ID 1), got ID %d", all[0].Pokemon.ID())
	}
}

func TestFindAllPokemonWithForm_FormFilter(t *testing.T) {
	pokemonList := []Pokemon{
		*NewPokemonBuilder().ID(1).DbSymbol("bulbasaur").Forms([]PokemonForm{
			*NewPokemonFormBuilder().Form(0).Build(),
			*NewPokemonFormBuilder().Form(1).Build(),
			*NewPokemonFormBuilder().Form(2).Build(),
		}).Build(),
		*NewPokemonBuilder().ID(2).DbSymbol("ivysaur").Forms([]PokemonForm{
			*NewPokemonFormBuilder().Form(0).Build(),
			*NewPokemonFormBuilder().Form(1).Build(),
		}).Build(),
	}

	store := NewStore()
	for _, p := range pokemonList {
		store.AddPokemon(p)
	}

	filter := func(f PokemonForm) bool { return f.Form() == 0 || f.Form() == 2 }
	result := store.FindAllPokemonWithForm(WithFormFilter(filter))
	all := slices.Collect(result)

	if len(all) != 3 {
		t.Errorf("Expected 3 PokemonWithForm (bulbasaur: form 0,2; ivysaur: form 0), got %d", len(all))
	}
}

func TestFindAllPokemonWithForm_LastId(t *testing.T) {
	pokemonList := []Pokemon{
		*NewPokemonBuilder().ID(1).DbSymbol("bulbasaur").Forms([]PokemonForm{
			*NewPokemonFormBuilder().Form(0).Build(),
		}).Build(),
		*NewPokemonBuilder().ID(2).DbSymbol("ivysaur").Forms([]PokemonForm{
			*NewPokemonFormBuilder().Form(0).Build(),
		}).Build(),
		*NewPokemonBuilder().ID(3).DbSymbol("venusaur").Forms([]PokemonForm{
			*NewPokemonFormBuilder().Form(0).Build(),
		}).Build(),
	}

	store := NewStore()
	for _, p := range pokemonList {
		store.AddPokemon(p)
	}

	lastId := int32(2)
	result := store.FindAllPokemonWithForm(WithLastId(lastId))
	all := slices.Collect(result)

	if len(all) != 1 {
		t.Errorf("Expected 1 PokemonWithForm (venusaur), got %d", len(all))
	}
}

func TestFindAllPokemonWithForm_MainFormsOnly(t *testing.T) {
	pokemonList := []Pokemon{
		*NewPokemonBuilder().ID(1).DbSymbol("bulbasaur").Forms([]PokemonForm{
			*NewPokemonFormBuilder().Form(0).Build(),
			*NewPokemonFormBuilder().Form(1).Build(),
		}).Build(),
		*NewPokemonBuilder().ID(2).DbSymbol("ivysaur").Forms([]PokemonForm{
			*NewPokemonFormBuilder().Form(0).Build(),
			*NewPokemonFormBuilder().Form(1).Build(),
			*NewPokemonFormBuilder().Form(2).Build(),
		}).Build(),
	}

	store := NewStore()
	for _, p := range pokemonList {
		store.AddPokemon(p)
	}

	result := store.FindAllPokemonWithForm(WithMainFormsOnly())
	all := slices.Collect(result)

	if len(all) != 2 {
		t.Errorf("Expected 2 PokemonWithForm (one per pokemon, form 0), got %d", len(all))
	}
	for i, p := range all {
		if p.FormId != 0 {
			t.Errorf("Entry %d: expected FormId 0, got %d", i, p.FormId)
		}
	}
}


func TestFindAllPokemonWithForm_CombinedFilters(t *testing.T) {
	pokemonList := []Pokemon{
		*NewPokemonBuilder().ID(1).DbSymbol("bulbasaur").Forms([]PokemonForm{
			*NewPokemonFormBuilder().Form(0).Build(),
			*NewPokemonFormBuilder().Form(1).Build(),
		}).Build(),
		*NewPokemonBuilder().ID(2).DbSymbol("ivysaur").Forms([]PokemonForm{
			*NewPokemonFormBuilder().Form(0).Build(),
			*NewPokemonFormBuilder().Form(1).Build(),
			*NewPokemonFormBuilder().Form(2).Build(),
		}).Build(),
	}

	store := NewStore()
	for _, p := range pokemonList {
		store.AddPokemon(p)
	}

	pokemonFilter := func(p Pokemon) bool { return p.ID() == 2 }
	formFilter := func(f PokemonForm) bool { return f.Form() >= 1 }
	result := store.FindAllPokemonWithForm(WithPokemonFilter(pokemonFilter), WithFormFilter(formFilter))
	all := slices.Collect(result)

	if len(all) != 2 {
		t.Errorf("Expected 2 PokemonWithForm (ivysaur forms 1 and 2), got %d", len(all))
	}
	for _, p := range all {
		if p.Pokemon.ID() != 2 {
			t.Errorf("Expected ivysaur (ID 2), got ID %d", p.Pokemon.ID())
		}
		if p.FormId < 1 {
			t.Errorf("Expected form >= 1, got %d", p.FormId)
		}
	}
}

func TestFindAllPokemonWithForm_LastIdAndMainFormsOnly(t *testing.T) {
	pokemonList := []Pokemon{
		*NewPokemonBuilder().ID(1).DbSymbol("bulbasaur").Forms([]PokemonForm{
			*NewPokemonFormBuilder().Form(0).Build(),
			*NewPokemonFormBuilder().Form(1).Build(),
		}).Build(),
		*NewPokemonBuilder().ID(2).DbSymbol("ivysaur").Forms([]PokemonForm{
			*NewPokemonFormBuilder().Form(0).Build(),
			*NewPokemonFormBuilder().Form(1).Build(),
		}).Build(),
		*NewPokemonBuilder().ID(3).DbSymbol("venusaur").Forms([]PokemonForm{
			*NewPokemonFormBuilder().Form(0).Build(),
			*NewPokemonFormBuilder().Form(1).Build(),
		}).Build(),
	}

	store := NewStore()
	for _, p := range pokemonList {
		store.AddPokemon(p)
	}

	lastId := int32(2)
	result := store.FindAllPokemonWithForm(WithLastId(lastId), WithMainFormsOnly())
	all := slices.Collect(result)

	if len(all) != 1 {
		t.Errorf("Expected 1 PokemonWithForm (venusaur form 0), got %d", len(all))
	}
}

func TestFindAllPokemonWithForm_EmptyStore(t *testing.T) {
	store := NewStore()

	result := store.FindAllPokemonWithForm()
	all := slices.Collect(result)

	if len(all) != 0 {
		t.Errorf("Expected 0 PokemonWithForm, got %d", len(all))
	}
}

func TestFindAllPokemonWithForm_NoForms(t *testing.T) {
	pokemonList := []Pokemon{
		*NewPokemonBuilder().ID(1).DbSymbol("bulbasaur").Build(),
		*NewPokemonBuilder().ID(2).DbSymbol("ivysaur").Build(),
	}

	store := NewStore()
	for _, p := range pokemonList {
		store.AddPokemon(p)
	}

	result := store.FindAllPokemonWithForm()
	all := slices.Collect(result)

	if len(all) != 0 {
		t.Errorf("Expected 0 PokemonWithForm (no forms), got %d", len(all))
	}
}

func TestFindAllPokemonWithForm_PokemonWithAndWithoutForms(t *testing.T) {
	pokemonList := []Pokemon{
		*NewPokemonBuilder().ID(1).DbSymbol("bulbasaur").Forms([]PokemonForm{
			*NewPokemonFormBuilder().Form(0).Build(),
		}).Build(),
		*NewPokemonBuilder().ID(2).DbSymbol("ivysaur").Build(),
		*NewPokemonBuilder().ID(3).DbSymbol("venusaur").Forms([]PokemonForm{
			*NewPokemonFormBuilder().Form(0).Build(),
			*NewPokemonFormBuilder().Form(1).Build(),
		}).Build(),
	}

	store := NewStore()
	for _, p := range pokemonList {
		store.AddPokemon(p)
	}

	result := store.FindAllPokemonWithForm()
	all := slices.Collect(result)

	if len(all) != 3 {
		t.Errorf("Expected 3 PokemonWithForm (bulbasaur form 0, venusaur forms 0,1), got %d", len(all))
	}
}

func TestFindAllPokemonWithForm_FormFilterExcludesAll(t *testing.T) {
	pokemonList := []Pokemon{
		*NewPokemonBuilder().ID(1).DbSymbol("bulbasaur").Forms([]PokemonForm{
			*NewPokemonFormBuilder().Form(0).Build(),
		}).Build(),
	}

	store := NewStore()
	for _, p := range pokemonList {
		store.AddPokemon(p)
	}

	filter := func(f PokemonForm) bool { return false }
	result := store.FindAllPokemonWithForm(WithFormFilter(filter))
	all := slices.Collect(result)

	if len(all) != 0 {
		t.Errorf("Expected 0 PokemonWithForm, got %d", len(all))
	}
}

func TestFindAllPokemonWithForm_LastIdNotFound(t *testing.T) {
	pokemonList := []Pokemon{
		*NewPokemonBuilder().ID(1).DbSymbol("bulbasaur").Forms([]PokemonForm{
			*NewPokemonFormBuilder().Form(0).Build(),
		}).Build(),
	}

	store := NewStore()
	for _, p := range pokemonList {
		store.AddPokemon(p)
	}

	lastId := int32(999)
	result := store.FindAllPokemonWithForm(WithLastId(lastId))
	all := slices.Collect(result)

	if len(all) != 0 {
		t.Errorf("Expected 0 PokemonWithForm (lastId not found), got %d", len(all))
	}
}

