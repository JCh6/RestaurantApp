package response

import (
	"encoding/json"
)

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

func GetResponseBody(code int, errorMessage string, body string) []byte {
	var resp *Response

	resp = New(code, errorMessage, body)
	resBodyBytes, err := json.Marshal(resp)

	if err != nil {
		return []byte(err.Error())
	}

	return resBodyBytes
}
