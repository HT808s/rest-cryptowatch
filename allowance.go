package cryptowatch

// Allowance holds information regarding usage limit.
type allowance struct {
	Cost          uint   `json:"cost"`
	Remaining     uint64 `json:"remaining"`
	RemainingPaid uint64 `json:"remainingPaid"`
	Account       string `json:"account"`
}
