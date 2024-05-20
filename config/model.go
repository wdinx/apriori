package config

type Config struct {
	Database Database
}

type Database struct {
	DBUser string
	DBPass string
	DBHost string
	DBPort string
	DBName string
}
