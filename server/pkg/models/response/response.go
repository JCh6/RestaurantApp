package response

type Response struct {
	Code         int    `json:"code"`
	ErrorMessage string `json:"errorMessage"`
	Body         string `json:"body"`
}

func New(code int, errorMessage string, body string) *Response {
	return &Response{
		Code:         code,
		ErrorMessage: errorMessage,
		Body:         body,
	}
}
