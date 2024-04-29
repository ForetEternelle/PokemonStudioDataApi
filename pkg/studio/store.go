package studio

type Store struct {
	PokemonStore *PokemonStore
	TypeStore    *TypeStore
}

func NewStore() *Store {
	return &Store{
		NewPokemonStore(),
		NewTypeStore(),
	}
}
