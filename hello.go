package main

import (
	"fmt"
	"log"

	"github.com/FahrezaIsnanto/golanggreetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Muhammad", "Fahreza", "Isnanto"}

	// Request greeting messages for the names.
	messages, err := golanggreetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	// If no error was returned, print the returned map of
	// messages to the console.
	fmt.Println(messages)
}
