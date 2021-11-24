package business

import (
	"net/http"
	"sync"
)

var oAuthHandlerOnce = &sync.Once{}
var oAuthHandlerSingleton IOAuthHandler

// IOAuthHandler represents the business layer for OAuth requests.
type IOAuthHandler interface {
	// Redirect redirects the caller to the specified provider's auth page.
	Redirect(provider string, writer http.ResponseWriter) (*ResponseDTO, error)

	// HandleCallback handles the callback received from the provider.
	HandleCallback(provider string, code string, writer http.ResponseWriter) (*ResponseDTO, error)

	// init can be used to initialize the implementation.
	init()
}

// GetOAuthHandler provides the IOAuthHandler singleton.
func GetOAuthHandler() IOAuthHandler {
	oAuthHandlerOnce.Do(func() {
		oAuthHandlerSingleton = &implOAuthHandler{}
		oAuthHandlerSingleton.init()
	})

	return oAuthHandlerSingleton
}
