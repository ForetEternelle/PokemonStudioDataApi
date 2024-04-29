package studio

type TypeStore struct {
	pokemonTypesBySymbol map[string]PokemonType
	types                []PokemonType
}

// NewTypeStore Create a new in memory type store
// types The list of types to store
func NewTypeStore() *TypeStore {
	return &TypeStore{
		pokemonTypesBySymbol: make(map[string]PokemonType),
		types:                make([]PokemonType, 0),
	}
}

// Add add a pokemon type to the store
// pokemonType the type to add
func (s *TypeStore) Add(pokemonType PokemonType) {
	s.types = append(s.types, pokemonType)
	s.pokemonTypesBySymbol[pokemonType.DbSymbol] = pokemonType
}

// FindBySymbol Find a type by its symbol
// symbol The symbol to find
func (s *TypeStore) FindBySymbol(symbol string) *PokemonType {
	pokemonType, ok := s.pokemonTypesBySymbol[symbol]
	if ok {
		copy := pokemonType
		return &copy
	} else {
		return nil
	}
}

// FindAll Find all types in the store
func (s *TypeStore) FindAll() []PokemonType {
	return s.types
}
