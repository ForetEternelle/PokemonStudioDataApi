package studioapi

import (
	"log/slog"

	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/studio"
)

type PokemonMapper struct {
	typeMapper    *TypeMapper
	abilityMapper *AbilityMapper
	store         *studio.Store
}

// NewPokemonMapper Create a new pokemon mapper
// typeMapper the mapper for pokemon types
// typeStore the store for pokemon types
func NewPokemonMapper(typeMapper *TypeMapper, abilityMapper *AbilityMapper, store *studio.Store) *PokemonMapper {
	return &PokemonMapper{
		typeMapper,
		abilityMapper,
		store,
	}
}

// PokemonToThumbnail map a pokemon to a thumbnail transfer object
// p the pokemon to map
// lang the language expected
func (m PokemonMapper) PokemonToThumbnail(p studio.Pokemon, lang string) *PokemonThumbnail {
	slog.Debug("Mapping pokemon to thumbnail")
	return &PokemonThumbnail{
		Symbol: p.DbSymbol,
		Number: p.Id,
		Image:  p.Forms[0].Resources.Front,
		Name:   p.Forms[0].Name[lang],
	}
}

// PokemonToDetail map a pokemon to a details transfer object
// p the pokemon to map
// lang the language expected
func (m PokemonMapper) PokemonToDetail(p studio.Pokemon, lang string) *PokemonDetails {
	slog.Debug("Mapping pokemon to details")
	return &PokemonDetails{
		Symbol:   p.DbSymbol,
		Number:   p.Id,
		MainForm: m.FormToPokemonFormDetails(p.Forms[0], lang),
	}
}

// FormToPokemonFormDetails map a pokemon form to a form details transfer object
// p the pokemon form to map
// lang the language expected
func (m PokemonMapper) FormToPokemonFormDetails(f studio.PokemonForm, lang string) FormDetails {
	slog.Debug("Mapping pokemon form to form details")
	var breedGroups []string
	for _, breedGroup := range f.BreedGroups {
		breedGroups = append(breedGroups, studio.BreedMap[breedGroup])
	}

	var partialType2 TypePartial

	type1 := m.store.FindTypeBySymbol(f.Type1)
	partialType1 := m.typeMapper.ToTypePartial(*type1, lang)

	if f.Type2 != nil {
		type2 := m.store.FindTypeBySymbol(*f.Type2)
		if type2 != nil {
			partialType2 = m.typeMapper.ToTypePartial(*type2, lang)
		}
	}

	var abilityPartials []AbilityPartial
	for _, abilitySymbol := range f.Abilities {
		ability := m.store.FindAbilityBySymbol(abilitySymbol)
		if ability != nil && !contains(abilityPartials, ability.DbSymbol) {
			abilityPartials = append(abilityPartials, m.abilityMapper.ToAbilityPartial(*ability, lang))
		}
	}

	return FormDetails{
		Form:        &f.Form,
		Name:        f.Name[lang],
		Description: f.Description[lang],

		Height: f.Height,
		Weight: f.Weight,

		Type1: &partialType1,
		Type2: &partialType2,

		BaseHp:  f.BaseHp,
		BaseAtk: f.BaseAtk,
		BaseDfe: f.BaseDfe,
		BaseSpd: f.BaseSpd,
		BaseAts: f.BaseAts,
		BaseDfs: f.BaseDfs,

		EvHp:  &f.EvHp,
		EvAtk: &f.EvAtk,
		EvDfe: &f.EvDfe,
		EvSpd: &f.EvSpd,
		EvAts: &f.EvAts,
		EvDfs: &f.EvDfs,

		ExperienceType: studio.ExperienceTypeMap[f.ExperienceType],
		BaseExperience: f.BaseExperience,
		BaseLoyalty:    f.BaseLoyalty,
		CatchRate:      f.CatchRate,
		FemaleRate:     f.FemaleRate,
		BreedGroups:    breedGroups,
		HatchSteps:     f.HatchSteps,
		BabyDbSymbol:   f.BabyDbSymbol,
		BabyForm:       &f.BabyForm,
		Abilities:      abilityPartials,
	}
}

func contains(partials []AbilityPartial, symbol string) bool {
	for _, p := range partials {
		if p.Symbol == symbol {
			return true
		}
	}
	return false
}
