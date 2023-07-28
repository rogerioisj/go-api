package main

import (
	"fmt"
	"log"
	"net/http"

	"roger.com/api/src/controllers"
)

func main() {
	http.HandleFunc("/", controllers.HelloHandler)

	fmt.Print("Server running on port 3001")

	if err := http.ListenAndServe(":3001", nil); err != nil {
		log.Fatal(err)
	}
}
