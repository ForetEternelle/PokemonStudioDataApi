package move

// ExecutionMethod represents how a move is executed
type ExecutionMethod string

const (
	ExecutionMethodBasic      ExecutionMethod = "s_basic"
	ExecutionMethodStat       ExecutionMethod = "s_stat"
	ExecutionMethodStatus     ExecutionMethod = "s_status"
	ExecutionMethodMultiHit   ExecutionMethod = "s_multi_hit"
	ExecutionMethod2Hits      ExecutionMethod = "s_2hits"
	ExecutionMethodOHKO       ExecutionMethod = "s_ohko"
	ExecutionMethod2Turns     ExecutionMethod = "s_2turns"
	ExecutionMethodSelfStat   ExecutionMethod = "s_self_stat"
	ExecutionMethodSelfStatus ExecutionMethod = "s_self_status"
)

// MoveExecution contains execution information for a move
type MoveExecution struct {
	Method   ExecutionMethod
	Charge   bool
	Recharge bool
}
