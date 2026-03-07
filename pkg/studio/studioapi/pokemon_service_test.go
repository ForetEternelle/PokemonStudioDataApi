package studioapi

import (
	"context"
	"testing"

	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/pagination"
	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/studio"
)

func setupPokemonService() (*studio.Store, PokemonAPIServicer) {
	store := studio.NewStore()
	normalType := studio.NewPokemonType(studio.WithPokemonTypeDbSymbol("normal"), studio.WithTypeName(studio.Translation{"en": "Normal"}))
	electricType := studio.NewPokemonType(studio.WithPokemonTypeDbSymbol("electric"), studio.WithTypeName(studio.Translation{"en": "Electric"}))
	store.AddType(*normalType)
	store.AddType(*electricType)

	form0 := studio.NewPokemonForm(
		studio.WithForm(0),
		studio.WithType1(electricType),
		studio.WithBaseHp(35),
		studio.WithBaseAtk(55),
	)
	form1 := studio.NewPokemonForm(
		studio.WithForm(1),
		studio.WithType1(electricType),
		studio.WithBaseHp(20),
		studio.WithBaseAtk(40),
	)
	pikachu := studio.NewPokemon(
		studio.WithID(25),
		studio.WithDbSymbol("pikachu"),
		studio.WithName(studio.Translation{"en": "Pikachu"}),
		studio.WithDescription(studio.Translation{"en": "Electric mouse"}),
		studio.WithForms(map[int32]studio.PokemonForm{0: *form0, 1: *form1}),
	)
	store.AddPokemon(*pikachu)

	bulbasaurForm := studio.NewPokemonForm(
		studio.WithForm(0),
		studio.WithType1(normalType),
		studio.WithBaseHp(45),
		studio.WithBaseAtk(49),
	)
	bulbasaur := studio.NewPokemon(
		studio.WithID(1),
		studio.WithDbSymbol("bulbasaur"),
		studio.WithName(studio.Translation{"en": "Bulbasaur"}),
		studio.WithDescription(studio.Translation{"en": "Grass starter"}),
		studio.WithForms(map[int32]studio.PokemonForm{0: *bulbasaurForm}),
	)
	store.AddPokemon(*bulbasaur)

	typeMapper := NewTypeMapper()
	abilityMapper := NewAbilityMapper()
	pokemonMapper := NewPokemonMapper(typeMapper, abilityMapper, store)

	accessPolicyFactory := func(ctx context.Context) *AccessPolicy {
		return NewAccessPolicy()
	}

	service := NewPokemonService(store, pokemonMapper, accessPolicyFactory)
	return store, service
}

func TestPokemonService_GetPokemonDetails(t *testing.T) {
	_, service := setupPokemonService()

	resp, err := service.GetPokemonDetails(context.Background(), "pikachu", "en")
	if err != nil {
		t.Error("Expected no error, got", err)
	}
	if resp.Code != 200 {
		t.Error("Expected status 200, got", resp.Code)
	}

	details := resp.Body.(*PokemonDetails)
	if details.Symbol != "pikachu" {
		t.Error("Expected symbol pikachu, got", details.Symbol)
	}
	if details.Name != "Pikachu" {
		t.Error("Expected name Pikachu, got", details.Name)
	}
}

func TestPokemonService_GetPokemonDetails_NotFound(t *testing.T) {
	_, service := setupPokemonService()

	resp, err := service.GetPokemonDetails(context.Background(), "mewtwo", "en")
	if err != nil {
		t.Error("Expected no error, got", err)
	}
	if resp.Code != 404 {
		t.Error("Expected status 404, got", resp.Code)
	}
	if resp.Body != nil {
		t.Error("Expected nil body for non-existent pokemon")
	}
}

func TestPokemonService_GetPokemon(t *testing.T) {
	_, service := setupPokemonService()

	resp, err := service.GetPokemon(context.Background(), 0, 10, "en")
	if err != nil {
		t.Error("Expected no error, got", err)
	}
	if resp.Code != 200 {
		t.Error("Expected status 200, got", resp.Code)
	}

	page := resp.Body.(pagination.Page[*PokemonThumbnail])
	if page.Content == nil {
		t.Error("Expected non-nil content")
	}
	if len(page.Content) != 2 {
		t.Error("Expected 2 pokemon, got", len(page.Content))
	}
}

func TestPokemonService_GetPokemon_Pagination(t *testing.T) {
	store := studio.NewStore()
	normalType := studio.NewPokemonType(studio.WithPokemonTypeDbSymbol("normal"))
	store.AddType(*normalType)

	for i := 1; i <= 15; i++ {
		form := studio.NewPokemonForm(
			studio.WithForm(0),
			studio.WithType1(normalType),
		)
		pokemon := studio.NewPokemon(
			studio.WithID(int32(i)),
			studio.WithDbSymbol("pokemon_"+string(rune(i))),
			studio.WithName(studio.Translation{"en": "Pokemon"}),
			studio.WithForms(map[int32]studio.PokemonForm{0: *form}),
		)
		store.AddPokemon(*pokemon)
	}

	typeMapper := NewTypeMapper()
	abilityMapper := NewAbilityMapper()
	pokemonMapper := NewPokemonMapper(typeMapper, abilityMapper, store)

	accessPolicyFactory := func(ctx context.Context) *AccessPolicy {
		return NewAccessPolicy()
	}

	service := NewPokemonService(store, pokemonMapper, accessPolicyFactory)

	resp, _ := service.GetPokemon(context.Background(), 0, 5, "en")
	page := resp.Body.(pagination.Page[*PokemonThumbnail])

	if len(page.Content) != 5 {
		t.Error("Expected 5 items per page, got", len(page.Content))
	}
	if page.Total != 15 {
		t.Error("Expected total 15, got", page.Total)
	}
}

func TestPokemonService_GetPokemonForm(t *testing.T) {
	_, service := setupPokemonService()

	resp, err := service.GetPokemonForm(context.Background(), "pikachu", 0, "en")
	if err != nil {
		t.Error("Expected no error, got", err)
	}
	if resp.Code != 200 {
		t.Error("Expected status 200, got", resp.Code)
	}

	form := resp.Body.(*FormDetails)
	if form.Form == nil {
		t.Error("Expected non-nil form")
	}
}

func TestPokemonService_GetPokemonForm_NotFound(t *testing.T) {
	_, service := setupPokemonService()

	resp, err := service.GetPokemonForm(context.Background(), "pikachu", 99, "en")
	if err != nil {
		t.Error("Expected no error, got", err)
	}
	if resp.Code != 404 {
		t.Error("Expected status 404 for non-existent form, got", resp.Code)
	}
}
