package repository

import "apriori-backend/model/domain"

type TransactionRepository interface {
	Create(transaction *domain.Transaction) error
	Update(transaction *domain.Transaction) error
	Delete(id int) error
	FindById(id int) (*domain.Transaction, error)
	FindAll() (*[]domain.Transaction, error)
}
