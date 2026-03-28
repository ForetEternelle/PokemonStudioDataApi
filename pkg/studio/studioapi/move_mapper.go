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
	slog.Debug("Mapping move to details", "move", move.DbSymbol(), "lang", lang)

	details := MoveDetails{
		Symbol:       move.DbSymbol(),
		Name:         move.Name(lang),
		Description:  move.Description(lang),
		Category:     string(move.Category()),
		Power:        int32(move.Power()),
		Accuracy:     int32(move.Accuracy()),
		Pp:           int32(move.PP()),
		CriticalRate: int32(move.CriticalRate()),
		Priority:     int32(move.Priority()),
	}

	// Map type
	t := move.Type()
	if t.DbSymbol() != "" {
		typePartial := m.typeMapper.ToTypePartial(move.Type(), lang, policy)
		details.Type = typePartial
	}

	// Map targeting
	targeting := move.Targeting()
	details.Targeting = map[string]interface{}{
		"aimedTarget": string(targeting.AimedTarget),
		"contactType": string(targeting.ContactType),
	}

	// Map execution
	execution := move.Execution()
	details.Execution = map[string]interface{}{
		"method":   string(execution.Method),
		"charge":   execution.Charge,
		"recharge": execution.Recharge,
	}

	// Map mechanical tags
	mechanicalTags := move.MechanicalTags()
	if len(mechanicalTags) > 0 {
		tags := make([]map[string]interface{}, len(mechanicalTags))
		for i, tag := range mechanicalTags {
			tags[i] = map[string]interface{}{"tag": string(tag)}
		}
		details.MechanicalTags = tags
	}

	// Map interactions
	interactions := move.Interactions()
	if len(interactions) > 0 {
		interactionList := make([]string, len(interactions))
		for i, interaction := range interactions {
			interactionList[i] = string(interaction)
		}
		details.Interactions = map[string]interface{}{"list": interactionList}
	}

	// Map secondary effects
	secondaryEffects := move.SecondaryEffects()
	secondaryEffectsMap := map[string]interface{}{
		"chance": secondaryEffects.Chance,
	}

	if len(secondaryEffects.StatusEffects) > 0 {
		statusEffects := make([]map[string]interface{}, len(secondaryEffects.StatusEffects))
		for i, effect := range secondaryEffects.StatusEffects {
			statusEffects[i] = map[string]interface{}{
				"status":   effect.Status,
				"luckRate": effect.LuckRate,
			}
		}
		secondaryEffectsMap["statusEffects"] = statusEffects
	}

	if len(secondaryEffects.StatStageChanges) > 0 {
		statStageChanges := make([]map[string]interface{}, len(secondaryEffects.StatStageChanges))
		for i, change := range secondaryEffects.StatStageChanges {
			statStageChanges[i] = map[string]interface{}{
				"battleStage": string(change.BattleStage),
				"modificator": change.Modificator,
			}
		}
		secondaryEffectsMap["statStageChanges"] = statStageChanges
	}

	details.SecondaryEffects = secondaryEffectsMap

	return details
}

// ToMovePartial map a move to a move partial transfer object
// move the move to map
// lang the language expected
func (m MoveMapper) ToMovePartial(move studio.Move, lang string, policy *AccessPolicy) MovePartial {
	slog.Debug("Mapping move to partial", "move", move.DbSymbol(), "lang", lang)
	partial := MovePartial{
		Symbol:   move.DbSymbol(),
		Name:     move.Name(lang),
		Category: string(move.Category()),
		Pp:       int32(move.PP()),
		Power:    int32(move.Power()),
		Accuracy: int32(move.Accuracy()),
	}

	t := move.Type()
	if t.DbSymbol() != "" {
		typePartial := m.typeMapper.ToTypePartial(move.Type(), lang, policy)
		partial.Type = typePartial
	}

	return partial
}
