package studio

type ExperienceType int32

const (
	ExperienceErraticNum     ExperienceType = 0
	ExperienceFastNum        ExperienceType = 1
	ExperienceMediumFastNum  ExperienceType = 2
	ExperienceMediumSlowNum  ExperienceType = 3
	ExperienceSlowNum        ExperienceType = 4
	ExperienceFluctuatingNum ExperienceType = 5
)

const (
	ExperienceErratic     = "erratic"
	ExperienceFast        = "fast"
	ExperienceMediumFast  = "medium_fast"
	ExperienceMediumSlow  = "medium_slow"
	ExperienceSlow        = "slow"
	ExperienceFluctuating = "fluctuating"
)

type BreedGroup int32

const (
	BreedMonsterNum      BreedGroup = 1
	BreedWater1Num       BreedGroup = 2
	BreedBugNum          BreedGroup = 3
	BreedFlyingNum       BreedGroup = 4
	BreedFieldNum        BreedGroup = 5
	BreedFairyNum        BreedGroup = 6
	BreedGrassNum        BreedGroup = 7
	BreedHumanNum        BreedGroup = 8
	BreedWater3Num       BreedGroup = 9
	BreedMineralNum      BreedGroup = 10
	BreedAmorphousNum    BreedGroup = 11
	BreedWater2Num       BreedGroup = 12
	BreedDittoNum        BreedGroup = 13
	BreedDragonNum       BreedGroup = 14
	BreedUndiscoveredNum BreedGroup = 15
)

const (
	BreedMonster      = "monster"
	BreedWater1       = "water1"
	BreedBug          = "bug"
	BreedFlying       = "flying"
	BreedField        = "field"
	BreedFairy        = "fairy"
	BreedGrass        = "grass"
	BreedHuman        = "human-like"
	BreedWater3       = "water3"
	BreedMineral      = "mineral"
	BreedAmorphous    = "amorphous"
	BreedWater2       = "water2"
	BreedDitto        = "ditto"
	BreedDragon       = "dragon"
	BreedUndiscovered = "undiscovered"
)

var BreedMap = map[BreedGroup]string{
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

var ExperienceTypeMap = map[ExperienceType]string{
	ExperienceErraticNum:     ExperienceErratic,
	ExperienceFastNum:        ExperienceFast,
	ExperienceMediumFastNum:  ExperienceMediumFast,
	ExperienceMediumSlowNum:  ExperienceMediumSlow,
	ExperienceSlowNum:        ExperienceSlow,
	ExperienceFluctuatingNum: ExperienceFluctuating,
}

type Pokemon struct {
	Id       int32
	DbSymbol string
	Forms    []PokemonForm
}

type PokemonForm struct {
	Form           int32
	Height         float32
	Weight         float32
	Type1          string
	Type2          *string
	BaseHp         int32
	BaseAtk        int32
	BaseDfe        int32
	BaseSpd        int32
	BaseAts        int32
	BaseDfs        int32
	EvHp           int32
	EvAtk          int32
	EvDfe          int32
	EvSpd          int32
	EvAts          int32
	EvDfs          int32
	Evolutions     []Evolution
	ExperienceType ExperienceType
	BaseExperience int32
	BaseLoyalty    int32
	CatchRate      int32
	FemaleRate     float32
	BreedGroups    []BreedGroup
	HatchSteps     int32
	BabyDbSymbol   string
	BabyForm       int32
	ItemHeld       []ItemHeld
	Abilities      []string
	FrontOffsetY   int32
	Resources      Resources
	FormTextId     FormTextId

	Name        Translation // from translation file
	Description Translation // from translation file
}

type FormTextId struct {
	Name        int
	Description int
}

type Resources struct {
	Icon           string
	IconShiny      string
	Front          string
	FrontF         string
	FrontShiny     string
	FrontShinyF    string
	Back           string
	BackShiny      string
	Footprint      string
	Character      string
	CharacterShiny string
	Cry            string
	HasFemale      bool
}

type Evolution struct {
	DbSymbol   string
	Form       int32
	Conditions []Condition
}

type Condition struct {
	Type string
}

type ItemHeld struct {
	DbSymbol string
	Chance   int32
}

// ComparePokemonId compare 2 pokemon by their ids
// p1 The first pokemon
// p2 The second pokemon
func ComparePokemonId(p1, p2 *Pokemon) int {
	if p1.Id >= p2.Id {
		return 1
	} else {
		return -1
	}
}
