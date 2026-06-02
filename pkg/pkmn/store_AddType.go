package pkmn

import "log/slog"

type AddTypeDto struct {
	DbSymbol string
	Color    string
	TextId   int
	Name     Translation
	DamageTo map[string]float32
}

func (s *Store) AddType(dto AddTypeDto) *PokemonType {
	pokemonType := &PokemonType{
		dbSymbol: dto.DbSymbol,
		color:    dto.Color,
		textId:   dto.TextId,
		name:     dto.Name,
		damageTo: dto.DamageTo,
	}

	s.types = append(s.types, pokemonType)
	s.pokemonTypesBySymbol[pokemonType.dbSymbol] = pokemonType
	slog.Info("Adding pokemon type", "symbol", pokemonType.dbSymbol)
	return pokemonType
}
