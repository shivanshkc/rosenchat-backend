package middlewares

import (
	"context"
	"net/http"
	"rosenchat/src/utils/ctxutils"
	"rosenchat/src/utils/httputils"
	"time"
)

// AccessLogger middlewares logs the request and its execution time.
func AccessLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		ctxData := ctxutils.GetRequestInfo(req.Context())
		if ctxData == nil {
			log.Warnf(context.Background(), "Request Context Data could not be type-asserted.")
			ctxData.ID = "unknown"
			ctxData.Arrival = time.Now()
		}

		log.Infof(req.Context(), "%s %s | ID: %s | Timestamp: %+v | Client-IP: %s",
			req.Method, req.URL, ctxData.ID, ctxData.Arrival, httputils.GetIPAddrFromRequest(req.Context(), req),
		)

		next.ServeHTTP(writer, req)
		log.Infof(req.Context(), "Request %s took %dms to execute.", ctxData.ID, time.Since(ctxData.Arrival).Milliseconds())
	})
}
