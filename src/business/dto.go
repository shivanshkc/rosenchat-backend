package business

// ResponseDTO contains all data that may be required by the
// controller layer to send back a response to the external user.
type ResponseDTO struct {
	StatusCode int
	Headers    map[string]string
	Body       *ResponseBodyDTO
}

// ResponseBodyDTO is the schema for all response bodies.
type ResponseBodyDTO struct {
	StatusCode int         `json:"status_code"`
	CustomCode string      `json:"custom_code"`
	Data       interface{} `json:"data"`
	Errors     []string    `json:"errors"`
}
