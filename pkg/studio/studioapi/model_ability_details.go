package studioapi

type AbilityDetails struct {
	Symbol      string `json:"symbol"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// AssertAbilityDetailsRequired checks if the required fields are not zero-ed
func AssertAbilityDetailsRequired(obj AbilityDetails) error {
	return nil
}

// AssertAbilityDetailsConstraints checks if the values respects the defined constraints
func AssertAbilityDetailsConstraints(obj AbilityDetails) error {
	return nil
}
