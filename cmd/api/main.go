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

	logger := zerolog.New(os.Stdout).Level(zerolog.DebugLevel)
	srvCfg, err := config.NewConfig()
	if err != nil {
		logger.Error().Msgf("Error: %v", err)
		os.Exit(1)
	}

	app := &application{
		cfg:    *srvCfg,
		logger: logger,
	}

	mux := http.NewServeMux()
	//mux.HandleFunc("/v2/healhcheck", app.healthcheckHandler)

	srv := &http.Server{
		Addr:         app.cfg.Address(),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger.Debug().Msgf("Starting: '%v' server listening on -> %s", app.cfg.Environment, srv.Addr)
	err = srv.ListenAndServe()
	logger.Fatal()
}
