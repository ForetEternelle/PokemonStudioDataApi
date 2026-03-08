package studio

// TypeBuilder builds PokemonType entities.
type TypeBuilder struct {
	pokemonType *PokemonType
}

// NewTypeBuilder creates a new TypeBuilder.
func NewTypeBuilder() *TypeBuilder {
	return &TypeBuilder{pokemonType: &PokemonType{}}
}

// DbSymbol sets the database symbol of the PokemonType.
func (b *TypeBuilder) DbSymbol(dbSymbol string) *TypeBuilder {
	b.pokemonType.dbSymbol = dbSymbol
	return b
}

// Color sets the color of the PokemonType.
func (b *TypeBuilder) Color(color string) *TypeBuilder {
	b.pokemonType.color = color
	return b
}

// TextId sets the text ID of the PokemonType.
func (b *TypeBuilder) TextId(id int) *TypeBuilder {
	b.pokemonType.textId = id
	return b
}

// Name sets the name translations of the PokemonType.
func (b *TypeBuilder) Name(name Translation) *TypeBuilder {
	b.pokemonType.name = name
	return b
}

// DamageTo sets the damage relations of the PokemonType.
func (b *TypeBuilder) DamageTo(damageTo []TypeDamage) *TypeBuilder {
	b.pokemonType.damageTo = damageTo
	return b
}

// Build returns a copy of the built PokemonType.
func (b *TypeBuilder) Build() *PokemonType {
	_copy := *b.pokemonType
	return &_copy
}
