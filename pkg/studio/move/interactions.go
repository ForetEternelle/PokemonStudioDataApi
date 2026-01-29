package move

// MoveInteraction represents an interaction property of a move
type MoveInteraction string

const (
	InteractionBlocable          MoveInteraction = "BLOCABLE"
	InteractionMirrorMove        MoveInteraction = "MIRROR_MOVE"
	InteractionSnatchable        MoveInteraction = "SNATCHABLE"
	InteractionMagicCoatAffected MoveInteraction = "MAGIC_COAT_AFFECTED"
	InteractionKingRockUtility   MoveInteraction = "KING_ROCK_UTILITY"
	InteractionAffectedByGravity MoveInteraction = "AFFECTED_BY_GRAVITY"
	InteractionNonSkyBattle      MoveInteraction = "NON_SKY_BATTLE"
)
