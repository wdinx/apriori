package service

import (
	"apriori-backend/constant"
	"apriori-backend/repository"
	"mime/multipart"
)

type ImageServiceImpl struct {
	imageRepository repository.ImageRepository
}

func NewImageService(imageRepository repository.ImageRepository) ImageService {
	return &ImageServiceImpl{imageRepository: imageRepository}
}

func (service *ImageServiceImpl) UploadImage(image *multipart.FileHeader, filename string) error {
	file, err := image.Open()
	defer file.Close()

	if err != nil {
		return constant.ErrInternalServer
	}

	err = service.imageRepository.UploadImage(file, filename)
	if err != nil {
		return err
	}

	return nil
}
