package api

type ApiError struct {
	Message string `json:"message"`
	code    int
	Data    interface{} `jsong:"data"`
}

func (e *ApiError) Error() string {
	return e.Message
}

func NewError(message string, args ...interface{}) *ApiError {
	err := &ApiError{
		Message: message,
		code:    400,
		Data:    nil,
	}
	if len(args) > 0 {
		if code, ok := args[0].(int); ok {
			err.code = code
		}
	}

	if len(args) > 1 {
		err.Data = args[1]
	}

	return err
}
