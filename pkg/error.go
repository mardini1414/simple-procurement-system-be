package pkg

type ApiError struct {
	Code    int
	Message string
	Detail  any
}

func (e *ApiError) Error() string {
	return e.Message
}

func NewApiError(code int, message string, detail any) *ApiError {
	return &ApiError{
		Code:    code,
		Message: message,
		Detail:  detail,
	}
}
