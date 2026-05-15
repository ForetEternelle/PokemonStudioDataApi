package studio

import "github.com/ForetEternelle/PokemonStudioDataApi/pkg/pkmn"

// MoveMapper handles mapping of Move descriptors using a store
type MoveMapper struct {
	store *pkmn.Store
}

// NewMoveMapper creates a new MoveMapper with the given store
func NewMoveMapper(store *pkmn.Store) *MoveMapper {
	return &MoveMapper{store: store}
}

// MapMoveDescriptorToMove converts a MoveDescriptor to a Move domain struct
func (m *MoveMapper) MapMoveDescriptorToMove(desc MoveDescriptor) *pkmn.Move {
	moveObj := pkmn.NewMoveBuilder().
		ID(desc.Id).
		DbSymbol(desc.DbSymbol).
		Type(m.store.FindTypeBySymbol(desc.Type)).
		Category(pkmn.MoveCategory(desc.Category)).
		Power(desc.Power).
		Accuracy(desc.Accuracy).
		PP(desc.PP).
		CriticalRate(desc.MoveCriticalRate).
		Priority(desc.Priority).
		MapUse(desc.MapUse).
		Targeting(m.mapTargeting(desc)).
		Execution(m.mapExecution(desc)).
		MechanicalTags(m.mapMechanicalTags(desc)).
		Interactions(m.mapInteractions(desc)).
		SecondaryEffects(m.mapSecondaryEffects(desc)).
		Name(desc.Name).
		Description(desc.Description).
		Build()

	return moveObj
}

func (m *MoveMapper) mapTargeting(desc MoveDescriptor) pkmn.MoveTargeting {
	targeting := pkmn.MoveTargeting{
		AimedTarget: pkmn.AimedTarget(desc.BattleEngineAimedTarget),
	}

	if desc.IsDirect {
		targeting.ContactType = pkmn.ContactTypeDirect
	} else if desc.IsDistance {
		targeting.ContactType = pkmn.ContactTypeDistant
	} else {
		targeting.ContactType = pkmn.ContactTypeNone
	}

	return targeting
}

func (m *MoveMapper) mapExecution(desc MoveDescriptor) pkmn.MoveExecution {
	return pkmn.MoveExecution{
		Method:   pkmn.ExecutionMethod(desc.BattleEngineMethod),
		Charge:   desc.IsCharge,
		Recharge: desc.IsRecharge,
	}
}

func (m *MoveMapper) mapMechanicalTags(desc MoveDescriptor) []pkmn.MoveMechanicalTag {
	tags := make([]pkmn.MoveMechanicalTag, 0)

	if desc.IsAuthentic {
		tags = append(tags, pkmn.MechanicalTagAuthentic)
	}
	if desc.IsBallistics {
		tags = append(tags, pkmn.MechanicalTagBallistic)
	}
	if desc.IsBite {
		tags = append(tags, pkmn.MechanicalTagBite)
	}
	if desc.IsDance {
		tags = append(tags, pkmn.MechanicalTagDance)
	}
	if desc.IsPunch {
		tags = append(tags, pkmn.MechanicalTagPunch)
	}
	if desc.IsSlicingAttack {
		tags = append(tags, pkmn.MechanicalTagSlice)
	}
	if desc.IsSoundAttack {
		tags = append(tags, pkmn.MechanicalTagSound)
	}
	if desc.IsWind {
		tags = append(tags, pkmn.MechanicalTagWind)
	}
	if desc.IsPulse {
		tags = append(tags, pkmn.MechanicalTagPulse)
	}
	if desc.IsPowder {
		tags = append(tags, pkmn.MechanicalTagPowder)
	}
	if desc.IsMental {
		tags = append(tags, pkmn.MechanicalTagMental)
	}

	return tags
}

func (m *MoveMapper) mapInteractions(desc MoveDescriptor) []pkmn.MoveInteraction {
	interactions := make([]pkmn.MoveInteraction, 0)

	if desc.IsBlocable {
		interactions = append(interactions, pkmn.InteractionBlocable)
	}
	if desc.IsMirrorMove {
		interactions = append(interactions, pkmn.InteractionMirrorMove)
	}
	if desc.IsSnatchable {
		interactions = append(interactions, pkmn.InteractionSnatchable)
	}
	if desc.IsMagicCoatAffected {
		interactions = append(interactions, pkmn.InteractionMagicCoatAffected)
	}
	if desc.IsKingRockUtility {
		interactions = append(interactions, pkmn.InteractionKingRockUtility)
	}
	if desc.IsGravity {
		interactions = append(interactions, pkmn.InteractionAffectedByGravity)
	}
	if desc.IsNonSkyBattle {
		interactions = append(interactions, pkmn.InteractionNonSkyBattle)
	}

	return interactions
}

func (m *MoveMapper) mapSecondaryEffects(desc MoveDescriptor) pkmn.MoveSecondaryEffects {
	effects := pkmn.MoveSecondaryEffects{
		Chance: desc.EffectChance,
	}

	if len(desc.MoveStatus) > 0 {
		effects.StatusEffects = make([]pkmn.MoveStatusEffect, len(desc.MoveStatus))
		for i, status := range desc.MoveStatus {
			effects.StatusEffects[i] = pkmn.MoveStatusEffect{
				Status:   status.Status,
				LuckRate: status.LuckRate,
			}
		}
	}

	if len(desc.BattleStageMod) > 0 {
		effects.StatStageChanges = make([]pkmn.MoveStatStageChange, len(desc.BattleStageMod))
		for i, stageMod := range desc.BattleStageMod {
			effects.StatStageChanges[i] = pkmn.MoveStatStageChange{
				BattleStage: pkmn.BattleStage(stageMod.BattleStage),
				Modificator: stageMod.Modificator,
			}
		}
	}

	return effects
}
