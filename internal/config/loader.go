package config

import (
	"flag"
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

func LoadConfig() (*Config, error) {
	const ap = "config.loadConfig"

	var (
		configPath string
		cfg        Config
		err        error
	)
	configPath = fetchConfigPath()
	if configPath == "" {
		return nil, ErrEmptyConfigPath
	}
	err = cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", ap, err)
	}
	return &cfg, nil
}

func fetchConfigPath() string {
	var path string
	flag.StringVar(&path, "config", "", "path to config file")
	flag.Parse()
	return path
}
