package plaud

import "net/http"

type Error struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	code    int
}

func (e *Error) Error() string {
	return e.Message
}

func NewError(message string) *Error {
	return &Error{
		Message: message,
		Data:    nil,
		code:    http.StatusBadRequest,
	}
}

func (e *Error) SetCode(code int) *Error {
	e.code = code
	return e
}

func (e *Error) SetData(data interface{}) *Error {
	e.Data = data
	return e
}
