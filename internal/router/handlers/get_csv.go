package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"org.freethegnomes.csv_api/internal"
	"org.freethegnomes.csv_api/internal/csv"
	"org.freethegnomes.csv_api/internal/helpers"
)

func GetCSV(app *internal.Application) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		const sortBy = "sortBy"
		const filterBy = "filterBy"
		const urlPath = "url"
		var data *csv.CSVData
		var err error
		if req.URL.Query().Has(urlPath) {
			app.Log.Debug("Query Param ", "url", req.URL.Query().Get(urlPath))
			path := req.URL.Query().Get(urlPath)
			data, err = csv.ParseFromUrl(path)
			if err != nil {
				helpers.ClientError(res, fmt.Sprintf("Error parsing csv from url %s, error: %s", path, err.Error()), http.StatusInternalServerError)
				return
			}
		} else {
			helpers.ClientError(res, "Missing required query parameter 'url'.", http.StatusBadRequest)
			return
		}

		res.Header().Set("Content-Type", "application/json; charset=utf-8")

		if req.URL.Query().Has(filterBy) {
			param := req.URL.Query().Get(filterBy)

			app.Log.Debug("Query Param ", "filterBy", param)
			params := strings.Split(param, "==")
			if len(params) != 2 {
				helpers.ClientError(res, "Invalid filter format. Use 'key==value'", http.StatusBadRequest)
				return
			}
			key, value := params[0], params[1]
			data.Filter(key, value)
		}

		if req.URL.Query().Has(sortBy) {
			param := req.URL.Query().Get(sortBy)
			app.Log.Debug("Query Param ", "sortBy", param)
			data.Sort(param)
		}

		jsonData, err := data.JSON()
		if err != nil {
			helpers.ServerError(app.Log, res, *req, fmt.Errorf("%s", "Error converting data to JSON: "+err.Error()))
			return
		}

		res.Header().Set("Content-Length", strconv.Itoa(len(jsonData)))
		res.Header().Set("Cache-Control", "public, max-age=3600")
		res.Header().Set("Content-Size", fmt.Sprintf("%.3f Kb", float64(len(jsonData))/1024.0))
		res.WriteHeader(http.StatusOK)
		if _, err := res.Write(jsonData); err != nil {
			helpers.ServerError(app.Log, res, *req, err)
			return
		}

	}
}
