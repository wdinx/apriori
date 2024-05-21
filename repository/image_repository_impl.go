package repository

import (
	"apriori-backend/config"
	"apriori-backend/constant"
	"apriori-backend/util"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"mime/multipart"
)

type ImageRepositoryImpl struct {
	cfg config.DigitalOceanSpaces
}

func NewImageRepository(cfg config.DigitalOceanSpaces) ImageRepository {
	return &ImageRepositoryImpl{cfg: cfg}
}

func (repository *ImageRepositoryImpl) UploadImage(image multipart.File, filename string) error {
	newSession, err := session.NewSession(config.DigitalOceanSpacesConfig(&repository.cfg))
	if err != nil {
		return constant.ErrInternalServer
	}
	s3Client := s3.New(newSession)

	contentType, err := util.GetContentType(filename)
	if err != nil {
		return err
	}

	object := s3.PutObjectInput{
		Bucket:      aws.String(repository.cfg.Name),
		Key:         aws.String(filename),
		Body:        image,
		ACL:         aws.String("public-read"),
		ContentType: aws.String(contentType),
	}

	_, err = s3Client.PutObject(&object)
	if err != nil {
		return constant.ErrInternalServer
	}
	return nil
}
