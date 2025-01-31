package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// Http config hosts only server host and port.
type HTTPConfig struct {
	Host string `default:"0.0.0.0"`
	Port string `default:"8080"`
}

// Config is an extensible config data. any further config variable should go into it
type Config struct {
	Http   *HTTPConfig
	ApiKey string
}

// Load config data at start up
func Load() *Config {
	var httpConfig HTTPConfig
	err := godotenv.Load()
	if err != nil {
		panic("Server can not start")
	}
	// any env data with SERVER_ must be processed
	envconfig.MustProcess("SERVER", &httpConfig)
	ApiKey := os.Getenv("ApiKey")

	return &Config{
		Http:   &httpConfig,
		ApiKey: ApiKey,
	}
}
