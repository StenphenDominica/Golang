package main

import (
	"fmt"
	"log"

	"rsc.io/quote"

	"example.com/greetings"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())

	greet("Stephen")

	names := []string{"Gladys", "Simmons", "Darwin"}
	greets(names)
}

func greet(name string) {
	message, err := greetings.Hello(name)
	if err == nil {
		fmt.Println(message)
	} else {
		log.Fatal(err)
	}
}

func greets(names []string) {
	messages, err := greetings.Hellos(names)
	if err == nil {
		fmt.Println(messages)
	} else {
		log.Fatal(err)
	}
}
