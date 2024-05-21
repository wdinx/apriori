package converter

import (
	"apriori-backend/model/domain"
	"apriori-backend/model/web"
	"apriori-backend/util"
)

func CreateToTransactionModel(transaction *web.CreateTransactionRequest) *domain.Transaction {
	transactionDate := util.StringToDate(transaction.Date)
	return &domain.Transaction{
		Items: transaction.Items,
		Date:  transactionDate,
	}
}

func UpdateToTransactionModel(transaction *web.UpdateTransactionRequest) *domain.Transaction {
	transactionDate := util.StringToDate(transaction.Date)
	return &domain.Transaction{
		ID:    transaction.ID,
		Items: transaction.Items,
		Date:  transactionDate,
	}
}

func ToTransactionResponse(transaction *domain.Transaction) *web.TransactionResponse {
	return &web.TransactionResponse{
		ID:    transaction.ID,
		Date:  transaction.Date.Format("2006-01-02"),
		Items: transaction.Items,
	}
}
