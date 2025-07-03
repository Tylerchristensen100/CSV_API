package docs

import (
	"html/template"
	"log/slog"
	"net/http"
	"path/filepath"

	"org.freethegnomes.csv_api/internal/helpers"
)

type OpenAPI struct {
	Page   func(http.ResponseWriter, *http.Request)
	Config func(http.ResponseWriter, *http.Request)
}

func Server(logger *slog.Logger) *OpenAPI {
	html, err := filepath.Abs("etc/swagger/index.html")
	if err != nil {
		logger.Error("docs/Server: " + err.Error())
	}
	yaml, err := filepath.Abs("etc/swagger/openapi.yaml")
	if err != nil {
		logger.Error("docs/Server: " + err.Error())
	}

	return &OpenAPI{
		Page: func(res http.ResponseWriter, req *http.Request) {
			http.ServeFile(res, req, html)
		},
		Config: func(res http.ResponseWriter, req *http.Request) {
			tmpl, err := template.ParseFiles(yaml)
			if err != nil {
				helpers.ServerError(logger, res, *req, err)
				return
			}

			res.Header().Set("Content-Type", "application/yaml")
			err = tmpl.Execute(res, nil)
			if err != nil {
				helpers.ServerError(logger, res, *req, err)
				return
			}
			res.WriteHeader(http.StatusOK)
		},
	}
}
