package api

type ApiData struct {
	Message string `json:"message"`
	code    int
	Data    interface{} `jsong:"data"`
}

func NewApiData(message string, args ...interface{}) *ApiData {
	data := &ApiData{
		Message: message,
		code:    200,
		Data:    nil,
	}
	if len(args) > 0 {
		if code, ok := args[0].(int); ok {
			data.code = code
		}
	}

	if len(args) > 1 {
		data.Data = args[1]
	}

	return data
}
