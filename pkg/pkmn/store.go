package pkmn

import (
	"iter"
	"log/slog"
	"slices"
	"strings"

	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/iter2"
)

const (
	StudioFolder   = "Studio"
	LanguageFolder = "Text/Dialogs"

	UndefType = "__undef__"
)

// Translation is a map of language codes to translated strings.
type Translation map[string]string

// Store is an in-memory store for a studio project
type Store struct {
	pokemonList []*Pokemon
	types       []*PokemonType
	abilities   []*Ability
	moves       []*Move

	pokemonBySymbol      map[string]*Pokemon
	pokemonTypesBySymbol map[string]*PokemonType
	abilitiesBySymbol    map[string]*Ability
	movesBySymbol        map[string]*Move

	pokemonNameTranslations []Translation
}

func NewStore() *Store {
	pokemonBySymbol := make(map[string]*Pokemon)
	pokemonTypesBySymbol := make(map[string]*PokemonType)
	abilitiesBySymbol := make(map[string]*Ability)
	movesBySymbol := make(map[string]*Move)

	return &Store{
		pokemonList:             []*Pokemon{},
		types:                   []*PokemonType{},
		abilities:               []*Ability{},
		moves:                   []*Move{},
		pokemonBySymbol:         pokemonBySymbol,
		pokemonTypesBySymbol:    pokemonTypesBySymbol,
		abilitiesBySymbol:       abilitiesBySymbol,
		movesBySymbol:           movesBySymbol,
		pokemonNameTranslations: []Translation{},
	}
}

// AddTypeDto describes a PokemonType to add to the store.
type AddTypeDto struct {
	DbSymbol string
	Color    string
	TextId   int
	Name     Translation
	DamageTo map[string]float32
}

// AddType adds a PokemonType to the store.
func (s *Store) AddType(dto AddTypeDto) *PokemonType {
	pokemonType := &PokemonType{
		DbSymbol: dto.DbSymbol,
		Color:    dto.Color,
		TextId:   dto.TextId,
		Name:     dto.Name,
		DamageTo: dto.DamageTo,
	}

	s.types = append(s.types, pokemonType)
	s.pokemonTypesBySymbol[pokemonType.DbSymbol] = pokemonType
	slog.Info("Adding pokemon type", "symbol", pokemonType.DbSymbol)
	return pokemonType
}

// AddAbilityDto describes an Ability to add to the store.
type AddAbilityDto struct {
	DbSymbol    string
	ID          int
	TextID      int
	Name        Translation
	Description Translation
}

// AddAbility adds an Ability to the store.
func (s *Store) AddAbility(dto AddAbilityDto) *Ability {
	ability := &Ability{
		DbSymbol:    dto.DbSymbol,
		ID:          dto.ID,
		TextID:      dto.TextID,
		Name:        dto.Name,
		Description: dto.Description,
	}

	s.abilities = append(s.abilities, ability)
	s.abilitiesBySymbol[ability.DbSymbol] = ability
	slog.Info("Adding ability", "symbol", ability.DbSymbol)
	return ability
}

// AddMoveDto describes a Move to add to the store.
type AddMoveDto struct {
	ID               int
	DbSymbol         string
	Type             string
	Category         MoveCategory
	Power            int
	Accuracy         int
	PP               int
	CriticalRate     int
	Priority         int
	MapUse           int
	Targeting        MoveTargeting
	Execution        MoveExecution
	MechanicalTags   []MoveMechanicalTag
	Interactions     []MoveInteraction
	SecondaryEffects MoveSecondaryEffects
	Name             Translation
	Description      Translation
}

// AddMove adds a Move to the store.
func (s *Store) AddMove(dto AddMoveDto) *Move {
	move := &Move{
		ID:               dto.ID,
		DbSymbol:         dto.DbSymbol,
		MoveType:         s.FindTypeBySymbol(dto.Type),
		Category:         dto.Category,
		Power:            dto.Power,
		Accuracy:         dto.Accuracy,
		PP:               dto.PP,
		CriticalRate:     dto.CriticalRate,
		Priority:         dto.Priority,
		MapUse:           dto.MapUse,
		Targeting:        dto.Targeting,
		Execution:        dto.Execution,
		MechanicalTags:   dto.MechanicalTags,
		Interactions:     dto.Interactions,
		SecondaryEffects: dto.SecondaryEffects,
		Name:             dto.Name,
		Description:      dto.Description,
	}

	s.moves = append(s.moves, move)
	s.movesBySymbol[move.DbSymbol] = move
	slog.Info("Adding move", "symbol", move.DbSymbol)
	return move
}

// AddPokemonDto describes a Pokemon to add to the store.
type AddPokemonDto struct {
	ID               int32
	DbSymbol         string
	Forms            []AddPokemonFormDto
	CustomProperties map[string]any
	Tags             []string
}

// AddPokemonFormDto describes a PokemonForm to add to the store.
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
	AbilitySymbols   []string
	FrontOffsetY     int32
	Name             Translation
	Description      Translation
	CustomProperties map[string]any
	Tags             []string
	Resources        PokemonResources
}

// AddPokemon adds a Pokemon to the store, keeping the list sorted by ID.
func (s *Store) AddPokemon(dto AddPokemonDto) *Pokemon {
	forms := make([]*PokemonForm, len(dto.Forms))
	for i, f := range dto.Forms {
		var type2 *PokemonType
		if f.Type2 != nil {
			type2 = s.FindTypeBySymbol(*f.Type2)
		}

		abilities := make([]*Ability, 0, len(f.AbilitySymbols))
		for _, sym := range f.AbilitySymbols {
			if a := s.FindAbilityBySymbol(sym); a != nil {
				abilities = append(abilities, a)
			}
		}

		customProperties := f.CustomProperties
		if customProperties == nil {
			customProperties = make(map[string]any)
		}

		tags := f.Tags
		if tags == nil {
			tags = make([]string, 0)
		}

		forms[i] = &PokemonForm{
			Form:             f.Form,
			Type1:            s.FindTypeBySymbol(f.Type1),
			Type2:            type2,
			Height:           f.Height,
			Weight:           f.Weight,
			BaseHp:           f.BaseHp,
			BaseAtk:          f.BaseAtk,
			BaseDfe:          f.BaseDfe,
			BaseSpd:          f.BaseSpd,
			BaseAts:          f.BaseAts,
			BaseDfs:          f.BaseDfs,
			EvHp:             f.EvHp,
			EvAtk:            f.EvAtk,
			EvDfe:            f.EvDfe,
			EvSpd:            f.EvSpd,
			EvAts:            f.EvAts,
			EvDfs:            f.EvDfs,
			Evolutions:       f.Evolutions,
			ExperienceType:   f.ExperienceType,
			BaseExperience:   f.BaseExperience,
			BaseLoyalty:      f.BaseLoyalty,
			CatchRate:        f.CatchRate,
			FemaleRate:       f.FemaleRate,
			BreedGroups:      f.BreedGroups,
			HatchSteps:       f.HatchSteps,
			BabyDbSymbol:     f.BabyDbSymbol,
			BabyForm:         f.BabyForm,
			ItemHeld:         f.ItemHeld,
			AbilitySymbols:   f.AbilitySymbols,
			Abilities:        abilities,
			FrontOffsetY:     f.FrontOffsetY,
			Name:             f.Name,
			Description:      f.Description,
			CustomProperties: customProperties,
			Tags:             tags,
			Resources:        f.Resources,
		}
	}

	slices.SortFunc(forms, func(a, b *PokemonForm) int {
		return int(a.Form - b.Form)
	})

	customProperties := dto.CustomProperties
	if customProperties == nil {
		customProperties = make(map[string]any)
	}

	tags := dto.Tags
	if tags == nil {
		tags = make([]string, 0)
	}

	pokemon := &Pokemon{
		ID:               dto.ID,
		DbSymbol:         dto.DbSymbol,
		Forms:            forms,
		CustomProperties: customProperties,
		Tags:             tags,
	}

	insertIndex, _ := slices.BinarySearchFunc(s.pokemonList, pokemon, func(a, b *Pokemon) int {
		return int(a.ID) - int(b.ID)
	})

	s.pokemonList = slices.Insert(s.pokemonList, insertIndex, pokemon)
	s.pokemonBySymbol[pokemon.DbSymbol] = pokemon

	translationIt := iter2.Map(slices.Values(forms), func(form *PokemonForm) Translation {
		return form.Name
	})
	s.pokemonNameTranslations = append(s.pokemonNameTranslations, slices.Collect(translationIt)...)

	slog.Info("Adding pokemon", "symbol", pokemon.DbSymbol)
	return pokemon
}

func (s *Store) FindAllPokemon(filters ...iter2.FilterFunc[*Pokemon]) iter.Seq[*Pokemon] {
	return func(yield func(*Pokemon) bool) {
		for _, p := range s.pokemonList {
			if iter2.And(filters...)(p) {
				if !yield(p) {
					return
				}
			}
		}
	}
}

func (s *Store) FindPokemonBySymbol(symbol string, filters ...iter2.FilterFunc[*Pokemon]) *Pokemon {
	pokemon, ok := s.pokemonBySymbol[symbol]
	if !ok {
		return nil
	}

	if !iter2.And(filters...)(pokemon) {
		return nil
	}

	return pokemon
}

func (s *Store) FindPokemonByName(name string, filters ...iter2.FilterFunc[*Pokemon]) *Pokemon {
	normalizedName := strings.ToLower(strings.TrimSpace(name))

	pokemonIter := s.FindAllPokemon(filters...)
	found, foundOk := iter2.First(iter2.Filter(pokemonIter, func(p *Pokemon) bool {
		// Check symbol
		if strings.ToLower(p.DbSymbol) == normalizedName {
			return true
		}

		// Check all names in forms
		for _, form := range p.Forms {
			for _, val := range form.Name {
				if strings.ToLower(strings.TrimSpace(val)) == normalizedName {
					return true
				}
			}
		}

		// Check the name translations as a fallback
		id := int(p.ID)
		if id >= 0 && id < len(s.pokemonNameTranslations) {
			translation := s.pokemonNameTranslations[id]
			for _, val := range translation {
				if strings.ToLower(strings.TrimSpace(val)) == normalizedName {
					return true
				}
			}
		}

		return false
	}))

	if foundOk {
		return s.pokemonBySymbol[found.DbSymbol]
	}

	return nil
}

func (s *Store) FindAllTypes(filters ...iter2.FilterFunc[*PokemonType]) iter.Seq[*PokemonType] {
	typeIt := slices.Values(s.types)
	return iter2.Filter(typeIt, iter2.And(filters...))
}

func (s *Store) FindTypeBySymbol(symbol string, filters ...iter2.FilterFunc[*PokemonType]) *PokemonType {
	pokemonType, ok := s.pokemonTypesBySymbol[symbol]
	if !ok {
		return nil
	}

	if !iter2.And(filters...)(pokemonType) {
		return nil
	}

	return pokemonType
}

// Resistances calculates the type resistances of the PokemonForm based on its types.
func (s *Store) Resistances(type1, type2 string, filters ...iter2.FilterFunc[*PokemonType]) iter.Seq2[string, float32] {
	typeIt := s.FindAllTypes(filters...)
	ptResIT := iter2.ToSeq2(typeIt, func(pt *PokemonType) float32 {
		return pt.DamageToTypes(type1, type2)
	})

	return iter2.Map2(ptResIT, func(pt *PokemonType, res float32) (string, float32) {
		return pt.DbSymbol, res
	})
}

func (s *Store) FindAllAbilities(filters ...iter2.FilterFunc[*Ability]) iter.Seq[*Ability] {
	abilityIt := slices.Values(s.abilities)
	return iter2.Filter(abilityIt, iter2.And(filters...))
}

func (s *Store) FindAbilityBySymbol(symbol string, filters ...iter2.FilterFunc[*Ability]) *Ability {
	ability, ok := s.abilitiesBySymbol[symbol]
	if !ok {
		return nil
	}

	if !iter2.And(filters...)(ability) {
		return nil
	}

	return ability
}

func (s *Store) FindAllMoves(filters ...iter2.FilterFunc[*Move]) iter.Seq[*Move] {
	moveIt := slices.Values(s.moves)
	return iter2.Filter(moveIt, iter2.And(filters...))
}

func (s *Store) FindMoveBySymbol(symbol string, filters ...iter2.FilterFunc[*Move]) *Move {
	move, ok := s.movesBySymbol[symbol]
	if !ok {
		return nil
	}

	if !iter2.And(filters...)(move) {
		return nil
	}

	return move
}
