package config

import (
	"fmt"
	"os"

	"github.com/psmey/time_tracker/internal/database"
	"github.com/psmey/time_tracker/internal/logging"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Logger   *logging.Config  `yaml:"logger"`
	Database *database.Config `yaml:"database"`
}

func Load(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	config := &Config{}

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(config); err != nil {
		return nil, fmt.Errorf("failed to decode yaml config: %w", err)
	}

	return config, nil
}
