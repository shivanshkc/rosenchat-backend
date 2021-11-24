package router

// ResponseDTO is the schema for any response sent to an external client.
type ResponseDTO struct {
	StatusCode int         `json:"status_code"`
	CustomCode string      `json:"custom_code"`
	Data       interface{} `json:"data"`
	Errors     []string    `json:"errors"`
}
