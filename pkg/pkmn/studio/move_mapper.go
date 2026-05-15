package studio

// MoveMapper handles mapping of Move descriptors using a store
type MoveMapper struct {
	store *Store
}

// NewMoveMapper creates a new MoveMapper with the given store
func NewMoveMapper(store *Store) *MoveMapper {
	return &MoveMapper{store: store}
}

// MapMoveDescriptorToMove converts a MoveDescriptor to a Move domain struct
func (m *MoveMapper) MapMoveDescriptorToMove(desc MoveDescriptor) *Move {
	moveObj := &Move{
		id:           desc.Id,
		dbSymbol:     desc.DbSymbol,
		category:     MoveCategory(desc.Category),
		power:        desc.Power,
		accuracy:     desc.Accuracy,
		pp:           desc.PP,
		criticalRate: desc.MoveCriticalRate,
		priority:     desc.Priority,
		mapUse:       desc.MapUse,
		name:         desc.Name,
		description:  desc.Description,
	}

	// Resolve type reference
	moveObj.moveType = m.store.FindTypeBySymbol(desc.Type)

	// Map targeting
	moveObj.targeting = m.mapTargeting(desc)

	// Map execution
	moveObj.execution = m.mapExecution(desc)

	// Map mechanical tags
	moveObj.mechanicalTags = m.mapMechanicalTags(desc)

	// Map interactions
	moveObj.interactions = m.mapInteractions(desc)

	// Map secondary effects
	moveObj.secondaryEffects = m.mapSecondaryEffects(desc)

	return moveObj
}

func (m *MoveMapper) mapTargeting(desc MoveDescriptor) MoveTargeting {
	targeting := MoveTargeting{
		AimedTarget: AimedTarget(desc.BattleEngineAimedTarget),
	}

	// Determine contact type from isDirect and isDistance flags
	if desc.IsDirect {
		targeting.ContactType = ContactTypeDirect
	} else if desc.IsDistance {
		targeting.ContactType = ContactTypeDistant
	} else {
		targeting.ContactType = ContactTypeNone
	}

	return targeting
}

func (m *MoveMapper) mapExecution(desc MoveDescriptor) MoveExecution {
	return MoveExecution{
		Method:   ExecutionMethod(desc.BattleEngineMethod),
		Charge:   desc.IsCharge,
		Recharge: desc.IsRecharge,
	}
}

func (m *MoveMapper) mapMechanicalTags(desc MoveDescriptor) []MoveMechanicalTag {
	tags := make([]MoveMechanicalTag, 0)

	if desc.IsAuthentic {
		tags = append(tags, MechanicalTagAuthentic)
	}
	if desc.IsBallistics {
		tags = append(tags, MechanicalTagBallistic)
	}
	if desc.IsBite {
		tags = append(tags, MechanicalTagBite)
	}
	if desc.IsDance {
		tags = append(tags, MechanicalTagDance)
	}
	if desc.IsPunch {
		tags = append(tags, MechanicalTagPunch)
	}
	if desc.IsSlicingAttack {
		tags = append(tags, MechanicalTagSlice)
	}
	if desc.IsSoundAttack {
		tags = append(tags, MechanicalTagSound)
	}
	if desc.IsWind {
		tags = append(tags, MechanicalTagWind)
	}
	if desc.IsPulse {
		tags = append(tags, MechanicalTagPulse)
	}
	if desc.IsPowder {
		tags = append(tags, MechanicalTagPowder)
	}
	if desc.IsMental {
		tags = append(tags, MechanicalTagMental)
	}

	return tags
}

func (m *MoveMapper) mapInteractions(desc MoveDescriptor) []MoveInteraction {
	interactions := make([]MoveInteraction, 0)

	if desc.IsBlocable {
		interactions = append(interactions, InteractionBlocable)
	}
	if desc.IsMirrorMove {
		interactions = append(interactions, InteractionMirrorMove)
	}
	if desc.IsSnatchable {
		interactions = append(interactions, InteractionSnatchable)
	}
	if desc.IsMagicCoatAffected {
		interactions = append(interactions, InteractionMagicCoatAffected)
	}
	if desc.IsKingRockUtility {
		interactions = append(interactions, InteractionKingRockUtility)
	}
	if desc.IsGravity {
		interactions = append(interactions, InteractionAffectedByGravity)
	}
	if desc.IsNonSkyBattle {
		interactions = append(interactions, InteractionNonSkyBattle)
	}

	return interactions
}

func (m *MoveMapper) mapSecondaryEffects(desc MoveDescriptor) MoveSecondaryEffects {
	effects := MoveSecondaryEffects{
		Chance: desc.EffectChance,
	}

	// Map status effects
	if len(desc.MoveStatus) > 0 {
		effects.StatusEffects = make([]MoveStatusEffect, len(desc.MoveStatus))
		for i, status := range desc.MoveStatus {
			effects.StatusEffects[i] = MoveStatusEffect{
				Status:   status.Status,
				LuckRate: status.LuckRate,
			}
		}
	}

	// Map stat stage changes
	if len(desc.BattleStageMod) > 0 {
		effects.StatStageChanges = make([]MoveStatStageChange, len(desc.BattleStageMod))
		for i, stageMod := range desc.BattleStageMod {
			effects.StatStageChanges[i] = MoveStatStageChange{
				BattleStage: BattleStage(stageMod.BattleStage),
				Modificator: stageMod.Modificator,
			}
		}
	}

	return effects
}
