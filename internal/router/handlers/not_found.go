package handlers

import (
	"net/http"

	"org.freethegnomes.csv_api/internal"
	"org.freethegnomes.csv_api/internal/helpers"
)

func NotFound(app *internal.Application) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		app.Log.Debug("Not Found", "method", req.Method, "url", req.URL.Path)
		helpers.ClientError(res, "The requested resource was not found.", http.StatusNotFound)
	}
}
