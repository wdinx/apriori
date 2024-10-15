package repository

import (
	"apriori-backend/config"
	"context"
	"errors"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"mime/multipart"
)

type ImageRepositoryImpl struct {
	cfg config.Cloudinary
}

func NewImageRepository(cfg config.Cloudinary) ImageRepository {
	return &ImageRepositoryImpl{cfg: cfg}
}

func (repository *ImageRepositoryImpl) UploadImage(file multipart.File) (string, error) {
	if repository.cfg.CloudinaryURL == "" {
		return "", errors.New("cloudinary url is empty")
	}

	cld, err := cloudinary.NewFromURL(repository.cfg.CloudinaryURL)
	if err != nil {
		return "", err
	}

	uploadResult, err := cld.Upload.Upload(context.Background(), file, uploader.UploadParams{})
	if err != nil {
		return "", err
	}

	return uploadResult.SecureURL, nil
}
