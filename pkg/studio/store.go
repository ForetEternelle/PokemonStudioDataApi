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

// Import a pokemon studio folder into a store
// folder the studio project folder
// store the store to import data to
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

// FindAllPokemon Find a page of pokemon corresponding to the page request
// filters iterator function to filter the pokemon
func (s *Store) FindAllPokemon(filters ...iter2.FilterFunc[Pokemon]) iter.Seq[Pokemon] {
	it := slices.Values(s.pokemonList)
	for _, filter := range filters {
		it = iter2.Filter(filter, it)
	}
	return it
}

// FindPokemonBySymbol Find pokemon by symbol
// symbol The symbol of the pokemon to find
func (s *Store) FindPokemonBySymbol(symbol string) *Pokemon {
	pokemon, ok := s.pokemonBySymbol[symbol]
	if ok {
		return pokemon
	} else {
		return nil
	}
}

// FindTypeBySymbol Find a type by its symbol
// symbol The symbol to find
func (s *Store) FindTypeBySymbol(symbol string) *PokemonType {
	pokemonType, ok := s.pokemonTypesBySymbol[symbol]
	if ok {
		return pokemonType
	} else {
		return nil
	}
}

// FindAllTypes Find all types in the store
func (s *Store) FindAllTypes() []PokemonType {
	return s.types
}

// FindAbilityBySymbol Find an ability by its symbol
// symbol The symbol to find
func (s *Store) FindAbilityBySymbol(symbol string) *Ability {
	ability, ok := s.abilitiesBySymbol[symbol]
	if ok {
		return ability
	} else {
		return nil
	}
}

// FindAllAbilities Find all abilities in the store
func (s *Store) FindAllAbilities() []Ability {
	return s.abilities
}

// FindMoveBySymbol Find a move by its symbol
// symbol The symbol to find
func (s *Store) FindMoveBySymbol(symbol string) *Move {
	move, ok := s.movesBySymbol[symbol]
	if ok {
		return move
	} else {
		return nil
	}
}

// FindAllMoves Find all moves in the store
func (s *Store) FindAllMoves() []Move {
	return s.moves
}
