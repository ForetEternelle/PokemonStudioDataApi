package studioapi_test

import (
	"testing"

	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/studio"
	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/studio/studioapi"
)

func TestGetRouter_WithoutStore(t *testing.T) {
	_, err := studioapi.GetRouter()
	if err == nil {
		t.Error("Expected error when store is not provided")
	}
}

func TestGetRouter_WithStore(t *testing.T) {
	store := studio.NewStore()
	router, err := studioapi.GetRouter(studioapi.WithStore(store))
	if err != nil {
		t.Error("Expected no error, got", err)
	}
	if router == nil {
		t.Error("Expected router to be not nil")
	}
}
