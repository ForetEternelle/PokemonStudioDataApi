package studio

import (
	"os"
	"slices"
	"testing"

	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/pkmn"
)

const (
	TechItemValid = "../../../test/test_resources/valid-data/Studio/items/tm01.json"
)

func TestMapLearnableMoves_AllMethods(t *testing.T) {
	mapped := mapLearnableMoves([]LearnableMoveDescriptor{
		{Klass: "LevelLearnableMove", Move: "tackle", Level: 1},
		{Klass: "TutorLearnableMove", Move: "ice_punch"},
		{Klass: "TechLearnableMove", Move: "swords_dance"},
		{Klass: "BreedLearnableMove", Move: "fire_punch"},
		{Klass: "EvolutionLearnableMove", Move: "psychic"},
	})

	expected := []struct {
		method pkmn.LearnMethod
		level  int32
	}{
		{pkmn.LearnMethodLevelUp, 1},
		{pkmn.LearnMethodTutor, 0},
		{pkmn.LearnMethodMachine, 0},
		{pkmn.LearnMethodBreed, 0},
		{pkmn.LearnMethodEvolution, 0},
	}

	if len(mapped) != len(expected) {
		t.Fatalf("mapped count = %d, want %d", len(mapped), len(expected))
	}

	for i, exp := range expected {
		if mapped[i].Method != exp.method {
			t.Errorf("mapped[%d].Method = %q, want %q", i, mapped[i].Method, exp.method)
		}
		if mapped[i].Level != exp.level {
			t.Errorf("mapped[%d].Level = %d, want %d", i, mapped[i].Level, exp.level)
		}
	}
}

func TestMapLearnableMoves_PreservesOrderAndDuplicates(t *testing.T) {
	mapped := mapLearnableMoves([]LearnableMoveDescriptor{
		{Klass: "LevelLearnableMove", Move: "tackle", Level: 1},
		{Klass: "LevelLearnableMove", Move: "tackle", Level: 1},
		{Klass: "BreedLearnableMove", Move: "tackle"},
	})

	if len(mapped) != 3 {
		t.Fatalf("mapped count = %d, want 3", len(mapped))
	}
	if mapped[0].MoveSymbol != "tackle" || mapped[2].Method != pkmn.LearnMethodBreed {
		t.Error("mapLearnableMoves should preserve source order and duplicates")
	}
}

func TestMapLearnableMoves_UnknownClass(t *testing.T) {
	mapped := mapLearnableMoves([]LearnableMoveDescriptor{
		{Klass: "UnknownLearnableMove", Move: "tackle"},
	})

	if len(mapped) != 0 {
		t.Error("mapLearnableMoves should skip unsupported classes")
	}
}

func TestMapLearnableMoves_Empty(t *testing.T) {
	if mapped := mapLearnableMoves(nil); mapped != nil {
		t.Error("mapLearnableMoves should return nil for no moves")
	}
}

func TestUnmarshalPokemonDescriptor_LearnableMoves(t *testing.T) {
	content, err := os.ReadFile(PokemonValid)
	if err != nil {
		t.Fatal("Error reading test file", "file", PokemonValid)
	}

	descriptor, err := UnmarshalPokemonDescriptor(content)
	if err != nil {
		t.Fatal("Unmarshal valid pokemon descriptor should not return error", err)
	}

	methods := make([]string, len(descriptor.Forms[0].MoveSet))
	for i, move := range descriptor.Forms[0].MoveSet {
		methods[i] = move.Klass
	}

	for _, method := range []string{
		"LevelLearnableMove",
		"TutorLearnableMove",
		"TechLearnableMove",
		"BreedLearnableMove",
		"EvolutionLearnableMove",
	} {
		if !slices.Contains(methods, method) {
			t.Errorf("UnmarshalPokemonDescriptor should preserve %s entries", method)
		}
	}
}

func TestUnmarshalTechItemDescriptor_Ok(t *testing.T) {
	content, err := os.ReadFile(TechItemValid)
	if err != nil {
		t.Fatal("Error reading test file", "file", TechItemValid)
	}

	descriptor, err := UnmarshalTechItemDescriptor(content)
	if err != nil {
		t.Fatal("Unmarshal valid tech item descriptor should not return error", err)
	}
	if descriptor.DbSymbol != "tm01" || descriptor.Move != "swords_dance" || descriptor.IsHm {
		t.Error("UnmarshalTechItemDescriptor should preserve item data")
	}
}

func TestImportTechItems_OnlyTechItems(t *testing.T) {
	descriptors, err := ImportTechItems(DataFolder+"/Studio", DataFolder+"/Text/Dialogs")
	if err != nil {
		t.Fatal("ImportTechItems should succeed", err)
	}

	items := slices.Collect(descriptors)
	if len(items) != 3 {
		t.Fatalf("ImportTechItems should import 3 tech items, has %d", len(items))
	}

	var hm, tm *TechItemDescriptor
	for _, item := range items {
		switch item.DbSymbol {
		case "hm01":
			hm = item
		case "tm99":
			tm = item
		}
	}

	if hm == nil || !hm.IsHm || hm.Move != "cut" {
		t.Error("ImportTechItems should import HM data")
	}
	if hm != nil && (hm.Name["en"] != "HM01" || hm.Name["fr"] != "CS01") {
		t.Error("ImportTechItems should import localized item names")
	}
	if tm == nil || tm.IsHm || tm.Move != "ice_beam" || tm.ID != 2 {
		t.Error("ImportTechItems should import TM data")
	}
}

func TestImportTechItems_MissingFolder(t *testing.T) {
	_, err := ImportTechItems(InvalidPath, InvalidPath)
	if err == nil {
		t.Error("ImportTechItems at an invalid path should return an error")
	}
}
