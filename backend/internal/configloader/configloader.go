package configloader

import (
	"fmt"
	"os"

	"github.com/psmey/time_tracker/internal/database"
	"gopkg.in/yaml.v3"
)

type ConfigLoader struct{}

type Config struct {
	Database *database.Config `yaml:"database"`
}

func New() *ConfigLoader {
	return &ConfigLoader{}
}

func (configLoader *ConfigLoader) Load(path string) (*Config, error) {
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
