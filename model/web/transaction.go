package web

import (
	"mime/multipart"
)

type CreateTransactionRequest struct {
	Date  string `json:"date" form:"date" validate:"required"`
	Items string `json:"items" form:"items" validate:"required"`
}

type UpdateTransactionRequest struct {
	ID    int
	Date  string `json:"date" form:"date" validate:"required"`
	Items string `json:"items" form:"items" validate:"required"`
}

type TransactionResponse struct {
	ID    int    `json:"id"`
	Date  string `json:"date"`
	Items string `json:"items"`
}

type InsertByExcelRequest struct {
	Excel *multipart.FileHeader `json:"excel" form:"excel" validate:"required"`
}
