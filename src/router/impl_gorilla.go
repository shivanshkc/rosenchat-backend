package router

import (
	"net/http"
	"rosenchat/src/business"
	"rosenchat/src/configs"
	"rosenchat/src/middlewares"

	"github.com/gorilla/mux"
)

var conf = configs.Get()

// implGorilla implements IRouter using gorilla/mux package.
type implGorilla struct {
	router *mux.Router

	oAuthHandler business.IOAuthHandler
}

func (i *implGorilla) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	i.router.ServeHTTP(writer, req)
}

func (i *implGorilla) init() {
	// Router dependencies.
	i.oAuthHandler = business.GetOAuthHandler()

	// The REST API router.
	i.router = mux.NewRouter()

	// Adding middlewares.
	i.router.Use(middlewares.Recovery)
	i.router.Use(middlewares.RequestContext)
	i.router.Use(middlewares.AccessLogger)
	i.router.Use(middlewares.CORS)

	i.router.HandleFunc("/api", introHandler).Methods(http.MethodOptions, http.MethodGet)

	// OAuth routes.
	i.router.HandleFunc("/api/auth/{provider}", oAuthRedirectHandler).Methods(http.MethodOptions, http.MethodGet)
	i.router.HandleFunc("/api/auth/{provider}/callback", oAuthCallbackHandler).Methods(http.MethodOptions, http.MethodGet)

	i.router.HandleFunc("/api/user/{user_id}", getUserHandler).Methods(http.MethodOptions, http.MethodGet)
}
