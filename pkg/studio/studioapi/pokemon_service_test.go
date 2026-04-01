package studioapi

import (
	"context"
	"testing"

	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/pagination"
	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/studio"
)

func setupPokemonService() (*studio.Store, PokemonAPIServicer) {
	store := studio.NewStore()
	normalType := studio.NewTypeBuilder().DbSymbol("normal").Name(studio.Translation{"en": "Normal"}).Build()
	electricType := studio.NewTypeBuilder().DbSymbol("electric").Name(studio.Translation{"en": "Electric"}).Build()
	store.AddType(*normalType)
	store.AddType(*electricType)

	form0 := studio.NewPokemonFormBuilder().
		Form(0).
		Type1(electricType).
		BaseHp(35).
		BaseAtk(55).
		BaseSpd(90).
		Name(studio.Translation{"en": "Pikachu", "fr": "PikachuFR"}).
		Description(studio.Translation{"en": "Electric mouse"}).
		Build()
	form1 := studio.NewPokemonFormBuilder().
		Form(1).
		Type1(electricType).
		Name(studio.Translation{"en": "Pikachu"}).
		Description(studio.Translation{"en": "Electric mouse"}).
		BaseHp(20).
		BaseAtk(40).
		Build()
	pikachu := studio.NewPokemonBuilder().
		ID(25).
		DbSymbol("pikachu").
		Forms(map[int32]studio.PokemonForm{0: *form0, 1: *form1}).
		Build()
	store.AddPokemon(*pikachu)

	bulbasaurForm := studio.NewPokemonFormBuilder().
		Form(0).
		Name(studio.Translation{"en": "Bulbasaur"}).
		Description(studio.Translation{"en": "Grass starter"}).
		Type1(normalType).
		BaseHp(45).
		BaseAtk(49).
		Name(studio.Translation{"en": "Bulbasaur"}).
		Description(studio.Translation{"en": "Grass starter"}).
		Build()
	bulbasaur := studio.NewPokemonBuilder().
		ID(1).
		DbSymbol("bulbasaur").
		Forms(map[int32]studio.PokemonForm{0: *bulbasaurForm}).
		Build()
	store.AddPokemon(*bulbasaur)

	smettleForm := studio.NewPokemonFormBuilder().
		Form(0).
		Name(studio.Translation{"en": "Smettle", "fr": "Malortie"}).
		Description(studio.Translation{"en": "Smettle is a mischievous Pokémon.", "fr": "Malortie est un Pokémon malicieux."}).
		Type1(normalType).
		BaseHp(45).
		BaseAtk(49).
		Build()
	smettle := studio.NewPokemonBuilder().
		ID(724).
		DbSymbol("smettle").
		Forms(map[int32]studio.PokemonForm{0: *smettleForm}).
		Build()
	store.AddPokemon(*smettle)

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
	if details.MainForm.Name != "Pikachu" {
		t.Error("Expected name Pikachu, got", details.MainForm.Name)
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
	if len(page.Content) != 3 {
		t.Error("Expected 3 pokemon, got", len(page.Content))
	}
}

func TestPokemonService_GetPokemon_Pagination(t *testing.T) {
	store := studio.NewStore()
	normalType := studio.NewTypeBuilder().DbSymbol("normal").Build()
	store.AddType(*normalType)

	for i := 1; i <= 15; i++ {
		form := studio.NewPokemonFormBuilder().
			Form(0).
			Type1(normalType).
			Name(studio.Translation{"en": "Pokemon"}).
			Build()
		pokemon := studio.NewPokemonBuilder().
			ID(int32(i)).
			DbSymbol("pokemon_" + string(rune(i))).
			Forms(map[int32]studio.PokemonForm{0: *form}).
			Build()
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

func TestPokemonService_GetPokemonDetailsByName(t *testing.T) {
	_, service := setupPokemonService()

	// Test English name
	resp, err := service.GetPokemonDetailsByName(context.Background(), "Pikachu", "en")
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

	// Test French name
	resp, err = service.GetPokemonDetailsByName(context.Background(), "PikachuFR", "en")
	if err != nil {
		t.Error("Expected no error, got", err)
	}
	if resp.Code != 200 {
		t.Error("Expected status 200, got", resp.Code)
	}
	details = resp.Body.(*PokemonDetails)
	if details.Symbol != "pikachu" {
		t.Error("Expected symbol pikachu, got", details.Symbol)
	}

	// Test Case Insensitivity
	resp, err = service.GetPokemonDetailsByName(context.Background(), "pikachu", "en")
	if err != nil {
		t.Error("Expected no error, got", err)
	}
	if resp.Code != 200 {
		t.Error("Expected status 200 for lowercase name, got", resp.Code)
	}

	// Test Symbol search (new fallback)
	resp, err = service.GetPokemonDetailsByName(context.Background(), "pikachu", "en")
	if err != nil {
		t.Error("Expected no error, got", err)
	}
	if resp.Code != 200 {
		t.Error("Expected status 200 for symbol search, got", resp.Code)
	}

	resp, err = service.GetPokemonDetailsByName(context.Background(), "PIKACHU", "en")
	if err != nil {
		t.Error("Expected no error, got", err)
	}
	if resp.Code != 200 {
		t.Error("Expected status 200 for uppercase name, got", resp.Code)
	}

	// Test with spaces
	resp, err = service.GetPokemonDetailsByName(context.Background(), "  Pikachu  ", "en")
	if err != nil {
		t.Error("Expected no error, got", err)
	}
	if resp.Code != 200 {
		t.Error("Expected status 200 for name with spaces, got", resp.Code)
	}

	// Test Bulbasaur
	resp, err = service.GetPokemonDetailsByName(context.Background(), "Bulbasaur", "en")
	if err != nil {
		t.Error("Expected no error, got", err)
	}
	if resp.Code != 200 {
		t.Error("Expected status 200, got", resp.Code)
	}

	// Test Smettle (English)
	resp, err = service.GetPokemonDetailsByName(context.Background(), "Smettle", "en")
	if err != nil {
		t.Error("Expected no error, got", err)
	}
	if resp.Code != 200 {
		t.Error("Expected status 200 for Smettle, got", resp.Code)
	}
	details = resp.Body.(*PokemonDetails)
	if details.Symbol != "smettle" {
		t.Error("Expected symbol smettle, got", details.Symbol)
	}

	// Test Malortie (French name for Smettle)
	resp, err = service.GetPokemonDetailsByName(context.Background(), "Malortie", "fr")
	if err != nil {
		t.Error("Expected no error, got", err)
	}
	if resp.Code != 200 {
		t.Error("Expected status 200 for Malortie, got", resp.Code)
	}
	details = resp.Body.(*PokemonDetails)
	if details.Symbol != "smettle" {
		t.Error("Expected symbol smettle for name Malortie, got", details.Symbol)
	}

	// Test Not Found
	resp, err = service.GetPokemonDetailsByName(context.Background(), "Mewtwo", "en")
	if err != nil {
		t.Error("Expected no error, got", err)
	}
	if resp.Code != 404 {
		t.Error("Expected status 404, got", resp.Code)
	}
}
