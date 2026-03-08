package studio

// PokemonBuilder builds Pokemon entities.
type PokemonBuilder struct {
	pokemon *Pokemon
}

// NewPokemonBuilder creates a new PokemonBuilder.
func NewPokemonBuilder() *PokemonBuilder {
	return &PokemonBuilder{pokemon: &Pokemon{
		customProperties: make(map[string]any),
	}}
}

// ID sets the ID of the Pokemon.
func (b *PokemonBuilder) ID(id int32) *PokemonBuilder {
	b.pokemon.id = id
	return b
}

// DbSymbol sets the database symbol of the Pokemon.
func (b *PokemonBuilder) DbSymbol(dbSymbol string) *PokemonBuilder {
	b.pokemon.dbSymbol = dbSymbol
	return b
}

// Name sets the name translations of the Pokemon.
func (b *PokemonBuilder) Name(name Translation) *PokemonBuilder {
	b.pokemon.name = name
	return b
}

// Description sets the description translations of the Pokemon.
func (b *PokemonBuilder) Description(desc Translation) *PokemonBuilder {
	b.pokemon.description = desc
	return b
}

// Forms sets the forms of the Pokemon.
func (b *PokemonBuilder) Forms(forms map[int32]PokemonForm) *PokemonBuilder {
	b.pokemon.forms = forms
	return b
}

// CustomProperties sets custom properties for the Pokemon.
func (b *PokemonBuilder) CustomProperties(props map[string]any) *PokemonBuilder {
	b.pokemon.customProperties = props
	return b
}

// Build returns a copy of the built Pokemon.
func (b *PokemonBuilder) Build() *Pokemon {
	_copy := *b.pokemon
	return &_copy
}

// PokemonFormBuilder builds PokemonForm entities.
type PokemonFormBuilder struct {
	form *PokemonForm
}

// NewPokemonFormBuilder creates a new PokemonFormBuilder.
func NewPokemonFormBuilder() *PokemonFormBuilder {
	return &PokemonFormBuilder{form: &PokemonForm{
		customProperties: make(map[string]any),
	}}
}

// Form sets the form number of the PokemonForm.
func (b *PokemonFormBuilder) Form(form int32) *PokemonFormBuilder {
	b.form.form = form
	return b
}

// Type1 sets the primary type of the PokemonForm.
func (b *PokemonFormBuilder) Type1(t *PokemonType) *PokemonFormBuilder {
	b.form.type1 = t
	return b
}

// Type2 sets the secondary type of the PokemonForm.
func (b *PokemonFormBuilder) Type2(t *PokemonType) *PokemonFormBuilder {
	b.form.type2 = t
	return b
}

// Height sets the height of the PokemonForm.
func (b *PokemonFormBuilder) Height(h float32) *PokemonFormBuilder {
	b.form.height = h
	return b
}

// Weight sets the weight of the PokemonForm.
func (b *PokemonFormBuilder) Weight(w float32) *PokemonFormBuilder {
	b.form.weight = w
	return b
}

// BaseHp sets the base HP of the PokemonForm.
func (b *PokemonFormBuilder) BaseHp(hp int32) *PokemonFormBuilder {
	b.form.baseHp = hp
	return b
}

// BaseAtk sets the base Attack of the PokemonForm.
func (b *PokemonFormBuilder) BaseAtk(atk int32) *PokemonFormBuilder {
	b.form.baseAtk = atk
	return b
}

// BaseDfe sets the base Defense of the PokemonForm.
func (b *PokemonFormBuilder) BaseDfe(dfe int32) *PokemonFormBuilder {
	b.form.baseDfe = dfe
	return b
}

// BaseSpd sets the base Speed of the PokemonForm.
func (b *PokemonFormBuilder) BaseSpd(spd int32) *PokemonFormBuilder {
	b.form.baseSpd = spd
	return b
}

// BaseAts sets the base Special Attack of the PokemonForm.
func (b *PokemonFormBuilder) BaseAts(ats int32) *PokemonFormBuilder {
	b.form.baseAts = ats
	return b
}

// BaseDfs sets the base Special Defense of the PokemonForm.
func (b *PokemonFormBuilder) BaseDfs(dfs int32) *PokemonFormBuilder {
	b.form.baseDfs = dfs
	return b
}

// EvHp sets the EV yield for HP of the PokemonForm.
func (b *PokemonFormBuilder) EvHp(hp int32) *PokemonFormBuilder {
	b.form.evHp = hp
	return b
}

// EvAtk sets the EV yield for Attack of the PokemonForm.
func (b *PokemonFormBuilder) EvAtk(atk int32) *PokemonFormBuilder {
	b.form.evAtk = atk
	return b
}

// EvDfe sets the EV yield for Defense of the PokemonForm.
func (b *PokemonFormBuilder) EvDfe(dfe int32) *PokemonFormBuilder {
	b.form.evDfe = dfe
	return b
}

// EvSpd sets the EV yield for Speed of the PokemonForm.
func (b *PokemonFormBuilder) EvSpd(spd int32) *PokemonFormBuilder {
	b.form.evSpd = spd
	return b
}

// EvAts sets the EV yield for Special Attack of the PokemonForm.
func (b *PokemonFormBuilder) EvAts(ats int32) *PokemonFormBuilder {
	b.form.evAts = ats
	return b
}

// EvDfs sets the EV yield for Special Defense of the PokemonForm.
func (b *PokemonFormBuilder) EvDfs(dfs int32) *PokemonFormBuilder {
	b.form.evDfs = dfs
	return b
}

// Evolutions sets the evolutions of the PokemonForm.
func (b *PokemonFormBuilder) Evolutions(evolutions []Evolution) *PokemonFormBuilder {
	b.form.evolutions = evolutions
	return b
}

// ExperienceType sets the experience type of the PokemonForm.
func (b *PokemonFormBuilder) ExperienceType(expType string) *PokemonFormBuilder {
	b.form.experienceType = expType
	return b
}

// BaseExperience sets the base experience of the PokemonForm.
func (b *PokemonFormBuilder) BaseExperience(exp int32) *PokemonFormBuilder {
	b.form.baseExperience = exp
	return b
}

// BaseLoyalty sets the base loyalty of the PokemonForm.
func (b *PokemonFormBuilder) BaseLoyalty(loyalty int32) *PokemonFormBuilder {
	b.form.baseLoyalty = loyalty
	return b
}

// CatchRate sets the catch rate of the PokemonForm.
func (b *PokemonFormBuilder) CatchRate(rate int32) *PokemonFormBuilder {
	b.form.catchRate = rate
	return b
}

// FemaleRate sets the female rate of the PokemonForm.
func (b *PokemonFormBuilder) FemaleRate(rate float32) *PokemonFormBuilder {
	b.form.femaleRate = rate
	return b
}

// BreedGroups sets the breed groups of the PokemonForm.
func (b *PokemonFormBuilder) BreedGroups(groups []string) *PokemonFormBuilder {
	b.form.breedGroups = groups
	return b
}

// HatchSteps sets the hatch steps of the PokemonForm.
func (b *PokemonFormBuilder) HatchSteps(steps int32) *PokemonFormBuilder {
	b.form.hatchSteps = steps
	return b
}

// BabyDbSymbol sets the baby Pokemon database symbol of the PokemonForm.
func (b *PokemonFormBuilder) BabyDbSymbol(symbol *string) *PokemonFormBuilder {
	b.form.babyDbSymbol = symbol
	return b
}

// BabyForm sets the baby form number of the PokemonForm.
func (b *PokemonFormBuilder) BabyForm(form int32) *PokemonFormBuilder {
	b.form.babyForm = form
	return b
}

// ItemHeld sets the items held by the PokemonForm.
func (b *PokemonFormBuilder) ItemHeld(items []*ItemHeld) *PokemonFormBuilder {
	b.form.itemHeld = items
	return b
}

// AbilitySymbols sets the ability symbols of the PokemonForm.
func (b *PokemonFormBuilder) AbilitySymbols(symbols []string) *PokemonFormBuilder {
	b.form.abilitySymbols = symbols
	return b
}

// Abilities sets the abilities of the PokemonForm.
func (b *PokemonFormBuilder) Abilities(abilities []*Ability) *PokemonFormBuilder {
	b.form.abilities = abilities
	return b
}

// FrontOffsetY sets the front offset Y of the PokemonForm.
func (b *PokemonFormBuilder) FrontOffsetY(y int32) *PokemonFormBuilder {
	b.form.frontOffsetY = y
	return b
}

// CustomProperties sets custom properties for the PokemonForm.
func (b *PokemonFormBuilder) CustomProperties(props map[string]any) *PokemonFormBuilder {
	b.form.customProperties = props
	return b
}

// Build returns a copy of the built PokemonForm.
func (b *PokemonFormBuilder) Build() *PokemonForm {
	_copy := *b.form
	return &_copy
}
