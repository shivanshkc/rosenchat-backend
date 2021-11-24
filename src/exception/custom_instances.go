package exception

import (
	"net/http"
)

// ProviderNotFound is for request that supply invalid OAuth provider.
var ProviderNotFound = func() *Exception {
	return &Exception{StatusCode: http.StatusNotFound, CustomCode: "PROVIDER_NOT_FOUND"}
}
