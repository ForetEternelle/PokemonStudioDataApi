package move

// MechanicalTag represents a mechanical classification of a move
type MechanicalTag string

const (
	MechanicalTagAuthentic MechanicalTag = "AUTHENTIC"
	MechanicalTagBallistic MechanicalTag = "BALLISTIC"
	MechanicalTagBite      MechanicalTag = "BITE"
	MechanicalTagDance     MechanicalTag = "DANCE"
	MechanicalTagPunch     MechanicalTag = "PUNCH"
	MechanicalTagSlice     MechanicalTag = "SLICE"
	MechanicalTagSound     MechanicalTag = "SOUND"
	MechanicalTagWind      MechanicalTag = "WIND"
	MechanicalTagPulse     MechanicalTag = "PULSE"
	MechanicalTagPowder    MechanicalTag = "POWDER"
	MechanicalTagMental    MechanicalTag = "MENTAL"
)
