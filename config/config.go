package config

import (
	"errors"
	"os"

	"github.com/BurntSushi/toml"
)

var ErrNoConfFile = errors.New("no configuration file specified")

type Config struct {
	Environment string `toml:"environment"`
	Bind        string `toml:"bind"`
	ApiUrl      string `toml:"api_url"`

	// [logging]
	Logging struct {
		Level string `toml:"level"`
	} `toml:"logging"`

	// [db]
	// DB data.DBConf `toml:"db"`
}

func New() *Config {
	return &Config{}
}

func NewFromFile(file string, env string) (*Config, error) {
	if file == "" {
		file = env
	}

	if _, err := os.Stat(file); os.IsNotExist(err) {
		return nil, ErrNoConfFile
	}

	var conf Config
	if _, err := toml.DecodeFile(file, &conf); err != nil {
		return nil, err
	}

	return &conf, nil
}
