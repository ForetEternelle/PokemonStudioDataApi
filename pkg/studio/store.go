package studio

import (
	"iter"
	"log/slog"
	"path"
	"slices"
	"sort"

	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/iter2"
)

type Store struct {
	pokemonBySymbol      map[string]*Pokemon
	pokemonList          []Pokemon
	pokemonTypesBySymbol map[string]*PokemonType
	types                []PokemonType
	abilities            []Ability
	abilitiesBySymbol    map[string]*Ability
}

type Translation map[string]string

const (
	StudioFolder   = "Studio"
	LanguageFolder = "Text/Dialogs"

	UndefType = "__undef__"
)

func NewStore(pokemonList []Pokemon, types []PokemonType, abilities []Ability) (*Store, error) {
	pokemonBySymbol := make(map[string]*Pokemon)
	pokemonTypesBySymbol := make(map[string]*PokemonType)
	abilitiesBySymbol := make(map[string]*Ability)

	sort.Slice(pokemonList, func(i, j int) bool {
		return pokemonList[i].Id < pokemonList[j].Id
	})

	for _, p := range pokemonList {
		pokemonBySymbol[p.DbSymbol] = &p
	}

	sort.Slice(types, func(i, j int) bool {
		return types[i].TextId < types[j].TextId
	})

	for _, t := range types {
		pokemonTypesBySymbol[t.DbSymbol] = &t
	}

	sort.Slice(abilities, func(i, j int) bool { return abilities[i].Id < abilities[j].Id })

	for _, t := range abilities {
		abilitiesBySymbol[t.DbSymbol] = &t
	}

	return &Store{
		pokemonBySymbol:      pokemonBySymbol,
		pokemonList:          pokemonList,
		pokemonTypesBySymbol: pokemonTypesBySymbol,
		types:                types,
		abilities:            abilities,
		abilitiesBySymbol:    abilitiesBySymbol,
	}, nil
}

// Import import a pokemon studio folder into a store
// folder the studio project folder
// store the store to import data to
func Load(folder string) (*Store, error) {
	translationFolder := path.Join(folder, LanguageFolder)
	studioFolder := path.Join(folder, StudioFolder)

	pokemonList, err := ImportPokemon(studioFolder, translationFolder)
	if err != nil {
		slog.Error("Failed to create store")
		return nil, err
	}

	types, err := ImportTypes(studioFolder, translationFolder)
	if err != nil {
		slog.Error("Failed to create store")
		return nil, err
	}

	abilities, err := ImportAbility(studioFolder, translationFolder)
	if err != nil {
		slog.Error("Failed to create store")
		return nil, err
	}

	return NewStore(pokemonList, types, abilities)
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
