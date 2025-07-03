package helpers

import (
	"log/slog"
	"os"

	"gopkg.in/yaml.v3"
	"org.freethegnomes.csv_api/internal"
)

func OpenConfigYaml(logger *slog.Logger, path string) *internal.ServerConfig {
	file, err := os.Open(path)
	if err != nil {
		logger.Error("helpers/OpenConfigYaml: " + err.Error())
		os.Exit(1)
	}
	defer file.Close()

	var config internal.ServerConfig

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		logger.Error("helpers/OpenConfigYaml: " + err.Error())
		os.Exit(1)
	}

	return &config
}
