package studio

// MoveCategory represents a move category (Physical, Special, Status).
type MoveCategory string

// Move represents a Pokemon move.
type Move struct {
	id           int
	dbSymbol     string
	moveType     *PokemonType
	category     MoveCategory
	power        int
	accuracy     int
	pp           int
	criticalRate int
	priority     int
	mapUse       int

	targeting        MoveTargeting
	execution        MoveExecution
	mechanicalTags   []MoveMechanicalTag
	interactions     []MoveInteraction
	secondaryEffects MoveSecondaryEffects

	name        Translation
	description Translation
}

// ID returns the ID of the Move.
func (m *Move) ID() int {
	return m.id
}

// DbSymbol returns the database symbol of the Move.
func (m *Move) DbSymbol() string {
	return m.dbSymbol
}

// Type returns the type of the Move.
func (m *Move) Type() PokemonType {
	return *m.moveType
}

// Category returns the category of the Move.
func (m *Move) Category() MoveCategory {
	return m.category
}

// Power returns the power of the Move.
func (m *Move) Power() int {
	return m.power
}

// Accuracy returns the accuracy of the Move.
func (m *Move) Accuracy() int {
	return m.accuracy
}

// PP returns the PP of the Move.
func (m *Move) PP() int {
	return m.pp
}

// CriticalRate returns the critical rate of the Move.
func (m *Move) CriticalRate() int {
	return m.criticalRate
}

// Priority returns the priority of the Move.
func (m *Move) Priority() int {
	return m.priority
}

// MapUse returns the map use of the Move.
func (m *Move) MapUse() int {
	return m.mapUse
}

// Name returns the localized name of the Move for the given language.
func (m *Move) Name(lang string) string {
	return m.name[lang]
}

// Description returns the localized description of the Move for the given language.
func (m *Move) Description(lang string) string {
	return m.description[lang]
}

// Targeting returns the targeting of the Move.
func (m *Move) Targeting() MoveTargeting {
	return m.targeting
}

// Execution returns the execution of the Move.
func (m *Move) Execution() MoveExecution {
	return m.execution
}

// MechanicalTags returns the mechanical tags of the Move.
func (m *Move) MechanicalTags() []MoveMechanicalTag {
	return m.mechanicalTags
}

// Interactions returns the interactions of the Move.
func (m *Move) Interactions() []MoveInteraction {
	return m.interactions
}

// SecondaryEffects returns the secondary effects of the Move.
func (m *Move) SecondaryEffects() MoveSecondaryEffects {
	return m.secondaryEffects
}

// AimedTarget represents the target of a move
type AimedTarget string

const (
	AimedTargetAdjacentPokemon    AimedTarget = "adjacent_pokemon"
	AimedTargetAdjacentFoe        AimedTarget = "adjacent_foe"
	AimedTargetAdjacentAllFoe     AimedTarget = "adjacent_all_foe"
	AimedTargetAllFoe             AimedTarget = "all_foe"
	AimedTargetAdjacentAllPokemon AimedTarget = "adjacent_all_pokemon"
	AimedTargetAllPokemon         AimedTarget = "all_pokemon"
	AimedTargetUser               AimedTarget = "user"
	AimedTargetUserOrAdjacentAlly AimedTarget = "user_or_adjacent_ally"
	AimedTargetAdjacentAlly       AimedTarget = "adjacent_ally"
	AimedTargetAllAlly            AimedTarget = "all_ally"
	AimedTargetAllAllyButUser     AimedTarget = "all_ally_but_user"
	AimedTargetAnyOtherPokemon    AimedTarget = "any_other_pokemon"
	AimedTargetRandomFoe          AimedTarget = "random_foe"
)

// ContactType represents how a move makes contact
type ContactType string

const (
	ContactTypeDirect  ContactType = "DIRECT"
	ContactTypeDistant ContactType = "DISTANT"
	ContactTypeNone    ContactType = "NONE"
)

// MoveTargeting contains targeting information for a move
type MoveTargeting struct {
	AimedTarget AimedTarget
	ContactType ContactType
}

// ExecutionMethod represents how a move is executed
type ExecutionMethod string

const (
	ExecutionMethodBasic      ExecutionMethod = "s_basic"
	ExecutionMethodStat       ExecutionMethod = "s_stat"
	ExecutionMethodStatus     ExecutionMethod = "s_status"
	ExecutionMethodMultiHit   ExecutionMethod = "s_multi_hit"
	ExecutionMethod2Hits      ExecutionMethod = "s_2hits"
	ExecutionMethodOHKO       ExecutionMethod = "s_ohko"
	ExecutionMethod2Turns     ExecutionMethod = "s_2turns"
	ExecutionMethodSelfStat   ExecutionMethod = "s_self_stat"
	ExecutionMethodSelfStatus ExecutionMethod = "s_self_status"
)

// MoveExecution contains execution information for a move
type MoveExecution struct {
	Method   ExecutionMethod
	Charge   bool
	Recharge bool
}

type MoveMechanicalTag string

const (
	MechanicalTagAuthentic MoveMechanicalTag = "AUTHENTIC"
	MechanicalTagBallistic MoveMechanicalTag = "BALLISTIC"
	MechanicalTagBite      MoveMechanicalTag = "BITE"
	MechanicalTagDance     MoveMechanicalTag = "DANCE"
	MechanicalTagPunch     MoveMechanicalTag = "PUNCH"
	MechanicalTagSlice     MoveMechanicalTag = "SLICE"
	MechanicalTagSound     MoveMechanicalTag = "SOUND"
	MechanicalTagWind      MoveMechanicalTag = "WIND"
	MechanicalTagPulse     MoveMechanicalTag = "PULSE"
	MechanicalTagPowder    MoveMechanicalTag = "POWDER"
	MechanicalTagMental    MoveMechanicalTag = "MENTAL"
)

// MoveInteraction represents an interaction property of a move
type MoveInteraction string

const (
	InteractionBlocable          MoveInteraction = "BLOCABLE"
	InteractionMirrorMove        MoveInteraction = "MIRROR_MOVE"
	InteractionSnatchable        MoveInteraction = "SNATCHABLE"
	InteractionMagicCoatAffected MoveInteraction = "MAGIC_COAT_AFFECTED"
	InteractionKingRockUtility   MoveInteraction = "KING_ROCK_UTILITY"
	InteractionAffectedByGravity MoveInteraction = "AFFECTED_BY_GRAVITY"
	InteractionNonSkyBattle      MoveInteraction = "NON_SKY_BATTLE"
)

// BattleStage represents a stat stage in battle
type BattleStage string

const (
	BattleStageATK BattleStage = "ATK_STAGE"
	BattleStageDFE BattleStage = "DFE_STAGE"
	BattleStageATS BattleStage = "ATS_STAGE"
	BattleStageDFS BattleStage = "DFS_STAGE"
	BattleStageSPD BattleStage = "SPD_STAGE"
	BattleStageEVA BattleStage = "EVA_STAGE"
	BattleStageACC BattleStage = "ACC_STAGE"
)

// MoveStatStageChange represents a stat stage modification
type MoveStatStageChange struct {
	BattleStage BattleStage
	Modificator int
}

// MoveStatusEffect represents a status effect applied by a move
type MoveStatusEffect struct {
	Status   string
	LuckRate int
}

// MoveSecondaryEffects contains secondary effects of a move
type MoveSecondaryEffects struct {
	Chance           int
	StatusEffects    []MoveStatusEffect
	StatStageChanges []MoveStatStageChange
}
