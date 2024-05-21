package repository

import "mime/multipart"

type ImageRepository interface {
	UploadImage(image multipart.File, filename string) error
}
