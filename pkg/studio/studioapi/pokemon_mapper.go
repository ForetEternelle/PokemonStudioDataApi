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
	slog.Debug("Mapping pokemon to thumbnail", "lang", lang)
	form := p.Forms[0]
	var type2 *TypePartial
	if form.Type2 != nil {
		type2 = m.typeMapper.ToTypePartial(*form.Type2, lang)
	}
	return &PokemonThumbnail{
		Symbol: p.DbSymbol,
		Number: p.Id,
		Image:  p.DbSymbol,
		Type1:  m.typeMapper.ToTypePartial(*form.Type1, lang),
		Type2:  type2,
		Name:   p.Name[lang],
	}
}

// PokemonToDetail map a pokemon to a details transfer object
// p the pokemon to map
// lang the language expected
func (m PokemonMapper) PokemonToDetail(p studio.Pokemon, lang string) *PokemonDetails {
	slog.Debug("Mapping pokemon to details", "pokemon", p, "lang", lang)
	return &PokemonDetails{
		Symbol:      p.DbSymbol,
		Number:      p.Id,
		Name:        p.Name[lang],
		Description: p.Description[lang],
		MainForm:    *m.FormToPokemonFormDetails(p.Forms[0], lang),
	}
}

// FormToPokemonFormDetails map a pokemon form to a form details transfer object
// f the pokemon form to map
// lang the language expected
func (m PokemonMapper) FormToPokemonFormDetails(f studio.PokemonForm, lang string) *FormDetails {
	slog.Debug("Mapping pokemon form to form details", "form", f, "lang", lang)

	partialType1 := m.typeMapper.ToTypePartial(*f.Type1, lang)
	var partialType2 *TypePartial

	if f.Type2 != nil {
		partialType2 = m.typeMapper.ToTypePartial(*f.Type2, lang)
	}

	var abilityPartials []AbilityPartial
	for _, ability := range f.Abilities {
		abilityPartials = append(abilityPartials, m.abilityMapper.ToAbilityPartial(*ability, lang))
	}

	return &FormDetails{
		Form: &f.Form,

		Height: f.Height,
		Weight: f.Weight,

		Type1: partialType1,
		Type2: partialType2,

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

		ExperienceType: f.ExperienceType,
		BaseExperience: f.BaseExperience,
		BaseLoyalty:    f.BaseLoyalty,
		CatchRate:      f.CatchRate,
		FemaleRate:     f.FemaleRate,
		BreedGroups:    f.BreedGroups,
		HatchSteps:     f.HatchSteps,
		BabyDbSymbol:   f.BabyDbSymbol,
		BabyForm:       &f.BabyForm,
		Abilities:      abilityPartials,
	}
}
