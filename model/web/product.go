package web

import "mime/multipart"

type ProductCreateRequest struct {
	Name  string                `form:"name" validate:"required,min=1,max=200"`
	Price int                   `form:"price" validate:"required,min=1"`
	Image *multipart.FileHeader `form:"image"`
}

type ProductUpdateRequest struct {
	ID    int
	Name  string                `form:"name" validate:"required,min=1,max=200"`
	Price int                   `form:"price" validate:"required,min=1"`
	Image *multipart.FileHeader `form:"image"`
}

type ProductResponse struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
	Image string `json:"image"`
}
