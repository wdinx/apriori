package service

import "mime/multipart"

type ImageService interface {
	UploadImage(image *multipart.FileHeader, filename string) error
}
