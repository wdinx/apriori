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

// Mengirim data gambar ke repository
func (service *ImageServiceImpl) UploadImage(image *multipart.FileHeader) (string, error) {
	file, err := image.Open()
	defer file.Close()

	if err != nil {
		return "", constant.ErrInternalServer
	}

	filename, err := service.imageRepository.UploadImage(file)
	if err != nil {
		return "", err
	}

	return filename, nil
}
