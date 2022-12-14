package main

import (
	"fmt"
	"log"
	"tutorial/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names to get greetings for
	names := []string{"Luke", "Alex", "Josh", "Mike"}

	// Request greeting messages for the slice of names
	messages, err := greetings.MultiHello(names)

	// Request a greeting message and print it.
	//message, err := greetings.Hello("Luke")

	// If an error was returned, print it to the console and exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error, prints map of names and greetings
	fmt.Println(messages)

	// If no error was returned, print the message to the console.
	//fmt.Println(message)
}
