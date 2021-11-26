package logger

import (
	"context"
	"sync"
)

var loggerOnce = &sync.Once{}
var loggerSingleton ILogger

// ILogger represents the application logger.
type ILogger interface {
	// Debugf logs at debug level.
	Debugf(ctx context.Context, format string, a ...interface{})
	// Infof logs at info level.
	Infof(ctx context.Context, format string, a ...interface{})
	// Warnf logs at warm level.
	Warnf(ctx context.Context, format string, a ...interface{})
	// Errorf logs at error level.
	Errorf(ctx context.Context, format string, a ...interface{})

	// Close closes the logger (required when logger logs over the network).
	Close(ctx context.Context)

	// init can be used to initialize the implementation.
	init()
}

// Get provides the ILogger singleton.
func Get() ILogger {
	loggerOnce.Do(func() {
		loggerSingleton = &implGCP{}
		loggerSingleton.init()
	})
	return loggerSingleton
}
