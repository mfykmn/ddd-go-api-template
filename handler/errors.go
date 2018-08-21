package handler

type responseError struct {
	RequestID string      `json:"request_id"`
	Reason    string      `json:"reason"`
	Error     interface{} `json:"error"`
}
