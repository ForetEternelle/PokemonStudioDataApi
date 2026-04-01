package studio

// ExperienceTypeDescriptor is a numeric descriptor for experience types.
type ExperienceTypeDescriptor int32

// BreedGroupDescriptor is a numeric descriptor for breed groups.
type BreedGroupDescriptor int32

const (
	ExperienceFastNum        ExperienceTypeDescriptor = 0
	ExperienceMediumFastNum  ExperienceTypeDescriptor = 1
	ExperienceMediumSlowNum  ExperienceTypeDescriptor = 2
	ExperienceSlowNum        ExperienceTypeDescriptor = 3
	ExperienceErraticNum     ExperienceTypeDescriptor = 4
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

// ExperienceTypeMap maps experience type descriptors to strings.
var ExperienceTypeMap = map[ExperienceTypeDescriptor]string{
	ExperienceErraticNum:     ExperienceErratic,
	ExperienceFastNum:        ExperienceFast,
	ExperienceMediumFastNum:  ExperienceMediumFast,
	ExperienceMediumSlowNum:  ExperienceMediumSlow,
	ExperienceSlowNum:        ExperienceSlow,
	ExperienceFluctuatingNum: ExperienceFluctuating,
}

// BreedMap maps breed group descriptors to strings.
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

// PokemonDescriptor is the JSON descriptor for a Pokemon.
type PokemonDescriptor struct {
	ID          int32            `json:"id"`
	DbSymbol    string           `json:"dbSymbol"`
	Forms       []FormDescriptor `json:"forms"`
	Name        Translation
	Description Translation
}

// FormDescriptor is the JSON descriptor for a Pokemon form.
type FormDescriptor struct {
	Form           int32                    `json:"form"`
	Height         float32                  `json:"height"`
	Weight         float32                  `json:"weight"`
	Type1          string                   `json:"type1"`
	Type2          *string                  `json:"type2"`
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
	BreedGroups    []int32                  `json:"breedGroups"`
	HatchSteps     int32                    `json:"hatchSteps"`
	BabyDbSymbol   *string                  `json:"babyDbSymbol"`
	BabyForm       int32                    `json:"babyForm"`
	ItemHeld       []ItemHeldDescriptor     `json:"itemHeld"`
	Abilities      []string                 `json:"abilities"`
	FrontOffsetY   int32                    `json:"frontOffsetY"`
	FormTextId     FormTextIdDescriptor     `json:"formTextId"`
	Name           Translation
	Description    Translation
}

// EvolutionDescriptor is the JSON descriptor for an evolution.
type EvolutionDescriptor struct {
	DbSymbol   string                `json:"dbSymbol"`
	Form       int32                 `json:"form"`
	Conditions []ConditionDescriptor `json:"conditions"`
}

// ConditionDescriptor is the JSON descriptor for an evolution condition.
type ConditionDescriptor struct {
	Type string `json:"type"`
}

// ItemHeldDescriptor is the JSON descriptor for a held item.
type ItemHeldDescriptor struct {
	DbSymbol string `json:"dbSymbol"`
	Chance   int32  `json:"chance"`
}

// FormTextIdDescriptor is the JSON descriptor for form text IDs.
type FormTextIdDescriptor struct {
	Name        int `json:"name"`
	Description int `json:"description"`
}
