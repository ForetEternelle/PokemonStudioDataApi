package studio

import (
	"slices"
	"sort"

	"github.com/rcharre/psapi/pkg/utils/iter2"
	"github.com/rcharre/psapi/pkg/utils/pagination"
)

type PokemonStore struct {
	pokemonBySymbol map[string]Pokemon
	pokemonList     []Pokemon
}

// NewPokemonStore Create an in memory pokemon store
func NewPokemonStore() *PokemonStore {
	pokemonBySymbol := make(map[string]Pokemon)
	pokemonList := make([]Pokemon, 0)

	return &PokemonStore{
		pokemonBySymbol: pokemonBySymbol,
		pokemonList:     pokemonList,
	}
}

func (s *PokemonStore) Add(pokemon Pokemon) {
	s.pokemonBySymbol[pokemon.DbSymbol] = pokemon
	index := sort.Search(len(s.pokemonList), func(i int) bool {
		return s.pokemonList[i].Id >= pokemon.Id
	})

	s.pokemonList = slices.Insert(s.pokemonList, index, pokemon)
}

// FindAll Find a page of pokemon corresponding to the page request
// pageRequest the page request
func (s *PokemonStore) FindAll(pageRequest pagination.PageRequest, filters ...iter2.FilterFunc[Pokemon]) pagination.Page[Pokemon] {

	it := slices.Values(s.pokemonList)
	for _, filter := range filters {
		it = iter2.Filter(filter, it)
	}

	return pagination.Collect(it, pageRequest)
}

// FindBySymbol Find pokemon by symbol
// symbol The symbol of the pokemon to find
func (s *PokemonStore) FindBySymbol(symbol string) *Pokemon {
	pokemon, ok := s.pokemonBySymbol[symbol]
	if ok {
		copy := pokemon
		return &copy
	} else {
		return nil
	}
}
