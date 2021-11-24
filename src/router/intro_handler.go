package router

import (
	"net/http"
	"rosenchat/src/utils/httputils"
)

// introHandler delivers basic information about the application.
func introHandler(writer http.ResponseWriter, _ *http.Request) {
	data := map[string]interface{}{
		"name":    conf.Application.Name,
		"version": conf.Application.Version,
	}

	responseBody := ResponseDTO{
		StatusCode: http.StatusOK,
		CustomCode: "OK",
		Data:       data,
	}

	httputils.WriteJSON(writer, responseBody, nil, http.StatusOK)
}
