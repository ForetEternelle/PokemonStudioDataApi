package studio

import (
	"iter"
	"log/slog"
	"maps"
	"path"
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
	pokemonList []Pokemon
	types       []PokemonType
	abilities   []Ability
	moves       []Move

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
		pokemonList:             []Pokemon{},
		types:                   []PokemonType{},
		abilities:               []Ability{},
		moves:                   []Move{},
		pokemonBySymbol:         pokemonBySymbol,
		pokemonTypesBySymbol:    pokemonTypesBySymbol,
		abilitiesBySymbol:       abilitiesBySymbol,
		movesBySymbol:           movesBySymbol,
		pokemonNameTranslations: []Translation{},
	}
}

func Load(folder string) (*Store, error) {
	store := NewStore()
	translationFolder := path.Join(folder, LanguageFolder)
	studioFolder := path.Join(folder, StudioFolder)

	typeMapper := NewTypeMapper(store)
	abilityMapper := NewAbilityMapper(store)
	pokemonMapper := NewPokemonMapper(store)
	moveMapper := NewMoveMapper(store)

	typeIterator, err := ImportTypes(studioFolder, translationFolder)
	if err != nil {
		slog.Error("Failed to load pokemon types")
		return nil, err
	}
	for descriptor := range typeIterator {
		pokemonType := typeMapper.MapPokemonTypeDescriptorToPokemonType(*descriptor)
		store.AddType(*pokemonType)
	}

	abilityIterator, err := ImportAbility(studioFolder, translationFolder)
	if err != nil {
		slog.Error("Failed to load abilities")
		return nil, err
	}
	for descriptor := range abilityIterator {
		ability := abilityMapper.MapAbilityDescriptorToAbility(*descriptor)
		store.AddAbility(*ability)
	}

	moveIterator, err := ImportMoves(studioFolder, translationFolder)
	if err != nil {
		slog.Error("Failed to load moves")
		return nil, err
	}
	for descriptor := range moveIterator {
		move := moveMapper.MapMoveDescriptorToMove(*descriptor)
		store.AddMove(*move)
	}

	pokemonIterator, err := ImportPokemon(studioFolder, translationFolder)
	if err != nil {
		slog.Error("Failed to load pokemon")
		return nil, err
	}
	for descriptor := range pokemonIterator {
		pokemon := pokemonMapper.MapPokemonDescriptorToPokemon(*descriptor)
		store.AddPokemon(*pokemon)
	}

	return store, nil
}

func (s *Store) AddPokemon(pokemon Pokemon) *Pokemon {
	insertIndex := len(s.pokemonList)
	for i, existingPokemon := range s.pokemonList {
		if pokemon.ID() < existingPokemon.ID() {
			insertIndex = i
			break
		}
	}

	s.pokemonList = slices.Insert(s.pokemonList, insertIndex, pokemon)
	s.pokemonBySymbol[pokemon.DbSymbol()] = &pokemon

	formIt := iter2.Values(pokemon.Forms())
	translationIt := iter2.Map(func(form PokemonForm) Translation {
		return form.name
	}, formIt)

	s.pokemonNameTranslations = append(s.pokemonNameTranslations, slices.Collect(translationIt)...)

	slog.Info("Adding pokemon", "symbol", pokemon.DbSymbol())
	return &pokemon
}

func (s *Store) AddType(pokemonType PokemonType) *PokemonType {
	s.types = append(s.types, pokemonType)
	s.pokemonTypesBySymbol[pokemonType.DbSymbol()] = &pokemonType
	slog.Info("Adding pokemon type", "symbol", pokemonType.DbSymbol())
	return &pokemonType
}

func (s *Store) AddAbility(ability Ability) *Ability {
	s.abilities = append(s.abilities, ability)
	s.abilitiesBySymbol[ability.DbSymbol()] = &ability
	slog.Info("Adding ability", "symbol", ability.DbSymbol())
	return &ability
}

func (s *Store) AddMove(move Move) *Move {
	s.moves = append(s.moves, move)
	s.movesBySymbol[move.DbSymbol()] = &move
	slog.Info("Adding move", "symbol", move.DbSymbol())
	return &move
}

func (s *Store) FindAllPokemon(filters ...iter2.FilterFunc[Pokemon]) iter.Seq[Pokemon] {
	it := slices.Values(s.pokemonList)
	return iter2.Filter(iter2.And(filters...), it)
}

func (s *Store) FindPokemonBySymbol(symbol string, filters ...iter2.FilterFunc[Pokemon]) *Pokemon {
	pokemon, ok := s.pokemonBySymbol[symbol]
	if !ok {
		return nil
	}

	if !iter2.And(filters...)(*pokemon) {
		return nil
	}

	return pokemon
}

func (s *Store) FindPokemonByName(name string, filters ...iter2.FilterFunc[Pokemon]) *Pokemon {
	normalizedName := strings.ToLower(strings.TrimSpace(name))

	pokemonIter := s.FindAllPokemon(filters...)
	found, foundOk := iter2.First(iter2.Filter(func(p Pokemon) bool {
		// Check symbol
		if strings.ToLower(p.DbSymbol()) == normalizedName {
			return true
		}

		// Check all names in forms
		for form := range iter2.Values(p.Forms()) {
			for val := range maps.Values(form.name) {
				if strings.ToLower(strings.TrimSpace(val)) == normalizedName {
					return true
				}
			}
		}

		// Check the name translations as a fallback
		id := int(p.ID())
		if id >= 0 && id < len(s.pokemonNameTranslations) {
			translation := s.pokemonNameTranslations[id]
			for val := range maps.Values(translation) {
				if strings.ToLower(strings.TrimSpace(val)) == normalizedName {
					return true
				}
			}
		}

		return false
	}, pokemonIter))

	if foundOk {
		return s.pokemonBySymbol[found.DbSymbol()]
	}

	return nil
}

func (s *Store) FindAllTypes(filters ...iter2.FilterFunc[PokemonType]) iter.Seq[PokemonType] {
	it := slices.Values(s.types)
	return iter2.Filter(iter2.And(filters...), it)
}

func (s *Store) FindTypeBySymbol(symbol string, filters ...iter2.FilterFunc[PokemonType]) *PokemonType {
	pokemonType, ok := s.pokemonTypesBySymbol[symbol]
	if !ok {
		return nil
	}

	if !iter2.And(filters...)(*pokemonType) {
		return nil
	}

	return pokemonType
}

// Resistances calculates the type resistances of the PokemonForm based on its types.
func (s *Store) Resistances(type1, type2 string, filters ...iter2.FilterFunc[PokemonType]) iter.Seq2[string, float32] {
	typeIt := s.FindAllTypes(filters...)
	ptResIT := iter2.ToSeq2(typeIt, func(pt PokemonType) float32 {
		return pt.DamageToTypes(type1, type2)
	})

	return iter2.Map2(func(pt PokemonType, res float32) (string, float32) {
		return pt.DbSymbol(), res
	}, ptResIT)
}

func (s *Store) FindAllAbilities(filters ...iter2.FilterFunc[Ability]) iter.Seq[Ability] {
	it := slices.Values(s.abilities)
	return iter2.Filter(iter2.And(filters...), it)
}

func (s *Store) FindAbilityBySymbol(symbol string, filters ...iter2.FilterFunc[Ability]) *Ability {
	ability, ok := s.abilitiesBySymbol[symbol]
	if !ok {
		return nil
	}

	if !iter2.And(filters...)(*ability) {
		return nil
	}

	return ability
}

func (s *Store) FindAllMoves(filters ...iter2.FilterFunc[Move]) iter.Seq[Move] {
	it := slices.Values(s.moves)
	return iter2.Filter(iter2.And(filters...), it)
}

func (s *Store) FindMoveBySymbol(symbol string, filters ...iter2.FilterFunc[Move]) *Move {
	move, ok := s.movesBySymbol[symbol]
	if !ok {
		return nil
	}

	if !iter2.And(filters...)(*move) {
		return nil
	}

	return move
}
