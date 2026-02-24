package studio

import "github.com/ForetEternelle/PokemonStudioDataApi/pkg/studio/move"

var ExperienceTypeMap = map[ExperienceTypeDescriptor]string{
	ExperienceErraticNum:     ExperienceErratic,
	ExperienceFastNum:        ExperienceFast,
	ExperienceMediumFastNum:  ExperienceMediumFast,
	ExperienceMediumSlowNum:  ExperienceMediumSlow,
	ExperienceSlowNum:        ExperienceSlow,
	ExperienceFluctuatingNum: ExperienceFluctuating,
}

var BreedMap = map[BreedGroupDescriptor]string{
	BreedMonsterNum:      BreedMonster,
	BreedWater1Num:       BreedWater1,
	BreedBugNum:          BreedBug,
	BreedFlyingNum:       BreedFlying,
	BreedFieldNum:        BreedField,
	BreedFairyNum:        BreedFairy,
	BreedGrassNum:        BreedGrass,
	BreedHumanNum:        BreedHuman,
	BreedWater3Num:       BreedWater3,
	BreedMineralNum:      BreedMineral,
	BreedAmorphousNum:    BreedAmorphous,
	BreedWater2Num:       BreedWater2,
	BreedDittoNum:        BreedDitto,
	BreedDragonNum:       BreedDragon,
	BreedUndiscoveredNum: BreedUndiscovered,
}

// PokemonMapper handles mapping of Pokemon descriptors with reference resolution using a store
type PokemonMapper struct {
	store *Store
}

// NewPokemonMapper creates a new PokemonMapper with the given store
func NewPokemonMapper(store *Store) *PokemonMapper {
	return &PokemonMapper{store: store}
}

// MapPokemonDescriptorToPokemon converts a PokemonDescriptor to a Pokemon domain struct with reference resolution
func (m *PokemonMapper) MapPokemonDescriptorToPokemon(desc PokemonDescriptor) *Pokemon {
	forms := make(map[int32]PokemonForm, len(desc.Forms))
	for _, formDesc := range desc.Forms {
		forms[formDesc.Form] = *m.MapFormDescriptorToPokemonForm(formDesc)
	}

	pokemon := &Pokemon{
		Id:               desc.ID,
		DbSymbol:         desc.DbSymbol,
		Forms:            forms,
		Name:             desc.Name,
		Description:      desc.Description,
		CustomProperties: make(map[string]any),
	}

	return pokemon
}

// MapFormDescriptorToPokemonForm converts a FormDescriptor to a PokemonForm domain struct with reference resolution
func (m *PokemonMapper) MapFormDescriptorToPokemonForm(desc FormDescriptor) *PokemonForm {
	form := &PokemonForm{
		Form:           desc.Form,
		Height:         desc.Height,
		Weight:         desc.Weight,
		BaseHp:         desc.BaseHp,
		BaseAtk:        desc.BaseAtk,
		BaseDfe:        desc.BaseDfe,
		BaseSpd:        desc.BaseSpd,
		BaseAts:        desc.BaseAts,
		BaseDfs:        desc.BaseDfs,
		EvHp:           desc.EvHp,
		EvAtk:          desc.EvAtk,
		EvDfe:          desc.EvDfe,
		EvSpd:          desc.EvSpd,
		EvAts:          desc.EvAts,
		EvDfs:          desc.EvDfs,
		ExperienceType: ExperienceTypeMap[desc.ExperienceType],
		BaseExperience: desc.BaseExperience,
		BaseLoyalty:    desc.BaseLoyalty,
		CatchRate:      desc.CatchRate,
		FemaleRate:     desc.FemaleRate,
		BreedGroups:    MapBreedGroups(desc.BreedGroups),
		HatchSteps:     desc.HatchSteps,
		BabyForm:       desc.BabyForm,
		FrontOffsetY:   desc.FrontOffsetY,
	}

	if desc.BabyDbSymbol != nil {
		form.BabyDbSymbol = desc.BabyDbSymbol
	}

	form.Evolutions = MapEvolutions(desc.Evolutions)
	form.ItemHeld = MapItemHelds(desc.ItemHeld)

	form.Type1 = m.store.FindTypeBySymbol(desc.Type1)
	if desc.Type2 != nil && *desc.Type2 != "" && *desc.Type2 != UndefType {
		form.Type2 = m.store.FindTypeBySymbol(*desc.Type2)
	}

	form.Abilities = make([]*Ability, 0, len(desc.Abilities))
	for _, abilitySymbol := range desc.Abilities {
		if ability := m.store.FindAbilityBySymbol(abilitySymbol); ability != nil {
			form.Abilities = append(form.Abilities, ability)
		}
	}

	return form
}

// TypeMapper handles mapping of PokemonType descriptors using a store
type TypeMapper struct {
	store *Store
}

// NewTypeMapper creates a new TypeMapper with the given store
func NewTypeMapper(store *Store) *TypeMapper {
	return &TypeMapper{store: store}
}

// MapPokemonTypeDescriptorToPokemonType converts a PokemonTypeDescriptor to a PokemonType domain struct
func (m *TypeMapper) MapPokemonTypeDescriptorToPokemonType(desc PokemonTypeDescriptor) *PokemonType {
	pokemonType := &PokemonType{
		DbSymbol: desc.DbSymbol,
		Color:    desc.Color,
		TextId:   desc.TextId,
		Name:     desc.Name,
	}

	pokemonType.DamageTo = MapTypeDamages(desc.DamageTo)
	return pokemonType
}

// AbilityMapper handles mapping of Ability descriptors using a store
type AbilityMapper struct {
	store *Store
}

// NewAbilityMapper creates a new AbilityMapper with the given store
func NewAbilityMapper(store *Store) *AbilityMapper {
	return &AbilityMapper{store: store}
}

// MapAbilityDescriptorToAbility converts an AbilityDescriptor to an Ability domain struct
func (m *AbilityMapper) MapAbilityDescriptorToAbility(desc AbilityDescriptor) *Ability {
	ability := &Ability{
		DbSymbol:    desc.DbSymbol,
		Id:          desc.Id,
		TextID:      desc.TextID,
		Description: desc.Description,
		Name:        desc.Name,
	}

	return ability
}

func MapBreedGroups(breedGroupInts []int32) []string {
	breedGroups := make([]string, len(breedGroupInts))
	for i, bgInt := range breedGroupInts {
		breedGroups[i] = BreedMap[BreedGroupDescriptor(bgInt)]
	}
	return breedGroups
}

func MapEvolutions(evolutions []EvolutionDescriptor) []Evolution {
	if len(evolutions) == 0 {
		return nil
	}

	mapped := make([]Evolution, len(evolutions))
	for i, evoDesc := range evolutions {
		mapped[i] = Evolution{
			DbSymbol:   evoDesc.DbSymbol,
			Form:       evoDesc.Form,
			Conditions: MapConditions(evoDesc.Conditions),
		}
	}
	return mapped
}

func MapConditions(conditions []ConditionDescriptor) []Condition {
	if len(conditions) == 0 {
		return nil
	}

	mapped := make([]Condition, len(conditions))
	for i, condDesc := range conditions {
		mapped[i] = Condition{Type: condDesc.Type}
	}
	return mapped
}

func MapItemHelds(itemHelds []ItemHeldDescriptor) []*ItemHeld {
	if len(itemHelds) == 0 {
		return nil
	}

	mapped := make([]*ItemHeld, len(itemHelds))
	for i, ihDesc := range itemHelds {
		mapped[i] = &ItemHeld{
			DbSymbol: ihDesc.DbSymbol,
			Chance:   ihDesc.Chance,
		}
	}
	return mapped
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
		Id:           desc.Id,
		DbSymbol:     desc.DbSymbol,
		Category:     MoveCategory(desc.Category),
		Power:        desc.Power,
		Accuracy:     desc.Accuracy,
		PP:           desc.PP,
		CriticalRate: desc.MoveCriticalRate,
		Priority:     desc.Priority,
		MapUse:       desc.MapUse,
		Name:         desc.Name,
		Description:  desc.Description,
	}

	// Resolve type reference
	moveObj.Type = m.store.FindTypeBySymbol(desc.Type)

	// Map targeting
	moveObj.Targeting = m.mapTargeting(desc)

	// Map execution
	moveObj.Execution = m.mapExecution(desc)

	// Map mechanical tags
	moveObj.MechanicalTags = m.mapMechanicalTags(desc)

	// Map interactions
	moveObj.Interactions = m.mapInteractions(desc)

	// Map secondary effects
	moveObj.SecondaryEffects = m.mapSecondaryEffects(desc)

	return moveObj
}

func (m *MoveMapper) mapTargeting(desc MoveDescriptor) move.MoveTargeting {
	targeting := move.MoveTargeting{
		AimedTarget: move.AimedTarget(desc.BattleEngineAimedTarget),
	}

	// Determine contact type from isDirect and isDistance flags
	if desc.IsDirect {
		targeting.ContactType = move.ContactTypeDirect
	} else if desc.IsDistance {
		targeting.ContactType = move.ContactTypeDistant
	} else {
		targeting.ContactType = move.ContactTypeNone
	}

	return targeting
}

func (m *MoveMapper) mapExecution(desc MoveDescriptor) move.MoveExecution {
	return move.MoveExecution{
		Method:   move.ExecutionMethod(desc.BattleEngineMethod),
		Charge:   desc.IsCharge,
		Recharge: desc.IsRecharge,
	}
}

func (m *MoveMapper) mapMechanicalTags(desc MoveDescriptor) []move.MechanicalTag {
	tags := make([]move.MechanicalTag, 0)

	if desc.IsAuthentic {
		tags = append(tags, move.MechanicalTagAuthentic)
	}
	if desc.IsBallistics {
		tags = append(tags, move.MechanicalTagBallistic)
	}
	if desc.IsBite {
		tags = append(tags, move.MechanicalTagBite)
	}
	if desc.IsDance {
		tags = append(tags, move.MechanicalTagDance)
	}
	if desc.IsPunch {
		tags = append(tags, move.MechanicalTagPunch)
	}
	if desc.IsSlicingAttack {
		tags = append(tags, move.MechanicalTagSlice)
	}
	if desc.IsSoundAttack {
		tags = append(tags, move.MechanicalTagSound)
	}
	if desc.IsWind {
		tags = append(tags, move.MechanicalTagWind)
	}
	if desc.IsPulse {
		tags = append(tags, move.MechanicalTagPulse)
	}
	if desc.IsPowder {
		tags = append(tags, move.MechanicalTagPowder)
	}
	if desc.IsMental {
		tags = append(tags, move.MechanicalTagMental)
	}

	return tags
}

func (m *MoveMapper) mapInteractions(desc MoveDescriptor) []move.MoveInteraction {
	interactions := make([]move.MoveInteraction, 0)

	if desc.IsBlocable {
		interactions = append(interactions, move.InteractionBlocable)
	}
	if desc.IsMirrorMove {
		interactions = append(interactions, move.InteractionMirrorMove)
	}
	if desc.IsSnatchable {
		interactions = append(interactions, move.InteractionSnatchable)
	}
	if desc.IsMagicCoatAffected {
		interactions = append(interactions, move.InteractionMagicCoatAffected)
	}
	if desc.IsKingRockUtility {
		interactions = append(interactions, move.InteractionKingRockUtility)
	}
	if desc.IsGravity {
		interactions = append(interactions, move.InteractionAffectedByGravity)
	}
	if desc.IsNonSkyBattle {
		interactions = append(interactions, move.InteractionNonSkyBattle)
	}

	return interactions
}

func (m *MoveMapper) mapSecondaryEffects(desc MoveDescriptor) move.MoveSecondaryEffects {
	effects := move.MoveSecondaryEffects{
		Chance: desc.EffectChance,
	}

	// Map status effects
	if len(desc.MoveStatus) > 0 {
		effects.StatusEffects = make([]move.MoveStatusEffect, len(desc.MoveStatus))
		for i, status := range desc.MoveStatus {
			effects.StatusEffects[i] = move.MoveStatusEffect{
				Status:   status.Status,
				LuckRate: status.LuckRate,
			}
		}
	}

	// Map stat stage changes
	if len(desc.BattleStageMod) > 0 {
		effects.StatStageChanges = make([]move.MoveStatStageChange, len(desc.BattleStageMod))
		for i, stageMod := range desc.BattleStageMod {
			effects.StatStageChanges[i] = move.MoveStatStageChange{
				BattleStage: move.BattleStage(stageMod.BattleStage),
				Modificator: stageMod.Modificator,
			}
		}
	}

	return effects
}
