package middlewares

import (
	"context"
	"net/http"
	"rosenchat/src/utils/ctxutils"
	"time"

	"github.com/google/uuid"
)

// RequestContext creates a new request context and attaches it to the request.
func RequestContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		ctxData := &ctxutils.RequestContextData{}
		ctxData.Arrival = time.Now()

		// Resolving the request ID.
		ctxData.ID = req.Header.Get("x-request-id")
		if ctxData.ID == "" {
			ctxData.ID = uuid.NewString()
		}

		newReqCtx := context.WithValue(req.Context(), ctxutils.RequestContextKey, ctxData)
		req = req.WithContext(newReqCtx)

		// Putting the same request ID in the response headers as well.
		writer.Header().Set("x-request-id", ctxData.ID)
		next.ServeHTTP(writer, req)
	})
}
