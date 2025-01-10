package api

type ApiError struct {
	Message string `json:"message"`
	code    int
}

func (e *ApiError) Error() string {
	return e.Message
}

func NewError(message string, args ...interface{}) *ApiError {
	err := &ApiError{
		Message: message,
		code:    400,
	}
	if len(args) > 0 {
		if code, ok := args[0].(int); ok {
			err.code = code
		}
	}
	return err
}
