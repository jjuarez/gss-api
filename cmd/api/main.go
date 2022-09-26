package main

import (
	"fmt"
	"net/http"
	"os"

	config "github.com/jjuarez/gss-api/internal/config"
	utils "github.com/jjuarez/gss-api/internal/utils"
	"github.com/labstack/echo/v4"
)

var (
	// Version the version of the GSS API release
	Version string
)

func main() {
	environment := utils.Getenv(config.GSSAPIEnvironmentEnvKey, config.DefaultEnvironment)
	config.SetupEnvironment(environment)

	srvCfg, err := config.NewConfig()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/v1/api/healthcheck", func(c echo.Context) error {
		return c.String(http.StatusOK, "Super healthy")
	})

	e.Logger.Fatal(e.Start(srvCfg.Address()))
	e.Logger.Info("Starting %v server (%s) listening on -> %s\n", srvCfg.Environment, Version, srvCfg.Address)
}
