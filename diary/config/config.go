package config

import (
	"fmt"
	"os"
)

type Config struct {
	DbDsn string
}

func Load() (*Config, error) {
	config := &Config{}

	dbDsn := os.Getenv("DATABASE_DSN")
	if dbDsn == "" {
		return nil, fmt.Errorf("Specify DATABASE_DSN")
	}
	config.DbDsn = dbDsn

	return config, nil
}
