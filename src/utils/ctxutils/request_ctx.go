package ctxutils

import (
	"context"
	"rosenchat/src/logger"
	"time"
)

var log = logger.Get()

// RequestContextKey is the key to which the context data will be mapped in the request.
const RequestContextKey RequestContextKeyType = iota

// RequestContextKeyType is the data-type of the key to which the context data will be mapped in the request.
type RequestContextKeyType int

// RequestContextData is the context data of a request.
type RequestContextData struct {
	ID      string
	Arrival time.Time
}

// GetRequestInfoForLog extracts and returns the request info from the provided context for logging purposes.
func GetRequestInfoForLog(ctx context.Context) interface{} {
	contextDataInterface := ctx.Value(RequestContextKey)
	if contextDataInterface == nil {
		log.Warnf("Provided context does not have any request data.")
		return nil
	}

	contextData, ok := contextDataInterface.(*RequestContextData)
	if !ok {
		log.Warnf("Provided context contains request data in unknown format.")
		return nil
	}

	return map[string]interface{}{
		"request_id": contextData.ID,
	}
}
