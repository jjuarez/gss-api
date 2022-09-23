package config

import (
	"fmt"
	"strconv"

	utils "github.com/jjuarez/gss-api/internal/utils"
	dotenv "github.com/joho/godotenv"
)

const (
	// GSSAPIEnvironmentEnvKey The configuration environment variables
	GSSAPIEnvironmentEnvKey = "GSSAPI_ENV"
	httpHostEnvKey          = "HTTP_HOST"
	httpPortEnvKey          = "HTTP_PORT"

	// DefaultEnvironment ...
	DefaultEnvironment = "development"
	// DefaultHTTPHost ...
	DefaultHTTPHost = "localhost"
	// DefaultHTTPPort ...
	DefaultHTTPPort = "8080"
)

// SetupEnvironment ...
func SetupEnvironment(environment string) {
	configFiles := []string{
		".env." + environment + ".local",
		".env." + environment,
	}
	for _, f := range configFiles {
		dotenv.Load(f)
	}
	dotenv.Load()
}

// Stringer interface
type Stringer interface {
	String() string
}

// Addresser interface
type Addresser interface {
	Address() string
}

// Config ...
type Config struct {
	Environment string
	Host        string `valid:"host"`
	Port        int    `valid:"port"`
}

// String implements the interface for config.Config
func (c Config) String() string {
	return fmt.Sprintf("{'environment':'%s','host':'%s','port':'%d'}", c.Environment, c.Host, c.Port)
}

// Address implements the interface for config.Config
func (c Config) Address() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

// NewConfig ...
func NewConfig() (*Config, error) {
	env := utils.Getenv(GSSAPIEnvironmentEnvKey, DefaultEnvironment)
	host := utils.Getenv(httpHostEnvKey, DefaultHTTPHost)
	port, err := strconv.Atoi(utils.Getenv(httpPortEnvKey, DefaultHTTPPort))

	if err != nil {
		return nil, err
	}

	return &Config{
		Environment: env,
		Host:        host,
		Port:        port,
	}, nil
}
