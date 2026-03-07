package studio

// Ability represents a Pokemon ability.
type Ability struct {
	dbSymbol    string
	id          int
	textId      int
	name        Translation
	description Translation
}

// AbilityOption is a functional option for configuring an Ability.
type AbilityOption func(*Ability)

// WithAbilityDbSymbol sets the database symbol of an Ability.
func WithAbilityDbSymbol(dbSymbol string) AbilityOption {
	return func(a *Ability) { a.dbSymbol = dbSymbol }
}

// WithAbilityID sets the ID of an Ability.
func WithAbilityID(id int) AbilityOption {
	return func(a *Ability) { a.id = id }
}

// WithAbilityTextID sets the text ID of an Ability.
func WithAbilityTextID(textId int) AbilityOption {
	return func(a *Ability) { a.textId = textId }
}

// WithAbilityName sets the name translations of an Ability.
func WithAbilityName(name Translation) AbilityOption {
	return func(a *Ability) { a.name = name }
}

// WithAbilityDescription sets the description translations of an Ability.
func WithAbilityDescription(desc Translation) AbilityOption {
	return func(a *Ability) { a.description = desc }
}

// NewAbility creates a new Ability with the given options.
func NewAbility(opts ...AbilityOption) *Ability {
	a := &Ability{}
	for _, opt := range opts {
		opt(a)
	}
	return a
}

// DbSymbol returns the database symbol of the Ability.
func (a Ability) DbSymbol() string {
	return a.dbSymbol
}

// ID returns the ID of the Ability.
func (a Ability) ID() int {
	return a.id
}

// TextID returns the text ID of the Ability.
func (a Ability) TextID() int {
	return a.textId
}

// Name returns the localized name of the Ability for the given language.
func (a Ability) Name(lang string) string {
	return a.name[lang]
}

// Description returns the localized description of the Ability for the given language.
func (a Ability) Description(lang string) string {
	return a.description[lang]
}

// AbilityDescriptor is the JSON descriptor for an Ability.
type AbilityDescriptor struct {
	DbSymbol    string `json:"dbSymbol"`
	Id          int    `json:"id"`
	TextID      int    `json:"textId"`
	Name        Translation
	Description Translation
}

// AbilityMapper maps Ability descriptors to Ability entities.
type AbilityMapper struct {
	store *Store
}

// NewAbilityMapper creates a new AbilityMapper.
func NewAbilityMapper(store *Store) *AbilityMapper {
	return &AbilityMapper{store: store}
}

// MapAbilityDescriptorToAbility maps an AbilityDescriptor to an Ability.
func (m *AbilityMapper) MapAbilityDescriptorToAbility(desc AbilityDescriptor) *Ability {
	ability := NewAbility(
		WithAbilityDbSymbol(desc.DbSymbol),
		WithAbilityID(desc.Id),
		WithAbilityTextID(desc.TextID),
		WithAbilityName(desc.Name),
		WithAbilityDescription(desc.Description),
	)

	return ability
}
