package converter

import (
	"apriori-backend/model/domain"
	"apriori-backend/model/web"
	"apriori-backend/util"
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
	"strings"
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
	defer file.Close()

	xlsx, err := excelize.OpenReader(file)
	if err != nil {
		return nil, errors.New("format file not supported")
	}

	rows, err := xlsx.GetRows("Transaksi")
	if err != nil {
		return nil, errors.New("no sheet named 'Transaksi' found in the file")
	}

	transactions := make([]domain.Transaction, 0)

	// supported date formats
	dateLayouts := []string{
		"02-01-06",   // dd-mm-yy
		"02-01-2006", // dd-mm-yyyy
		"01-02-06",   // mm-dd-yy
		"01-02-2006", // mm-dd-yyyy
	}

	for i, row := range rows {
		// skip header
		if i == 0 {
			continue
		}

		// skip baris kosong
		if len(row) < 2 || strings.TrimSpace(row[0]) == "" {
			continue
		}

		dateStr := strings.TrimSpace(row[0])
		items := strings.TrimSpace(row[1])

		var date time.Time
		var parseErr error

		for _, layout := range dateLayouts {
			date, parseErr = time.Parse(layout, dateStr)
			if parseErr == nil {
				break
			}
		}

		if parseErr != nil {
			return nil, fmt.Errorf(
				"invalid date format at row %d: %s (expected dd-mm-yy / dd-mm-yyyy)",
				i+1,
				dateStr,
			)
		}

		transactions = append(transactions, domain.Transaction{
			Date:  date,
			Items: items,
		})
	}

	return &transactions, nil
}
