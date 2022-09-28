package config

import (
	"fmt"
	"strconv"

	utils "github.com/jjuarez/gss-api/internal/utils"
	dotenv "github.com/joho/godotenv"
)

// Configuration environment variables
const (
	// GSSAPIEnvironmentEnvKey The configuration environment variables
	GSSAPI_ENV_ENVKEY = "GSSAPI_ENV"
	HTTP_HOST_ENVKEY  = "HTTP_HOST"
	HTTP_PORT_ENVKEY  = "HTTP_PORT"
)

// Some configuration default values
const (
	DEFAULT_ENV       = "development"
	DEFAULT_HTTP_HOST = "localhost"
	DEFAULT_HTTP_PORT = "8080"
	DE
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
func New() (*Config, error) {
	env := utils.Getenv(GSSAPI_ENV_ENVKEY, DEFAULT_ENV)
	host := utils.Getenv(HTTP_HOST_ENVKEY, DEFAULT_HTTP_HOST)
	port, err := strconv.Atoi(utils.Getenv(HTTP_PORT_ENVKEY, DEFAULT_HTTP_PORT))

	if err != nil {
		return nil, err
	}

	return &Config{
		Environment: env,
		Host:        host,
		Port:        port,
	}, nil
}
