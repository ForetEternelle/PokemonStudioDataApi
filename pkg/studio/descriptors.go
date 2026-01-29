package studio

const (
	ExperienceErraticNum     ExperienceTypeDescriptor = 0
	ExperienceFastNum        ExperienceTypeDescriptor = 1
	ExperienceMediumFastNum  ExperienceTypeDescriptor = 2
	ExperienceMediumSlowNum  ExperienceTypeDescriptor = 3
	ExperienceSlowNum        ExperienceTypeDescriptor = 4
	ExperienceFluctuatingNum ExperienceTypeDescriptor = 5
)

const (
	BreedMonsterNum      BreedGroupDescriptor = 1
	BreedWater1Num       BreedGroupDescriptor = 2
	BreedBugNum          BreedGroupDescriptor = 3
	BreedFlyingNum       BreedGroupDescriptor = 4
	BreedFieldNum        BreedGroupDescriptor = 5
	BreedFairyNum        BreedGroupDescriptor = 6
	BreedGrassNum        BreedGroupDescriptor = 7
	BreedHumanNum        BreedGroupDescriptor = 8
	BreedWater3Num       BreedGroupDescriptor = 9
	BreedMineralNum      BreedGroupDescriptor = 10
	BreedAmorphousNum    BreedGroupDescriptor = 11
	BreedWater2Num       BreedGroupDescriptor = 12
	BreedDittoNum        BreedGroupDescriptor = 13
	BreedDragonNum       BreedGroupDescriptor = 14
	BreedUndiscoveredNum BreedGroupDescriptor = 15
)

// BreedGroupDescriptor is the descriptor for breed groups in raw JSON
type BreedGroupDescriptor int32

// ExperienceTypeDescriptor is the descriptor for experience types in raw JSON
type ExperienceTypeDescriptor int32

// PokemonDescriptor represents the raw JSON structure for Pokemon data
type PokemonDescriptor struct {
	ID       int32            `json:"id"`
	DbSymbol string           `json:"dbSymbol"`
	Forms    []FormDescriptor `json:"forms"`
}

// FormDescriptor represents the raw JSON structure for Pokemon forms
type FormDescriptor struct {
	Form           int32                    `json:"form"`
	Height         float32                  `json:"height"`
	Weight         float32                  `json:"weight"`
	Type1          string                   `json:"type1"`
	Type2          *string                  `json:"type2"` // Nullable
	BaseHp         int32                    `json:"baseHp"`
	BaseAtk        int32                    `json:"baseAtk"`
	BaseDfe        int32                    `json:"baseDfe"`
	BaseSpd        int32                    `json:"baseSpd"`
	BaseAts        int32                    `json:"baseAts"`
	BaseDfs        int32                    `json:"baseDfs"`
	EvHp           int32                    `json:"evHp"`
	EvAtk          int32                    `json:"evAtk"`
	EvDfe          int32                    `json:"evDfe"`
	EvSpd          int32                    `json:"evSpd"`
	EvAts          int32                    `json:"evAts"`
	EvDfs          int32                    `json:"evDfs"`
	Evolutions     []EvolutionDescriptor    `json:"evolutions"`
	ExperienceType ExperienceTypeDescriptor `json:"experienceType"`
	BaseExperience int32                    `json:"baseExperience"`
	BaseLoyalty    int32                    `json:"baseLoyalty"`
	CatchRate      int32                    `json:"catchRate"`
	FemaleRate     float32                  `json:"femaleRate"`
	BreedGroups    []int32                  `json:"breedGroups"` // Int32 in JSON, will be mapped to BreedGroup enum
	HatchSteps     int32                    `json:"hatchSteps"`
	BabyDbSymbol   *string                  `json:"babyDbSymbol"` // Nullable
	BabyForm       int32                    `json:"babyForm"`
	ItemHeld       []ItemHeldDescriptor     `json:"itemHeld"`
	Abilities      []string                 `json:"abilities"`
	FrontOffsetY   int32                    `json:"frontOffsetY"`
	FormTextId     FormTextIdDescriptor     `json:"formTextId"`
	Name           Translation
	Description    Translation
}

// EvolutionDescriptor represents the raw JSON structure for evolutions
type EvolutionDescriptor struct {
	DbSymbol   string                `json:"dbSymbol"`
	Form       int32                 `json:"form"`
	Conditions []ConditionDescriptor `json:"conditions"`
}

// ConditionDescriptor represents the raw JSON structure for evolution conditions
type ConditionDescriptor struct {
	Type string `json:"type"`
}

// ItemHeldDescriptor represents the raw JSON structure for held items
type ItemHeldDescriptor struct {
	DbSymbol string `json:"dbSymbol"`
	Chance   int32  `json:"chance"`
}

// FormTextIdDescriptor represents the raw JSON structure for form text IDs
type FormTextIdDescriptor struct {
	Name        int `json:"name"`
	Description int `json:"description"`
}

// PokemonTypeDescriptor represents the raw JSON structure for Pokemon types
type PokemonTypeDescriptor struct {
	DbSymbol string                 `json:"dbSymbol"`
	Color    string                 `json:"color"`
	TextId   int                    `json:"textId"`
	DamageTo []TypeDamageDescriptor `json:"damageTo"`
	Name     Translation
}

// TypeDamageDescriptor represents the raw JSON structure for type damage effectiveness
type TypeDamageDescriptor struct {
	DefensiveType string  `json:"defensiveType"`
	Factor        float32 `json:"factor"`
}

// AbilityDescriptor represents the raw JSON structure for abilities
type AbilityDescriptor struct {
	DbSymbol    string `json:"dbSymbol"`
	Id          int    `json:"id"`
	TextID      int    `json:"textId"`
	Name        Translation
	Description Translation
}

// MoveDescriptor represents the raw JSON structure for moves
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

// BattleStageModDescriptor represents a stat stage modification in raw JSON
type BattleStageModDescriptor struct {
	BattleStage string `json:"battleStage"`
	Modificator int    `json:"modificator"`
}

// MoveStatusDescriptor represents a status effect in raw JSON
type MoveStatusDescriptor struct {
	Status   string `json:"status"`
	LuckRate int    `json:"luckRate"`
}
