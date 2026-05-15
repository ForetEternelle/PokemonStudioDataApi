package pkmnapi

import (
	"testing"

	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/pkmn"
)

func TestToAbilityDetail(t *testing.T) {
	lang := "test"
	ability := pkmn.NewAbilityBuilder().
		DbSymbol("testDbSymbol").
		Name(pkmn.Translation{lang: "testName"}).
		Description(pkmn.Translation{lang: "testDescription"}).
		Build()

	abilityMapper := NewAbilityMapper()
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
	ability := pkmn.NewAbilityBuilder().
		DbSymbol("testDbSymbol").
		Name(pkmn.Translation{lang: "testName"}).
		Description(pkmn.Translation{lang: "testDescription"}).
		Build()

	abilityMapper := NewAbilityMapper()
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
