package pkmn

// LearnMethod represents how a Pokemon learns a move.
type LearnMethod string

const (
	LearnMethodLevelUp   LearnMethod = "level_up"
	LearnMethodTutor     LearnMethod = "tutor"
	LearnMethodMachine   LearnMethod = "machine"
	LearnMethodBreed     LearnMethod = "breed"
	LearnMethodEvolution LearnMethod = "evolution"
)

// PokemonFormMove is a move a Pokemon form can learn and the way it learns it.
//
// Entities are immutable by convention only: fields are exported but must not
// be mutated after the entity has been registered in a store.
type PokemonFormMove struct {
	// MoveSymbol is the database symbol of the learnable move, as authored in Studio.
	MoveSymbol string
	// Move is the resolved move, nil when the symbol is unknown.
	Move *Move
	// Method is the acquisition method.
	Method LearnMethod
	// Level is the learning level for levels moves, 0 otherwise.
	Level int32
	// Machine is the technical machine teaching the move, only for machine-learned moves.
	Machine *Machine
}

// Machine is a technical machine item (HM or TM) teaching a move.
//
// Entities are immutable by convention only: fields are exported but must not
// be mutated after the entity has been registered in a store.
type Machine struct {
	// DbSymbol is the database symbol of the item.
	// Immutable: used to index machines in the store.
	DbSymbol   string
	ID         int
	Name       Translation
	IsHm       bool
	MoveSymbol string
}
