package studio

type Ability struct {
	dbSymbol    string
	id          int
	textId      int
	name        Translation
	description Translation
}

type AbilityOption func(*Ability)

func WithAbilityDbSymbol(dbSymbol string) AbilityOption {
	return func(a *Ability) { a.dbSymbol = dbSymbol }
}

func WithAbilityID(id int) AbilityOption {
	return func(a *Ability) { a.id = id }
}

func WithAbilityTextID(textId int) AbilityOption {
	return func(a *Ability) { a.textId = textId }
}

func WithAbilityName(name Translation) AbilityOption {
	return func(a *Ability) { a.name = name }
}

func WithAbilityDescription(desc Translation) AbilityOption {
	return func(a *Ability) { a.description = desc }
}

func NewAbility(opts ...AbilityOption) *Ability {
	a := &Ability{}
	for _, opt := range opts {
		opt(a)
	}
	return a
}

func (a Ability) DbSymbol() string {
	return a.dbSymbol
}

func (a Ability) ID() int {
	return a.id
}

func (a Ability) TextID() int {
	return a.textId
}

func (a Ability) Name(lang string) string {
	return a.name[lang]
}

func (a Ability) Description(lang string) string {
	return a.description[lang]
}

type AbilityDescriptor struct {
	DbSymbol    string `json:"dbSymbol"`
	Id          int    `json:"id"`
	TextID      int    `json:"textId"`
	Name        Translation
	Description Translation
}

type AbilityMapper struct {
	store *Store
}

func NewAbilityMapper(store *Store) *AbilityMapper {
	return &AbilityMapper{store: store}
}

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
