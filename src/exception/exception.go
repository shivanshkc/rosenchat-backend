package exception

// Exception is a custom error type that implements the error interface.
type Exception struct {
	StatusCode int      `json:"status_code"`
	CustomCode string   `json:"custom_code"`
	Errors     []string `json:"errors"`
}

func (e *Exception) Error() string {
	return e.CustomCode
}

// AddMessages is a chainable method to add string messages to the Exception.
func (e *Exception) AddMessages(message ...string) *Exception {
	e.Errors = append(e.Errors, message...)
	return e
}

// AddErrors is a chainable method to add error messages to the Exception.
func (e *Exception) AddErrors(errors ...error) *Exception {
	for _, err := range errors {
		e.Errors = append(e.Errors, err.Error())
	}
	return e
}

// ToException converts any value to an appropriate Exception.
func ToException(err interface{}) *Exception {
	switch asserted := err.(type) {
	case *Exception:
		return asserted
	case error:
		return InternalServerError().AddErrors(err.(error))
	case string:
		return InternalServerError().AddMessages(err.(string))
	default:
		return InternalServerError()
	}
}
