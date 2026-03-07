package studio

import (
	"iter"
)

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

type ExperienceTypeDescriptor int32
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

type Pokemon struct {
	id               int32
	dbSymbol         string
	forms            map[int32]PokemonForm
	name             Translation
	description      Translation
	customProperties map[string]any
}

type PokemonOption func(*Pokemon)

func WithID(id int32) PokemonOption {
	return func(p *Pokemon) { p.id = id }
}

func WithDbSymbol(dbSymbol string) PokemonOption {
	return func(p *Pokemon) { p.dbSymbol = dbSymbol }
}

func WithName(name Translation) PokemonOption {
	return func(p *Pokemon) { p.name = name }
}

func WithDescription(desc Translation) PokemonOption {
	return func(p *Pokemon) { p.description = desc }
}

func WithForms(forms map[int32]PokemonForm) PokemonOption {
	return func(p *Pokemon) { p.forms = forms }
}

func WithCustomProperties(props map[string]any) PokemonOption {
	return func(p *Pokemon) { p.customProperties = props }
}

func NewPokemon(opts ...PokemonOption) *Pokemon {
	p := &Pokemon{
		customProperties: make(map[string]any),
	}
	for _, opt := range opts {
		opt(p)
	}
	return p
}

func (p *Pokemon) ID() int32 {
	return p.id
}

func (p *Pokemon) DbSymbol() string {
	return p.dbSymbol
}

func (p *Pokemon) Name(lang string) string {
	return p.name[lang]
}

func (p *Pokemon) Description(lang string) string {
	return p.description[lang]
}

func (p *Pokemon) Forms() iter.Seq2[int32, PokemonForm] {
	return func(yield func(int32, PokemonForm) bool) {
		for k, v := range p.forms {
			if !yield(k, v) {
				return
			}
		}
	}
}

func (p *Pokemon) Form(form int32) (PokemonForm, bool) {
	f, ok := p.forms[form]
	return f, ok
}

func (p *Pokemon) CustomProperties() map[string]any {
	return p.customProperties
}

func ComparePokemonId(p1, p2 *Pokemon) int {
	if p1.id >= p2.id {
		return 1
	}
	return -1
}

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

type PokemonFormOption func(*PokemonForm)

func WithForm(form int32) PokemonFormOption {
	return func(f *PokemonForm) { f.form = form }
}

func WithType1(t *PokemonType) PokemonFormOption {
	return func(f *PokemonForm) { f.type1 = t }
}

func WithType2(t *PokemonType) PokemonFormOption {
	return func(f *PokemonForm) { f.type2 = t }
}

func WithHeight(h float32) PokemonFormOption {
	return func(f *PokemonForm) { f.height = h }
}

func WithWeight(w float32) PokemonFormOption {
	return func(f *PokemonForm) { f.weight = w }
}

func WithBaseHp(hp int32) PokemonFormOption {
	return func(f *PokemonForm) { f.baseHp = hp }
}

func WithBaseAtk(atk int32) PokemonFormOption {
	return func(f *PokemonForm) { f.baseAtk = atk }
}

func WithBaseDfe(dfe int32) PokemonFormOption {
	return func(f *PokemonForm) { f.baseDfe = dfe }
}

func WithBaseSpd(spd int32) PokemonFormOption {
	return func(f *PokemonForm) { f.baseSpd = spd }
}

func WithBaseAts(ats int32) PokemonFormOption {
	return func(f *PokemonForm) { f.baseAts = ats }
}

func WithBaseDfs(dfs int32) PokemonFormOption {
	return func(f *PokemonForm) { f.baseDfs = dfs }
}

func WithEvHp(hp int32) PokemonFormOption {
	return func(f *PokemonForm) { f.evHp = hp }
}

func WithEvAtk(atk int32) PokemonFormOption {
	return func(f *PokemonForm) { f.evAtk = atk }
}

func WithEvDfe(dfe int32) PokemonFormOption {
	return func(f *PokemonForm) { f.evDfe = dfe }
}

func WithEvSpd(spd int32) PokemonFormOption {
	return func(f *PokemonForm) { f.evSpd = spd }
}

func WithEvAts(ats int32) PokemonFormOption {
	return func(f *PokemonForm) { f.evAts = ats }
}

func WithEvDfs(dfs int32) PokemonFormOption {
	return func(f *PokemonForm) { f.evDfs = dfs }
}

func WithEvolutions(evolutions []Evolution) PokemonFormOption {
	return func(f *PokemonForm) { f.evolutions = evolutions }
}

func WithExperienceType(expType string) PokemonFormOption {
	return func(f *PokemonForm) { f.experienceType = expType }
}

func WithBaseExperience(exp int32) PokemonFormOption {
	return func(f *PokemonForm) { f.baseExperience = exp }
}

func WithBaseLoyalty(loyalty int32) PokemonFormOption {
	return func(f *PokemonForm) { f.baseLoyalty = loyalty }
}

func WithCatchRate(rate int32) PokemonFormOption {
	return func(f *PokemonForm) { f.catchRate = rate }
}

func WithFemaleRate(rate float32) PokemonFormOption {
	return func(f *PokemonForm) { f.femaleRate = rate }
}

func WithBreedGroups(groups []string) PokemonFormOption {
	return func(f *PokemonForm) { f.breedGroups = groups }
}

func WithHatchSteps(steps int32) PokemonFormOption {
	return func(f *PokemonForm) { f.hatchSteps = steps }
}

func WithBabyDbSymbol(symbol *string) PokemonFormOption {
	return func(f *PokemonForm) { f.babyDbSymbol = symbol }
}

func WithBabyForm(form int32) PokemonFormOption {
	return func(f *PokemonForm) { f.babyForm = form }
}

func WithItemHeld(items []*ItemHeld) PokemonFormOption {
	return func(f *PokemonForm) { f.itemHeld = items }
}

func WithAbilitySymbols(symbols []string) PokemonFormOption {
	return func(f *PokemonForm) { f.abilitySymbols = symbols }
}

func WithAbilities(abilities []*Ability) PokemonFormOption {
	return func(f *PokemonForm) { f.abilities = abilities }
}

func WithFrontOffsetY(y int32) PokemonFormOption {
	return func(f *PokemonForm) { f.frontOffsetY = y }
}

func WithFormCustomProperties(props map[string]any) PokemonFormOption {
	return func(f *PokemonForm) { f.customProperties = props }
}

func NewPokemonForm(opts ...PokemonFormOption) *PokemonForm {
	f := &PokemonForm{
		customProperties: make(map[string]any),
	}
	for _, opt := range opts {
		opt(f)
	}
	return f
}

func (f *PokemonForm) Form() int32 {
	return f.form
}

func (f *PokemonForm) Height() float32 {
	return f.height
}

func (f *PokemonForm) Weight() float32 {
	return f.weight
}

func (f *PokemonForm) Type1() PokemonType {
	return *f.type1
}

func (f *PokemonForm) Type2() PokemonType {
	if f.type2 == nil {
		return PokemonType{}
	}
	return *f.type2
}

func (f *PokemonForm) BaseHp() int32 {
	return f.baseHp
}

func (f *PokemonForm) BaseAtk() int32 {
	return f.baseAtk
}

func (f *PokemonForm) BaseDfe() int32 {
	return f.baseDfe
}

func (f *PokemonForm) BaseSpd() int32 {
	return f.baseSpd
}

func (f *PokemonForm) BaseAts() int32 {
	return f.baseAts
}

func (f *PokemonForm) BaseDfs() int32 {
	return f.baseDfs
}

func (f *PokemonForm) EvHp() int32 {
	return f.evHp
}

func (f *PokemonForm) EvAtk() int32 {
	return f.evAtk
}

func (f *PokemonForm) EvDfe() int32 {
	return f.evDfe
}

func (f *PokemonForm) EvSpd() int32 {
	return f.evSpd
}

func (f *PokemonForm) EvAts() int32 {
	return f.evAts
}

func (f *PokemonForm) EvDfs() int32 {
	return f.evDfs
}

func (f *PokemonForm) ExperienceType() string {
	return f.experienceType
}

func (f *PokemonForm) BaseExperience() int32 {
	return f.baseExperience
}

func (f *PokemonForm) BaseLoyalty() int32 {
	return f.baseLoyalty
}

func (f *PokemonForm) CatchRate() int32 {
	return f.catchRate
}

func (f *PokemonForm) FemaleRate() float32 {
	return f.femaleRate
}

func (f *PokemonForm) HatchSteps() int32 {
	return f.hatchSteps
}

func (f *PokemonForm) BabyDbSymbol() *string {
	return f.babyDbSymbol
}

func (f *PokemonForm) BabyForm() int32 {
	return f.babyForm
}

func (f *PokemonForm) FrontOffsetY() int32 {
	return f.frontOffsetY
}

func (f *PokemonForm) CustomProperties() map[string]any {
	return f.customProperties
}

func (f *PokemonForm) Evolutions() iter.Seq2[int, Evolution] {
	return func(yield func(int, Evolution) bool) {
		for i, e := range f.evolutions {
			if !yield(i, e) {
				return
			}
		}
	}
}

func (f *PokemonForm) Evolution(i int) (Evolution, bool) {
	if i < 0 || i >= len(f.evolutions) {
		return Evolution{}, false
	}
	return f.evolutions[i], true
}

func (f *PokemonForm) ItemHeld() iter.Seq2[int, *ItemHeld] {
	return func(yield func(int, *ItemHeld) bool) {
		for i, item := range f.itemHeld {
			if !yield(i, item) {
				return
			}
		}
	}
}

func (f *PokemonForm) Item(i int) (*ItemHeld, bool) {
	if i < 0 || i >= len(f.itemHeld) {
		return nil, false
	}
	return f.itemHeld[i], true
}

func (f *PokemonForm) Abilities() iter.Seq2[int, Ability] {
	return func(yield func(int, Ability) bool) {
		for i, a := range f.abilities {
			if !yield(i, *a) {
				return
			}
		}
	}
}

func (f *PokemonForm) Ability(i int) (Ability, bool) {
	if i < 0 || i >= len(f.abilities) {
		return Ability{}, false
	}
	return *f.abilities[i], true
}

func (f *PokemonForm) AbilitySymbols() iter.Seq2[int, string] {
	return func(yield func(int, string) bool) {
		for i, s := range f.abilitySymbols {
			if !yield(i, s) {
				return
			}
		}
	}
}

func (f *PokemonForm) AbilitySymbol(i int) (string, bool) {
	if i < 0 || i >= len(f.abilitySymbols) {
		return "", false
	}
	return f.abilitySymbols[i], true
}

func (f *PokemonForm) BreedGroups() iter.Seq2[int, string] {
	return func(yield func(int, string) bool) {
		for i, g := range f.breedGroups {
			if !yield(i, g) {
				return
			}
		}
	}
}

func (f *PokemonForm) BreedGroup(i int) (string, bool) {
	if i < 0 || i >= len(f.breedGroups) {
		return "", false
	}
	return f.breedGroups[i], true
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

type PokemonDescriptor struct {
	ID          int32            `json:"id"`
	DbSymbol    string           `json:"dbSymbol"`
	Forms       []FormDescriptor `json:"forms"`
	Name        Translation
	Description Translation
}

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

type EvolutionDescriptor struct {
	DbSymbol   string                `json:"dbSymbol"`
	Form       int32                 `json:"form"`
	Conditions []ConditionDescriptor `json:"conditions"`
}

type ConditionDescriptor struct {
	Type string `json:"type"`
}

type ItemHeldDescriptor struct {
	DbSymbol string `json:"dbSymbol"`
	Chance   int32  `json:"chance"`
}

type FormTextIdDescriptor struct {
	Name        int `json:"name"`
	Description int `json:"description"`
}

var ExperienceTypeMap = map[ExperienceTypeDescriptor]string{
	ExperienceErraticNum:     ExperienceErratic,
	ExperienceFastNum:        ExperienceFast,
	ExperienceMediumFastNum:  ExperienceMediumFast,
	ExperienceMediumSlowNum:  ExperienceMediumSlow,
	ExperienceSlowNum:        ExperienceSlow,
	ExperienceFluctuatingNum: ExperienceFluctuating,
}

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

type PokemonMapper struct {
	store *Store
}

func NewPokemonMapper(store *Store) *PokemonMapper {
	return &PokemonMapper{store: store}
}

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

func (m *PokemonMapper) MapBreedGroups(breedGroupInts []int32) []string {
	breedGroups := make([]string, len(breedGroupInts))
	for i, bgInt := range breedGroupInts {
		breedGroups[i] = BreedMap[BreedGroupDescriptor(bgInt)]
	}
	return breedGroups
}

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
