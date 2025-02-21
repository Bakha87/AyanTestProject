package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type (
	Configuration struct {
		App `yaml:"app"`
		PG  `yaml:"postgres"`
	}

	App struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
	}

	PG struct {
		URL          string `yaml:"url"`
		PoolMax      int    `yaml:"poolMax"`
		ConnAttempts int    `yaml:"connAttempts"`
	}
)

func New(configPath string) (*Configuration, error) {
	cfg := &Configuration{}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)

	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
