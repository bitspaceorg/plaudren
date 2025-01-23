package plaud

type Error struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	code    int
}

func (e *Error) Error() string {
	return e.Message
}

func NewError(message string) *Error {
	err := &Error{
		Message: message,
		Data:    nil,
		code:    400,
	}
	return err
}

func (e *Error) SetCode(code int) *Error {
	e.code = code
	return e
}

func (e *Error) SetData(data interface{}) *Error{
	e.Data=data
	return e
}
