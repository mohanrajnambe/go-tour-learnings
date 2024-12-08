package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	//Get a greeting message and print it
	message := greetings.Hello("Mohan")
	printMessageWithError(message, nil)
	message, err := greetings.RandomHello("Mohan")
	printMessageWithError(message, err)
	// message, err = greetings.HelloWithError("")
	// printMessageWithError(message, err)

	// slice of names
	names := []string{"Gladys", "Samantha", "Darrin"}
	messages, err := greetings.Hellos(names)

	printMessagesWithError(messages, err)

}

func printMessageWithError(message string, err error) {

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}

func printMessagesWithError(messages map[string]string, err error) {

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
