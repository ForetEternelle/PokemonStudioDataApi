package studio

import (
	"iter"
)

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

// Translation is a map of language codes to translated strings.
type Translation map[string]string

// Pokemon represents a Pokémon species with all its forms.
type Pokemon struct {
	id               int32
	dbSymbol         string
	forms            map[int32]PokemonForm
	name             Translation
	description      Translation
	customProperties map[string]any
}

// ID returns the national ID of the Pokemon.
func (p *Pokemon) ID() int32 {
	return p.id
}

// DbSymbol returns the database symbol of the Pokemon.
func (p *Pokemon) DbSymbol() string {
	return p.dbSymbol
}

// Name returns the localized name of the Pokemon for the given language.
func (p *Pokemon) Name(lang string) string {
	return p.name[lang]
}

// Description returns the localized description of the Pokemon for the given language.
func (p *Pokemon) Description(lang string) string {
	return p.description[lang]
}

// Forms returns an iterator over all forms of the Pokemon.
func (p *Pokemon) Forms() iter.Seq2[int32, PokemonForm] {
	return func(yield func(int32, PokemonForm) bool) {
		for k, v := range p.forms {
			if !yield(k, v) {
				return
			}
		}
	}
}

// Form returns a specific form of the Pokemon by its form number.
func (p *Pokemon) Form(form int32) (PokemonForm, bool) {
	f, ok := p.forms[form]
	return f, ok
}

// CustomProperties returns the custom properties of the Pokemon.
func (p *Pokemon) CustomProperties() map[string]any {
	return p.customProperties
}

// ComparePokemonId compares two Pokemon by their ID.
func ComparePokemonId(p1, p2 *Pokemon) int {
	if p1.id >= p2.id {
		return 1
	}
	return -1
}

// PokemonForm represents a specific form of a Pokemon species.
type PokemonForm struct {
	form             int32
	height           float32
	weight           float32
	type1            *PokemonType
	type2            *PokemonType
	baseHp           int32
	baseAtk          int32
	baseDfe          int32
	baseSpd          int32
	baseAts          int32
	baseDfs          int32
	evHp             int32
	evAtk            int32
	evDfe            int32
	evSpd            int32
	evAts            int32
	evDfs            int32
	evolutions       []Evolution
	experienceType   string
	baseExperience   int32
	baseLoyalty      int32
	catchRate        int32
	femaleRate       float32
	breedGroups      []string
	hatchSteps       int32
	babyDbSymbol     *string
	babyForm         int32
	itemHeld         []*ItemHeld
	abilitySymbols   []string
	abilities        []*Ability
	frontOffsetY     int32
	customProperties map[string]any
}

// Form returns the form number of the PokemonForm.
func (f *PokemonForm) Form() int32 {
	return f.form
}

// Height returns the height of the PokemonForm in meters.
func (f *PokemonForm) Height() float32 {
	return f.height
}

// Weight returns the weight of the PokemonForm in hectograms.
func (f *PokemonForm) Weight() float32 {
	return f.weight
}

// Type1 returns the primary type of the PokemonForm.
func (f *PokemonForm) Type1() PokemonType {
	return *f.type1
}

// Type2 returns the secondary type of the PokemonForm.
func (f *PokemonForm) Type2() (PokemonType, bool) {
	if f.type2 == nil {
		return PokemonType{}, false
	}
	return *f.type2, true
}

// BaseHp returns the base HP of the PokemonForm.
func (f *PokemonForm) BaseHp() int32 {
	return f.baseHp
}

// BaseAtk returns the base Attack of the PokemonForm.
func (f *PokemonForm) BaseAtk() int32 {
	return f.baseAtk
}

// BaseDfe returns the base Defense of the PokemonForm.
func (f *PokemonForm) BaseDfe() int32 {
	return f.baseDfe
}

// BaseSpd returns the base Speed of the PokemonForm.
func (f *PokemonForm) BaseSpd() int32 {
	return f.baseSpd
}

// BaseAts returns the base Special Attack of the PokemonForm.
func (f *PokemonForm) BaseAts() int32 {
	return f.baseAts
}

// BaseDfs returns the base Special Defense of the PokemonForm.
func (f *PokemonForm) BaseDfs() int32 {
	return f.baseDfs
}

// EvHp returns the EV yield for HP of the PokemonForm.
func (f *PokemonForm) EvHp() int32 {
	return f.evHp
}

// EvAtk returns the EV yield for Attack of the PokemonForm.
func (f *PokemonForm) EvAtk() int32 {
	return f.evAtk
}

// EvDfe returns the EV yield for Defense of the PokemonForm.
func (f *PokemonForm) EvDfe() int32 {
	return f.evDfe
}

// EvSpd returns the EV yield for Speed of the PokemonForm.
func (f *PokemonForm) EvSpd() int32 {
	return f.evSpd
}

// EvAts returns the EV yield for Special Attack of the PokemonForm.
func (f *PokemonForm) EvAts() int32 {
	return f.evAts
}

// EvDfs returns the EV yield for Special Defense of the PokemonForm.
func (f *PokemonForm) EvDfs() int32 {
	return f.evDfs
}

// ExperienceType returns the experience type of the PokemonForm.
func (f *PokemonForm) ExperienceType() string {
	return f.experienceType
}

// BaseExperience returns the base experience of the PokemonForm.
func (f *PokemonForm) BaseExperience() int32 {
	return f.baseExperience
}

// BaseLoyalty returns the base loyalty of the PokemonForm.
func (f *PokemonForm) BaseLoyalty() int32 {
	return f.baseLoyalty
}

// CatchRate returns the catch rate of the PokemonForm.
func (f *PokemonForm) CatchRate() int32 {
	return f.catchRate
}

// FemaleRate returns the female rate of the PokemonForm.
func (f *PokemonForm) FemaleRate() float32 {
	return f.femaleRate
}

// HatchSteps returns the hatch steps of the PokemonForm.
func (f *PokemonForm) HatchSteps() int32 {
	return f.hatchSteps
}

// BabyDbSymbol returns the baby Pokemon database symbol of the PokemonForm.
func (f *PokemonForm) BabyDbSymbol() *string {
	return f.babyDbSymbol
}

// BabyForm returns the baby form number of the PokemonForm.
func (f *PokemonForm) BabyForm() int32 {
	return f.babyForm
}

// FrontOffsetY returns the front offset Y of the PokemonForm.
func (f *PokemonForm) FrontOffsetY() int32 {
	return f.frontOffsetY
}

// CustomProperties returns the custom properties of the PokemonForm.
func (f *PokemonForm) CustomProperties() map[string]any {
	return f.customProperties
}

// Evolutions returns an iterator over the evolutions of the PokemonForm.
func (f *PokemonForm) Evolutions() iter.Seq[Evolution] {
	return func(yield func(Evolution) bool) {
		for _, e := range f.evolutions {
			if !yield(e) {
				return
			}
		}
	}
}

// Evolution returns a specific evolution by index.
func (f *PokemonForm) Evolution(i int) (Evolution, bool) {
	if i < 0 || i >= len(f.evolutions) {
		return Evolution{}, false
	}
	return f.evolutions[i], true
}

// ItemHeld returns an iterator over the items held by the PokemonForm.
func (f *PokemonForm) ItemHeld() iter.Seq[*ItemHeld] {
	return func(yield func(*ItemHeld) bool) {
		for _, item := range f.itemHeld {
			if !yield(item) {
				return
			}
		}
	}
}

// Item returns a specific held item by index.
func (f *PokemonForm) Item(i int) (*ItemHeld, bool) {
	if i < 0 || i >= len(f.itemHeld) {
		return nil, false
	}
	return f.itemHeld[i], true
}

// Abilities returns an iterator over the abilities of the PokemonForm.
func (f *PokemonForm) Abilities() iter.Seq[Ability] {
	return func(yield func(Ability) bool) {
		for _, a := range f.abilities {
			if !yield(*a) {
				return
			}
		}
	}
}

// Ability returns a specific ability by index.
func (f *PokemonForm) Ability(i int) (Ability, bool) {
	if i < 0 || i >= len(f.abilities) {
		return Ability{}, false
	}
	return *f.abilities[i], true
}

// AbilitySymbols returns an iterator over the ability symbols of the PokemonForm.
func (f *PokemonForm) AbilitySymbols() iter.Seq[string] {
	return func(yield func(string) bool) {
		for _, s := range f.abilitySymbols {
			if !yield(s) {
				return
			}
		}
	}
}

// AbilitySymbol returns a specific ability symbol by index.
func (f *PokemonForm) AbilitySymbol(i int) (string, bool) {
	if i < 0 || i >= len(f.abilitySymbols) {
		return "", false
	}
	return f.abilitySymbols[i], true
}

// BreedGroups returns an iterator over the breed groups of the PokemonForm.
func (f *PokemonForm) BreedGroups() iter.Seq[string] {
	return func(yield func(string) bool) {
		for _, g := range f.breedGroups {
			if !yield(g) {
				return
			}
		}
	}
}

// BreedGroup returns a specific breed group by index.
func (f *PokemonForm) BreedGroup(i int) (string, bool) {
	if i < 0 || i >= len(f.breedGroups) {
		return "", false
	}
	return f.breedGroups[i], true
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
