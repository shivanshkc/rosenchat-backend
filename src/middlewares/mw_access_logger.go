package middlewares

import (
	"net/http"
	"rosenchat/src/utils/ctxutils"
	"rosenchat/src/utils/httputils"
	"time"
)

// AccessLogger middlewares logs the request and its execution time.
func AccessLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		ctxData, ok := req.Context().Value(ctxutils.RequestContextKey).(*ctxutils.RequestContextData)
		if !ok {
			log.Warnf("Request Context Data could not be type-asserted.")
			ctxData.ID = "unknown"
			ctxData.Arrival = time.Now()
		}

		log.Infof(
			"%s %s | ID: %s | Timestamp: %+v | Client-IP: %s",
			req.Method, req.URL, ctxData.ID, ctxData.Arrival, httputils.GetIPAddrFromRequest(req),
		)

		next.ServeHTTP(writer, req)
		log.Infof("Request %s took %dms to execute.", ctxData.ID, time.Since(ctxData.Arrival).Milliseconds())
	})
}
