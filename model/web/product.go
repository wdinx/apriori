package web

import "mime/multipart"

type ProductCreateRequest struct {
	Name  string                `form:"name" validate:"required,min=1,max=200"`
	Price int                   `form:"price" validate:"required,min=1"`
	Image *multipart.FileHeader `form:"image"`
	Stock int                   `form:"stock" validate:"required,min=1"`
}

type ProductUpdateRequest struct {
	ID    int
	Name  string                `form:"name" validate:"required,min=1,max=200"`
	Price int                   `form:"price" validate:"required,min=1"`
	Image *multipart.FileHeader `form:"image"`
	Stock int                   `form:"stock" validate:"required,min=1"`
}

type ProductResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Image string `json:"image"`
	Stock int    `json:"stock"`
}
