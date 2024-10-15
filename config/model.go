package config

type Config struct {
	Database           Database
	DigitalOceanSpaces DigitalOceanSpaces
	Cloudinary         Cloudinary
}

type Database struct {
	DBUser string
	DBPass string
	DBHost string
	DBPort string
	DBName string
}
type DigitalOceanSpaces struct {
	AccessToken string
	SecretKey   string
	Region      string
	Name        string
	Endpoint    string
}
type Cloudinary struct {
	CloudinaryURL string
}
