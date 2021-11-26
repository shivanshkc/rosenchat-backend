package router

import (
	"net/http"
	"rosenchat/src/business"
	"rosenchat/src/utils/httputils"
)

// introHandler delivers basic information about the application.
func introHandler(writer http.ResponseWriter, req *http.Request) {
	responseBody := business.ResponseBodyDTO{
		StatusCode: http.StatusOK,
		CustomCode: "OK",
		Data: map[string]interface{}{
			"name":    conf.Application.Name,
			"version": conf.Application.Version,
		},
	}

	httputils.WriteJSON(req.Context(), writer, responseBody, nil, http.StatusOK)
}
