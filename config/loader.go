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
		DigitalOceanSpaces: DigitalOceanSpaces{
			AccessToken: os.Getenv("DO_ACCESS_TOKEN"),
			SecretKey:   os.Getenv("DO_SECRET_KEY"),
			Region:      os.Getenv("DO_REGION"),
			Name:        os.Getenv("DO_BUCKET"),
			Endpoint:    os.Getenv("DO_ENDPOINT"),
		},
		Cloudinary: Cloudinary{
			CloudinaryURL: os.Getenv("CLOUDINARY_URL"),
		},
	}
}
