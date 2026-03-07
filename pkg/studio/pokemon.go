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

// ExperienceTypeDescriptor is a numeric descriptor for experience types.
type ExperienceTypeDescriptor int32

// BreedGroupDescriptor is a numeric descriptor for breed groups.
type BreedGroupDescriptor int32

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

// Pokemon represents a Pokémon species with all its forms.
type Pokemon struct {
	id               int32
	dbSymbol         string
	forms            map[int32]PokemonForm
	name             Translation
	description      Translation
	customProperties map[string]any
}

// PokemonOption is a functional option for configuring a Pokemon.
type PokemonOption func(*Pokemon)

// WithID sets the ID of a Pokemon.
func WithID(id int32) PokemonOption {
	return func(p *Pokemon) { p.id = id }
}

// WithDbSymbol sets the database symbol of a Pokemon.
func WithDbSymbol(dbSymbol string) PokemonOption {
	return func(p *Pokemon) { p.dbSymbol = dbSymbol }
}

// WithName sets the name translations of a Pokemon.
func WithName(name Translation) PokemonOption {
	return func(p *Pokemon) { p.name = name }
}

// WithDescription sets the description translations of a Pokemon.
func WithDescription(desc Translation) PokemonOption {
	return func(p *Pokemon) { p.description = desc }
}

// WithForms sets the forms of a Pokemon.
func WithForms(forms map[int32]PokemonForm) PokemonOption {
	return func(p *Pokemon) { p.forms = forms }
}

// WithCustomProperties sets custom properties for a Pokemon.
func WithCustomProperties(props map[string]any) PokemonOption {
	return func(p *Pokemon) { p.customProperties = props }
}

// NewPokemon creates a new Pokemon with the given options.
func NewPokemon(opts ...PokemonOption) *Pokemon {
	p := &Pokemon{
		customProperties: make(map[string]any),
	}
	for _, opt := range opts {
		opt(p)
	}
	return p
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

// PokemonFormOption is a functional option for configuring a PokemonForm.
type PokemonFormOption func(*PokemonForm)

// WithForm sets the form number of a PokemonForm.
func WithForm(form int32) PokemonFormOption {
	return func(f *PokemonForm) { f.form = form }
}

// WithType1 sets the primary type of a PokemonForm.
func WithType1(t *PokemonType) PokemonFormOption {
	return func(f *PokemonForm) { f.type1 = t }
}

// WithType2 sets the secondary type of a PokemonForm.
func WithType2(t *PokemonType) PokemonFormOption {
	return func(f *PokemonForm) { f.type2 = t }
}

// WithHeight sets the height of a PokemonForm.
func WithHeight(h float32) PokemonFormOption {
	return func(f *PokemonForm) { f.height = h }
}

// WithWeight sets the weight of a PokemonForm.
func WithWeight(w float32) PokemonFormOption {
	return func(f *PokemonForm) { f.weight = w }
}

// WithBaseHp sets the base HP of a PokemonForm.
func WithBaseHp(hp int32) PokemonFormOption {
	return func(f *PokemonForm) { f.baseHp = hp }
}

// WithBaseAtk sets the base Attack of a PokemonForm.
func WithBaseAtk(atk int32) PokemonFormOption {
	return func(f *PokemonForm) { f.baseAtk = atk }
}

// WithBaseDfe sets the base Defense of a PokemonForm.
func WithBaseDfe(dfe int32) PokemonFormOption {
	return func(f *PokemonForm) { f.baseDfe = dfe }
}

// WithBaseSpd sets the base Speed of a PokemonForm.
func WithBaseSpd(spd int32) PokemonFormOption {
	return func(f *PokemonForm) { f.baseSpd = spd }
}

// WithBaseAts sets the base Special Attack of a PokemonForm.
func WithBaseAts(ats int32) PokemonFormOption {
	return func(f *PokemonForm) { f.baseAts = ats }
}

// WithBaseDfs sets the base Special Defense of a PokemonForm.
func WithBaseDfs(dfs int32) PokemonFormOption {
	return func(f *PokemonForm) { f.baseDfs = dfs }
}

// WithEvHp sets the EV yield for HP of a PokemonForm.
func WithEvHp(hp int32) PokemonFormOption {
	return func(f *PokemonForm) { f.evHp = hp }
}

// WithEvAtk sets the EV yield for Attack of a PokemonForm.
func WithEvAtk(atk int32) PokemonFormOption {
	return func(f *PokemonForm) { f.evAtk = atk }
}

// WithEvDfe sets the EV yield for Defense of a PokemonForm.
func WithEvDfe(dfe int32) PokemonFormOption {
	return func(f *PokemonForm) { f.evDfe = dfe }
}

// WithEvSpd sets the EV yield for Speed of a PokemonForm.
func WithEvSpd(spd int32) PokemonFormOption {
	return func(f *PokemonForm) { f.evSpd = spd }
}

// WithEvAts sets the EV yield for Special Attack of a PokemonForm.
func WithEvAts(ats int32) PokemonFormOption {
	return func(f *PokemonForm) { f.evAts = ats }
}

// WithEvDfs sets the EV yield for Special Defense of a PokemonForm.
func WithEvDfs(dfs int32) PokemonFormOption {
	return func(f *PokemonForm) { f.evDfs = dfs }
}

// WithEvolutions sets the evolutions of a PokemonForm.
func WithEvolutions(evolutions []Evolution) PokemonFormOption {
	return func(f *PokemonForm) { f.evolutions = evolutions }
}

// WithExperienceType sets the experience type of a PokemonForm.
func WithExperienceType(expType string) PokemonFormOption {
	return func(f *PokemonForm) { f.experienceType = expType }
}

// WithBaseExperience sets the base experience of a PokemonForm.
func WithBaseExperience(exp int32) PokemonFormOption {
	return func(f *PokemonForm) { f.baseExperience = exp }
}

// WithBaseLoyalty sets the base loyalty of a PokemonForm.
func WithBaseLoyalty(loyalty int32) PokemonFormOption {
	return func(f *PokemonForm) { f.baseLoyalty = loyalty }
}

// WithCatchRate sets the catch rate of a PokemonForm.
func WithCatchRate(rate int32) PokemonFormOption {
	return func(f *PokemonForm) { f.catchRate = rate }
}

// WithFemaleRate sets the female rate of a PokemonForm.
func WithFemaleRate(rate float32) PokemonFormOption {
	return func(f *PokemonForm) { f.femaleRate = rate }
}

// WithBreedGroups sets the breed groups of a PokemonForm.
func WithBreedGroups(groups []string) PokemonFormOption {
	return func(f *PokemonForm) { f.breedGroups = groups }
}

// WithHatchSteps sets the hatch steps of a PokemonForm.
func WithHatchSteps(steps int32) PokemonFormOption {
	return func(f *PokemonForm) { f.hatchSteps = steps }
}

// WithBabyDbSymbol sets the baby Pokemon database symbol of a PokemonForm.
func WithBabyDbSymbol(symbol *string) PokemonFormOption {
	return func(f *PokemonForm) { f.babyDbSymbol = symbol }
}

// WithBabyForm sets the baby form number of a PokemonForm.
func WithBabyForm(form int32) PokemonFormOption {
	return func(f *PokemonForm) { f.babyForm = form }
}

// WithItemHeld sets the items held by a PokemonForm.
func WithItemHeld(items []*ItemHeld) PokemonFormOption {
	return func(f *PokemonForm) { f.itemHeld = items }
}

// WithAbilitySymbols sets the ability symbols of a PokemonForm.
func WithAbilitySymbols(symbols []string) PokemonFormOption {
	return func(f *PokemonForm) { f.abilitySymbols = symbols }
}

// WithAbilities sets the abilities of a PokemonForm.
func WithAbilities(abilities []*Ability) PokemonFormOption {
	return func(f *PokemonForm) { f.abilities = abilities }
}

// WithFrontOffsetY sets the front offset Y of a PokemonForm.
func WithFrontOffsetY(y int32) PokemonFormOption {
	return func(f *PokemonForm) { f.frontOffsetY = y }
}

// WithFormCustomProperties sets custom properties for a PokemonForm.
func WithFormCustomProperties(props map[string]any) PokemonFormOption {
	return func(f *PokemonForm) { f.customProperties = props }
}

// NewPokemonForm creates a new PokemonForm with the given options.
func NewPokemonForm(opts ...PokemonFormOption) *PokemonForm {
	f := &PokemonForm{
		customProperties: make(map[string]any),
	}
	for _, opt := range opts {
		opt(f)
	}
	return f
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
func (f *PokemonForm) Evolutions() iter.Seq2[int, Evolution] {
	return func(yield func(int, Evolution) bool) {
		for i, e := range f.evolutions {
			if !yield(i, e) {
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
func (f *PokemonForm) ItemHeld() iter.Seq2[int, *ItemHeld] {
	return func(yield func(int, *ItemHeld) bool) {
		for i, item := range f.itemHeld {
			if !yield(i, item) {
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
func (f *PokemonForm) Abilities() iter.Seq2[int, Ability] {
	return func(yield func(int, Ability) bool) {
		for i, a := range f.abilities {
			if !yield(i, *a) {
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
func (f *PokemonForm) AbilitySymbols() iter.Seq2[int, string] {
	return func(yield func(int, string) bool) {
		for i, s := range f.abilitySymbols {
			if !yield(i, s) {
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
func (f *PokemonForm) BreedGroups() iter.Seq2[int, string] {
	return func(yield func(int, string) bool) {
		for i, g := range f.breedGroups {
			if !yield(i, g) {
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

// PokemonMapper maps Pokemon descriptors to Pokemon entities.
type PokemonMapper struct {
	store *Store
}

// NewPokemonMapper creates a new PokemonMapper.
func NewPokemonMapper(store *Store) *PokemonMapper {
	return &PokemonMapper{store: store}
}

// MapPokemonDescriptorToPokemon maps a PokemonDescriptor to a Pokemon.
func (m *PokemonMapper) MapPokemonDescriptorToPokemon(desc PokemonDescriptor) *Pokemon {
	forms := make(map[int32]PokemonForm, len(desc.Forms))
	for _, formDesc := range desc.Forms {
		forms[formDesc.Form] = *m.MapFormDescriptorToPokemonForm(formDesc)
	}

	pokemon := NewPokemon(
		WithID(desc.ID),
		WithDbSymbol(desc.DbSymbol),
		WithName(desc.Name),
		WithDescription(desc.Description),
		WithForms(forms),
	)

	return pokemon
}

// MapFormDescriptorToPokemonForm maps a FormDescriptor to a PokemonForm.
func (m *PokemonMapper) MapFormDescriptorToPokemonForm(desc FormDescriptor) *PokemonForm {
	form := NewPokemonForm(
		WithForm(desc.Form),
		WithType1(m.store.FindTypeBySymbol(desc.Type1)),
		WithHeight(desc.Height),
		WithWeight(desc.Weight),
		WithBaseHp(desc.BaseHp),
		WithBaseAtk(desc.BaseAtk),
		WithBaseDfe(desc.BaseDfe),
		WithBaseSpd(desc.BaseSpd),
		WithBaseAts(desc.BaseAts),
		WithBaseDfs(desc.BaseDfs),
		WithEvHp(desc.EvHp),
		WithEvAtk(desc.EvAtk),
		WithEvDfe(desc.EvDfe),
		WithEvSpd(desc.EvSpd),
		WithEvAts(desc.EvAts),
		WithEvDfs(desc.EvDfs),
		WithExperienceType(ExperienceTypeMap[desc.ExperienceType]),
		WithBaseExperience(desc.BaseExperience),
		WithBaseLoyalty(desc.BaseLoyalty),
		WithCatchRate(desc.CatchRate),
		WithFemaleRate(desc.FemaleRate),
		WithBreedGroups(m.MapBreedGroups(desc.BreedGroups)),
		WithHatchSteps(desc.HatchSteps),
		WithBabyForm(desc.BabyForm),
		WithBabyDbSymbol(desc.BabyDbSymbol),
		WithFrontOffsetY(desc.FrontOffsetY),
		WithEvolutions(m.MapEvolutions(desc.Evolutions)),
		WithItemHeld(m.MapItemHelds(desc.ItemHeld)),
		WithFormCustomProperties(make(map[string]any)),
	)

	if desc.Type2 != nil && *desc.Type2 != "" && *desc.Type2 != UndefType {
		form.type2 = m.store.FindTypeBySymbol(*desc.Type2)
	}

	abilities := make([]*Ability, 0, len(desc.Abilities))
	for _, abilitySymbol := range desc.Abilities {
		if ability := m.store.FindAbilityBySymbol(abilitySymbol); ability != nil {
			abilities = append(abilities, ability)
		}
	}
	form.abilities = abilities

	return form
}

// MapBreedGroups maps breed group descriptors to breed group strings.
func (m *PokemonMapper) MapBreedGroups(breedGroupInts []int32) []string {
	breedGroups := make([]string, len(breedGroupInts))
	for i, bgInt := range breedGroupInts {
		breedGroups[i] = BreedMap[BreedGroupDescriptor(bgInt)]
	}
	return breedGroups
}

// MapEvolutions maps evolution descriptors to evolutions.
func (m *PokemonMapper) MapEvolutions(evolutions []EvolutionDescriptor) []Evolution {
	if len(evolutions) == 0 {
		return nil
	}

	mapped := make([]Evolution, len(evolutions))
	for i, evoDesc := range evolutions {
		mapped[i] = Evolution{
			DbSymbol:   evoDesc.DbSymbol,
			Form:       evoDesc.Form,
			Conditions: m.MapConditions(evoDesc.Conditions),
		}
	}
	return mapped
}

// MapConditions maps condition descriptors to conditions.
func (m *PokemonMapper) MapConditions(conditions []ConditionDescriptor) []Condition {
	if len(conditions) == 0 {
		return nil
	}

	mapped := make([]Condition, len(conditions))
	for i, condDesc := range conditions {
		mapped[i] = Condition{Type: condDesc.Type}
	}
	return mapped
}

// MapItemHelds maps item held descriptors to item held entities.
func (m *PokemonMapper) MapItemHelds(itemHelds []ItemHeldDescriptor) []*ItemHeld {
	if len(itemHelds) == 0 {
		return nil
	}

	mapped := make([]*ItemHeld, len(itemHelds))
	for i, ihDesc := range itemHelds {
		mapped[i] = &ItemHeld{
			DbSymbol: ihDesc.DbSymbol,
			Chance:   ihDesc.Chance,
		}
	}
	return mapped
}
