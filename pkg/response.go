package pkg

type ApiResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Error   any    `json:"error,omitempty"`
}

func NewErrorResponse(message string, error any) *ApiResponse {
	return &ApiResponse{
		Success: false,
		Message: message,
		Error:   error,
	}
}

func NewSuccessResponse(message string, data any) *ApiResponse {
	return &ApiResponse{
		Success: true,
		Message: message,
		Data:    data,
	}
}
