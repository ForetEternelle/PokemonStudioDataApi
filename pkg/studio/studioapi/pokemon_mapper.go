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

	mainForm, okMainForm := p.Form(0)

	if !okMainForm {
		return nil
	}

	thumbnail := &PokemonThumbnail{
		Symbol: p.DbSymbol(),
		Number: p.ID(),
		Image:  p.DbSymbol(),
		Type1:  m.typeMapper.ToTypePartial(mainForm.Type1(), lang, policy),
		Name:   mainForm.Name(lang),
	}

	var type2, okType2 = mainForm.Type2()
	if okType2 {
		thumbnail.Type2 = m.typeMapper.ToTypePartial(type2, lang, policy)
	}

	return thumbnail
}

func (m PokemonMapper) PokemonToDetail(p studio.Pokemon, lang string, policy *AccessPolicy) *PokemonDetails {
	slog.Debug("Mapping pokemon to details", "pokemon", p.DbSymbol(), "lang", lang)

	var mainForm studio.PokemonForm
	hasForm := false
	formFilter := iter2.And(policy.FormFilter)
	for _, form := range p.Forms() {
		if formFilter(form) {
			mainForm = form
			hasForm = true
			break
		}
	}

	if !hasForm {
		return nil
	}

	return &PokemonDetails{
		Symbol:   p.DbSymbol(),
		Number:   p.ID(),
		MainForm: *m.FormToPokemonFormDetails(mainForm, lang, policy),
	}
}

func (m PokemonMapper) FormToPokemonFormDetails(f studio.PokemonForm, lang string, policy *AccessPolicy) *FormDetails {
	slog.Debug("Mapping pokemon form to form details", "form", f.Form(), "lang", lang)

	abilityIt := f.Abilities()
	abilityIt = iter2.Filter(policy.AbilityFilter, abilityIt)

	abilityPartialIt := iter2.Map(func(a studio.Ability) AbilityPartial {
		return m.abilityMapper.ToAbilityPartial(a, lang)
	}, abilityIt)

	partialType1 := m.typeMapper.ToTypePartial(f.Type1(), lang, policy)
	var partialType2 *TypePartial
	type2, ok := f.Type2()
	if ok {
		partialType2 = m.typeMapper.ToTypePartial(type2, lang, policy)
	}

	form := f.Form()
	babyForm := f.BabyForm()

	return &FormDetails{
		Form: &form,

		Name:        f.Name(lang),
		Description: f.Description(lang),
		Height:      f.Height(),
		Weight:      f.Weight(),

		Type1: partialType1,
		Type2: partialType2,

		BaseHp:  f.BaseHp(),
		BaseAtk: f.BaseAtk(),
		BaseDfe: f.BaseDfe(),
		BaseSpd: f.BaseSpd(),
		BaseAts: f.BaseAts(),
		BaseDfs: f.BaseDfs(),

		EvHp:  f.EvHp(),
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
		Abilities:      slices.Collect(abilityPartialIt),
	}
}

func (m PokemonMapper) FormToPokemonFormPartial(f studio.PokemonForm, lang string, policy *AccessPolicy) *FormPartial {
	slog.Debug("Mapping pokemon form to form partial", "form", f.Form(), "lang", lang)

	partialType1 := m.typeMapper.ToTypePartial(f.Type1(), lang, policy)
	var partialType2 *TypePartial
	type2, ok := f.Type2()
	if ok {
		partialType2 = m.typeMapper.ToTypePartial(type2, lang, policy)
	}

	form := f.Form()
	return &FormPartial{
		Form: &form,

		Name: f.Name(lang),

		Type1: partialType1,
		Type2: partialType2,

		BaseHp:  f.BaseHp(),
		BaseAtk: f.BaseAtk(),
		BaseDfe: f.BaseDfe(),
		BaseSpd: f.BaseSpd(),
		BaseAts: f.BaseAts(),
		BaseDfs: f.BaseDfs(),
	}
}
