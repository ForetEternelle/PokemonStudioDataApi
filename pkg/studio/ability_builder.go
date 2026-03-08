package studio

// AbilityBuilder builds Ability entities.
type AbilityBuilder struct {
	ability *Ability
}

// NewAbilityBuilder creates a new AbilityBuilder.
func NewAbilityBuilder() *AbilityBuilder {
	return &AbilityBuilder{ability: &Ability{}}
}

// DbSymbol sets the database symbol of the Ability.
func (b *AbilityBuilder) DbSymbol(dbSymbol string) *AbilityBuilder {
	b.ability.dbSymbol = dbSymbol
	return b
}

// ID sets the ID of the Ability.
func (b *AbilityBuilder) ID(id int) *AbilityBuilder {
	b.ability.id = id
	return b
}

// TextId sets the text ID of the Ability.
func (b *AbilityBuilder) TextId(textId int) *AbilityBuilder {
	b.ability.textId = textId
	return b
}

// Name sets the name translations of the Ability.
func (b *AbilityBuilder) Name(name Translation) *AbilityBuilder {
	b.ability.name = name
	return b
}

// Description sets the description translations of the Ability.
func (b *AbilityBuilder) Description(desc Translation) *AbilityBuilder {
	b.ability.description = desc
	return b
}

// Build returns a copy of the built Ability.
func (b *AbilityBuilder) Build() *Ability {
	_copy := *b.ability
	return &_copy
}
