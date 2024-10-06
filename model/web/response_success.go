package web

type BaseSuccessResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewBaseSuccessResponse(message string, data interface{}) *BaseSuccessResponse {
	return &BaseSuccessResponse{
		Error:   false,
		Message: message,
		Data:    data,
	}
}

type BaseSuccessPaginationResponse struct {
	Error      bool        `json:"error"`
	Message    string      `json:"message"`
	Pagination Metadata    `json:"pagination"`
	Data       interface{} `json:"data"`
}

func NewBaseSuccessPaginationResponse(message string, pagination Metadata, data interface{}) *BaseSuccessPaginationResponse {
	pagination.Limit = 15
	return &BaseSuccessPaginationResponse{
		Error:      false,
		Message:    message,
		Pagination: pagination,
		Data:       data,
	}
}
