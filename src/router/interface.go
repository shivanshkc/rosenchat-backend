package router

import (
	"net/http"
	"sync"
)

var routerOnce = &sync.Once{}
var routerSingleton IRouter

// IRouter represents an HTTP router.
type IRouter interface {
	http.Handler

	// init can be used to initialize the implementation.
	init()
}

// Get provides the IRouter singleton.
func Get() IRouter {
	routerOnce.Do(func() {
		routerSingleton = &implGorilla{}
		routerSingleton.init()
	})

	return routerSingleton
}
