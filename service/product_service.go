package service

import (
	"apriori-backend/model/web"
)

type ProductService interface {
	GetAll(metadata *web.Metadata) (*[]web.ProductResponse, error)
	GetByID(id int) (*web.ProductResponse, error)
	Create(product *web.ProductCreateRequest) error
	Update(product *web.ProductUpdateRequest) error
	Delete(id int) error
}
