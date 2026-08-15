package pkmn

// MoveCategory represents a move category (Physical, Special, Status).
type MoveCategory string

// Move represents a Pokemon move.
//
// Entities are immutable by convention only: fields are exported but must not
// be mutated after the entity has been registered in a store.
type Move struct {
	// DbSymbol is the database symbol of the Move.
	// Immutable: used to index moves in the store.
	DbSymbol         string
	ID               int
	MoveType         *PokemonType
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
