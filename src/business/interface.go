package business

import (
	"net/http"
	"sync"
)

var oAuthHandlerOnce = &sync.Once{}
var oAuthHandlerSingleton IOAuthHandler

var userHandlerOnce = &sync.Once{}
var userHandlerSingleton IUserHandler

// IOAuthHandler represents the business layer for OAuth requests.
type IOAuthHandler interface {
	// Redirect redirects the caller to the specified provider's auth page.
	Redirect(provider string, writer http.ResponseWriter) (*ResponseDTO, error)

	// HandleCallback handles the callback received from the provider.
	HandleCallback(provider string, code string, writer http.ResponseWriter) (*ResponseDTO, error)

	// init can be used to initialize the implementation.
	init()
}

// IUserHandler represents the business layer for the User entity.
type IUserHandler interface {
	// GetUser provides the profile of the user with the specified ID.
	GetUser(userID string) (*ResponseDTO, error)

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

// GetUserHandler provides the IUserHandler singleton.
func GetUserHandler() IUserHandler {
	userHandlerOnce.Do(func() {
		userHandlerSingleton = &implUserHandler{}
		userHandlerSingleton.init()
	})

	return userHandlerSingleton
}
