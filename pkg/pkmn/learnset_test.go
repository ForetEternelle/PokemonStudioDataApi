package pkmn

import "testing"

func TestAddMachine(t *testing.T) {
	store := NewStore()

	first := store.AddMachine(AddMachineDto{
		ID:         0,
		DbSymbol:   "hm01",
		Name:       Translation{"en": "HM01", "fr": "CS01"},
		IsHm:       true,
		MoveSymbol: "cut",
	})
	store.AddMachine(AddMachineDto{
		ID:         1,
		DbSymbol:   "tm01",
		Name:       Translation{"en": "TM01", "fr": "CT01"},
		IsHm:       false,
		MoveSymbol: "swords_dance",
	})

	if found := store.FindMachineBySymbol("hm01"); found != first {
		t.Error("FindMachineBySymbol should return the added machine")
	}
	if found := store.FindMachineByMoveSymbol("cut"); found != first {
		t.Error("FindMachineByMoveSymbol should return the added machine")
	}
	if found := store.FindMachineBySymbol("unknown"); found != nil {
		t.Error("FindMachineBySymbol should return nil for an unknown symbol")
	}
	if found := store.FindMachineByMoveSymbol("unknown"); found != nil {
		t.Error("FindMachineByMoveSymbol should return nil for an unknown move")
	}
}

func TestAddMachine_IndexesLatestMachineForMove(t *testing.T) {
	store := NewStore()

	store.AddMachine(AddMachineDto{DbSymbol: "hm01", MoveSymbol: "cut"})
	last := store.AddMachine(AddMachineDto{DbSymbol: "tm01", MoveSymbol: "cut"})

	if found := store.FindMachineByMoveSymbol("cut"); found != last {
		t.Error("FindMachineByMoveSymbol should return the latest machine teaching the move")
	}
}

func TestAddPokemon_ResolvesLearnableMoves(t *testing.T) {
	store := NewStore()
	store.AddMove(AddMoveDto{ID: 1, DbSymbol: "tackle"})
	store.AddMove(AddMoveDto{ID: 2, DbSymbol: "growl"})
	store.AddMachine(AddMachineDto{DbSymbol: "tm01", Name: Translation{"en": "TM01"}, MoveSymbol: "growl"})

	store.AddPokemon(AddPokemonDto{
		ID:       1,
		DbSymbol: "bulbasaur",
		Forms: []AddPokemonFormDto{
			{
				Form: 0,
				Learnset: []PokemonFormMove{
					{MoveSymbol: "tackle", Method: LearnMethodLevelUp, Level: 1},
					{MoveSymbol: "tackle", Method: LearnMethodBreed},
					{MoveSymbol: "growl", Method: LearnMethodMachine},
					{MoveSymbol: "unknown", Method: LearnMethodTutor},
				},
			},
		},
	})

	pokemon := store.FindPokemonBySymbol("bulbasaur")
	if pokemon == nil {
		t.Fatal("FindPokemonBySymbol should return the added pokemon")
	}

	learnset := pokemon.Forms[0].Learnset
	if len(learnset) != 4 {
		t.Fatalf("learnset count = %d, want 4", len(learnset))
	}

	if learnset[0].Move == nil || learnset[0].Move.DbSymbol != "tackle" {
		t.Error("level-up move should be resolved")
	}
	if learnset[0].Level != 1 || learnset[0].Method != LearnMethodLevelUp {
		t.Error("level-up move should preserve method and level")
	}
	if learnset[1].Move != learnset[0].Move {
		t.Error("the same move learned through two methods should be preserved twice")
	}
	if learnset[2].Machine == nil || learnset[2].Machine.DbSymbol != "tm01" {
		t.Error("machine move should be resolved to its machine item")
	}
	if learnset[3].Move != nil {
		t.Error("unknown move symbol should not resolve to a move")
	}
}
