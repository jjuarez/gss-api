package main

import (
	"fmt"
	"os"

	"github.com/jjuarez/gss-api/internal/config"
	"github.com/jjuarez/gss-api/internal/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

var (
	// Version the version of the GSS API release
	Version string
	// GitCommit commit information
	GitCommit string
)

const (
	CONFIGURATION_ERROR = 1
	SERVER_ERROR        = 2
)

func main() {
	environment := utils.Getenv(config.GSSAPI_ENV_ENVKEY, config.DEFAULT_ENV)
	config.SetupEnvironment(environment)

	cfg, err := config.New()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(CONFIGURATION_ERROR)
	}

	e := echo.New()
	e.HideBanner = true
	if logger, ok := e.Logger.(*log.Logger); ok {
		logger.SetHeader("${time_rfc3339} ${level}")
	}

	// Middleware

	// Routes
	e.GET("/api/v1/healthcheck", healthCheck)

	// Server
	e.Logger.Info("Starting %v server (%s) listening: %s\n", cfg.Environment, Version, cfg.Address)
	if err = e.Start(cfg.Address()); err != nil {
		e.Logger.Fatalf("Error: %v", err)
		os.Exit(SERVER_ERROR)
	}
}
