package studio

type PokemonType struct {
	DbSymbol string
	Color    string
	TextId   int
	Name     Translation
	DamageTo []TypeDamage
}

type TypeDamage struct {
	DefensiveType string
	Factor        float32
}
