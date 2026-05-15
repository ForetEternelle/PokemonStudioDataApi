package studio

// AbilityDescriptor is the JSON descriptor for an Ability.
type AbilityDescriptor struct {
	DbSymbol    string `json:"dbSymbol"`
	Id          int    `json:"id"`
	TextID      int    `json:"textId"`
	Name        Translation
	Description Translation
}

