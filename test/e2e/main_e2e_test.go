package e2e

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func NewServer(handler http.Handler) *httptest.Server {
	return httptest.NewServer(handler)
}

func TestMain(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:3001", nil)
	response := httptest.NewRecorder()

	fmt.Print("Request event: ", request)
	fmt.Print("response event: ", response)

	if response.Code != 201 {
		t.Errorf("Expected status code 201, got %d", response.Code)
	}
}
