package studio

type MoveCategory string

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

type MoveOption func(*Move)

func WithMoveID(id int) MoveOption {
	return func(m *Move) { m.id = id }
}

func WithMoveDbSymbol(dbSymbol string) MoveOption {
	return func(m *Move) { m.dbSymbol = dbSymbol }
}

func WithMoveType(t *PokemonType) MoveOption {
	return func(m *Move) { m.moveType = t }
}

func WithMoveCategory(cat MoveCategory) MoveOption {
	return func(m *Move) { m.category = cat }
}

func WithMovePower(power int) MoveOption {
	return func(m *Move) { m.power = power }
}

func WithMoveAccuracy(acc int) MoveOption {
	return func(m *Move) { m.accuracy = acc }
}

func WithMovePP(pp int) MoveOption {
	return func(m *Move) { m.pp = pp }
}

func WithMoveCriticalRate(rate int) MoveOption {
	return func(m *Move) { m.criticalRate = rate }
}

func WithMovePriority(priority int) MoveOption {
	return func(m *Move) { m.priority = priority }
}

func WithMoveMapUse(mapUse int) MoveOption {
	return func(m *Move) { m.mapUse = mapUse }
}

func WithMoveName(name Translation) MoveOption {
	return func(m *Move) { m.name = name }
}

func WithMoveDescription(desc Translation) MoveOption {
	return func(m *Move) { m.description = desc }
}

func WithMoveTargeting(targeting MoveTargeting) MoveOption {
	return func(m *Move) { m.targeting = targeting }
}

func WithMoveExecution(exec MoveExecution) MoveOption {
	return func(m *Move) { m.execution = exec }
}

func WithMoveMechanicalTags(tags []MoveMechanicalTag) MoveOption {
	return func(m *Move) { m.mechanicalTags = tags }
}

func WithMoveInteractions(interactions []MoveInteraction) MoveOption {
	return func(m *Move) { m.interactions = interactions }
}

func WithMoveSecondaryEffects(effects MoveSecondaryEffects) MoveOption {
	return func(m *Move) { m.secondaryEffects = effects }
}

func NewMove(opts ...MoveOption) *Move {
	m := &Move{}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

func (m Move) ID() int {
	return m.id
}

func (m Move) DbSymbol() string {
	return m.dbSymbol
}

func (m Move) Type() PokemonType {
	if m.moveType == nil {
		return PokemonType{}
	}
	return *m.moveType
}

func (m Move) Category() MoveCategory {
	return m.category
}

func (m Move) Power() int {
	return m.power
}

func (m Move) Accuracy() int {
	return m.accuracy
}

func (m Move) PP() int {
	return m.pp
}

func (m Move) CriticalRate() int {
	return m.criticalRate
}

func (m Move) Priority() int {
	return m.priority
}

func (m Move) MapUse() int {
	return m.mapUse
}

func (m Move) Name(lang string) string {
	return m.name[lang]
}

func (m Move) Description(lang string) string {
	return m.description[lang]
}

func (m Move) Targeting() MoveTargeting {
	return m.targeting
}

func (m Move) Execution() MoveExecution {
	return m.execution
}

func (m Move) MechanicalTags() []MoveMechanicalTag {
	return m.mechanicalTags
}

func (m Move) Interactions() []MoveInteraction {
	return m.interactions
}

func (m Move) SecondaryEffects() MoveSecondaryEffects {
	return m.secondaryEffects
}

type MoveTargeting struct {
	AimedTarget MoveTarget
	ContactType MoveContactType
}

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

type MoveContactType string

const (
	MoveContactTypeNone        MoveContactType = "none"
	MoveContactTypeMelee       MoveContactType = "melee"
	MoveContactTypeRanged      MoveContactType = "ranged"
	MoveContactTypeMeleeRanged MoveContactType = "melee_and_ranged"
)

type MoveExecution struct {
	Method   MoveMethod
	Charge   bool
	Recharge bool
}

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

type MoveSecondaryEffects struct {
	Chance           int
	StatusEffects    []MoveStatusEffect
	StatStageChanges []MoveStatStageChange
}

type MoveStatusEffect struct {
	Status   MoveStatus
	LuckRate int
}

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

type MoveStatStageChange struct {
	BattleStage BattleStage
	Modificator int
}

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

type MoveMapper struct {
	store *Store
}

func NewMoveMapper(store *Store) *MoveMapper {
	return &MoveMapper{store: store}
}

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

func (m *MoveMapper) mapExecution(desc MoveDescriptor) MoveExecution {
	return MoveExecution{
		Method:   MoveMethod(desc.BattleEngineMethod),
		Charge:   desc.IsCharge,
		Recharge: desc.IsRecharge,
	}
}

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
