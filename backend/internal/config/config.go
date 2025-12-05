package config

import (
	"fmt"
	"os"

	"github.com/psmey/time_tracker/internal/auth/keycloak"
	"github.com/psmey/time_tracker/internal/database"
	"github.com/psmey/time_tracker/internal/logger"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Logger   *logger.Config   `yaml:"logger"`
	Database *database.Config `yaml:"database"`
	Keycloak *keycloak.Config `yaml:"keycloak"`
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
