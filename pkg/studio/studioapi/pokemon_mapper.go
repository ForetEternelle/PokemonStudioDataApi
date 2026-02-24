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

func NewPokemonMapper(
	typeMapper *TypeMapper,
	abilityMapper *AbilityMapper,
	store *studio.Store,
) *PokemonMapper {
	return &PokemonMapper{
		typeMapper,
		abilityMapper,
		store,
	}
}

func (m PokemonMapper) PokemonToThumbnail(p studio.Pokemon, lang string, policy *AccessPolicy) *PokemonThumbnail {
	slog.Debug("Mapping pokemon to thumbnail", "lang", lang)

	// Get the main form of the pokemon
	form, ok := p.Forms[0]
	if !ok {
		slog.Warn("Pokemon has no main form", "pokemon", p)
		return nil
	}

	// Check if the form is allowed by the policy
	for _, filter := range policy.FormFilters {
		if !filter(form) {
			ok = false
			return nil
		}
	}

	var type2 *TypePartial
	if form.Type2 != nil {
		type2 = m.typeMapper.ToTypePartial(*form.Type2, lang, policy)
	}

	return &PokemonThumbnail{
		Symbol:           p.DbSymbol,
		Number:           p.Id,
		Image:            p.DbSymbol,
		Type1:            m.typeMapper.ToTypePartial(*form.Type1, lang, policy),
		Type2:            type2,
		Name:             p.Name[lang],
		CustomProperties: p.CustomProperties,
	}
}

func (m PokemonMapper) PokemonToDetail(p studio.Pokemon, lang string, policy *AccessPolicy) *PokemonDetails {
	slog.Debug("Mapping pokemon to details", "pokemon", p, "lang", lang)

	// Get the main form of the pokemon
	form, ok := p.Forms[0]
	if !ok {
		slog.Warn("Pokemon has no main form", "pokemon", p)
		return nil
	}

	// Check if the form is allowed by the policy
	for _, filter := range policy.FormFilters {
		if !filter(form) {
			ok = false
			return nil
		}
	}

	return &PokemonDetails{
		Symbol:      p.DbSymbol,
		Number:      p.Id,
		Name:        p.Name[lang],
		Description: p.Description[lang],
		MainForm:    *m.FormToPokemonFormDetails(form, lang, policy),
	}
}

func (m PokemonMapper) FormToPokemonFormDetails(f studio.PokemonForm, lang string, policy *AccessPolicy) *FormDetails {
	slog.Debug("Mapping pokemon form to form details", "form", f, "lang", lang)

	partialType1 := m.typeMapper.ToTypePartial(*f.Type1, lang, policy)
	var partialType2 *TypePartial

	if f.Type2 != nil {
		partialType2 = m.typeMapper.ToTypePartial(*f.Type2, lang, policy)
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
