package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Print("HelloHandle")
	fmt.Print("Request event: ", request)

	fmt.Fprintf(writer, "BOM DIA FILHA DA PUTA!")
}

func main() {
	http.HandleFunc("/", helloHandler)

	fmt.Print("Server running on port 3001")

	if err := http.ListenAndServe(":3001", nil); err != nil {
		log.Fatal(err)
	}
}
