package studio

import (
	"iter"
)

// PokemonType represents a Pokemon type (e.g., Fire, Water, Grass).
type PokemonType struct {
	dbSymbol string
	color    string
	textId   int
	name     Translation
	damageTo []TypeDamage
}

// PokemonTypeOption is a functional option for configuring a PokemonType.
type PokemonTypeOption func(*PokemonType)

// WithPokemonTypeDbSymbol sets the database symbol of a PokemonType.
func WithPokemonTypeDbSymbol(dbSymbol string) PokemonTypeOption {
	return func(t *PokemonType) { t.dbSymbol = dbSymbol }
}

// WithTypeColor sets the color of a PokemonType.
func WithTypeColor(color string) PokemonTypeOption {
	return func(t *PokemonType) { t.color = color }
}

// WithTypeTextId sets the text ID of a PokemonType.
func WithTypeTextId(id int) PokemonTypeOption {
	return func(t *PokemonType) { t.textId = id }
}

// WithTypeName sets the name translations of a PokemonType.
func WithTypeName(name Translation) PokemonTypeOption {
	return func(t *PokemonType) { t.name = name }
}

// WithDamageTo sets the damage relations of a PokemonType.
func WithDamageTo(damageTo []TypeDamage) PokemonTypeOption {
	return func(t *PokemonType) { t.damageTo = damageTo }
}

// NewPokemonType creates a new PokemonType with the given options.
func NewPokemonType(opts ...PokemonTypeOption) *PokemonType {
	t := &PokemonType{}
	for _, opt := range opts {
		opt(t)
	}
	return t
}

// DbSymbol returns the database symbol of the PokemonType.
func (t PokemonType) DbSymbol() string {
	return t.dbSymbol
}

// Color returns the color of the PokemonType.
func (t PokemonType) Color() string {
	return t.color
}

// TextId returns the text ID of the PokemonType.
func (t PokemonType) TextId() int {
	return t.textId
}

// Name returns the localized name of the PokemonType for the given language.
func (t PokemonType) Name(lang string) string {
	return t.name[lang]
}

// DamageTo returns an iterator over the type damage relations.
func (t PokemonType) DamageTo() iter.Seq2[int, TypeDamage] {
	return func(yield func(int, TypeDamage) bool) {
		for i, d := range t.damageTo {
			if !yield(i, d) {
				return
			}
		}
	}
}

// Damage returns the type damage relation for a defending type.
func (t PokemonType) Damage(defType string) (TypeDamage, bool) {
	for _, d := range t.damageTo {
		if d.DefensiveType == defType {
			return d, true
		}
	}
	return TypeDamage{}, false
}

// TypeDamage represents a type damage relation.
type TypeDamage struct {
	DefensiveType string
	Factor        float32
}

// PokemonTypeDescriptor is the JSON descriptor for a Pokemon type.
type PokemonTypeDescriptor struct {
	DbSymbol string                 `json:"dbSymbol"`
	Color    string                 `json:"color"`
	TextId   int                    `json:"textId"`
	DamageTo []TypeDamageDescriptor `json:"damageTo"`
	Name     Translation
}

// TypeDamageDescriptor is the JSON descriptor for a type damage relation.
type TypeDamageDescriptor struct {
	DefensiveType string  `json:"defensiveType"`
	Factor        float32 `json:"factor"`
}

// TypeMapper maps Type descriptors to PokemonType entities.
type TypeMapper struct {
	store *Store
}

// NewTypeMapper creates a new TypeMapper.
func NewTypeMapper(store *Store) *TypeMapper {
	return &TypeMapper{store: store}
}

// MapPokemonTypeDescriptorToPokemonType maps a PokemonTypeDescriptor to a PokemonType.
func (m *TypeMapper) MapPokemonTypeDescriptorToPokemonType(desc PokemonTypeDescriptor) *PokemonType {
	pokemonType := NewPokemonType(
		WithPokemonTypeDbSymbol(desc.DbSymbol),
		WithTypeColor(desc.Color),
		WithTypeTextId(desc.TextId),
		WithTypeName(desc.Name),
		WithDamageTo(MapTypeDamages(desc.DamageTo)),
	)

	return pokemonType
}

// MapTypeDamages maps type damage descriptors to type damage entities.
func MapTypeDamages(damages []TypeDamageDescriptor) []TypeDamage {
	if len(damages) == 0 {
		return nil
	}

	mapped := make([]TypeDamage, len(damages))
	for i, tdDesc := range damages {
		mapped[i] = TypeDamage{
			DefensiveType: tdDesc.DefensiveType,
			Factor:        tdDesc.Factor,
		}
	}
	return mapped
}
