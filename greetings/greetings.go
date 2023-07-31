package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {

	if name=="" {
		return "", errors.New("gimme a name bro")
	}

    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	formatsLength := len(formats)
	
	return formats[rand.Intn(formatsLength)]
}