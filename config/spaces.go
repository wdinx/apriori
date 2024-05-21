package config

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

func DigitalOceanSpacesConfig(cfg *DigitalOceanSpaces) *aws.Config {
	return &aws.Config{
		Credentials:      credentials.NewStaticCredentials(cfg.AccessToken, cfg.SecretKey, ""),
		Endpoint:         aws.String(cfg.Endpoint),
		S3ForcePathStyle: aws.Bool(false),
		Region:           aws.String(cfg.Region),
	}
}
