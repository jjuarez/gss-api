package main

import (
	"net/http"
	"os"
	"time"

	config "github.com/jjuarez/gss-api/internal/config"
	utils "github.com/jjuarez/gss-api/internal/utils"
	"github.com/rs/zerolog"
)

var (
	// Version the version of the GSS API release
	Version string
	// GitCommit information about the CVS
	GitCommit string
)

type application struct {
	cfg    config.Config
	logger zerolog.Logger
}

func main() {
	environment := utils.Getenv(config.GSSAPIEnvironmentEnvKey, config.DefaultEnvironment)
	config.SetupEnvironment(environment)

	logger := zerolog.New(os.Stdout).
		Level(zerolog.TraceLevel).
		With().
		Timestamp().
		Caller().
		Logger()

	srvCfg, err := config.NewConfig()
	if err != nil {
		logger.Error().Msgf("Error: %v", err)
		os.Exit(1)
	}

	app := &application{
		cfg:    *srvCfg,
		logger: logger,
	}

	srv := &http.Server{
		Addr:         app.cfg.Address(),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger.Debug().Msgf("Starting %v server (%s) listening on -> %s", app.cfg.Environment, Version, srv.Addr)
	err = srv.ListenAndServe()
	logger.Error().Msgf("Error: %v", err)
}
