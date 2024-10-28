package converter

import (
	"apriori-backend/model/domain"
	"apriori-backend/model/web"
	"apriori-backend/util"
	"errors"
	"github.com/xuri/excelize/v2"
	"time"
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

func ToTransactionModelByExcel(r *web.InsertByExcelRequest) (*[]domain.Transaction, error) {
	file, err := r.Excel.Open()
	if err != nil {
		return nil, errors.New("failed to open file")
	}

	xlsx, err := excelize.OpenReader(file)
	if err != nil {
		return nil, errors.New("format file not supported")
	}

	rows, err := xlsx.GetRows("Transaksi")
	if err != nil {
		return nil, errors.New("no sheet named 'Transaksi' found in the file")
	}

	transactions := make([]domain.Transaction, 0)
	for i, row := range rows {
		if i == 0 {
			continue
		}
		date, err := time.Parse("01-02-06", row[0])
		if err != nil {
			return nil, errors.New("invalid date format")
		}
		transactions = append(transactions, domain.Transaction{
			Date:  date,
			Items: row[1],
		})
	}
	return &transactions, nil
}
