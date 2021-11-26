package ctxutils

import (
	"context"
	"time"
)

// RequestContextKey is the key to which the context data will be mapped in the request.
const RequestContextKey RequestContextKeyType = iota

// RequestContextKeyType is the data-type of the key to which the context data will be mapped in the request.
type RequestContextKeyType int

// RequestContextData is the context data of a request.
type RequestContextData struct {
	ID      string
	Arrival time.Time
}

// GetRequestInfo provides the RequestContextData from the provided context.
// If there is no data found, nil is returned.
func GetRequestInfo(ctx context.Context) *RequestContextData {
	contextDataInterface := ctx.Value(RequestContextKey)
	if contextDataInterface == nil {
		return nil
	}

	contextData, ok := contextDataInterface.(*RequestContextData)
	if !ok {
		return nil
	}

	return contextData
}
