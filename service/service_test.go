package service

import (
	"net/http/httptest"
	"testing"
)

func TestDidcommHandlerDisallowedMethods(t *testing.T) {
	disallowedMethods := [4]string{"GET", "OPTIONS", "PUT", "DELETE"}

	for _, method := range disallowedMethods {
		writer := httptest.NewRecorder()
		req := httptest.NewRequest(method, "/didcomm", nil)
		didcommHandler(writer, req)

		if writer.Code != 405 {
			t.FailNow()
		}
	}

}
