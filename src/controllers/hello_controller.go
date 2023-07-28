package controllers

import (
	"fmt"
	"net/http"
)

func HelloHandler(writer http.ResponseWriter, request *http.Request) {

	fmt.Fprintf(writer, "BOM DIA FILHA DA PUTA!")
}
