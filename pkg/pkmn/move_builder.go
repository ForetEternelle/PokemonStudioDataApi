package studio

// MoveBuilder builds Move entities.
type MoveBuilder struct {
	move *Move
}

// NewMoveBuilder creates a new MoveBuilder.
func NewMoveBuilder() *MoveBuilder {
	return &MoveBuilder{
		move: &Move{},
	}
}

// ID sets the ID of a Move.
func (b *MoveBuilder) ID(id int) *MoveBuilder {
	b.move.id = id
	return b
}

// DbSymbol sets the database symbol of a Move.
func (b *MoveBuilder) DbSymbol(dbSymbol string) *MoveBuilder {
	b.move.dbSymbol = dbSymbol
	return b
}

// Type sets the type of a Move.
func (b *MoveBuilder) Type(t *PokemonType) *MoveBuilder {
	b.move.moveType = t
	return b
}

// Category sets the category of a Move.
func (b *MoveBuilder) Category(cat MoveCategory) *MoveBuilder {
	b.move.category = cat
	return b
}

// Power sets the power of a Move.
func (b *MoveBuilder) Power(power int) *MoveBuilder {
	b.move.power = power
	return b
}

// Accuracy sets the accuracy of a Move.
func (b *MoveBuilder) Accuracy(acc int) *MoveBuilder {
	b.move.accuracy = acc
	return b
}

// PP sets the PP of a Move.
func (b *MoveBuilder) PP(pp int) *MoveBuilder {
	b.move.pp = pp
	return b
}

// CriticalRate sets the critical rate of a Move.
func (b *MoveBuilder) CriticalRate(rate int) *MoveBuilder {
	b.move.criticalRate = rate
	return b
}

// Priority sets the priority of a Move.
func (b *MoveBuilder) Priority(priority int) *MoveBuilder {
	b.move.priority = priority
	return b
}

// MapUse sets the map use of a Move.
func (b *MoveBuilder) MapUse(mapUse int) *MoveBuilder {
	b.move.mapUse = mapUse
	return b
}

// Name sets the name translations of a Move.
func (b *MoveBuilder) Name(name Translation) *MoveBuilder {
	b.move.name = name
	return b
}

// Description sets the description translations of a Move.
func (b *MoveBuilder) Description(desc Translation) *MoveBuilder {
	b.move.description = desc
	return b
}

// Targeting sets the targeting of a Move.
func (b *MoveBuilder) Targeting(targeting MoveTargeting) *MoveBuilder {
	b.move.targeting = targeting
	return b
}

// Execution sets the execution of a Move.
func (b *MoveBuilder) Execution(exec MoveExecution) *MoveBuilder {
	b.move.execution = exec
	return b
}

// MechanicalTags sets the mechanical tags of a Move.
func (b *MoveBuilder) MechanicalTags(tags []MoveMechanicalTag) *MoveBuilder {
	b.move.mechanicalTags = tags
	return b
}

// Interactions sets the interactions of a Move.
func (b *MoveBuilder) Interactions(interactions []MoveInteraction) *MoveBuilder {
	b.move.interactions = interactions
	return b
}

// SecondaryEffects sets the secondary effects of a Move.
func (b *MoveBuilder) SecondaryEffects(effects MoveSecondaryEffects) *MoveBuilder {
	b.move.secondaryEffects = effects
	return b
}

// Build returns a copy of the built Move.
func ( b *MoveBuilder) Build() *Move {
	_copy := *b.move
	return &_copy
}
