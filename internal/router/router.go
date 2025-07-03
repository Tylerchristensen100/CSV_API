package router

import (
	"expvar"
	"net/http"

	"org.freethegnomes.csv_api/internal"
	"org.freethegnomes.csv_api/internal/docs"
	"org.freethegnomes.csv_api/internal/helpers"
	"org.freethegnomes.csv_api/internal/router/handlers"
	"org.freethegnomes.csv_api/internal/router/middleware"
)

func CreateRouter(app *internal.Application) http.Handler {
	server := http.NewServeMux()

	documentation := docs.Server(app.Log)

	server.HandleFunc("GET "+helpers.DocsRoute, documentation.Page)
	server.HandleFunc("GET "+helpers.DocsRoute+"/config", documentation.Config)

	server.HandleFunc("GET /healthz", handlers.HealthCheck(app))
	server.Handle("GET /server-status", expvar.Handler())

	server.HandleFunc("GET /", handlers.GetCSV(app))

	server.HandleFunc("/", handlers.NotFound(app))
	server.HandleFunc("GET /favicon.ico", handlers.NotFound(app))

	return createMiddlewareStack(app, server)
}

func createMiddlewareStack(app *internal.Application, server http.Handler) http.Handler {
	return middleware.PanicRecovery(app,
		middleware.RequestLogging(app,
			middleware.HandleBaseUrl(app,
				middleware.ApplyBaseHeaders(app,
					server))))
}
