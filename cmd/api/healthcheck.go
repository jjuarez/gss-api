package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// HealthCheck ...
type HealthCheck struct {
	Status  string `json:"status"`
	Version string `json:"version"`
}

func healthCheck(c echo.Context) error {
	status := HealthCheck{
		Status:  "OK",
		Version: Version,
	}
	return c.JSONPretty(http.StatusOK, status, "  ")
}
