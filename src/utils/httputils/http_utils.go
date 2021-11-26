package httputils

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"rosenchat/src/exception"
	"rosenchat/src/logger"
	"strings"
)

var log = logger.Get()

// WriteJSON marshals the provided data into JSON and sends it as the HTTP response along with the provided status.
func WriteJSON(ctx context.Context, writer http.ResponseWriter, body interface{}, headers map[string]string, status int) {
	// Setting all the headers.
	writer.Header().Set("content-type", "application/json")
	for key, value := range headers {
		writer.Header().Set(key, value)
	}

	// Setting the status code.
	writer.WriteHeader(status)

	// Marshalling the body.
	response, err := json.Marshal(body)
	if err != nil {
		log.Warnf(ctx, "Failed to marshal HTTP response: %+v", err)
		return
	}

	// Writing the body to the response.
	if _, err := writer.Write(response); err != nil {
		log.Warnf(ctx, "Failed to write HTTP response: %+v", err)
	}
}

// GetIPAddrFromRequest extracts the client IP Address from the given HTTP request.
func GetIPAddrFromRequest(ctx context.Context, req *http.Request) string {
	// Using x-real-ip header.
	ip := req.Header.Get("x-real-ip")
	if parsedIP := net.ParseIP(ip); parsedIP != nil {
		return parsedIP.String()
	}

	// Using x-forwarded-for header.
	ips := req.Header.Get("x-forwarded-for")
	ipArr := strings.Split(ips, ",")
	if len(ipArr) > 0 {
		if parsedIP := net.ParseIP(ipArr[0]); parsedIP != nil {
			return parsedIP.String()
		}
	}

	// Using RemoteAddr property.
	ip, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		log.Warnf(ctx, "Error in SplitHostPort call: %s", err.Error())
		return "unknown"
	}

	if parsedIP := net.ParseIP(ip); parsedIP != nil {
		return parsedIP.String()
	}

	log.Warnf(ctx, "Failed to obtain IP address of client.")
	return "unknown"
}

// UnmarshalBody reads the body of the given HTTP request and decodes it into the provided interface.
func UnmarshalBody(ctx context.Context, req *http.Request, target interface{}) error {
	defer func() {
		_ = req.Body.Close()
	}()

	bodyBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Errorf(ctx, "Failed to read request body because: %+v", err)
		return exception.BadRequest().AddErrors(err)
	}

	if err := json.Unmarshal(bodyBytes, target); err != nil {
		log.Errorf(ctx, "Failed to unmarshal request body because: %+v", err)
		return exception.BadRequest().AddErrors(err)
	}

	return nil
}
