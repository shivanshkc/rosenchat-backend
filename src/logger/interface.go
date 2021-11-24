package logger

import "sync"

var loggerOnce = &sync.Once{}
var loggerSingleton ILogger

// ILogger represents the application logger.
type ILogger interface {
	// Debugf logs at debug level.
	Debugf(format string, a ...interface{})
	// Infof logs at info level.
	Infof(format string, a ...interface{})
	// Warnf logs at warm level.
	Warnf(format string, a ...interface{})
	// Errorf logs at error level.
	Errorf(format string, a ...interface{})

	// init can be used to initialize the implementation.
	init()
}

// Get provides the ILogger singleton.
func Get() ILogger {
	loggerOnce.Do(func() {
		loggerSingleton = &implZap{}
		loggerSingleton.init()
	})
	return loggerSingleton
}
