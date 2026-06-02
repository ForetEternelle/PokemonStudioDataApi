package pkmn

import (
	"log/slog"
	"slices"

	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/iter2"
)

type AddPokemonDto struct {
	ID       int32
	DbSymbol string
	Forms    []AddPokemonFormDto
	CustomProperties map[string]any
}

type AddPokemonFormDto struct {
	Form             int32
	Type1            string
	Type2            *string
	Height           float32
	Weight           float32
	BaseHp           int32
	BaseAtk          int32
	BaseDfe          int32
	BaseSpd          int32
	BaseAts          int32
	BaseDfs          int32
	EvHp             int32
	EvAtk            int32
	EvDfe            int32
	EvSpd            int32
	EvAts            int32
	EvDfs            int32
	Evolutions       []Evolution
	ExperienceType   string
	BaseExperience   int32
	BaseLoyalty      int32
	CatchRate        int32
	FemaleRate       float32
	BreedGroups      []string
	HatchSteps       int32
	BabyDbSymbol     *string
	BabyForm         int32
	ItemHeld         []*ItemHeld
	Abilities        []string
	FrontOffsetY     int32
	Name             Translation
	Description      Translation
	CustomProperties map[string]any
	Resources        PokemonResources
}

func (s *Store) AddPokemon(dto AddPokemonDto) *Pokemon {
	forms := make([]*PokemonForm, len(dto.Forms))
	for i, f := range dto.Forms {
		var type2 *PokemonType
		if f.Type2 != nil {
			type2 = s.FindTypeBySymbol(*f.Type2)
		}

		abilities := make([]*Ability, 0, len(f.Abilities))
		for _, sym := range f.Abilities {
			if a := s.FindAbilityBySymbol(sym); a != nil {
				abilities = append(abilities, a)
			}
		}

		forms[i] = &PokemonForm{
			form:             f.Form,
			type1:            s.FindTypeBySymbol(f.Type1),
			type2:            type2,
			height:           f.Height,
			weight:           f.Weight,
			baseHp:           f.BaseHp,
			baseAtk:          f.BaseAtk,
			baseDfe:          f.BaseDfe,
			baseSpd:          f.BaseSpd,
			baseAts:          f.BaseAts,
			baseDfs:          f.BaseDfs,
			evHp:             f.EvHp,
			evAtk:            f.EvAtk,
			evDfe:            f.EvDfe,
			evSpd:            f.EvSpd,
			evAts:            f.EvAts,
			evDfs:            f.EvDfs,
			evolutions:       f.Evolutions,
			experienceType:   f.ExperienceType,
			baseExperience:   f.BaseExperience,
			baseLoyalty:      f.BaseLoyalty,
			catchRate:        f.CatchRate,
			femaleRate:       f.FemaleRate,
			breedGroups:      f.BreedGroups,
			hatchSteps:       f.HatchSteps,
			babyDbSymbol:     f.BabyDbSymbol,
			babyForm:         f.BabyForm,
			itemHeld:         f.ItemHeld,
			abilitySymbols:   f.Abilities,
			abilities:        abilities,
			frontOffsetY:     f.FrontOffsetY,
			name:             f.Name,
			description:      f.Description,
			customProperties: f.CustomProperties,
			resources:        f.Resources,
		}
	}

	pokemon := &Pokemon{
		id:               dto.ID,
		dbSymbol:         dto.DbSymbol,
		forms:            forms,
		customProperties: dto.CustomProperties,
	}

	insertIndex, _ := slices.BinarySearchFunc(s.pokemonList, pokemon, func(a, b *Pokemon) int {
		return int(a.id) - int(b.id)
	})

	s.pokemonList = slices.Insert(s.pokemonList, insertIndex, pokemon)
	s.pokemonBySymbol[pokemon.dbSymbol] = pokemon

	translationIt := iter2.Map(slices.Values(forms), func(form *PokemonForm) Translation {
		return form.name
	})
	s.pokemonNameTranslations = append(s.pokemonNameTranslations, slices.Collect(translationIt)...)

	slog.Info("Adding pokemon", "symbol", pokemon.dbSymbol)
	return pokemon
}
