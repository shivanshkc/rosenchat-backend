package router

import (
	"net/http"
	"rosenchat/src/business"
	"rosenchat/src/exception"
	"rosenchat/src/utils/httputils"

	"github.com/gorilla/mux"
)

// getUserHandler handles the GetUserByID calls.
func getUserHandler(writer http.ResponseWriter, req *http.Request) {
	userID := mux.Vars(req)["user_id"]

	response, err := business.GetUserHandler().GetUser(req.Context(), userID)
	if err != nil {
		exc := exception.ToException(err)
		httputils.WriteJSON(req.Context(), writer, exc, nil, exc.StatusCode)
		return
	}

	httputils.WriteJSON(req.Context(), writer, response.Body, response.Headers, response.StatusCode)
}
