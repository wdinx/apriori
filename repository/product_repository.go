package repository

import (
	"apriori-backend/model/domain"
	"apriori-backend/model/web"
)

type ProductRepository interface {
	GetAll(metadata *web.Metadata) (*[]domain.Product, error)
	GetByID(id int) (*domain.Product, error)
	Create(product *domain.Product) error
	Update(product *domain.Product) error
	Delete(id int) error
	GetTotalPage(limit int) (int, error)
}
