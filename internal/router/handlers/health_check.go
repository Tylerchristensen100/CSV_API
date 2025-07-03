package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"org.freethegnomes.csv_api/internal"
	"org.freethegnomes.csv_api/internal/helpers"
)

func HealthCheck(app *internal.Application) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

		duration := time.Since(app.Config.UpTime)
		hours := int(duration.Hours())
		minutes := int(duration.Minutes()) % 60
		seconds := int(duration.Seconds()) % 60

		status := map[string]interface{}{
			"status": "ok",
			"uptime": fmt.Sprintf("%d hours, %d minutes, %d seconds", hours, minutes, seconds),
		}

		data, err := json.Marshal(status)
		if err != nil {
			helpers.ServerError(app.Log, res, *req, err)
			return
		}

		res.Header().Set("Content-Type", "application/json; charset=utf-8")
		res.Write(data)
		res.WriteHeader(http.StatusOK)
	}
}
