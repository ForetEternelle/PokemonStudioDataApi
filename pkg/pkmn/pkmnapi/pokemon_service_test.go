package pkmnapi

import (
	"context"
	"testing"

	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/pkmn"
	. "github.com/ForetEternelle/PokemonStudioDataApi/pkg/pkmn/pkmnapispec"
	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/scroll"
)

func setupPokemonService(t *testing.T) (*pkmn.Store, PokemonAPIServicer) {
	t.Helper()
	store := pkmn.NewStore()

	store.AddType(pkmn.AddTypeDto{
		DbSymbol: "normal",
		Name:     pkmn.Translation{"en": "Normal"},
	})
	store.AddType(pkmn.AddTypeDto{
		DbSymbol: "electric",
		Name:     pkmn.Translation{"en": "Electric"},
	})

	store.AddPokemon(pkmn.AddPokemonDto{
		ID: 25, DbSymbol: "pikachu",
		Forms: []pkmn.AddPokemonFormDto{
			{Form: 0, Type1: "electric", Name: pkmn.Translation{"en": "Pikachu", "fr": "PikachuFR"}, Description: pkmn.Translation{"en": "Electric mouse"}},
			{Form: 1, Type1: "electric", Name: pkmn.Translation{"en": "Pikachu"}, Description: pkmn.Translation{"en": "Electric mouse"}},
		},
	})

	store.AddPokemon(pkmn.AddPokemonDto{
		ID: 1, DbSymbol: "bulbasaur",
		Forms: []pkmn.AddPokemonFormDto{
			{Form: 0, Type1: "normal", Name: pkmn.Translation{"en": "Bulbasaur"}, Description: pkmn.Translation{"en": "Grass starter"}},
		},
	})

	store.AddPokemon(pkmn.AddPokemonDto{
		ID: 724, DbSymbol: "smettle",
		Forms: []pkmn.AddPokemonFormDto{
			{Form: 0, Type1: "normal", Name: pkmn.Translation{"en": "Smettle", "fr": "Malortie"}, Description: pkmn.Translation{"en": "Smettle is a mischievous Pokémon.", "fr": "Malortie est un Pokémon malicieux."}},
		},
	})

	typeMapper := NewTypeMapper()
	abilityMapper := NewAbilityMapper()
	pokemonMapper := NewPokemonMapper(typeMapper, abilityMapper, store)

	accessPolicyFactory := func(ctx context.Context) *PokemonFilterPolicy {
		return NewPokemonFilterPolicy()
	}

	service := NewPokemonService(store, pokemonMapper, accessPolicyFactory)
	return store, service
}

func TestPokemonService_GetPokemonDetails(t *testing.T) {
	_, service := setupPokemonService(t)

	resp, err := service.GetPokemonDetails(context.Background(), "pikachu", 0, "en")
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
	if details.Form.Name != "Pikachu" {
		t.Error("Expected name Pikachu, got", details.Form.Name)
	}
}

func TestPokemonService_GetPokemonDetails_NotFound(t *testing.T) {
	_, service := setupPokemonService(t)

	resp, err := service.GetPokemonDetails(context.Background(), "mewtwo", 0, "en")
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

func TestPokemonService_GetPokemonDetails_FormNotFound(t *testing.T) {
	_, service := setupPokemonService(t)

	resp, err := service.GetPokemonDetails(context.Background(), "pikachu", 2, "en")
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
	_, service := setupPokemonService(t)

	resp, err := service.GetPokemon(context.Background(), "en", 20, nil, nil, true, nil, []string{}, []string{}, []string{})
	if err != nil {
		t.Error("Expected no error, got", err)
	}
	if resp.Code != 200 {
		t.Error("Expected status 200, got", resp.Code)
	}

	scroll := resp.Body.(scroll.Scroll[*PokemonThumbnail])
	if scroll.Content == nil {
		t.Error("Expected non-nil content")
	}
	if len(scroll.Content) != 3 {
		t.Error("Expected 3 pokemon, got", len(scroll.Content))
	}
}

// func TestPokemonService_GetPokemonDetailsByName(t *testing.T) {
// 	_, service := setupPokemonService(t)

// 	// Test English name
// 	resp, err := service.GetPokemonDetailsByName(context.Background(), "Pikachu", "en")
// 	if err != nil {
// 		t.Error("Expected no error, got", err)
// 	}
// 	if resp.Code != 200 {
// 		t.Error("Expected status 200, got", resp.Code)
// 	}
// 	details := resp.Body.(*PokemonDetails)
// 	if details.Symbol != "pikachu" {
// 		t.Error("Expected symbol pikachu, got", details.Symbol)
// 	}

// 	// Test French name
// 	resp, err = service.GetPokemonDetailsByName(context.Background(), "PikachuFR", "en")
// 	if err != nil {
// 		t.Error("Expected no error, got", err)
// 	}
// 	if resp.Code != 200 {
// 		t.Error("Expected status 200, got", resp.Code)
// 	}
// 	details = resp.Body.(*PokemonDetails)
// 	if details.Symbol != "pikachu" {
// 		t.Error("Expected symbol pikachu, got", details.Symbol)
// 	}

// 	// Test Case Insensitivity
// 	resp, err = service.GetPokemonDetailsByName(context.Background(), "pikachu", "en")
// 	if err != nil {
// 		t.Error("Expected no error, got", err)
// 	}
// 	if resp.Code != 200 {
// 		t.Error("Expected status 200 for lowercase name, got", resp.Code)
// 	}

// 	// Test Symbol search (new fallback)
// 	resp, err = service.GetPokemonDetailsByName(context.Background(), "pikachu", "en")
// 	if err != nil {
// 		t.Error("Expected no error, got", err)
// 	}
// 	if resp.Code != 200 {
// 		t.Error("Expected status 200 for symbol search, got", resp.Code)
// 	}

// 	resp, err = service.GetPokemonDetailsByName(context.Background(), "PIKACHU", "en")
// 	if err != nil {
// 		t.Error("Expected no error, got", err)
// 	}
// 	if resp.Code != 200 {
// 		t.Error("Expected status 200 for uppercase name, got", resp.Code)
// 	}

// 	// Test with spaces
// 	resp, err = service.GetPokemonDetailsByName(context.Background(), "  Pikachu  ", "en")
// 	if err != nil {
// 		t.Error("Expected no error, got", err)
// 	}
// 	if resp.Code != 200 {
// 		t.Error("Expected status 200 for name with spaces, got", resp.Code)
// 	}

// 	// Test Bulbasaur
// 	resp, err = service.GetPokemonDetailsByName(context.Background(), "Bulbasaur", "en")
// 	if err != nil {
// 		t.Error("Expected no error, got", err)
// 	}
// 	if resp.Code != 200 {
// 		t.Error("Expected status 200, got", resp.Code)
// 	}

// 	// Test Smettle (English)
// 	resp, err = service.GetPokemonDetailsByName(context.Background(), "Smettle", "en")
// 	if err != nil {
// 		t.Error("Expected no error, got", err)
// 	}
// 	if resp.Code != 200 {
// 		t.Error("Expected status 200 for Smettle, got", resp.Code)
// 	}
// 	details = resp.Body.(*PokemonDetails)
// 	if details.Symbol != "smettle" {
// 		t.Error("Expected symbol smettle, got", details.Symbol)
// 	}

// 	// Test Malortie (French name for Smettle)
// 	resp, err = service.GetPokemonDetailsByName(context.Background(), "Malortie", "fr")
// 	if err != nil {
// 		t.Error("Expected no error, got", err)
// 	}
// 	if resp.Code != 200 {
// 		t.Error("Expected status 200 for Malortie, got", resp.Code)
// 	}
// 	details = resp.Body.(*PokemonDetails)
// 	if details.Symbol != "smettle" {
// 		t.Error("Expected symbol smettle for name Malortie, got", details.Symbol)
// 	}

// 	// Test Not Found
// 	resp, err = service.GetPokemonDetailsByName(context.Background(), "Mewtwo", "en")
// 	if err != nil {
// 		t.Error("Expected no error, got", err)
// 	}
// 	if resp.Code != 404 {
// 		t.Error("Expected status 404, got", resp.Code)
// 	}
// }
