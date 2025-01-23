package plaud

type Data struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	code    int
}

func NewApiData(message string) *Data {
	data := &Data{
		Message: message,
		code:    200,
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
