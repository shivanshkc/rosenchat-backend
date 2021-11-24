package middlewares

import (
	"net/http"
)

// CORS middlewares handled the CORS issues.
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Headers", "*")
		writer.Header().Set("Access-Control-Allow-Methods", "*")
		writer.Header().Set("Access-Control-Allow-Credentials", "*")

		if req.Method == http.MethodOptions {
			writer.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(writer, req)
	})
}
