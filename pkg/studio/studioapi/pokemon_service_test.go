package studioapi_test

import (
	"context"
	"testing"

	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/iter2"
	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/pagination"
	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/studio"
	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/studio/studioapi"
)

func setupPokemonService(
	pokemonCtxFilter studioapi.ContextFilter[studio.Pokemon],
	formCtxFilter studioapi.ContextFilter[studio.PokemonForm],
) (*studio.Store, studioapi.PokemonAPIServicer) {
	store := studio.NewStore()
	normalType := studio.PokemonType{DbSymbol: "normal", Name: studio.Translation{"en": "Normal"}}
	electricType := studio.PokemonType{DbSymbol: "electric", Name: studio.Translation{"en": "Electric"}}
	store.AddType(normalType)
	store.AddType(electricType)

	pikachu := studio.Pokemon{
		Id:          25,
		DbSymbol:    "pikachu",
		Name:        studio.Translation{"en": "Pikachu"},
		Description: studio.Translation{"en": "Electric mouse"},
		Forms: map[int32]studio.PokemonForm{
			0: {Form: 0, Type1: &electricType, BaseHp: 35, BaseAtk: 55},
			1: {Form: 1, Type1: &electricType, BaseHp: 20, BaseAtk: 40},
		},
	}
	store.AddPokemon(pikachu)

	bulbasaur := studio.Pokemon{
		Id:          1,
		DbSymbol:    "bulbasaur",
		Name:        studio.Translation{"en": "Bulbasaur"},
		Description: studio.Translation{"en": "Grass starter"},
		Forms: map[int32]studio.PokemonForm{
			0: {Form: 0, Type1: &normalType, BaseHp: 45, BaseAtk: 49},
		},
	}
	store.AddPokemon(bulbasaur)

	typeMapper := studioapi.NewTypeMapper()
	abilityMapper := studioapi.NewAbilityMapper()
	pokemonMapper := studioapi.NewPokemonMapper(typeMapper, abilityMapper, store)

	if pokemonCtxFilter == nil {
		pokemonCtxFilter = studioapi.DefaultPokemonContextFilter
	}
	if formCtxFilter == nil {
		formCtxFilter = studioapi.DefaultFormContextFilter
	}

	service := studioapi.NewPokemonService(store, pokemonMapper, pokemonCtxFilter, formCtxFilter)
	return store, service
}

func TestPokemonService_GetPokemonDetails(t *testing.T) {
	_, service := setupPokemonService(studioapi.DefaultPokemonContextFilter, studioapi.DefaultFormContextFilter)

	resp, err := service.GetPokemonDetails(context.Background(), "pikachu", "en")
	if err != nil {
		t.Error("Expected no error, got", err)
	}
	if resp.Code != 200 {
		t.Error("Expected status 200, got", resp.Code)
	}

	details := resp.Body.(*studioapi.PokemonDetails)
	if details.Symbol != "pikachu" {
		t.Error("Expected symbol pikachu, got", details.Symbol)
	}
	if details.Name != "Pikachu" {
		t.Error("Expected name Pikachu, got", details.Name)
	}
}

func TestPokemonService_GetPokemonDetails_NotFound(t *testing.T) {
	_, service := setupPokemonService(studioapi.DefaultPokemonContextFilter, studioapi.DefaultFormContextFilter)

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
	_, service := setupPokemonService(studioapi.DefaultPokemonContextFilter, studioapi.DefaultFormContextFilter)

	resp, err := service.GetPokemon(context.Background(), 0, 10, "en")
	if err != nil {
		t.Error("Expected no error, got", err)
	}
	if resp.Code != 200 {
		t.Error("Expected status 200, got", resp.Code)
	}

	page := resp.Body.(pagination.Page[*studioapi.PokemonThumbnail])
	if page.Content == nil {
		t.Error("Expected non-nil content")
	}
	if len(page.Content) != 2 {
		t.Error("Expected 2 pokemon, got", len(page.Content))
	}
}

func TestPokemonService_GetPokemon_Pagination(t *testing.T) {
	store := studio.NewStore()
	normalType := studio.PokemonType{DbSymbol: "normal"}
	store.AddType(normalType)

	for i := 1; i <= 15; i++ {
		store.AddPokemon(studio.Pokemon{
			Id:       int32(i),
			DbSymbol: "pokemon_" + string(rune(i)),
			Name:     studio.Translation{"en": "Pokemon"},
			Forms: map[int32]studio.PokemonForm{
				0: {Form: 0, Type1: &normalType},
			},
		})
	}

	typeMapper := studioapi.NewTypeMapper()
	abilityMapper := studioapi.NewAbilityMapper()
	pokemonMapper := studioapi.NewPokemonMapper(typeMapper, abilityMapper, store)

	service := studioapi.NewPokemonService(store, pokemonMapper, studioapi.DefaultPokemonContextFilter, studioapi.DefaultFormContextFilter)

	resp, _ := service.GetPokemon(context.Background(), 0, 5, "en")
	page := resp.Body.(pagination.Page[*studioapi.PokemonThumbnail])

	if len(page.Content) != 5 {
		t.Error("Expected 5 items per page, got", len(page.Content))
	}
	if page.Total != 15 {
		t.Error("Expected total 15, got", page.Total)
	}
}

func TestPokemonService_GetPokemonForm(t *testing.T) {
	_, service := setupPokemonService(studioapi.DefaultPokemonContextFilter, studioapi.DefaultFormContextFilter)

	resp, err := service.GetPokemonForm(context.Background(), "pikachu", 0, "en")
	if err != nil {
		t.Error("Expected no error, got", err)
	}
	if resp.Code != 200 {
		t.Error("Expected status 200, got", resp.Code)
	}

	form := resp.Body.(*studioapi.FormDetails)
	if form.Form == nil {
		t.Error("Expected non-nil form")
	}
}

func TestPokemonService_GetPokemonForm_NotFound(t *testing.T) {
	_, service := setupPokemonService(studioapi.DefaultPokemonContextFilter, studioapi.DefaultFormContextFilter)

	resp, err := service.GetPokemonForm(context.Background(), "pikachu", 99, "en")
	if err != nil {
		t.Error("Expected no error, got", err)
	}
	if resp.Code != 404 {
		t.Error("Expected status 404 for non-existent form, got", resp.Code)
	}
}

func TestPokemonService_WithPokemonContextFilter(t *testing.T) {
	pokemonCtxFilter := func(ctx context.Context) iter2.FilterFunc[studio.Pokemon] {
		return func(p studio.Pokemon) bool {
			return p.DbSymbol == "pikachu"
		}
	}
	_, service := setupPokemonService(pokemonCtxFilter, studioapi.DefaultFormContextFilter)

	resp, _ := service.GetPokemon(context.Background(), 0, 10, "en")
	page := resp.Body.(pagination.Page[*studioapi.PokemonThumbnail])

	if len(page.Content) != 1 {
		t.Error("Expected 1 pokemon after filter, got", len(page.Content))
	}
}

func TestPokemonService_WithFormContextFilter(t *testing.T) {
	store := studio.NewStore()
	electricType := studio.PokemonType{DbSymbol: "electric"}
	store.AddType(electricType)

	store.AddPokemon(studio.Pokemon{
		Id:       1,
		DbSymbol: "pikachu",
		Name:     studio.Translation{"en": "Pikachu"},
		Forms: map[int32]studio.PokemonForm{
			0: {Form: 0, Type1: &electricType},
			1: {Form: 1, Type1: &electricType},
		},
	})

	typeMapper := studioapi.NewTypeMapper()
	abilityMapper := studioapi.NewAbilityMapper()
	pokemonMapper := studioapi.NewPokemonMapper(typeMapper, abilityMapper, store)

	formCtxFilter := func(ctx context.Context) iter2.FilterFunc[studio.PokemonForm] {
		return func(f studio.PokemonForm) bool {
			return f.Form == 0
		}
	}

	service := studioapi.NewPokemonService(store, pokemonMapper, studioapi.DefaultPokemonContextFilter, formCtxFilter)

	resp, _ := service.GetPokemonForm(context.Background(), "pikachu", 0, "en")
	if resp.Code != 200 {
		t.Error("Expected status 200 for form 0, got", resp.Code)
	}

	resp, _ = service.GetPokemonForm(context.Background(), "pikachu", 1, "en")
	if resp.Code != 404 {
		t.Error("Expected status 404 for filtered form 1, got", resp.Code)
	}
}
