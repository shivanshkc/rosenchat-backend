package configs

import (
	"sync"

	// Automatically loads the environment from the .env file.
	_ "github.com/joho/godotenv/autoload"
)

var configOnce = &sync.Once{}
var configSingleton *IConfig

// IConfig represents the configuration data. It is not an interface
// but would have been if Go allowed properties in interfaces.
type IConfig struct {
	// Application holds the Application configs.
	Application struct {
		// Name is the name of the application.
		Name string `default:"RosenChat" env:"APPLICATION_NAME" arg:"application-name"`
		// Version is the version of the application.
		Version string `default:"1.0.0" env:"APPLICATION_VERSION" arg:"application-version"`
	}

	// GRPCServer holds the gRPC server configs.
	GRPCServer struct {
		// Addr is the address where the gRPC server will listen.
		Addr string `default:"0.0.0.0:9090" env:"GRPC_SERVER_ADDR" arg:"grpc-server-addr"`
		// Enabled is a flag that enables/disables the gRPC server.
		Enabled bool `default:"true" env:"GRPC_SERVER_ENABLED" arg:"grpc-server-enabled"`
	}

	// HTTPServer holds the HTTP server configs.
	HTTPServer struct {
		// Addr is the address where the HTTP server will listen.
		Addr string `default:"0.0.0.0:8080" env:"HTTP_SERVER_ADDR" arg:"http-server-addr"`
		// Enabled is a flag that enables/disables the HTTP server.
		Enabled bool `default:"true" env:"HTTP_SERVER_ENABLED" arg:"http-server-enabled"`
		// ShutDownTimeoutSec is the timeout in seconds for HTTP server's graceful shutdown call.
		ShutDownTimeoutSec int `default:"60" env:"HTTP_SERVER_SHUTDOWN_TIMEOUT_SEC" arg:"http-server-shutdown-timeout-sec"`
	}

	// Logger holds the logger configs.
	Logger struct {
		// Level is the logging level.
		Level string `default:"info" env:"LOGGER_LEVEL" arg:"logger-level"`
		// FilePath is the path to the log file.
		FilePath string `default:"logs/service.log" env:"LOGGER_FILE_PATH" arg:"logger-file-path"`
	}
}

// Get provides the IConfig singleton.
func Get() IConfig {
	configOnce.Do(func() {
		configSingleton = &IConfig{}
		loadWithConfetti(configSingleton)
	})

	return *configSingleton
}
