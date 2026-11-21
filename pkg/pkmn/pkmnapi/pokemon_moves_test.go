package pkmnapi

import (
	"cmp"
	"context"
	"slices"
	"testing"

	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/pkmn"
	. "github.com/ForetEternelle/PokemonStudioDataApi/pkg/pkmn/pkmnapispec"
)

func setupMovesService(t *testing.T) (*pkmn.Store, PokemonAPIServicer) {
	t.Helper()

	store := pkmn.NewStore()
	store.AddType(pkmn.AddTypeDto{DbSymbol: "normal", Name: pkmn.Translation{"en": "Normal", "fr": "Normal"}})
	store.AddType(pkmn.AddTypeDto{DbSymbol: "psychic", Name: pkmn.Translation{"en": "Psychic", "fr": "Psy"}})

	store.AddMove(pkmn.AddMoveDto{ID: 1, DbSymbol: "tackle", Type: "normal", Category: "physical", Power: 40, Accuracy: 100, PP: 35, Name: pkmn.Translation{"en": "Tackle", "fr": "Charge"}, Description: pkmn.Translation{"en": "A physical attack.", "fr": "Une attaque physique."}})
	store.AddMove(pkmn.AddMoveDto{ID: 2, DbSymbol: "psychic", Type: "psychic", Category: "special", Power: 90, Accuracy: 100, PP: 10, Name: pkmn.Translation{"en": "Psychic", "fr": "Psyko"}, Description: pkmn.Translation{"en": "A psychic attack.", "fr": "Une attaque psy."}})
	store.AddMove(pkmn.AddMoveDto{ID: 3, DbSymbol: "hidden_move", Type: "normal", Category: "status", Name: pkmn.Translation{"en": "Hidden", "fr": "Cachée"}, Description: pkmn.Translation{"en": "A hidden move.", "fr": "Une capacité cachée."}})

	store.AddMachine(pkmn.AddMachineDto{
		ID:         0,
		DbSymbol:   "hm01",
		Name:       pkmn.Translation{"en": "HM01", "fr": "CS01"},
		IsHm:       true,
		MoveSymbol: "tackle",
	})
	store.AddMachine(pkmn.AddMachineDto{
		ID:         2,
		DbSymbol:   "tm99",
		Name:       pkmn.Translation{"en": "TM99", "fr": "CT99"},
		IsHm:       false,
		MoveSymbol: "psychic",
	})

	store.AddPokemon(pkmn.AddPokemonDto{
		ID:       65,
		DbSymbol: "alakazam",
		Forms: []pkmn.AddPokemonFormDto{
			{
				Form:     0,
				Type1:    "psychic",
				Name:     pkmn.Translation{"en": "Alakazam", "fr": "Alakazam"},
				Learnset: alakazamMoves(),
			},
			{
				Form:     1,
				Type1:    "psychic",
				Name:     pkmn.Translation{"en": "Alakazam Mega", "fr": "Alakazam Méga"},
				Learnset: []pkmn.PokemonFormMove{},
			},
		},
	})

	accessPolicyFactory := func(context.Context) *PokemonFilterPolicy {
		return NewPokemonFilterPolicy()
	}

	typeMapper := NewTypeMapper()
	pokemonMapper := NewPokemonMapper(typeMapper, NewAbilityMapper(), store)

	return store, NewPokemonService(store, pokemonMapper, accessPolicyFactory)
}

func alakazamMoves() []pkmn.PokemonFormMove {
	return []pkmn.PokemonFormMove{
		{MoveSymbol: "tackle", Method: pkmn.LearnMethodLevelUp, Level: 1},
		{MoveSymbol: "tackle", Method: pkmn.LearnMethodBreed},
		{MoveSymbol: "tackle", Method: pkmn.LearnMethodMachine},
		{MoveSymbol: "psychic", Method: pkmn.LearnMethodTutor},
		{MoveSymbol: "psychic", Method: pkmn.LearnMethodEvolution},
		{MoveSymbol: "unknown_move", Method: pkmn.LearnMethodTutor},
	}
}

func TestPokemonService_GetPokemonDetails_Learnset(t *testing.T) {
	_, service := setupMovesService(t)

	resp, err := service.GetPokemonDetails(context.Background(), "alakazam", 0, "fr")
	if err != nil {
		t.Fatal("Expected no error, got", err)
	}
	if resp.Code != 200 {
		t.Fatalf("Expected status 200, got %d", resp.Code)
	}

	details := resp.Body.(*PokemonDetails)
	learnset := details.Form.Learnset
	if len(learnset) != 5 {
		t.Fatalf("learnset = %d, want 5 (unknown move excluded)", len(learnset))
	}

	levelMove := learnset[0]
	if levelMove.Move.Symbol != "tackle" || levelMove.Method != string(pkmn.LearnMethodLevelUp) {
		t.Errorf("level move = %+v, want tackle through level up", levelMove)
	}
	if levelMove.Move.Name != "Charge" || levelMove.Move.Type.Symbol != "normal" {
		t.Errorf("level move = %+v, want localized tackle data", levelMove.Move)
	}
	if levelMove.Move.Description != "Une attaque physique." {
		t.Errorf("level move description = %q, want localized tackle description", levelMove.Move.Description)
	}
	if levelMove.Move.Category != "physical" || levelMove.Move.Pp != 35 || levelMove.Move.Power != 40 || levelMove.Move.Accuracy != 100 {
		t.Errorf("level move = %+v, want tackle battle fields", levelMove.Move)
	}
	if levelMove.Level == nil || *levelMove.Level != 1 {
		t.Error("level move should preserve its learning level")
	}
	if levelMove.Machine != nil {
		t.Error("level move should not have a machine")
	}

	breedMove := learnset[1]
	if breedMove.Method != string(pkmn.LearnMethodBreed) || breedMove.Level != nil {
		t.Errorf("breed move = %+v, want breed without level", breedMove)
	}
	if breedMove.Move.Symbol != "tackle" || breedMove.Move.Name != "Charge" {
		t.Errorf("breed move = %+v, want hydrated tackle", breedMove.Move)
	}

	machineMove := learnset[2]
	if machineMove.Method != string(pkmn.LearnMethodMachine) {
		t.Errorf("machine move method = %q, want machine", machineMove.Method)
	}
	if machineMove.Move.Symbol != "tackle" || machineMove.Move.Name != "Charge" {
		t.Errorf("machine move = %+v, want hydrated tackle", machineMove.Move)
	}
	if machineMove.Machine == nil {
		t.Fatal("machine move should reference its machine item")
	}
	machine := *machineMove.Machine
	if machine.Symbol != "hm01" || !machine.IsHm || machine.Name != "CS01" || machine.Id != 0 {
		t.Errorf("machine = %+v, want localized HM01 data", machine)
	}
	if machine.Number == nil || *machine.Number != 1 {
		t.Error("machine should expose the number derived from its stable symbol")
	}

	if learnset[3].Method != string(pkmn.LearnMethodTutor) || learnset[4].Method != string(pkmn.LearnMethodEvolution) {
		t.Error("tutor and evolution moves should be preserved")
	}
	if learnset[3].Move.Symbol != "psychic" || learnset[3].Move.Name != "Psyko" || learnset[3].Move.Type.Symbol != "psychic" {
		t.Errorf("tutor move = %+v, want localized psychic data", learnset[3].Move)
	}
	if learnset[3].Move.Description != "Une attaque psy." {
		t.Errorf("tutor move description = %q, want localized psychic description", learnset[3].Move.Description)
	}

	enResp, err := service.GetPokemonDetails(context.Background(), "alakazam", 0, "en")
	if err != nil {
		t.Fatal("Expected no error, got", err)
	}
	enLearnset := enResp.Body.(*PokemonDetails).Form.Learnset
	if enLearnset[0].Move.Name != "Tackle" || enLearnset[0].Move.Description != "A physical attack." {
		t.Errorf("en level move = %+v, want localized tackle data", enLearnset[0].Move)
	}
}

func TestPokemonService_GetPokemonDetails_Learnset_HiddenMove(t *testing.T) {
	store, _ := setupMovesService(t)
	policyFactory := func(context.Context) *PokemonFilterPolicy {
		return NewPokemonFilterPolicy(WithMoveFilter(func(m *pkmn.Move) bool {
			return m.DbSymbol != "tackle"
		}))
	}
	service := NewPokemonService(store, NewPokemonMapper(NewTypeMapper(), NewAbilityMapper(), store), policyFactory)

	resp, err := service.GetPokemonDetails(context.Background(), "alakazam", 0, "en")
	if err != nil {
		t.Fatal("Expected no error, got", err)
	}

	learnset := resp.Body.(*PokemonDetails).Form.Learnset
	if len(learnset) != 2 {
		t.Fatalf("learnset = %d, want 2 visible moves", len(learnset))
	}
	for _, move := range learnset {
		if move.Move.Symbol == "tackle" {
			t.Error("hidden moves should not be exposed")
		}
	}
}

func TestPokemonService_GetPokemonDetails_Learnset_MachineWithoutItem(t *testing.T) {
	store, _ := setupMovesService(t)
	store.AddPokemon(pkmn.AddPokemonDto{
		ID:       96,
		DbSymbol: "drowzee",
		Forms: []pkmn.AddPokemonFormDto{
			{
				Form:     0,
				Type1:    "psychic",
				Name:     pkmn.Translation{"en": "Drowzee"},
				Learnset: []pkmn.PokemonFormMove{{MoveSymbol: "hidden_move", Method: pkmn.LearnMethodMachine}},
			},
		},
	})

	service := NewPokemonService(store, NewPokemonMapper(NewTypeMapper(), NewAbilityMapper(), store), func(context.Context) *PokemonFilterPolicy {
		return NewPokemonFilterPolicy()
	})

	resp, err := service.GetPokemonDetails(context.Background(), "drowzee", 0, "en")
	if err != nil {
		t.Fatal("Expected no error, got", err)
	}

	learnset := resp.Body.(*PokemonDetails).Form.Learnset
	if len(learnset) != 1 {
		t.Fatalf("learnset = %d, want 1", len(learnset))
	}
	if learnset[0].Machine != nil {
		t.Error("machine move without a matching item should not invent machine data")
	}
}

func TestPokemonService_GetPokemonDetails_Learnset_Empty(t *testing.T) {
	_, service := setupMovesService(t)

	resp, err := service.GetPokemonDetails(context.Background(), "alakazam", 1, "en")
	if err != nil {
		t.Fatal("Expected no error, got", err)
	}
	if resp.Code != 200 {
		t.Fatalf("Expected status 200, got %d", resp.Code)
	}

	learnset := resp.Body.(*PokemonDetails).Form.Learnset
	if learnset == nil || len(learnset) != 0 {
		t.Error("Empty learnset should be an empty list")
	}
}

func TestPokemonService_GetPokemonDetails_Learnset_HiddenPokemonAndForm(t *testing.T) {
	store, _ := setupMovesService(t)

	pokemonPolicy := func(context.Context) *PokemonFilterPolicy {
		return NewPokemonFilterPolicy(WithPokemonFilter(func(p *pkmn.Pokemon) bool {
			return p.DbSymbol != "alakazam"
		}))
	}
	service := NewPokemonService(store, NewPokemonMapper(NewTypeMapper(), NewAbilityMapper(), store), pokemonPolicy)
	if resp, err := service.GetPokemonDetails(context.Background(), "alakazam", 0, "en"); err != nil || resp.Code != 404 {
		t.Errorf("hidden pokemon: code = %d, err = %v, want 404", resp.Code, err)
	}

	formPolicy := func(context.Context) *PokemonFilterPolicy {
		return NewPokemonFilterPolicy(WithFormFilter(func(f *pkmn.PokemonForm) bool {
			return f.Form != 0
		}))
	}
	service = NewPokemonService(store, NewPokemonMapper(NewTypeMapper(), NewAbilityMapper(), store), formPolicy)
	if resp, err := service.GetPokemonDetails(context.Background(), "alakazam", 0, "en"); err != nil || resp.Code != 404 {
		t.Errorf("hidden form: code = %d, err = %v, want 404", resp.Code, err)
	}
}

func TestMachineNumber(t *testing.T) {
	cases := map[string]struct {
		symbol string
		want   *int32
	}{
		"hm symbol":         {symbol: "hm01", want: int32Ptr(1)},
		"tm symbol":         {symbol: "tm128", want: int32Ptr(128)},
		"single digit":      {symbol: "tm5", want: int32Ptr(5)},
		"no number":         {symbol: "potion", want: nil},
		"number too large":  {symbol: "tm999999999999", want: nil},
		"number not suffix": {symbol: "tm01_beta", want: nil},
	}

	for label, tc := range cases {
		got := machineNumber(tc.symbol)
		if tc.want == nil {
			if got != nil {
				t.Errorf("%s: machineNumber(%q) = %d, want nil", label, tc.symbol, *got)
			}
			continue
		}
		if got == nil || *got != *tc.want {
			t.Errorf("%s: machineNumber(%q) = %v, want %d", label, tc.symbol, got, *tc.want)
		}
	}
}

func TestMachineToLearnsetMachine_LocalizationKeepsStableNumber(t *testing.T) {
	machines := []pkmn.Machine{
		{ID: 0, DbSymbol: "tm01", Name: pkmn.Translation{"en": "TM01", "fr": "CT01"}},
		{ID: 1, DbSymbol: "hm02", Name: pkmn.Translation{"en": "HM02", "fr": "CS02"}},
		{ID: 2, DbSymbol: "tm99", Name: pkmn.Translation{"en": "Technical Machine", "fr": "Machine Technique"}},
	}
	mapper := NewPokemonMapper(NewTypeMapper(), NewAbilityMapper(), pkmn.NewStore())

	wantNumbers := map[string]int32{"tm01": 1, "hm02": 2, "tm99": 99}
	wantOrder := []string{"tm01", "hm02", "tm99"}

	for _, lang := range []string{"en", "fr"} {
		partials := make([]LearnsetMachine, 0, len(machines))
		for _, machine := range machines {
			partials = append(partials, mapper.machineToLearnsetMachine(machine, lang))
		}

		for _, partial := range partials {
			if partial.Number == nil {
				t.Fatalf("%s: machine %q should expose a number", lang, partial.Symbol)
			}
			if want := wantNumbers[partial.Symbol]; *partial.Number != want {
				t.Errorf("%s: machine %q number = %d, want %d", lang, partial.Symbol, *partial.Number, want)
			}
		}

		slices.SortFunc(partials, func(a, b LearnsetMachine) int {
			return cmp.Compare(*a.Number, *b.Number)
		})
		for i, want := range wantOrder {
			if partials[i].Symbol != want {
				t.Errorf("%s: machine order[%d] = %q, want %q", lang, i, partials[i].Symbol, want)
			}
		}
	}
}

func TestMachineToLearnsetMachine_IgnoresLocalizedName(t *testing.T) {
	mapper := NewPokemonMapper(NewTypeMapper(), NewAbilityMapper(), pkmn.NewStore())
	machine := pkmn.Machine{
		DbSymbol: "tm_unnumbered",
		Name:     pkmn.Translation{"en": "TM 42", "fr": "CT 42"},
	}

	for _, lang := range []string{"en", "fr"} {
		if partial := mapper.machineToLearnsetMachine(machine, lang); partial.Number != nil {
			t.Errorf("%s: machine number = %d, want nil when the symbol has no number", lang, *partial.Number)
		}
	}
}

func int32Ptr(value int32) *int32 {
	return &value
}
