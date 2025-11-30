package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Connector struct {
	Config *Config
}

type Config struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	SSLMode  string `yaml:"sslmode"`
}

func NewConnector(config *Config) *Connector {
	return &Connector{Config: config}
}

func (connector *Connector) NewPostgresDB(dbname string) (*sql.DB, error) {
	connectionString := fmt.Sprintf(
		"dbname=%s user=%s password=%s host=%s port=%s sslmode=%s",
		dbname,
		connector.Config.User,
		connector.Config.Password,
		connector.Config.Host,
		connector.Config.Port,
		connector.Config.SSLMode,
	)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to open DB: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping DB: %w", err)
	}

	return db, nil
}
