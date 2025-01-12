package api

type ApiData struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	code    int
}

func NewApiData(message string) *ApiData {
	data := &ApiData{
		Message: message,
		code:    200,
		Data:    nil,
	}
	return data
}

func (e *ApiData) SetCode(code int) *ApiData {
	e.code = code
	return e
}

func (e *ApiData) SetData(data interface{}) *ApiData {
	e.Data = data
	return e
}
