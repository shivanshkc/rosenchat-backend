package exception

import (
	"net/http"
)

// BadRequest is for logically incorrect requests.
var BadRequest = func() *Exception {
	return &Exception{StatusCode: http.StatusBadRequest, CustomCode: "BAD_REQUEST"}
}

// Unauthorized is for requests with invalid credentials.
var Unauthorized = func() *Exception {
	return &Exception{StatusCode: http.StatusUnauthorized, CustomCode: "UNAUTHORIZED"}
}

// PaymentRequired is for requests that require payment completion.
var PaymentRequired = func() *Exception {
	return &Exception{StatusCode: http.StatusPaymentRequired, CustomCode: "PAYMENT_REQUIRED"}
}

// Forbidden is for requests that do not have enough authority to execute the operation.
var Forbidden = func() *Exception {
	return &Exception{StatusCode: http.StatusForbidden, CustomCode: "FORBIDDEN"}
}

// NotFound is for requests that try to access a non-existent resource.
var NotFound = func() *Exception {
	return &Exception{StatusCode: http.StatusNotFound, CustomCode: "NOT_FOUND"}
}

// RequestTimeout is for requests that take longer than a certain time limit to execute.
var RequestTimeout = func() *Exception {
	return &Exception{StatusCode: http.StatusRequestTimeout, CustomCode: "REQUEST_TIMEOUT"}
}

// Conflict is for requests that attempt paradoxical operations, such as re-creating the same resource.
var Conflict = func() *Exception {
	return &Exception{StatusCode: http.StatusConflict, CustomCode: "CONFLICT"}
}

// PreconditionFailed is for requests that do not satisfy pre-business layers of the application.
var PreconditionFailed = func() *Exception {
	return &Exception{StatusCode: http.StatusPreconditionFailed, CustomCode: "PRECONDITION_FAILED"}
}

// InternalServerError is for requests that cause an unexpected misbehaviour.
var InternalServerError = func() *Exception {
	return &Exception{StatusCode: http.StatusInternalServerError, CustomCode: "INTERNAL_SERVER_ERROR"}
}
