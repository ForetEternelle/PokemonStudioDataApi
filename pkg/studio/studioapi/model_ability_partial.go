package studioapi

type AbilityPartial struct {
	// The translated name of the ability
	Name string `json:"name,omitempty"`

	// The symbol of the ability
	Symbol string `json:"symbol,omitempty"`

	// The translated description of the ability
	Description string `json:"description,omitempty"`
}

// AssertAbilityPartialRequired checks if the required fields are not zero-ed
func AssertAbilityPartialRequired(obj AbilityPartial) error {
	return nil
}

// AssertAbilityPartialConstraints checks if the values respects the defined constraints
func AssertAbilityPartialConstraints(obj AbilityPartial) error {
	return nil
}
