package middlewares

import (
	"net/http"
	"rosenchat/src/exception"
	"rosenchat/src/logger"
	"rosenchat/src/utils/httputils"
)

var log = logger.Get()

// Recovery middlewares recovers any panics that happened during
// the request execution and returns a sanitized response.
func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		defer func() {
			err := recover()
			if err == nil {
				return
			}

			log.Warnf(req.Context(), "Panic occurred during request execution: %+v", err)
			exc := exception.ToException(err)
			httputils.WriteJSON(req.Context(), writer, exc, nil, exc.StatusCode)
		}()

		next.ServeHTTP(writer, req)
	})
}
