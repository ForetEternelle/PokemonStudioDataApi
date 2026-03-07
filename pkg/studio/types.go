package studio

import (
	"iter"
)

type PokemonType struct {
	dbSymbol string
	color    string
	textId   int
	name     Translation
	damageTo []TypeDamage
}

type PokemonTypeOption func(*PokemonType)

func WithPokemonTypeDbSymbol(dbSymbol string) PokemonTypeOption {
	return func(t *PokemonType) { t.dbSymbol = dbSymbol }
}

func WithTypeColor(color string) PokemonTypeOption {
	return func(t *PokemonType) { t.color = color }
}

func WithTypeTextId(id int) PokemonTypeOption {
	return func(t *PokemonType) { t.textId = id }
}

func WithTypeName(name Translation) PokemonTypeOption {
	return func(t *PokemonType) { t.name = name }
}

func WithDamageTo(damageTo []TypeDamage) PokemonTypeOption {
	return func(t *PokemonType) { t.damageTo = damageTo }
}

func NewPokemonType(opts ...PokemonTypeOption) *PokemonType {
	t := &PokemonType{}
	for _, opt := range opts {
		opt(t)
	}
	return t
}

func (t PokemonType) DbSymbol() string {
	return t.dbSymbol
}

func (t PokemonType) Color() string {
	return t.color
}

func (t PokemonType) TextId() int {
	return t.textId
}

func (t PokemonType) Name(lang string) string {
	return t.name[lang]
}

func (t PokemonType) DamageTo() iter.Seq2[int, TypeDamage] {
	return func(yield func(int, TypeDamage) bool) {
		for i, d := range t.damageTo {
			if !yield(i, d) {
				return
			}
		}
	}
}

func (t PokemonType) Damage(defType string) (TypeDamage, bool) {
	for _, d := range t.damageTo {
		if d.DefensiveType == defType {
			return d, true
		}
	}
	return TypeDamage{}, false
}

type TypeDamage struct {
	DefensiveType string
	Factor        float32
}

type PokemonTypeDescriptor struct {
	DbSymbol string                 `json:"dbSymbol"`
	Color    string                 `json:"color"`
	TextId   int                    `json:"textId"`
	DamageTo []TypeDamageDescriptor `json:"damageTo"`
	Name     Translation
}

type TypeDamageDescriptor struct {
	DefensiveType string  `json:"defensiveType"`
	Factor        float32 `json:"factor"`
}

type TypeMapper struct {
	store *Store
}

func NewTypeMapper(store *Store) *TypeMapper {
	return &TypeMapper{store: store}
}

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
