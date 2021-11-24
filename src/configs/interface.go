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

	// GeneralOAuth holds the general OAuth configs.
	GeneralOAuth struct {
		// ServerCallbackBaseURL is the base URL of the backend server where it shall receive OAuth callbacks.
		ServerCallbackBaseURL string `default:"http://localhost:8080" env:"GENERAL_OAUTH_SERVER_CALLBACK_BASE_URL" arg:"general-oauth-server-callback-base-url"`
		// ClientCallbackURL is where the frontend will receive the OAuth result.
		ClientCallbackURL string `default:"https://rosenchat.com/auth/callback" env:"GENERAL_OAUTH_CLIENT_CALLBACK_URL" arg:"general-oauth-client-callback-url"`
	}

	// GoogleOAuth holds the Google OAuth configs.
	GoogleOAuth struct {
		// RedirectURL is the authentication URL where the users are redirected.
		RedirectURL string `default:"https://accounts.google.com/o/oauth2/v2/auth" env:"GOOGLE_OAUTH_REDIRECT_URL" arg:"google-oauth-redirect-url"`
		// Scopes are the OAuth scopes.
		Scopes string `default:"https://www.googleapis.com/auth/userinfo.email+https://www.googleapis.com/auth/userinfo.profile" env:"GOOGLE_OAUTH_SCOPES" arg:"google-oauth-scopes"`
		// ClientID is the OAuth client ID.
		ClientID string `default:"" env:"GOOGLE_OAUTH_CLIENT_ID" arg:"google-oauth-client-id"`
		// ClientSecret is the OAuth client secret.
		ClientSecret string `default:"" env:"GOOGLE_OAUTH_CLIENT_SECRET" arg:"google-oauth-client-secret"`
		// TokenEndpoint is Google's endpoint to exchange OAuth-code with ID token.
		TokenEndpoint string `default:"https://oauth2.googleapis.com/token" env:"GOOGLE_OAUTH_TOKEN_ENDPOINT" arg:"google-oauth-token-endpoint"`
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

	// Mongo holds the MongoDB configs.
	Mongo struct {
		// Addr is the address of the MongoDB deployment.
		Addr string `default:"mongodb://dev:dev@localhost:27017" env:"MONGO_ADDR" arg:"mongo-addr"`
		// OperationTimeoutSec is the timeout in seconds for any MongoDB operation.
		OperationTimeoutSec int `default:"60" env:"MONGO_OPERATION_TIMEOUT_SEC" arg:"mongo-operation-timeout-sec"`
		// DatabaseName is the name of the MongoDB database.
		DatabaseName string `default:"rosenchat-dev" env:"MONGO_DATABASE_NAME" arg:"mongo-database-name"`
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
