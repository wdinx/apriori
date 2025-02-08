package service

import (
	"apriori-backend/model/web"
	"apriori-backend/repository"
	"apriori-backend/util/converter"

	"github.com/go-playground/validator/v10"
)

type TransactionServiceImpl struct {
	transactionRepository repository.TransactionRepository
	validator             *validator.Validate
}

func NewTransactionService(transactionRepository repository.TransactionRepository, validator *validator.Validate) TransactionService {
	return &TransactionServiceImpl{transactionRepository: transactionRepository, validator: validator}
}

// Melakukan request ke repository untuk membuat transaksi dan mengirimnya ke kontroller
func (service *TransactionServiceImpl) Create(request *web.CreateTransactionRequest) error {
	if err := service.validator.Struct(request); err != nil {
		return err
	}
	if err := service.transactionRepository.Create(converter.CreateToTransactionModel(request)); err != nil {
		return err
	}
	return nil
}

// Melakukan request ke repository untuk update data transaksi dan mengirimnya ke kontroller
func (service *TransactionServiceImpl) Update(request *web.UpdateTransactionRequest) error {
	if err := service.validator.Struct(request); err != nil {
		return err
	}
	_, err := service.transactionRepository.FindById(request.ID)
	if err != nil {
		return err
	}

	if err = service.transactionRepository.Update(converter.UpdateToTransactionModel(request)); err != nil {
		return err
	}
	return nil
}

// Melakukan request ke repository untuk menghapus transaksi dan mengirimnya ke kontroller
func (service *TransactionServiceImpl) Delete(id int) error {
	if err := service.transactionRepository.Delete(id); err != nil {
		return err
	}
	return nil
}

// Melakukan request ke repository untuk mengambil data transaksi by id dan mengirimnya ke kontroller
func (service *TransactionServiceImpl) FindById(id int) (*web.TransactionResponse, error) {
	transaction, err := service.transactionRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return converter.ToTransactionResponse(transaction), nil
}

// Melakukan request ke repository untuk mengambil semua data transaksi dan mengirimnya ke kontroller
func (service *TransactionServiceImpl) FindAll(metadata *web.Metadata) (*[]web.TransactionResponse, error) {
	transactions, err := service.transactionRepository.FindAll(metadata)
	if err != nil {
		return nil, err
	}

	var responses []web.TransactionResponse
	for _, transaction := range *transactions {
		responses = append(responses, *converter.ToTransactionResponse(&transaction))
	}
	return &responses, nil
}

// Melakukan request ke repository untuk membuat transaksi berdasarkan data excel dan mengirimnya ke kontroller
func (service *TransactionServiceImpl) InsertByExcel(request *web.InsertByExcelRequest) error {
	if err := service.validator.Struct(request); err != nil {
		return err
	}
	transaction, err := converter.ToTransactionModelByExcel(request)
	if err != nil {
		return err
	}
	if err = service.transactionRepository.InsertByExcel(transaction); err != nil {
		return err
	}
	return nil
}
