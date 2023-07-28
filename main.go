package main

import (
	"fmt"
)

func Hello(text string) string {
	message := fmt.Sprintf("Received message: %s", text)

	return message
}

func main() {
	var message string
	fmt.Println("Hello!")
	fmt.Println("Plesae, insert your message: ")
	fmt.Scanln(&message)

	fmt.Println(Hello(message))
}
