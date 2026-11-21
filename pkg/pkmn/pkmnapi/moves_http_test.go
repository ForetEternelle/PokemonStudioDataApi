package pkmnapi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/ForetEternelle/PokemonStudioDataApi/pkg/pkmn/pkmnapispec"
)

func TestGetPokemonDetails_Learnset_HTTP(t *testing.T) {
	store, _ := setupMovesService(t)
	router, err := GetRouter(WithStore(store))
	if err != nil {
		t.Fatal("GetRouter should succeed", err)
	}

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, httptest.NewRequest(http.MethodGet, "/pokemon/alakazam/forms/0?lang=fr", nil))

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", recorder.Code)
	}

	var details PokemonDetails
	if err := json.Unmarshal(recorder.Body.Bytes(), &details); err != nil {
		t.Fatal("Response should be pokemon details", err)
	}

	learnset := details.Form.Learnset
	if len(learnset) != 5 {
		t.Fatalf("learnset = %d, want 5", len(learnset))
	}

	for _, method := range []string{"level_up", "tutor", "machine", "breed", "evolution"} {
		found := false
		for _, move := range learnset {
			if move.Method == method {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("form moves should expose a %s move", method)
		}
	}

	if learnset[0].Move.Symbol != "tackle" || learnset[0].Move.Name != "Charge" || learnset[0].Move.Type.Symbol != "normal" {
		t.Errorf("learnset[0].move = %+v, want fully hydrated localized tackle", learnset[0].Move)
	}
	if learnset[0].Move.Category != "physical" || learnset[0].Move.Pp != 35 || learnset[0].Move.Power != 40 || learnset[0].Move.Accuracy != 100 {
		t.Errorf("learnset[0].move = %+v, want tackle battle fields", learnset[0].Move)
	}
	if learnset[0].Move.Description != "Une attaque physique." {
		t.Errorf("learnset[0].move.description = %q, want localized tackle description", learnset[0].Move.Description)
	}
	if learnset[0].Level == nil || *learnset[0].Level != 1 {
		t.Error("level-up move should expose its level")
	}
	if learnset[0].Machine != nil {
		t.Error("level-up move should not expose a machine")
	}
	if learnset[2].Machine == nil || learnset[2].Machine.Name != "CS01" || learnset[2].Machine.Number == nil || *learnset[2].Machine.Number != 1 {
		t.Errorf("machine move = %+v, want localized machine data", learnset[2])
	}
}

func TestGetMoves_HTTP_Removed(t *testing.T) {
	store, _ := setupMovesService(t)
	router, err := GetRouter(WithStore(store))
	if err != nil {
		t.Fatal("GetRouter should succeed", err)
	}

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, httptest.NewRequest(http.MethodGet, "/moves?lang=fr", nil))

	if recorder.Code != http.StatusNotFound {
		t.Errorf("status = %d, want 404 for the removed bulk moves route", recorder.Code)
	}
}
