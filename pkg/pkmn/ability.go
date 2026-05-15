package studio

// Ability represents a Pokemon ability.
type Ability struct {
	dbSymbol    string
	id          int
	textId      int
	name        Translation
	description Translation
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
