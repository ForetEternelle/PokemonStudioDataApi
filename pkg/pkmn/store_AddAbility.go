package pkmn

import "log/slog"

type AddAbilityDto struct {
	DbSymbol    string
	Id          int
	TextId      int
	Name        Translation
	Description Translation
}

func (s *Store) AddAbility(dto AddAbilityDto) *Ability {
	ability := &Ability{
		dbSymbol:    dto.DbSymbol,
		id:          dto.Id,
		textId:      dto.TextId,
		name:        dto.Name,
		description: dto.Description,
	}

	s.abilities = append(s.abilities, ability)
	s.abilitiesBySymbol[ability.dbSymbol] = ability
	slog.Info("Adding ability", "symbol", ability.dbSymbol)
	return ability
}
