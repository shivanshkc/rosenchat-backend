package exception

import (
	"net/http"
)

// ProviderNotFound is for request that supply invalid OAuth provider.
var ProviderNotFound = func() *Exception {
	return &Exception{StatusCode: http.StatusNotFound, CustomCode: "PROVIDER_NOT_FOUND"}
}

// UserNotFound is for request that requires a non-existent user.
var UserNotFound = func() *Exception {
	return &Exception{StatusCode: http.StatusNotFound, CustomCode: "USER_NOT_FOUND"}
}
