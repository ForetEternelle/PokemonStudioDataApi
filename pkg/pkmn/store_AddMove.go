package pkmn

import "log/slog"

type AddMoveDto struct {
	Id               int
	DbSymbol         string
	Type             string
	Category         MoveCategory
	Power            int
	Accuracy         int
	PP               int
	CriticalRate     int
	Priority         int
	MapUse           int
	Targeting        MoveTargeting
	Execution        MoveExecution
	MechanicalTags   []MoveMechanicalTag
	Interactions     []MoveInteraction
	SecondaryEffects MoveSecondaryEffects
	Name             Translation
	Description      Translation
}

func (s *Store) AddMove(dto AddMoveDto) *Move {
	move := &Move{
		id:               dto.Id,
		dbSymbol:         dto.DbSymbol,
		moveType:         s.FindTypeBySymbol(dto.Type),
		category:         dto.Category,
		power:            dto.Power,
		accuracy:         dto.Accuracy,
		pp:               dto.PP,
		criticalRate:     dto.CriticalRate,
		priority:         dto.Priority,
		mapUse:           dto.MapUse,
		targeting:        dto.Targeting,
		execution:        dto.Execution,
		mechanicalTags:   dto.MechanicalTags,
		interactions:     dto.Interactions,
		secondaryEffects: dto.SecondaryEffects,
		name:             dto.Name,
		description:      dto.Description,
	}

	s.moves = append(s.moves, move)
	s.movesBySymbol[move.dbSymbol] = move
	slog.Info("Adding move", "symbol", move.dbSymbol)
	return move
}
