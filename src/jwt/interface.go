package jwt

import (
	"sync"
)

var jwtManagerOnce = &sync.Once{}
var jwtManagerSingleton IJWTManager

// IJWTManager represents a generic service to manage JWTs.
type IJWTManager interface {
	// DecodeUnsafe decodes the JWT without validating it and without verifying the sign.
	DecodeUnsafe(token string) (map[string]interface{}, error)

	// init can be used to initialize the implementation.
	init()
}

// Get provides the IJWTManager singleton.
func Get() IJWTManager {
	jwtManagerOnce.Do(func() {
		jwtManagerSingleton = &implGolangJWT{}
		jwtManagerSingleton.init()
	})

	return jwtManagerSingleton
}
