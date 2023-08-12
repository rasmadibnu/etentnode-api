package helper

type Meta struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func Response(message string, status int, data interface{}) Meta {
	resp := Meta{
		StatusCode: status,
		Message:    message,
		Data:       data,
	}

	return resp
}
