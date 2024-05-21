package exception

import (
	"apriori-backend/constant"
	"errors"
	"net/http"
)

func ErrorHandler(err error) int {

	if errors.Is(err, constant.ErrEmptyInput) || errors.Is(err, constant.ErrUnsupportedFileFormat) || errors.Is(err, constant.ErrInputDate) {
		return http.StatusBadRequest
	} else if errors.Is(err, constant.ErrInvalidToken) || errors.Is(err, constant.ErrLogin) || errors.Is(err, constant.ErrRegister) {
		return http.StatusUnauthorized
	} else if errors.Is(err, constant.ErrDataNotFound) || errors.Is(err, constant.ErrPaymentNotFound) || errors.Is(err, constant.ErrUnauthorized) {
		return http.StatusNotFound
	} else if errors.Is(err, constant.ErrPaymentAlreadyConfirmed) {
		return http.StatusConflict
	} else {
		return http.StatusInternalServerError
	}
}
