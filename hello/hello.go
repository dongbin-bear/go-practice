package main

import (
	"fmt"
	"log"

	"go-practice/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	// log.SetFlags(0)

	message, err := greetings.Hello("Dongbin")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	names := []string{"Dongbin", "HJ"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	for _, greeting := range messages {
		fmt.Println(greeting)
	}
}
