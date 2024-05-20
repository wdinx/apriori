package web

type BaseErrorResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewBaseSuccessResponse(message string, data interface{}) *BaseErrorResponse {
	return &BaseErrorResponse{
		Error:   false,
		Message: message,
		Data:    data,
	}
}
