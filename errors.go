package cryptowatch

type errResponse struct {
	Error     string    `json:"error"`
	Allowance allowance `json:"allowance"`
}
