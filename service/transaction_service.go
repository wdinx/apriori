package service

import "apriori-backend/model/web"

type TransactionService interface {
	Create(request *web.CreateTransactionRequest) error
	Update(request *web.UpdateTransactionRequest) error
	Delete(id int) error
	FindById(id int) (*web.TransactionResponse, error)
	FindAll(metadata *web.Metadata) (*[]web.TransactionResponse, error)
	InsertByExcel(request *web.InsertByExcelRequest) error
}
