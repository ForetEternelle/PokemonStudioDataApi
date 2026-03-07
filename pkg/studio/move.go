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

// MoveOption is a functional option for configuring a Move.
type MoveOption func(*Move)

// WithMoveID sets the ID of a Move.
func WithMoveID(id int) MoveOption {
	return func(m *Move) { m.id = id }
}

// WithMoveDbSymbol sets the database symbol of a Move.
func WithMoveDbSymbol(dbSymbol string) MoveOption {
	return func(m *Move) { m.dbSymbol = dbSymbol }
}

// WithMoveType sets the type of a Move.
func WithMoveType(t *PokemonType) MoveOption {
	return func(m *Move) { m.moveType = t }
}

// WithMoveCategory sets the category of a Move.
func WithMoveCategory(cat MoveCategory) MoveOption {
	return func(m *Move) { m.category = cat }
}

// WithMovePower sets the power of a Move.
func WithMovePower(power int) MoveOption {
	return func(m *Move) { m.power = power }
}

// WithMoveAccuracy sets the accuracy of a Move.
func WithMoveAccuracy(acc int) MoveOption {
	return func(m *Move) { m.accuracy = acc }
}

// WithMovePP sets the PP of a Move.
func WithMovePP(pp int) MoveOption {
	return func(m *Move) { m.pp = pp }
}

// WithMoveCriticalRate sets the critical rate of a Move.
func WithMoveCriticalRate(rate int) MoveOption {
	return func(m *Move) { m.criticalRate = rate }
}

// WithMovePriority sets the priority of a Move.
func WithMovePriority(priority int) MoveOption {
	return func(m *Move) { m.priority = priority }
}

// WithMoveMapUse sets the map use of a Move.
func WithMoveMapUse(mapUse int) MoveOption {
	return func(m *Move) { m.mapUse = mapUse }
}

// WithMoveName sets the name translations of a Move.
func WithMoveName(name Translation) MoveOption {
	return func(m *Move) { m.name = name }
}

// WithMoveDescription sets the description translations of a Move.
func WithMoveDescription(desc Translation) MoveOption {
	return func(m *Move) { m.description = desc }
}

// WithMoveTargeting sets the targeting of a Move.
func WithMoveTargeting(targeting MoveTargeting) MoveOption {
	return func(m *Move) { m.targeting = targeting }
}

// WithMoveExecution sets the execution of a Move.
func WithMoveExecution(exec MoveExecution) MoveOption {
	return func(m *Move) { m.execution = exec }
}

// WithMoveMechanicalTags sets the mechanical tags of a Move.
func WithMoveMechanicalTags(tags []MoveMechanicalTag) MoveOption {
	return func(m *Move) { m.mechanicalTags = tags }
}

// WithMoveInteractions sets the interactions of a Move.
func WithMoveInteractions(interactions []MoveInteraction) MoveOption {
	return func(m *Move) { m.interactions = interactions }
}

// WithMoveSecondaryEffects sets the secondary effects of a Move.
func WithMoveSecondaryEffects(effects MoveSecondaryEffects) MoveOption {
	return func(m *Move) { m.secondaryEffects = effects }
}

// NewMove creates a new Move with the given options.
func NewMove(opts ...MoveOption) *Move {
	m := &Move{}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// ID returns the ID of the Move.
func (m Move) ID() int {
	return m.id
}

// DbSymbol returns the database symbol of the Move.
func (m Move) DbSymbol() string {
	return m.dbSymbol
}

// Type returns the type of the Move.
func (m Move) Type() PokemonType {
	if m.moveType == nil {
		return PokemonType{}
	}
	return *m.moveType
}

// Category returns the category of the Move.
func (m Move) Category() MoveCategory {
	return m.category
}

// Power returns the power of the Move.
func (m Move) Power() int {
	return m.power
}

// Accuracy returns the accuracy of the Move.
func (m Move) Accuracy() int {
	return m.accuracy
}

// PP returns the PP of the Move.
func (m Move) PP() int {
	return m.pp
}

// CriticalRate returns the critical rate of the Move.
func (m Move) CriticalRate() int {
	return m.criticalRate
}

// Priority returns the priority of the Move.
func (m Move) Priority() int {
	return m.priority
}

// MapUse returns the map use of the Move.
func (m Move) MapUse() int {
	return m.mapUse
}

// Name returns the localized name of the Move for the given language.
func (m Move) Name(lang string) string {
	return m.name[lang]
}

// Description returns the localized description of the Move for the given language.
func (m Move) Description(lang string) string {
	return m.description[lang]
}

// Targeting returns the targeting of the Move.
func (m Move) Targeting() MoveTargeting {
	return m.targeting
}

// Execution returns the execution of the Move.
func (m Move) Execution() MoveExecution {
	return m.execution
}

// MechanicalTags returns the mechanical tags of the Move.
func (m Move) MechanicalTags() []MoveMechanicalTag {
	return m.mechanicalTags
}

// Interactions returns the interactions of the Move.
func (m Move) Interactions() []MoveInteraction {
	return m.interactions
}

// SecondaryEffects returns the secondary effects of the Move.
func (m Move) SecondaryEffects() MoveSecondaryEffects {
	return m.secondaryEffects
}

// MoveTargeting represents the targeting of a Move.
type MoveTargeting struct {
	AimedTarget MoveTarget
	ContactType MoveContactType
}

// MoveTarget represents a move target.
type MoveTarget string

const (
	MoveTargetNone           MoveTarget = "none"
	MoveTargetOpponent       MoveTarget = "opponent"
	MoveTargetOpposingSide   MoveTarget = "opposing_side"
	MoveTargetAllySide       MoveTarget = "ally_side"
	MoveTargetAlly           MoveTarget = "ally"
	MoveTargetAllyTeam       MoveTarget = "ally_team"
	MoveTargetField          MoveTarget = "field"
	MoveTargetFieldSide      MoveTarget = "field_side"
	MoveTargetUser           MoveTarget = "user"
	MoveTargetUserOrAlly     MoveTarget = "user_or_ally"
	MoveTargetRandomOpponent MoveTarget = "random_opponent"
)

// MoveContactType represents a move contact type.
type MoveContactType string

const (
	MoveContactTypeNone        MoveContactType = "none"
	MoveContactTypeMelee       MoveContactType = "melee"
	MoveContactTypeRanged      MoveContactType = "ranged"
	MoveContactTypeMeleeRanged MoveContactType = "melee_and_ranged"
)

// MoveExecution represents the execution of a Move.
type MoveExecution struct {
	Method   MoveMethod
	Charge   bool
	Recharge bool
}

// MoveMethod represents a move method.
type MoveMethod string

const (
	MoveMethodStandard  MoveMethod = "standard"
	MoveMethodOneTurn   MoveMethod = "one_turn"
	MoveMethodCharge    MoveMethod = "charge"
	MoveMethodRecharge  MoveMethod = "recharge"
	MoveMethodMinimize  MoveMethod = "minimize"
	MoveMethodConsume   MoveMethod = "consume"
	MoveMethodUnconsume MoveMethod = "unconsume"
	MoveMethodInstant   MoveMethod = "instant"
	MoveMethodPursuit   MoveMethod = "pursuit"
	MoveMethodMultiTurn MoveMethod = "multi_turn"
)

// MoveMechanicalTag represents a move mechanical tag.
type MoveMechanicalTag string

const (
	MoveMechanicalTagNone      MoveMechanicalTag = ""
	MoveMechanicalTagAuthentic MoveMechanicalTag = "authentic"
	MoveMechanicalTagBallistic MoveMechanicalTag = "ballistic"
	MoveMechanicalTagBite      MoveMechanicalTag = "bite"
	MoveMechanicalTagDance     MoveMechanicalTag = "dance"
	MoveMechanicalTagPunch     MoveMechanicalTag = "punch"
	MoveMechanicalTagSlice     MoveMechanicalTag = "slice"
	MoveMechanicalTagSound     MoveMechanicalTag = "sound"
	MoveMechanicalTagWind      MoveMechanicalTag = "wind"
	MoveMechanicalTagPulse     MoveMechanicalTag = "pulse"
	MoveMechanicalTagPowder    MoveMechanicalTag = "powder"
	MoveMechanicalTagMental    MoveMechanicalTag = "mental"
)

// MoveInteraction represents a move interaction.
type MoveInteraction string

const (
	MoveInteractionNone              MoveInteraction = ""
	MoveInteractionBlocable          MoveInteraction = "BLOCABLE"
	MoveInteractionMirrorMove        MoveInteraction = "MIRROR_MOVE"
	MoveInteractionSnatchable        MoveInteraction = "SNATCHABLE"
	MoveInteractionMagicCoatAffected MoveInteraction = "MAGIC_COAT_AFFECTED"
	MoveInteractionKingRockUtility   MoveInteraction = "KING_ROCK_UTILITY"
	MoveInteractionAffectedByGravity MoveInteraction = "AFFECTED_BY_GRAVITY"
	MoveInteractionNonSkyBattle      MoveInteraction = "NON_SKY_BATTLE"
)

// MoveInteractionType represents a move interaction type.
type MoveInteractionType string

const (
	MoveInteractionTypeNone              MoveInteractionType = "none"
	MoveInteractionTypeProtect           MoveInteractionType = "protect"
	MoveInteractionTypeMist              MoveInteractionType = "mist"
	MoveInteractionTypeScreen            MoveInteractionType = "screen"
	MoveInteractionTypeReflect           MoveInteractionType = "reflect"
	MoveInteractionTypeLightScreen       MoveInteractionType = "lightscreen"
	MoveInteractionTypeSafeguard         MoveInteractionType = "safeguard"
	MoveInteractionTypeBlocable          MoveInteractionType = "blocable"
	MoveInteractionTypeMirrorMove        MoveInteractionType = "mirror_move"
	MoveInteractionTypeSnatchable        MoveInteractionType = "snatchable"
	MoveInteractionTypeMagicCoatAffected MoveInteractionType = "magic_coat_affected"
	MoveInteractionTypeKingRockUtility   MoveInteractionType = "king_rock_utility"
	MoveInteractionTypeAffectedByGravity MoveInteractionType = "affected_by_gravity"
	MoveInteractionTypeNonSkyBattle      MoveInteractionType = "non_sky_battle"
)

// MoveSecondaryEffects represents the secondary effects of a Move.
type MoveSecondaryEffects struct {
	Chance           int
	StatusEffects    []MoveStatusEffect
	StatStageChanges []MoveStatStageChange
}

// MoveStatusEffect represents a status effect from a move.
type MoveStatusEffect struct {
	Status   MoveStatus
	LuckRate int
}

// MoveStatus represents a move status effect.
type MoveStatus string

const (
	MoveStatusNone      MoveStatus = "none"
	MoveStatusPoison    MoveStatus = "poison"
	MoveStatusToxic     MoveStatus = "toxic"
	MoveStatusParalysis MoveStatus = "paralysis"
	MoveStatusSleep     MoveStatus = "sleep"
	MoveStatusFreeze    MoveStatus = "freeze"
	MoveStatusBurn      MoveStatus = "burn"
	MoveStatusBadPoison MoveStatus = "bad_poison"
)

// MoveStatStageChange represents a stat stage change from a move.
type MoveStatStageChange struct {
	BattleStage BattleStage
	Modificator int
}

// BattleStage represents a battle stage.
type BattleStage string

const (
	BattleStageAtk       BattleStage = "atk"
	BattleStageDfe       BattleStage = "dfe"
	BattleStageSpd       BattleStage = "spd"
	BattleStageAts       BattleStage = "ats"
	BattleStageDfs       BattleStage = "dfs"
	BattleStageEva       BattleStage = "eva"
	BattleStageAcc       BattleStage = "acc"
	BattleStageAtkSpd    BattleStage = "atk_spd"
	BattleStageAtkDfe    BattleStage = "atk_dfe"
	BattleStageAtkSpdDfe BattleStage = "atk_spd_dfe"
)

// Type aliases for backward compatibility.
type AimedTarget = MoveTarget
type ContactTypeDirect = MoveContactType
type ContactTypeDistant = MoveContactType
type ContactTypeNone = MoveContactType
type ExecutionMethod = MoveMethod
type MechanicalTag = MoveMechanicalTag
type MechanicalTagAuthentic = MoveMechanicalTag
type MechanicalTagBallistic = MoveMechanicalTag
type MechanicalTagBite = MoveMechanicalTag
type MechanicalTagDance = MoveMechanicalTag
type MechanicalTagPunch = MoveMechanicalTag
type MechanicalTagSlice = MoveMechanicalTag
type MechanicalTagSound = MoveMechanicalTag
type MechanicalTagWind = MoveMechanicalTag
type MechanicalTagPulse = MoveMechanicalTag
type MechanicalTagPowder = MoveMechanicalTag
type MechanicalTagMental = MoveMechanicalTag
type InteractionBlocable = MoveInteractionType
type InteractionMirrorMove = MoveInteractionType
type InteractionSnatchable = MoveInteractionType
type InteractionMagicCoatAffected = MoveInteractionType
type InteractionKingRockUtility = MoveInteractionType
type InteractionAffectedByGravity = MoveInteractionType
type InteractionNonSkyBattle = MoveInteractionType

// MoveDescriptor is the JSON descriptor for a Move.
type MoveDescriptor struct {
	Klass                   string                     `json:"klass"`
	Id                      int                        `json:"id"`
	DbSymbol                string                     `json:"dbSymbol"`
	MapUse                  int                        `json:"mapUse"`
	BattleEngineMethod      string                     `json:"battleEngineMethod"`
	Type                    string                     `json:"type"`
	Power                   int                        `json:"power"`
	Accuracy                int                        `json:"accuracy"`
	PP                      int                        `json:"pp"`
	Category                string                     `json:"category"`
	MoveCriticalRate        int                        `json:"movecriticalRate"`
	Priority                int                        `json:"priority"`
	IsAuthentic             bool                       `json:"isAuthentic"`
	IsBallistics            bool                       `json:"isBallistics"`
	IsBite                  bool                       `json:"isBite"`
	IsBlocable              bool                       `json:"isBlocable"`
	IsCharge                bool                       `json:"isCharge"`
	IsDance                 bool                       `json:"isDance"`
	IsDirect                bool                       `json:"isDirect"`
	IsDistance              bool                       `json:"isDistance"`
	IsEffectChance          bool                       `json:"isEffectChance"`
	IsGravity               bool                       `json:"isGravity"`
	IsHeal                  bool                       `json:"isHeal"`
	IsKingRockUtility       bool                       `json:"isKingRockUtility"`
	IsMagicCoatAffected     bool                       `json:"isMagicCoatAffected"`
	IsMental                bool                       `json:"isMental"`
	IsMirrorMove            bool                       `json:"isMirrorMove"`
	IsNonSkyBattle          bool                       `json:"isNonSkyBattle"`
	IsPowder                bool                       `json:"isPowder"`
	IsPulse                 bool                       `json:"isPulse"`
	IsPunch                 bool                       `json:"isPunch"`
	IsRecharge              bool                       `json:"isRecharge"`
	IsSnatchable            bool                       `json:"isSnatchable"`
	IsSoundAttack           bool                       `json:"isSoundAttack"`
	IsUnfreeze              bool                       `json:"isUnfreeze"`
	BattleEngineAimedTarget string                     `json:"battleEngineAimedTarget"`
	BattleStageMod          []BattleStageModDescriptor `json:"battleStageMod"`
	MoveStatus              []MoveStatusDescriptor     `json:"moveStatus"`
	EffectChance            int                        `json:"effectChance"`
	IsSlicingAttack         bool                       `json:"isSlicingAttack"`
	IsWind                  bool                       `json:"isWind"`
	Name                    Translation
	Description             Translation
}

type BattleStageModDescriptor struct {
	BattleStage string `json:"battleStage"`
	Modificator int    `json:"modificator"`
}

type MoveStatusDescriptor struct {
	Status   string `json:"status"`
	LuckRate int    `json:"luckRate"`
}

// MoveMapper maps Move descriptors to Move entities.
type MoveMapper struct {
	store *Store
}

// NewMoveMapper creates a new MoveMapper.
func NewMoveMapper(store *Store) *MoveMapper {
	return &MoveMapper{store: store}
}

// MapMoveDescriptorToMove maps a MoveDescriptor to a Move.
func (m *MoveMapper) MapMoveDescriptorToMove(desc MoveDescriptor) *Move {
	moveObj := NewMove(
		WithMoveID(desc.Id),
		WithMoveDbSymbol(desc.DbSymbol),
		WithMoveCategory(MoveCategory(desc.Category)),
		WithMovePower(desc.Power),
		WithMoveAccuracy(desc.Accuracy),
		WithMovePP(desc.PP),
		WithMoveCriticalRate(desc.MoveCriticalRate),
		WithMovePriority(desc.Priority),
		WithMoveMapUse(desc.MapUse),
		WithMoveName(desc.Name),
		WithMoveDescription(desc.Description),
		WithMoveType(m.store.FindTypeBySymbol(desc.Type)),
		WithMoveTargeting(m.mapTargeting(desc)),
		WithMoveExecution(m.mapExecution(desc)),
		WithMoveMechanicalTags(m.mapMechanicalTags(desc)),
		WithMoveInteractions(m.mapInteractions(desc)),
		WithMoveSecondaryEffects(m.mapSecondaryEffects(desc)),
	)

	return moveObj
}

// mapTargeting maps a MoveDescriptor to MoveTargeting.
func (m *MoveMapper) mapTargeting(desc MoveDescriptor) MoveTargeting {
	targeting := MoveTargeting{
		AimedTarget: AimedTarget(desc.BattleEngineAimedTarget),
	}

	if desc.IsDirect {
		targeting.ContactType = MoveContactTypeMelee
	} else if desc.IsDistance {
		targeting.ContactType = MoveContactTypeRanged
	} else {
		targeting.ContactType = MoveContactTypeNone
	}

	return targeting
}

// mapExecution maps a MoveDescriptor to MoveExecution.
func (m *MoveMapper) mapExecution(desc MoveDescriptor) MoveExecution {
	return MoveExecution{
		Method:   MoveMethod(desc.BattleEngineMethod),
		Charge:   desc.IsCharge,
		Recharge: desc.IsRecharge,
	}
}

// mapMechanicalTags maps a MoveDescriptor to MoveMechanicalTags.
func (m *MoveMapper) mapMechanicalTags(desc MoveDescriptor) []MoveMechanicalTag {
	tags := make([]MoveMechanicalTag, 0)

	if desc.IsAuthentic {
		tags = append(tags, MoveMechanicalTagAuthentic)
	}
	if desc.IsBallistics {
		tags = append(tags, MoveMechanicalTagBallistic)
	}
	if desc.IsBite {
		tags = append(tags, MoveMechanicalTagBite)
	}
	if desc.IsDance {
		tags = append(tags, MoveMechanicalTagDance)
	}
	if desc.IsPunch {
		tags = append(tags, MoveMechanicalTagPunch)
	}
	if desc.IsSlicingAttack {
		tags = append(tags, MoveMechanicalTagSlice)
	}
	if desc.IsSoundAttack {
		tags = append(tags, MoveMechanicalTagSound)
	}
	if desc.IsWind {
		tags = append(tags, MoveMechanicalTagWind)
	}
	if desc.IsPulse {
		tags = append(tags, MoveMechanicalTagPulse)
	}
	if desc.IsPowder {
		tags = append(tags, MoveMechanicalTagPowder)
	}
	if desc.IsMental {
		tags = append(tags, MoveMechanicalTagMental)
	}

	return tags
}

// mapInteractions maps a MoveDescriptor to MoveInteractions.
func (m *MoveMapper) mapInteractions(desc MoveDescriptor) []MoveInteraction {
	interactions := make([]MoveInteraction, 0)

	if desc.IsBlocable {
		interactions = append(interactions, MoveInteractionBlocable)
	}
	if desc.IsMirrorMove {
		interactions = append(interactions, MoveInteractionMirrorMove)
	}
	if desc.IsSnatchable {
		interactions = append(interactions, MoveInteractionSnatchable)
	}
	if desc.IsMagicCoatAffected {
		interactions = append(interactions, MoveInteractionMagicCoatAffected)
	}
	if desc.IsKingRockUtility {
		interactions = append(interactions, MoveInteractionKingRockUtility)
	}
	if desc.IsGravity {
		interactions = append(interactions, MoveInteractionAffectedByGravity)
	}
	if desc.IsNonSkyBattle {
		interactions = append(interactions, MoveInteractionNonSkyBattle)
	}

	return interactions
}

// mapSecondaryEffects maps a MoveDescriptor to MoveSecondaryEffects.
func (m *MoveMapper) mapSecondaryEffects(desc MoveDescriptor) MoveSecondaryEffects {
	effects := MoveSecondaryEffects{
		Chance: desc.EffectChance,
	}

	if len(desc.MoveStatus) > 0 {
		effects.StatusEffects = make([]MoveStatusEffect, len(desc.MoveStatus))
		for i, status := range desc.MoveStatus {
			effects.StatusEffects[i] = MoveStatusEffect{
				Status:   MoveStatus(status.Status),
				LuckRate: status.LuckRate,
			}
		}
	}

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
