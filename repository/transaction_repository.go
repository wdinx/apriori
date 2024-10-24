package repository

import (
	"apriori-backend/model/domain"
	"apriori-backend/model/web"
)

type TransactionRepository interface {
	Create(transaction *domain.Transaction) error
	Update(transaction *domain.Transaction) error
	Delete(id int) error
	FindById(id int) (*domain.Transaction, error)
	FindAll(metadata *web.Metadata) (*[]domain.Transaction, error)
	GetALl() (*[]domain.Transaction, error)
	FindByDateRange(startDate string, endDate string) (*[]domain.Transaction, error)
}
