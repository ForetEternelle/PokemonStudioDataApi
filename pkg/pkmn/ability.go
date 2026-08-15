package pkmn

// Ability represents a Pokemon ability.
//
// Entities are immutable by convention only: fields are exported but must not
// be mutated after the entity has been registered in a store.
type Ability struct {
	// DbSymbol is the database symbol of the Ability.
	// Immutable: used to index abilities in the store.
	DbSymbol    string
	ID          int
	TextID      int
	Name        Translation
	Description Translation
}
