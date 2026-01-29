package move

// BattleStage represents a stat stage in battle
type BattleStage string

const (
	BattleStageATK BattleStage = "ATK_STAGE"
	BattleStageDFE BattleStage = "DFE_STAGE"
	BattleStageATS BattleStage = "ATS_STAGE"
	BattleStageDFS BattleStage = "DFS_STAGE"
	BattleStageSPD BattleStage = "SPD_STAGE"
	BattleStageEVA BattleStage = "EVA_STAGE"
	BattleStageACC BattleStage = "ACC_STAGE"
)

// MoveStatStageChange represents a stat stage modification
type MoveStatStageChange struct {
	BattleStage BattleStage
	Modificator int
}

// MoveStatusEffect represents a status effect applied by a move
type MoveStatusEffect struct {
	Status   string
	LuckRate int
}

// MoveSecondaryEffects contains secondary effects of a move
type MoveSecondaryEffects struct {
	Chance           int
	StatusEffects    []MoveStatusEffect
	StatStageChanges []MoveStatStageChange
}
