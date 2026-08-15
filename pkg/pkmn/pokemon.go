package pkmn

// ExperienceErratic is the erratic experience type.
const (
	ExperienceErratic     = "erratic"
	ExperienceFast        = "fast"
	ExperienceMediumFast  = "medium_fast"
	ExperienceMediumSlow  = "medium_slow"
	ExperienceSlow        = "slow"
	ExperienceFluctuating = "fluctuating"
)

// Breed groups constants.
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

// ExperienceType is a type alias for experience types.
type ExperienceType string

// Pokemon represents a Pokémon species with all its forms.
//
// Entities are immutable by convention only: fields are exported but must not
// be mutated after the entity has been registered in a store.
type Pokemon struct {
	// ID is the national ID of the Pokemon.
	// Immutable: used to order the pokemon list in the store.
	ID int32
	// DbSymbol is the database symbol of the Pokemon.
	// Immutable: used to index pokemon in the store.
	DbSymbol         string
	Forms            []*PokemonForm
	CustomProperties map[string]any
	Tags             []string
}

// Form returns a specific form of the Pokemon by its form number.
func (p *Pokemon) Form(form int32) (*PokemonForm, bool) {
	for _, f := range p.Forms {
		if f.Form == form {
			return f, true
		}
	}
	return nil, false
}

// ComparePokemonId compares two Pokemon by their ID.
func ComparePokemonId(p1, p2 Pokemon) int {
	if p1.ID >= p2.ID {
		return 1
	}
	return -1
}

// PokemonForm represents a specific form of a Pokemon species.
//
// Entities are immutable by convention only: fields are exported but must not
// be mutated after the entity has been registered in a store.
type PokemonForm struct {
	Form             int32
	Height           float32
	Weight           float32
	Type1            *PokemonType
	Type2            *PokemonType
	BaseHp           int32
	BaseAtk          int32
	BaseDfe          int32
	BaseSpd          int32
	BaseAts          int32
	BaseDfs          int32
	EvHp             int32
	EvAtk            int32
	EvDfe            int32
	EvSpd            int32
	EvAts            int32
	EvDfs            int32
	Evolutions       []Evolution
	ExperienceType   string
	BaseExperience   int32
	BaseLoyalty      int32
	CatchRate        int32
	FemaleRate       float32
	BreedGroups      []string
	HatchSteps       int32
	BabyDbSymbol     *string
	BabyForm         int32
	ItemHeld         []*ItemHeld
	AbilitySymbols   []string
	Abilities        []*Ability
	FrontOffsetY     int32
	Name             Translation
	Description      Translation
	CustomProperties map[string]any
	Tags             []string
	Resources        PokemonResources
}

// Evolution represents an evolution from one Pokemon to another.
type Evolution struct {
	DbSymbol   string
	Form       int32
	Conditions []Condition
}

// Condition represents a condition for evolution.
type Condition struct {
	Type string
}

// ItemHeld represents an item held by a Pokemon.
type ItemHeld struct {
	DbSymbol string
	Chance   int32
}

// PokemonResources represents resources for a Pokemon form.
type PokemonResources struct {
	Icon            string
	IconF           string
	IconShiny       string
	IconShinyF      string
	Front           string
	FrontF          string
	FrontShiny      string
	FrontShinyF     string
	Back            string
	BackF           string
	BackShiny       string
	BackShinyF      string
	Footprint       string
	Character       string
	CharacterF      string
	CharacterShiny  string
	CharacterShinyF string
	Cry             string
	HasFemale       bool
	Egg             string
	IconEgg         string
}

// MaxHp calculates the maximum possible HP value for a given base HP.
func MaxHp(base int32) int32 {
	res := float64(base*2 + 204)
	return int32(res)
}

// MinHp calculates the minimum possible HP value for a given base HP.
func MinHp(base int32) int32 {
	res := float64(base*2 + 110)
	return int32(res)
}

// MaxStat calculates the maximum possible stat value for a given base stat.
func MaxStat(base int32) int32 {
	res := 1.1 * float64(base*2+99)
	return int32(res)
}

// MinStat calculates the minimum possible stat value for a given base stat.
func MinStat(base int32) int32 {
	res := 0.9 * float64(base*2+5)
	return int32(res)
}
