package web

type ErrorResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

func NewBaseErrorResponse(message string) *ErrorResponse {
	return &ErrorResponse{
		Error:   true,
		Message: message,
	}
}
