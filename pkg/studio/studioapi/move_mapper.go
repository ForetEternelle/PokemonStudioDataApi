package studioapi

import (
	"log/slog"

	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/studio"
)

type MoveMapper struct {
	typeMapper *TypeMapper
}

// NewMoveMapper Create a new move mapper
func NewMoveMapper(typeMapper *TypeMapper) *MoveMapper {
	return &MoveMapper{typeMapper: typeMapper}
}

// ToMoveDetail map a move to a move details transfer object
// move the move to map
// lang the language expected
func (m MoveMapper) ToMoveDetail(move studio.Move, lang string, policy *AccessPolicy) MoveDetails {
	slog.Debug("Mapping move to details", "move", move.DbSymbol, "lang", lang)

	details := MoveDetails{
		Symbol:       move.DbSymbol,
		Name:         move.Name[lang],
		Description:  move.Description[lang],
		Category:     string(move.Category),
		Power:        int32(move.Power),
		Accuracy:     int32(move.Accuracy),
		Pp:           int32(move.PP),
		CriticalRate: int32(move.CriticalRate),
		Priority:     int32(move.Priority),
	}

	// Map type
	if move.Type != nil {
		typePartial := m.typeMapper.ToTypePartial(*move.Type, lang, policy)
		details.Type = typePartial
	}

	// Map targeting
	details.Targeting = map[string]interface{}{
		"aimedTarget": string(move.Targeting.AimedTarget),
		"contactType": string(move.Targeting.ContactType),
	}

	// Map execution
	details.Execution = map[string]interface{}{
		"method":   string(move.Execution.Method),
		"charge":   move.Execution.Charge,
		"recharge": move.Execution.Recharge,
	}

	// Map mechanical tags
	if len(move.MechanicalTags) > 0 {
		tags := make([]map[string]interface{}, len(move.MechanicalTags))
		for i, tag := range move.MechanicalTags {
			tags[i] = map[string]interface{}{"tag": string(tag)}
		}
		details.MechanicalTags = tags
	}

	// Map interactions
	if len(move.Interactions) > 0 {
		interactions := make([]string, len(move.Interactions))
		for i, interaction := range move.Interactions {
			interactions[i] = string(interaction)
		}
		details.Interactions = map[string]interface{}{"list": interactions}
	}

	// Map secondary effects
	secondaryEffects := map[string]interface{}{
		"chance": move.SecondaryEffects.Chance,
	}

	if len(move.SecondaryEffects.StatusEffects) > 0 {
		statusEffects := make([]map[string]interface{}, len(move.SecondaryEffects.StatusEffects))
		for i, effect := range move.SecondaryEffects.StatusEffects {
			statusEffects[i] = map[string]interface{}{
				"status":   effect.Status,
				"luckRate": effect.LuckRate,
			}
		}
		secondaryEffects["statusEffects"] = statusEffects
	}

	if len(move.SecondaryEffects.StatStageChanges) > 0 {
		statStageChanges := make([]map[string]interface{}, len(move.SecondaryEffects.StatStageChanges))
		for i, change := range move.SecondaryEffects.StatStageChanges {
			statStageChanges[i] = map[string]interface{}{
				"battleStage": string(change.BattleStage),
				"modificator": change.Modificator,
			}
		}
		secondaryEffects["statStageChanges"] = statStageChanges
	}

	details.SecondaryEffects = secondaryEffects

	return details
}

// ToMovePartial map a move to a move partial transfer object
// move the move to map
// lang the language expected
func (m MoveMapper) ToMovePartial(move studio.Move, lang string, policy *AccessPolicy) MovePartial {
	slog.Debug("Mapping move to partial", "move", move.DbSymbol, "lang", lang)
	partial := MovePartial{
		Symbol:   move.DbSymbol,
		Name:     move.Name[lang],
		Category: string(move.Category),
		Pp:       int32(move.PP),
		Power:    int32(move.Power),
		Accuracy: int32(move.Accuracy),
	}

	// Map type
	if move.Type != nil {
		typePartial := m.typeMapper.ToTypePartial(*move.Type, lang, policy)
		partial.Type = typePartial
	}

	return partial
}
