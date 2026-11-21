package pkmnapi

import (
	"log/slog"
	"slices"
	"strconv"

	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/iter2"
	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/pkmn"
	. "github.com/ForetEternelle/PokemonStudioDataApi/pkg/pkmn/pkmnapispec"
)

type PokemonMapper struct {
	typeMapper    *TypeMapper
	abilityMapper *AbilityMapper
	moveMapper    *MoveMapper
	store         *pkmn.Store
}

func NewPokemonMapper(
	typeMapper *TypeMapper,
	abilityMapper *AbilityMapper,
	store *pkmn.Store,
) *PokemonMapper {
	return &PokemonMapper{
		typeMapper:    typeMapper,
		abilityMapper: abilityMapper,
		moveMapper:    NewMoveMapper(typeMapper),
		store:         store,
	}
}

func (m PokemonMapper) PokemonToThumbnail(p pkmn.Pokemon, formId int32, lang string, policy PokemonFilterPolicy) *PokemonThumbnail {
	slog.Debug("Mapping pokemon to thumbnail", "lang", lang)

	form, okForm := p.Form(formId)

	if !okForm {
		return nil
	}

	thumbnail := &PokemonThumbnail{
		Symbol:           p.DbSymbol,
		Number:           p.ID,
		Form:             formId,
		Image:            p.DbSymbol,
		Type1:            m.typeMapper.ToTypePartial(*form.Type1, lang, policy),
		Name:             form.Name[lang],
		CustomProperties: form.CustomProperties,
		Tags:             form.Tags,
	}

	if form.Type2 != nil {
		thumbnail.Type2 = m.typeMapper.ToTypePartial(*form.Type2, lang, policy)
	}

	return thumbnail
}

func (m PokemonMapper) PokemonToDetail(p pkmn.Pokemon, formId int32, lang string, policy PokemonFilterPolicy) *PokemonDetails {
	slog.Debug("Mapping pokemon to details", "pokemon", p.DbSymbol, "lang", lang)

	formFilter := iter2.And(policy.FormFilter)
	f, ok := p.Form(formId)

	if !ok || !formFilter(f) {
		return nil
	}

	return &PokemonDetails{
		Symbol:           p.DbSymbol,
		Number:           p.ID,
		Form:             *m.FormToPokemonFormDetails(*f, lang, policy),
		CustomProperties: f.CustomProperties,
		Tags:             f.Tags,
	}
}

func (m PokemonMapper) FormToPokemonFormDetails(f pkmn.PokemonForm, lang string, policy PokemonFilterPolicy) *FormDetails {
	slog.Debug("Mapping pokemon form to form details", "form", f.Form, "lang", lang)

	abilityIt := slices.Values(f.Abilities)
	abilityIt = iter2.Filter(abilityIt, policy.AbilityFilter)

	abilityPartialIt := iter2.Map(abilityIt, func(a *pkmn.Ability) AbilityPartial {
		return m.abilityMapper.ToAbilityPartial(*a, lang)
	})

	partialType1 := m.typeMapper.ToTypePartial(*f.Type1, lang, policy)
	var partialType2 *TypePartial
	if f.Type2 != nil {
		partialType2 = m.typeMapper.ToTypePartial(*f.Type2, lang, policy)
	}

	babyForm := f.BabyForm

	return &FormDetails{
		Number:      f.Form,
		Name:        f.Name[lang],
		Description: f.Description[lang],
		Height:      f.Height,
		Weight:      f.Weight,

		Type1: partialType1,
		Type2: partialType2,

		BaseHp:  f.BaseHp,
		BaseAtk: f.BaseAtk,
		BaseDfe: f.BaseDfe,
		BaseSpd: f.BaseSpd,
		BaseAts: f.BaseAts,
		BaseDfs: f.BaseDfs,

		EvHp:  f.EvHp,
		EvAtk: f.EvAtk,
		EvDfe: f.EvDfe,
		EvSpd: f.EvSpd,
		EvAts: f.EvAts,
		EvDfs: f.EvDfs,

		ExperienceType:   f.ExperienceType,
		BaseExperience:   f.BaseExperience,
		BaseLoyalty:      f.BaseLoyalty,
		CatchRate:        f.CatchRate,
		FemaleRate:       f.FemaleRate,
		HatchSteps:       f.HatchSteps,
		BabyDbSymbol:     f.BabyDbSymbol,
		BabyForm:         &babyForm,
		Abilities:        slices.Collect(abilityPartialIt),
		Learnset:         m.FormToLearnset(f, lang, policy),
		CustomProperties: f.CustomProperties,
		Tags:             f.Tags,
	}
}

func (m PokemonMapper) FormToLearnset(f pkmn.PokemonForm, lang string, policy PokemonFilterPolicy) []LearnsetEntry {
	slog.Debug("Mapping pokemon form to moves", "form", f.Form, "lang", lang)

	entries := make([]LearnsetEntry, 0, len(f.Learnset))
	for _, learnable := range f.Learnset {
		if learnable.Move == nil || !policy.MoveFilter(learnable.Move) {
			continue
		}

		entry := LearnsetEntry{
			Move:   m.moveMapper.ToMovePartial(*learnable.Move, lang, policy),
			Method: string(learnable.Method),
		}

		if learnable.Method == pkmn.LearnMethodLevelUp {
			level := learnable.Level
			entry.Level = &level
		}

		if learnable.Method == pkmn.LearnMethodMachine && learnable.Machine != nil {
			machine := m.machineToLearnsetMachine(*learnable.Machine, lang)
			entry.Machine = &machine
		}

		entries = append(entries, entry)
	}

	return entries
}

func (m PokemonMapper) machineToLearnsetMachine(machine pkmn.Machine, lang string) LearnsetMachine {
	partial := LearnsetMachine{
		Id:     int32(machine.ID),
		Symbol: machine.DbSymbol,
		Name:   machine.Name[lang],
		IsHm:   machine.IsHm,
		Number: machineNumber(machine.DbSymbol),
	}

	return partial
}

// machineNumber extracts the machine number from the stable item symbol (for example tm01 or hm01), nil when there is none.
func machineNumber(symbol string) *int32 {
	end := len(symbol)
	start := end
	for start > 0 && symbol[start-1] >= '0' && symbol[start-1] <= '9' {
		start--
	}
	if start == end {
		return nil
	}

	number, err := strconv.ParseInt(symbol[start:end], 10, 32)
	if err != nil {
		return nil
	}

	machineNumber := int32(number)
	return &machineNumber
}

func (m PokemonMapper) FormToPokemonFormPartial(f pkmn.PokemonForm, lang string, policy PokemonFilterPolicy) *FormPartial {
	slog.Debug("Mapping pokemon form to form partial", "form", f.Form, "lang", lang)

	partialType1 := m.typeMapper.ToTypePartial(*f.Type1, lang, policy)
	var partialType2 *TypePartial
	if f.Type2 != nil {
		partialType2 = m.typeMapper.ToTypePartial(*f.Type2, lang, policy)
	}

	return &FormPartial{
		Number: f.Form,

		Name: f.Name[lang],

		Type1: partialType1,
		Type2: partialType2,

		BaseHp:           f.BaseHp,
		BaseAtk:          f.BaseAtk,
		BaseDfe:          f.BaseDfe,
		BaseSpd:          f.BaseSpd,
		BaseAts:          f.BaseAts,
		BaseDfs:          f.BaseDfs,
		CustomProperties: f.CustomProperties,
		Tags:             f.Tags,
	}
}
