package move

// AimedTarget represents the target of a move
type AimedTarget string

const (
	AimedTargetAdjacentPokemon    AimedTarget = "adjacent_pokemon"
	AimedTargetAdjacentFoe        AimedTarget = "adjacent_foe"
	AimedTargetAdjacentAllFoe     AimedTarget = "adjacent_all_foe"
	AimedTargetAllFoe             AimedTarget = "all_foe"
	AimedTargetAdjacentAllPokemon AimedTarget = "adjacent_all_pokemon"
	AimedTargetAllPokemon         AimedTarget = "all_pokemon"
	AimedTargetUser               AimedTarget = "user"
	AimedTargetUserOrAdjacentAlly AimedTarget = "user_or_adjacent_ally"
	AimedTargetAdjacentAlly       AimedTarget = "adjacent_ally"
	AimedTargetAllAlly            AimedTarget = "all_ally"
	AimedTargetAllAllyButUser     AimedTarget = "all_ally_but_user"
	AimedTargetAnyOtherPokemon    AimedTarget = "any_other_pokemon"
	AimedTargetRandomFoe          AimedTarget = "random_foe"
)

// ContactType represents how a move makes contact
type ContactType string

const (
	ContactTypeDirect  ContactType = "DIRECT"
	ContactTypeDistant ContactType = "DISTANT"
	ContactTypeNone    ContactType = "NONE"
)

// MoveTargeting contains targeting information for a move
type MoveTargeting struct {
	AimedTarget AimedTarget
	ContactType ContactType
}
