package studio

// MoveDescriptor is the JSON descriptor for a Move.
type MoveDescriptor struct {
	Klass                   string                     `json:"klass"`
	Id                      int                        `json:"id"`
	DbSymbol                string                     `json:"dbSymbol"`
	MapUse                  int                        `json:"mapUse"`
	BattleEngineMethod      string                     `json:"battleEngineMethod"`
	Type                    string                     `json:"type"`
	Power                   int                        `json:"power"`
	Accuracy                int                        `json:"accuracy"`
	PP                      int                        `json:"pp"`
	Category                string                     `json:"category"`
	MoveCriticalRate        int                        `json:"movecriticalRate"`
	Priority                int                        `json:"priority"`
	IsAuthentic             bool                       `json:"isAuthentic"`
	IsBallistics            bool                       `json:"isBallistics"`
	IsBite                  bool                       `json:"isBite"`
	IsBlocable              bool                       `json:"isBlocable"`
	IsCharge                bool                       `json:"isCharge"`
	IsDance                 bool                       `json:"isDance"`
	IsDirect                bool                       `json:"isDirect"`
	IsDistance              bool                       `json:"isDistance"`
	IsEffectChance          bool                       `json:"isEffectChance"`
	IsGravity               bool                       `json:"isGravity"`
	IsHeal                  bool                       `json:"isHeal"`
	IsKingRockUtility       bool                       `json:"isKingRockUtility"`
	IsMagicCoatAffected     bool                       `json:"isMagicCoatAffected"`
	IsMental                bool                       `json:"isMental"`
	IsMirrorMove            bool                       `json:"isMirrorMove"`
	IsNonSkyBattle          bool                       `json:"isNonSkyBattle"`
	IsPowder                bool                       `json:"isPowder"`
	IsPulse                 bool                       `json:"isPulse"`
	IsPunch                 bool                       `json:"isPunch"`
	IsRecharge              bool                       `json:"isRecharge"`
	IsSnatchable            bool                       `json:"isSnatchable"`
	IsSoundAttack           bool                       `json:"isSoundAttack"`
	IsUnfreeze              bool                       `json:"isUnfreeze"`
	BattleEngineAimedTarget string                     `json:"battleEngineAimedTarget"`
	BattleStageMod          []BattleStageModDescriptor `json:"battleStageMod"`
	MoveStatus              []MoveStatusDescriptor     `json:"moveStatus"`
	EffectChance            int                        `json:"effectChance"`
	IsSlicingAttack         bool                       `json:"isSlicingAttack"`
	IsWind                  bool                       `json:"isWind"`
	Name                    Translation
	Description             Translation
}

type BattleStageModDescriptor struct {
	BattleStage string `json:"battleStage"`
	Modificator int    `json:"modificator"`
}

type MoveStatusDescriptor struct {
	Status   string `json:"status"`
	LuckRate int    `json:"luckRate"`
}


