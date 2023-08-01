package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	// uncomment line below if you don't want to see date and time in the logs
    // log.SetFlags(0)
	
	names := []string{"Gladys", "Samantha", "Darrin"}
    message, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

    fmt.Println(message)
}