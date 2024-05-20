package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Get() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("error when loading env file: %s, Environment from Docker Will Be Load", err.Error())
	}
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
