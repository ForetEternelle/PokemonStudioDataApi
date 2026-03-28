package studioapi

import (
	"testing"

	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/studio"
)

func TestMoveToDetail(t *testing.T) {
	lang := "test"
	fireType := studio.NewTypeBuilder().DbSymbol("fire").Build()
	move := studio.NewMoveBuilder().
		DbSymbol("testDbSymbol").
		Name(studio.Translation{lang: "testName"}).
		Description(studio.Translation{lang: "testDescription"}).
		Power(80).
		Accuracy(100).
		PP(15).
		Type(fireType).
		Category("special").
		Build()

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

	pt := move.Type()
	if moveDetail.Type.Symbol != pt.DbSymbol() {
		t.Error("Mapper should map type symbol, expected", pt.DbSymbol(), ", has", moveDetail.Type.Symbol)
	}

	if moveDetail.Category != string(move.Category()) {
		t.Error("Mapper should map category, expected", move.Category(), ", has", moveDetail.Category)
	}
}

func TestMoveToPartial(t *testing.T) {
	lang := "test"
	fireType := studio.NewTypeBuilder().DbSymbol("fire").Build()
	move := studio.NewMoveBuilder().
		DbSymbol("testDbSymbol").
		Name(studio.Translation{lang: "testName"}).
		Description(studio.Translation{lang: "testDescription"}).
		Power(80).
		Accuracy(100).
		PP(15).
		Type(fireType).
		Category("special").
		Build()

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
