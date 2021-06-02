package response

import "testing"

func TestNewResponse(t *testing.T) {
	code := 200
	errorMessage := ""
	body := "OK"

	r := New(code, errorMessage, body)

	if r.Code != code || r.ErrorMessage != errorMessage || r.Body != body {
		t.Errorf("Response was not created correctly")
	}
}
