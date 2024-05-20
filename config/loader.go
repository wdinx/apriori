package config

import (
	"os"
)

func Get() *Config {
	return &Config{
		Database: Database{
			DBUser: os.Getenv("DBUSER"),
			DBPass: os.Getenv("DBPASS"),
			DBHost: os.Getenv("DBHOST"),
			DBPort: os.Getenv("DBPORT"),
			DBName: os.Getenv("DBNAME"),
		},
	}
}
