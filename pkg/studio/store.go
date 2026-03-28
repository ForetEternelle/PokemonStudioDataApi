package studio

import (
	"iter"
	"log/slog"
	"path"
	"slices"

	. "github.com/ForetEternelle/PokemonStudioDataApi/pkg/iter2"
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
}

func NewStore() *Store {
	pokemonBySymbol := make(map[string]*Pokemon)
	pokemonTypesBySymbol := make(map[string]*PokemonType)
	abilitiesBySymbol := make(map[string]*Ability)
	movesBySymbol := make(map[string]*Move)

	return &Store{
		pokemonList:          []Pokemon{},
		types:                []PokemonType{},
		abilities:            []Ability{},
		moves:                []Move{},
		pokemonBySymbol:      pokemonBySymbol,
		pokemonTypesBySymbol: pokemonTypesBySymbol,
		abilitiesBySymbol:    abilitiesBySymbol,
		movesBySymbol:        movesBySymbol,
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

func (s *Store) FindAllPokemon(filters ...FilterFunc[Pokemon]) iter.Seq[Pokemon] {
	it := slices.Values(s.pokemonList)
	return Filter(And(filters...), it)
}

func (s *Store) FindPokemonBySymbol(symbol string, filters ...FilterFunc[Pokemon]) *Pokemon {
	pokemon, ok := s.pokemonBySymbol[symbol]
	if !ok {
		return nil
	}

	if !And(filters...)(*pokemon) {
		return nil
	}

	return pokemon
}

func (s *Store) FindAllTypes(filters ...FilterFunc[PokemonType]) iter.Seq[PokemonType] {
	it := slices.Values(s.types)
	return Filter(And(filters...), it)
}

func (s *Store) FindTypeBySymbol(symbol string, filters ...FilterFunc[PokemonType]) *PokemonType {
	pokemonType, ok := s.pokemonTypesBySymbol[symbol]
	if !ok {
		return nil
	}

	if !And(filters...)(*pokemonType) {
		return nil
	}

	return pokemonType
}

// Resistances calculates the type resistances of the PokemonForm based on its types.
func (s *Store) Resistances(type1, type2 string, filters ...FilterFunc[PokemonType]) iter.Seq2[string, float32] {
	typeIt := s.FindAllTypes(filters...)
	ptResIT := ToSeq2(typeIt, func(pt PokemonType) float32 {
		return pt.DamageToTypes(type1, type2)
	})

	return Map2(func(pt PokemonType, res float32) (string, float32) {
		return pt.DbSymbol(), res
	}, ptResIT)
}

func (s *Store) FindAllAbilities(filters ...FilterFunc[Ability]) iter.Seq[Ability] {
	it := slices.Values(s.abilities)
	return Filter(And(filters...), it)
}

func (s *Store) FindAbilityBySymbol(symbol string, filters ...FilterFunc[Ability]) *Ability {
	ability, ok := s.abilitiesBySymbol[symbol]
	if !ok {
		return nil
	}

	if !And(filters...)(*ability) {
		return nil
	}

	return ability
}

func (s *Store) FindAllMoves(filters ...FilterFunc[Move]) iter.Seq[Move] {
	it := slices.Values(s.moves)
	return Filter(And(filters...), it)
}

func (s *Store) FindMoveBySymbol(symbol string, filters ...FilterFunc[Move]) *Move {
	move, ok := s.movesBySymbol[symbol]
	if !ok {
		return nil
	}

	if !And(filters...)(*move) {
		return nil
	}

	return move
}
