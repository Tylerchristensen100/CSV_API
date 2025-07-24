package internal

import (
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"time"
)

type Configuration struct {
	ServerConfig *ServerConfig
	UpTime       time.Time
}

type ServerConfig struct {
	Address        string   `yaml:"host"`
	Port           int      `yaml:"port"`
	TrustedOrigins []string `yaml:"trustedOrigins"`
}

type Secrets struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Application struct {
	Config *Configuration
	Log    *slog.Logger
	Router http.Handler
}

func (app *Application) Location() string {
	return fmt.Sprintf("%s:%d", app.Config.ServerConfig.Address, app.Config.ServerConfig.Port)
}

func (app *Application) Serve() error {

	server := &http.Server{
		Addr:     app.Location(),
		Handler:  app.Router,
		ErrorLog: slog.NewLogLogger(app.Log.Handler(), slog.LevelError),
	}

	listener, err := net.Listen("tcp", server.Addr)
	if err != nil {
		app.Log.Error("internal/Serve: Failed to Start Go Server. Could not make TCP connection", slog.String("location", server.Addr))
		return err
	}

	app.Log.Info(fmt.Sprintf("Starting Go Server on %d", app.Config.ServerConfig.Port))

	err = server.Serve(listener)

	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	app.Log.Info("Go Server Stopped", slog.String("location", server.Addr))

	return nil
}
