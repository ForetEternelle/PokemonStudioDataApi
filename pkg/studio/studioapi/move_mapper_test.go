package studioapi

import (
	"testing"

	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/studio"
)

func TestMoveToDetail(t *testing.T) {
	lang := "test"
	fireType := studio.NewPokemonType(studio.WithPokemonTypeDbSymbol("fire"))
	move := studio.NewMove(
		studio.WithMoveDbSymbol("testDbSymbol"),
		studio.WithMoveName(studio.Translation{lang: "testName"}),
		studio.WithMoveDescription(studio.Translation{lang: "testDescription"}),
		studio.WithMovePower(80),
		studio.WithMoveAccuracy(100),
		studio.WithMovePP(15),
		studio.WithMoveType(fireType),
		studio.WithMoveCategory("special"),
	)

	typeMapper := NewTypeMapper()
	moveMapper := NewMoveMapper(typeMapper)
	policy := NewAccessPolicy()
	moveDetail := moveMapper.ToMoveDetail(*move, lang, policy)

	if moveDetail.Name != move.Name(lang) {
		t.Error("Mapper should map name, expected", move.Name(lang), ", has", moveDetail.Name)
	}

	if moveDetail.Symbol != move.DbSymbol() {
		t.Error("Mapper should map db symbol, expected", move.DbSymbol(), ", has", moveDetail.Symbol)
	}

	if moveDetail.Description != move.Description(lang) {
		t.Error("Mapper should map description, expected", move.Description(lang), ", has", moveDetail.Description)
	}

	if moveDetail.Power != int32(move.Power()) {
		t.Error("Mapper should map power, expected", move.Power(), ", has", moveDetail.Power)
	}

	if moveDetail.Accuracy != int32(move.Accuracy()) {
		t.Error("Mapper should map accuracy, expected", move.Accuracy(), ", has", moveDetail.Accuracy)
	}

	if moveDetail.Pp != int32(move.PP()) {
		t.Error("Mapper should map pp, expected", move.PP(), ", has", moveDetail.Pp)
	}

	if moveDetail.Type.Symbol != move.Type().DbSymbol() {
		t.Error("Mapper should map type symbol, expected", move.Type().DbSymbol(), ", has", moveDetail.Type.Symbol)
	}

	if moveDetail.Category != string(move.Category()) {
		t.Error("Mapper should map category, expected", move.Category(), ", has", moveDetail.Category)
	}
}

func TestMoveToPartial(t *testing.T) {
	lang := "test"
	fireType := studio.NewPokemonType(studio.WithPokemonTypeDbSymbol("fire"))
	move := studio.NewMove(
		studio.WithMoveDbSymbol("testDbSymbol"),
		studio.WithMoveName(studio.Translation{lang: "testName"}),
		studio.WithMoveDescription(studio.Translation{lang: "testDescription"}),
		studio.WithMovePower(80),
		studio.WithMoveAccuracy(100),
		studio.WithMovePP(15),
		studio.WithMoveType(fireType),
		studio.WithMoveCategory("special"),
	)

	typeMapper := NewTypeMapper()
	moveMapper := NewMoveMapper(typeMapper)
	policy := NewAccessPolicy()
	movePartial := moveMapper.ToMovePartial(*move, lang, policy)

	if movePartial.Name != move.Name(lang) {
		t.Error("Mapper should map name, expected", move.Name(lang), ", has", movePartial.Name)
	}

	if movePartial.Symbol != move.DbSymbol() {
		t.Error("Mapper should map db symbol, expected", move.DbSymbol(), ", has", movePartial.Symbol)
	}
}
