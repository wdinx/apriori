package constant

import "errors"

var ErrInsertDatabase error = errors.New("Invalid Add Data")
var ErrEmptyInput error = errors.New("field cannot be empty")
var ErrDataNotFound error = errors.New("Data not found")

var ErrLogin error = errors.New("Invalid Username or Password")
var ErrRegister error = errors.New("Email already taken")

var ErrInsertData error = errors.New("Failed to input data")
var ErrUpdateData error = errors.New("Failed to update data")
var ErrDeleteData error = errors.New("Failed to delete data")

var ErrUnsupportedFileFormat error = errors.New("Unsupported File Format")
var ErrInternalServer error = errors.New("Internal Server Error")
var ErrInvalidToken error = errors.New("Invalid Token")

var ErrFailedToCreateSnapURL error = errors.New("Failed to create snap URL")
var ErrPaymentNotFound error = errors.New("Payment not found")
var ErrPaymentAlreadyConfirmed error = errors.New("Payment already confirmed")

var ErrInputDate error = errors.New("Invalid Date Input")
var ErrUnauthorized error = errors.New("Unauthorized")
