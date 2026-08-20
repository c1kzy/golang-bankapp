package core_logger

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Level  string `envconfig:"LEVEL" default:DEBUG`
	Folder string `envconfig:"FOLDER" required:"true"`
}

func NewConfig() (Config, error) {
	var config Config

	if err := envconfig.Process("LOGGER", &config); err != nil {
		panic(err)
	}

	return config, nil
}
