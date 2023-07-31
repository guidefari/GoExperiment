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
	
    message, err := greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}

    fmt.Println(message)
}