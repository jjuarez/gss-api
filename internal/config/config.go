package config

import (
	"fmt"
	"strconv"

	utils "github.com/jjuarez/gss-api/internal/utils"
)

const (
	httpHostEnvKey = "HTTP_HOST"
	httpPortEnvKey = "HTTP_PORT"
	// DefaultHTTPHost ...
	DefaultHTTPHost = "0.0.0.0"
	// DefaultHTTPPort ...
	DefaultHTTPPort = "8080"
)

type Stringer interface {
	String() string
}

// ServerConfig ...
type ServerConfig struct {
	host string
	port int
}

func (sc ServerConfig) String() string {
	return fmt.Sprintf("%s:%d", sc.host, sc.port)
}

// New ...
func New() (*ServerConfig, error) {
	host := utils.GetEnvWithDefault(httpHostEnvKey, DefaultHTTPHost)
	port := utils.GetEnvWithDefault(httpPortEnvKey, DefaultHTTPPort)

	numPort, err := strconv.Atoi(port)
	if err != nil {
		return nil, fmt.Errorf("invalid HTTP server port: %v", port)
	}

	if numPort < 1024 || numPort > 65535 {
		return nil, fmt.Errorf("HTTP server port out of valid range: %v", port)
	}

	return &ServerConfig{
		host: host,
		port: numPort,
	}, nil
}
