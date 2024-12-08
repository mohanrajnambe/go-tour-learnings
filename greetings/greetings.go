package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) string {
	// Return a greeting
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

// Return a greeting
func HelloWithError(name string) (string, error) {
	// if no name, return an error
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}

// Return random greeting
func RandomHello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	// Create a message using a random format.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := RandomHello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}
