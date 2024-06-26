package service

import "apriori-backend/model/web"

type TransactionService interface {
	Create(request *web.CreateTransactionRequest) error
	Update(request *web.UpdateTransactionRequest) error
	Delete(id int) error
	FindById(id int) (*web.TransactionResponse, error)
	FindAll() (*[]web.TransactionResponse, error)
}
