package main

import (
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"org.freethegnomes.csv_api/internal"
	"org.freethegnomes.csv_api/internal/helpers"
	"org.freethegnomes.csv_api/internal/router"
)

const configPath = "etc/.conf/config.yaml"

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	conf, err := filepath.Abs(configPath)
	if err != nil {
		logger.Error("server/main: " + err.Error())
		os.Exit(1)
	}

	config := &internal.Configuration{
		ServerConfig: helpers.OpenConfigYaml(logger, conf),
		UpTime:       time.Now(),
	}

	app := &internal.Application{
		Config: config,
		Log:    logger,
	}

	app.Router = router.CreateRouter(app)

	app.Serve()
}
