package plaud

import "net/http"

type Data struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	code    int
}

func NewData(message string) *Data {
	data := &Data{
		Message: message,
		code:    http.StatusOK,
		Data:    nil,
	}

	return data
}

func (e *Data) SetCode(code int) *Data {
	e.code = code
	return e
}

func (e *Data) SetData(data interface{}) *Data {
	e.Data = data
	return e
}
