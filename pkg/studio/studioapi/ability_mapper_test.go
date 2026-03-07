package studioapi_test

import (
	"testing"

	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/studio"
	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/studio/studioapi"
)

func TestToAbilityDetail(t *testing.T) {
	lang := "test"
	ability := studio.NewAbility(
		studio.WithAbilityDbSymbol("testDbSymbol"),
		studio.WithAbilityName(studio.Translation{lang: "testName"}),
		studio.WithAbilityDescription(studio.Translation{lang: "testDescription"}),
	)

	abilityMapper := studioapi.NewAbilityMapper()
	abilityDetail := abilityMapper.ToAbilityDetail(*ability, lang)

	if abilityDetail.Name != ability.Name(lang) {
		t.Error("Mapper should map name, expected", ability.Name(lang), ", has", abilityDetail.Name)
	}

	if abilityDetail.Symbol != ability.DbSymbol() {
		t.Error("Mapper should map db symbol, expected", ability.DbSymbol(), ", has", abilityDetail.Symbol)
	}

	if abilityDetail.Description != ability.Description(lang) {
		t.Error("Mapper should map description, expected", ability.Description(lang), ", has", abilityDetail.Description)
	}
}

func TestToAbilityPartial(t *testing.T) {
	lang := "test"
	ability := studio.NewAbility(
		studio.WithAbilityDbSymbol("testDbSymbol"),
		studio.WithAbilityName(studio.Translation{lang: "testName"}),
		studio.WithAbilityDescription(studio.Translation{lang: "testDescription"}),
	)

	abilityMapper := studioapi.NewAbilityMapper()
	abilityPartial := abilityMapper.ToAbilityPartial(*ability, lang)

	if abilityPartial.Name != ability.Name(lang) {
		t.Error("Mapper should map name, expected", ability.Name(lang), ", has", abilityPartial.Name)
	}

	if abilityPartial.Symbol != ability.DbSymbol() {
		t.Error("Mapper should map db symbol, expected", ability.DbSymbol(), ", has", abilityPartial.Symbol)
	}

	if abilityPartial.Description != ability.Description(lang) {
		t.Error("Mapper should map description, expected", ability.Description(lang), ", has", abilityPartial.Description)
	}
}
