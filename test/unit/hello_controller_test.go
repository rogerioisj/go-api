package unit

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"roger.com/api/src/controllers"
)

func TestHelloHandler(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:3001", nil)
	writer := httptest.NewRecorder()
	controllers.HelloHandler(writer, request)

	fmt.Println("Request event: \n", writer.Body)

	if writer.Body.String() != "BOM DIA FILHA DA PUTA!" {
		t.Errorf("Expected a differente message than was received, got %s", writer.Body.String())
	}
}
