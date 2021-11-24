package adapters

import (
	"context"
)

// IAdapter represents an adapter.
type IAdapter interface {
	// Name provides the name of the adapter.
	Name() string

	// Start method starts the adapter.
	Start(ctx context.Context) error
	// Stop method stops the adapter.
	Stop(ctx context.Context) error

	// init can be used to initialize the implementation.
	init()
}
