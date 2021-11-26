package adapters

import (
	"context"
	"rosenchat/src/configs"
	"rosenchat/src/logger"
	"sync"
)

var conf = configs.Get()
var log = logger.Get()

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

var grpcOnce = &sync.Once{}
var grpcSingleton IAdapter

// GetGRPC provides the gRPC adapter singleton.
func GetGRPC() IAdapter {
	grpcOnce.Do(func() {
		grpcSingleton = &implGRPC{}
		grpcSingleton.init()
	})

	return grpcSingleton
}

var httpOnce = &sync.Once{}
var httpSingleton IAdapter

// GetHTTP provides the HTTP adapter singleton.
func GetHTTP() IAdapter {
	httpOnce.Do(func() {
		httpSingleton = &implHTTP{}
		httpSingleton.init()
	})

	return httpSingleton
}

var cleanupOnce = &sync.Once{}
var cleanupSingleton IAdapter

// GetCleanup provides the Cleanup adapter singleton.
func GetCleanup() IAdapter {
	cleanupOnce.Do(func() {
		cleanupSingleton = &implCleanup{}
		cleanupSingleton.init()
	})

	return cleanupSingleton
}
