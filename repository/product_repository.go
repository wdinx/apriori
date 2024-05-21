package repository

import "apriori-backend/model/domain"

type ProductRepository interface {
	GetAll() (*[]domain.Product, error)
	GetByID(id int) (*domain.Product, error)
	Create(product *domain.Product) error
	Update(product *domain.Product) error
	Delete(id int) error
}
