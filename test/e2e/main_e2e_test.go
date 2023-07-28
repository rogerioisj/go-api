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

	fmt.Println("Request event: \n", request)
	fmt.Println("response event: \n", response)

	if response.Code != 200 {
		t.Errorf("Expected status code 200, got %d", response.Code)
	}
}
