package util

import (
	"apriori-backend/config"
	"context"
	"errors"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"mime/multipart"
)

func UploadImage(file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	cloudinaryURL := config.Get().Cloudinary.CloudinaryURL
	if cloudinaryURL == "" {
		return "", errors.New("cloudinary url is empty")
	}

	cld, err := cloudinary.NewFromURL(cloudinaryURL)
	if err != nil {
		return "", err
	}

	uploadResult, err := cld.Upload.Upload(context.Background(), src, uploader.UploadParams{})
	if err != nil {
		return "", err
	}

	return uploadResult.SecureURL, nil
}
