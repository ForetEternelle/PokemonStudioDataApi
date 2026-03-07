package studioapi

import (
	"log/slog"
	"slices"

	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/iter2"
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

	var mainForm studio.PokemonForm
	hasForm := false
	for _, form := range p.Forms() {
		passesFilter := true
		for _, filter := range policy.FormFilters {
			if !filter(form) {
				passesFilter = false
				break
			}
		}
		if passesFilter {
			mainForm = form
			hasForm = true
			break
		}
	}

	if !hasForm {
		return nil
	}

	thumbnail := &PokemonThumbnail{
		Symbol: p.DbSymbol(),
		Number: p.ID(),
		Image:  p.DbSymbol(),
		Type1:  m.typeMapper.ToTypePartial(mainForm.Type1(), lang, policy),
		Name:   p.Name(lang),
	}
	var type2, ok = mainForm.Type2()
	if ok {
		thumbnail.Type2 = m.typeMapper.ToTypePartial(type2, lang, policy)
	}

	return thumbnail
}

func (m PokemonMapper) PokemonToDetail(p studio.Pokemon, lang string, policy *AccessPolicy) *PokemonDetails {
	slog.Debug("Mapping pokemon to details", "pokemon", p, "lang", lang)

	var mainForm studio.PokemonForm
	hasForm := false
	for _, form := range p.Forms() {
		passesFilter := true
		for _, filter := range policy.FormFilters {
			if !filter(form) {
				passesFilter = false
				break
			}
		}
		if passesFilter {
			mainForm = form
			hasForm = true
			break
		}
	}

	if !hasForm {
		return nil
	}

	return &PokemonDetails{
		Symbol:      p.DbSymbol(),
		Number:      p.ID(),
		Name:        p.Name(lang),
		Description: p.Description(lang),
		MainForm:    *m.FormToPokemonFormDetails(mainForm, lang, policy),
	}
}

func (m PokemonMapper) FormToPokemonFormDetails(f studio.PokemonForm, lang string, policy *AccessPolicy) *FormDetails {
	slog.Debug("Mapping pokemon form to form details", "form", f, "lang", lang)

	var filteredAbilities []studio.Ability
	for _, a := range f.Abilities() {
		passesFilter := true
		for _, filter := range policy.AbilityFilters {
			if !filter(a) {
				passesFilter = false
				break
			}
		}
		if passesFilter {
			filteredAbilities = append(filteredAbilities, a)
		}
	}

	partialType1 := m.typeMapper.ToTypePartial(f.Type1(), lang, policy)
	var partialType2 *TypePartial
	type2, ok := f.Type2()
	if ok {
		partialType2 = m.typeMapper.ToTypePartial(type2, lang, policy)
	}

	abilityPartials := iter2.Map(func(a studio.Ability) AbilityPartial {
		return m.abilityMapper.ToAbilityPartial(a, lang)
	}, slices.Values(filteredAbilities))

	form := f.Form()
	babyForm := f.BabyForm()

	return &FormDetails{
		Form: &form,

		Height: f.Height(),
		Weight: f.Weight(),

		Type1: partialType1,
		Type2: partialType2,

		BaseHp:  f.BaseHp(),
		BaseAtk: f.BaseAtk(),
		BaseDfe: f.BaseDfe(),
		BaseSpd: f.BaseSpd(),
		BaseAts: f.BaseAts(),
		BaseDfs: f.BaseDfs(),

		EvHp: f.EvHp(),
		EvAtk: f.EvAtk(),
		EvDfe: f.EvDfe(),
		EvSpd: f.EvSpd(),
		EvAts: f.EvAts(),
		EvDfs: f.EvDfs(),

		ExperienceType: f.ExperienceType(),
		BaseExperience: f.BaseExperience(),
		BaseLoyalty:    f.BaseLoyalty(),
		CatchRate:      f.CatchRate(),
		FemaleRate:     f.FemaleRate(),
		HatchSteps:     f.HatchSteps(),
		BabyDbSymbol:   f.BabyDbSymbol(),
		BabyForm:       &babyForm,
		Abilities:      slices.Collect(abilityPartials),
	}
}
