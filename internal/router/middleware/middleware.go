package middleware

import (
	"fmt"
	"net/http"

	"org.freethegnomes.csv_api/internal"
	"org.freethegnomes.csv_api/internal/helpers"
)

func PanicRecovery(app *internal.Application, next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				res.Header().Set("Connection", "close")
				helpers.ServerError(app.Log, res, *req, fmt.Errorf("%s", err))
			}
		}()

		next.ServeHTTP(res, req)
	})
}

func isPreflightRequest(req *http.Request) bool {
	return req.Method == http.MethodOptions && req.Header.Get("Access-Control-Request-Method") != ""
}

func ApplyBaseHeaders(app *internal.Application, next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Add("Vary", "origin")
		res.Header().Add("Vary", "Access-Control-Request-Method")
		res.Header().Add("Vary", "Access-Control-Request-Headers")

		origin := req.Header.Get("Origin")

		if origin != "" {
			for i := range app.Config.ServerConfig.TrustedOrigins {
				if origin == app.Config.ServerConfig.TrustedOrigins[i] {
					res.Header().Set("Access-Control-Allow-Origin", origin)
					if isPreflightRequest(req) {

						res.Header().Set("Access-Control-Allow-Methods", "PATCH")

						res.WriteHeader(http.StatusOK)
						return
					}
					break
				}
			}
		}

		res.Header().Set("referrer-policy", "origin-when-cross-origin")
		res.Header().Set("Access-Control-Allow-Credentials", "true")
		res.Header().Add("Server", "Go - http.ServeMux")
		next.ServeHTTP(res, req)
	})
}

func HandleBaseUrl(app *internal.Application, next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if app.Config.ServerConfig.Protocol == "fcgi" && req.Proto == "INCLUDED" {
			next.ServeHTTP(res, req)
		} else {
			http.StripPrefix(app.Config.ServerConfig.BaseUrl, next).ServeHTTP(res, req)
		}
	})
}
