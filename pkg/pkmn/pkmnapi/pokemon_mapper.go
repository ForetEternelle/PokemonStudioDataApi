package pkmnapi

import (
	"log/slog"
	"slices"

	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/iter2"
	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/pkmn"
	. "github.com/ForetEternelle/PokemonStudioDataApi/pkg/pkmn/pkmnapispec"
)

type PokemonMapper struct {
	typeMapper    *TypeMapper
	abilityMapper *AbilityMapper
	store         *pkmn.Store
}

func NewPokemonMapper(
	typeMapper *TypeMapper,
	abilityMapper *AbilityMapper,
	store *pkmn.Store,
) *PokemonMapper {
	return &PokemonMapper{
		typeMapper,
		abilityMapper,
		store,
	}
}

func (m PokemonMapper) PokemonToThumbnail(p pkmn.Pokemon, formId int32, lang string, policy PokemonFilterPolicy) *PokemonThumbnail {
	slog.Debug("Mapping pokemon to thumbnail", "lang", lang)

	form, okForm := p.Form(formId)

	if !okForm {
		return nil
	}

	thumbnail := &PokemonThumbnail{
		Symbol: p.DbSymbol(),
		Number: p.ID(),
		Form: formId,
		Image:  p.DbSymbol(),
		Type1:  m.typeMapper.ToTypePartial(form.Type1(), lang, policy),
		Name:   form.Name(lang),
	}

	var type2, okType2 = form.Type2()
	if okType2 {
		thumbnail.Type2 = m.typeMapper.ToTypePartial(type2, lang, policy)
	}

	return thumbnail
}

func (m PokemonMapper) PokemonToDetail(p pkmn.Pokemon, formId int32, lang string, policy PokemonFilterPolicy) *PokemonDetails {
	slog.Debug("Mapping pokemon to details", "pokemon", p.DbSymbol(), "lang", lang)

	
	formFilter := iter2.And(policy.FormFilter)
	f, ok := p.Form(formId)

	if !ok || !formFilter(f){
		return nil
	}

	return &PokemonDetails{
		Symbol:   p.DbSymbol(),
		Number:   p.ID(),
		Form: *m.FormToPokemonFormDetails(*f, lang, policy),
	}
}

func (m PokemonMapper) FormToPokemonFormDetails(f pkmn.PokemonForm, lang string, policy PokemonFilterPolicy) *FormDetails {
	slog.Debug("Mapping pokemon form to form details", "form", f.Form(), "lang", lang)

	abilityIt := f.Abilities()
	abilityIt = iter2.Filter(abilityIt, policy.AbilityFilter)

	abilityPartialIt := iter2.Map(abilityIt, func(a *pkmn.Ability) AbilityPartial {
		return m.abilityMapper.ToAbilityPartial(*a, lang)
	})

	partialType1 := m.typeMapper.ToTypePartial(f.Type1(), lang, policy)
	var partialType2 *TypePartial
	type2, ok := f.Type2()
	if ok {
		partialType2 = m.typeMapper.ToTypePartial(type2, lang, policy)
	}

	form := f.Form()
	babyForm := f.BabyForm()

	return &FormDetails{
		Number: form,
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

func (m PokemonMapper) FormToPokemonFormPartial(f pkmn.PokemonForm, lang string, policy PokemonFilterPolicy) *FormPartial {
	slog.Debug("Mapping pokemon form to form partial", "form", f.Form(), "lang", lang)

	partialType1 := m.typeMapper.ToTypePartial(f.Type1(), lang, policy)
	var partialType2 *TypePartial
	type2, ok := f.Type2()
	if ok {
		partialType2 = m.typeMapper.ToTypePartial(type2, lang, policy)
	}

	form := f.Form()
	return &FormPartial{
		Number: form,

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
