package core_http_server

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Addr            string        `envconfig:"ADDR" required=true`
	ShutdownTimeOut time.Duration `envconfig:"SHUTDOWN_TIMEOUT" default=30s`
}

func NewConfig() (Config, error) {
	var config Config

	if err := envconfig.Process("HTTP", &config); err != nil {
		panic(err)
	}

	return config, nil
}
