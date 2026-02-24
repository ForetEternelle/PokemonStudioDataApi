package studio

import "github.com/ForetEternelle/PokemonStudioDataApi/pkg/studio/move"

const (
	ExperienceErratic     = "erratic"
	ExperienceFast        = "fast"
	ExperienceMediumFast  = "medium_fast"
	ExperienceMediumSlow  = "medium_slow"
	ExperienceSlow        = "slow"
	ExperienceFluctuating = "fluctuating"
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

type ExperienceType string
type Translation map[string]string

type Pokemon struct {
	Id               int32
	DbSymbol         string
	Forms            map[int32]PokemonForm
	Name             Translation
	Description      Translation
	CustomProperties map[string]any
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

type PokemonForm struct {
	Form           int32
	Height         float32
	Weight         float32
	Type1          *PokemonType
	Type2          *PokemonType
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
	ExperienceType string
	BaseExperience int32
	BaseLoyalty    int32
	CatchRate      int32
	FemaleRate     float32
	BreedGroups    []string
	HatchSteps     int32
	BabyDbSymbol   *string
	BabyForm       int32
	ItemHeld       []*ItemHeld
	AbilitySymbols []string
	Abilities      []*Ability
	FrontOffsetY   int32

	CustomProperties map[string]any
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

type Ability struct {
	DbSymbol string
	Id       int
	TextID   int

	Name        Translation
	Description Translation
}

type PokemonType struct {
	DbSymbol string
	Color    string
	TextId   int
	Name     Translation
	DamageTo []TypeDamage
}

type TypeDamage struct {
	DefensiveType string
	Factor        float32
}

// MoveCategory represents the category of a move
type MoveCategory string

// Move represents a battle move
type Move struct {
	Id           int
	DbSymbol     string
	Type         *PokemonType
	Category     MoveCategory
	Power        int
	Accuracy     int
	PP           int
	CriticalRate int
	Priority     int
	MapUse       int

	Targeting        move.MoveTargeting
	Execution        move.MoveExecution
	MechanicalTags   []move.MechanicalTag
	Interactions     []move.MoveInteraction
	SecondaryEffects move.MoveSecondaryEffects

	Name        Translation
	Description Translation
}
