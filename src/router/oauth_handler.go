package router

import (
	"net/http"
	"rosenchat/src/business"
	"rosenchat/src/exception"
	"rosenchat/src/utils/httputils"

	"github.com/gorilla/mux"
)

// oAuthRedirectHandler redirects the caller to the provider's auth page.
func oAuthRedirectHandler(writer http.ResponseWriter, req *http.Request) {
	provider := mux.Vars(req)["provider"]

	if _, err := business.GetOAuthHandler().Redirect(req.Context(), provider, writer); err != nil {
		exc := exception.ToException(err)
		httputils.WriteJSON(req.Context(), writer, exc, nil, exc.StatusCode)
		return
	}
}

// oAuthCallbackHandler handles the callback from the provider.
func oAuthCallbackHandler(writer http.ResponseWriter, req *http.Request) {
	provider := mux.Vars(req)["provider"]
	code := req.URL.Query().Get("code")

	if _, err := business.GetOAuthHandler().HandleCallback(req.Context(), provider, code, writer); err != nil {
		exc := exception.ToException(err)
		httputils.WriteJSON(req.Context(), writer, exc, nil, exc.StatusCode)
		return
	}
}
