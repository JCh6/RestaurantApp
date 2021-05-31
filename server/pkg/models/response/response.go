package response

import (
	"encoding/json"
)

type Response struct {
	Code         int         `json:"code"`
	ErrorMessage string      `json:"errorMessage"`
	Body         interface{} `json:"body"`
}

func New(code int, errorMessage string, body interface{}) *Response {
	return &Response{
		Code:         code,
		ErrorMessage: errorMessage,
		Body:         body,
	}
}

func ContentType() string {
	return "application/json; charset=utf-8"
}

func GetResponseBody(code int, errorMessage string, body interface{}) []byte {
	var resp *Response

	resp = New(code, errorMessage, body)
	resBodyBytes, err := json.Marshal(resp)

	if err != nil {
		return []byte(err.Error())
	}

	return resBodyBytes
}
