package responses

type ErrorResponse struct {
	StatusCode int    `json:"status_code"`
	StatusType string `json:"status_type"`
	Message    string `json:"message"`
}
