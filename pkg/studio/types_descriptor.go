package studio

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

