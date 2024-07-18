package myinvois

// StandErrorResponse inferred from sandbox response
// currently not aligned with the spec at https://sdk.myinvois.hasil.gov.my/standard-error-response/
type StandardErrorResponse struct {
	Error Error `json:"error"`
}

type Error struct {
	Code    string  `json:"code"`
	Message string  `json:"message"`
	Target  string  `json:"target"`
	Details []Error `json:"details,omitempty"`
}
