package helpers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"runtime/debug"
)

// The serverError helper writes a log entry at Error level (including the request
// method and URI as attributes), then sends a generic 500 Internal Server Error
// response to the user.
func ServerError(logger *slog.Logger, res http.ResponseWriter, req http.Request, err error) {
	var (
		method = req.Method
		uri    = req.URL.RequestURI()
		trace  = string(debug.Stack())
	)

	logger.Error("helpers/ServerError: "+err.Error(), slog.String("method", method), slog.String("uri", uri), slog.String("stacktrace", trace))
	data, err := writeHTTPError(http.StatusInternalServerError, err.Error(), &req.Host)
	if err != nil {
		slog.Error("helpers/ServerError: Failed to marshal error response", slog.String("error", err.Error()))
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if _, err := res.Write(data); err != nil {
		slog.Error("helpers/ServerError: Failed to write error response", slog.String("error", err.Error()))
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

// The clientError helper sends a specific status code and corresponding description
// to the user.
func ClientError(res http.ResponseWriter, message string, status int) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(status)
	data, err := writeHTTPError(status, message, nil)
	if err != nil {
		slog.Error("helpers/ClientError: Failed to marshal error response", slog.String("error", err.Error()))
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if _, err := res.Write(data); err != nil {
		slog.Error("helpers/ClientError: Failed to write error response", slog.String("error", err.Error()))
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

const DocsRoute = "/docs"

func writeHTTPError(status int, message string, host *string) ([]byte, error) {
	e := errorResponse{
		StatusCode: status,
		Message:    message,
	}
	if host != nil {
		e.Docs = *host + DocsRoute
	} else {
		e.Docs = "." + DocsRoute
	}

	data, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}
	return data, nil
}

type errorResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Docs       string `json:"docs"`
}
