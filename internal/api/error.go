package api

type ApiError struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	code    int
}

func (e *ApiError) Error() string {
	return e.Message
}

func NewError(message string) *ApiError {
	err := &ApiError{
		Message: message,
		Data:    nil,
		code:    400,
	}
	return err
}

func (e *ApiError) SetCode(code int) *ApiError {
	e.code = code
	return e
}

func (e *ApiError) SetData(data interface{}) *ApiError{
	e.Data=data
	return e
}
