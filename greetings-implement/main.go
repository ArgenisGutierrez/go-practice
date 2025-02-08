package main

import (
	"fmt"
	"log"

	"github.com/ArgenisGutierrez/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	message, err := greetings.Hello("Arfhel")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	names := []string{
		"Arfhle",
		"Samaerl",
		"Engrias",
	}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	for _, message := range messages {
		fmt.Println(message)
	}
}
