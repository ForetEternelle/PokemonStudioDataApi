package studio

import (
	"iter"
	"log/slog"
	"path"
	"slices"

	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/iter2"
)

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

const (
	StudioFolder   = "Studio"
	LanguageFolder = "Text/Dialogs"

	UndefType = "__undef__"
)

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
		if pokemon.Id < existingPokemon.Id {
			insertIndex = i
			break
		}
	}

	s.pokemonList = slices.Insert(s.pokemonList, insertIndex, pokemon)
	s.pokemonBySymbol[pokemon.DbSymbol] = &pokemon
	slog.Info("Adding pokemon", "symbol", pokemon.DbSymbol)
	return &pokemon
}

func (s *Store) AddType(pokemonType PokemonType) *PokemonType {
	s.types = append(s.types, pokemonType)
	s.pokemonTypesBySymbol[pokemonType.DbSymbol] = &pokemonType
	slog.Info("Adding pokemon type", "symbol", pokemonType.DbSymbol)
	return &pokemonType
}

func (s *Store) AddAbility(ability Ability) *Ability {
	s.abilities = append(s.abilities, ability)
	s.abilitiesBySymbol[ability.DbSymbol] = &ability
	slog.Info("Adding ability", "symbol", ability.DbSymbol)
	return &ability
}

func (s *Store) AddMove(move Move) *Move {
	s.moves = append(s.moves, move)
	s.movesBySymbol[move.DbSymbol] = &move
	slog.Info("Adding move", "symbol", move.DbSymbol)
	return &move
}

func (s *Store) FindAllPokemon(filters ...iter2.FilterFunc[Pokemon]) iter.Seq[Pokemon] {
	it := slices.Values(s.pokemonList)

	for _, filter := range filters {
		it = iter2.Filter(filter, it)
	}

	return it
}

func (s *Store) FindPokemonBySymbol(symbol string, filters ...iter2.FilterFunc[Pokemon]) *Pokemon {
	pokemon, ok := s.pokemonBySymbol[symbol]
	if !ok {
		return nil
	}

	for _, filter := range filters {
		if !filter(*pokemon) {
			return nil
		}
	}

	return pokemon
}

func (s *Store) FindAllTypes(filters ...iter2.FilterFunc[PokemonType]) iter.Seq[PokemonType] {
	it := slices.Values(s.types)

	for _, filter := range filters {
		it = iter2.Filter(filter, it)
	}

	return it
}

func (s *Store) FindTypeBySymbol(symbol string, filters ...iter2.FilterFunc[PokemonType]) *PokemonType {
	pokemonType, ok := s.pokemonTypesBySymbol[symbol]
	if !ok {
		return nil
	}

	for _, filter := range filters {
		if !filter(*pokemonType) {
			return nil
		}
	}

	return pokemonType
}

func (s *Store) FindAllAbilities(filters ...iter2.FilterFunc[Ability]) iter.Seq[Ability] {
	it := slices.Values(s.abilities)

	for _, filter := range filters {
		it = iter2.Filter(filter, it)
	}

	return it
}

func (s *Store) FindAbilityBySymbol(symbol string, filters ...iter2.FilterFunc[Ability]) *Ability {
	ability, ok := s.abilitiesBySymbol[symbol]
	if !ok {
		return nil
	}

	for _, filter := range filters {
		if !filter(*ability) {
			return nil
		}
	}

	return ability
}

func (s *Store) FindAllMoves(filters ...iter2.FilterFunc[Move]) iter.Seq[Move] {
	it := slices.Values(s.moves)

	for _, filter := range filters {
		it = iter2.Filter(filter, it)
	}

	return it
}

func (s *Store) FindMoveBySymbol(symbol string, filters ...iter2.FilterFunc[Move]) *Move {
	move, ok := s.movesBySymbol[symbol]
	if !ok {
		return nil
	}

	for _, filter := range filters {
		if !filter(*move) {
			return nil
		}
	}

	return move
}
