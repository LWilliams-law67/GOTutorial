package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// Return a greeting that embeds the name in a message.
	// var message string
	//message = fmt.Sprintf("Hi, %v. Welcome!", name)
	message := fmt.Sprintf(randomString(), name)
	return message, nil
}

// MultiHello : return a map that pairs a greeting for each person
func MultiHello(names []string) (map[string]string, error) {
    // A map to associate names with messages.
    messages := make(map[string]string)
    // Loop through the received slice of names, calling
    // the Hello function to get a message for each name.
    for _, name := range names {
        message, err := Hello(name)
        if err != nil {
            return nil, err
        }
        // In the map, associate the retrieved message with the name.
        messages[name] = message
    }
    return messages, nil
}

// init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns a random greeting.
func randomString() string {
	// A slice of message strings
	strings := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v",
		"Hail, %v! Well met!",
	}

	// Return a random message from strings
	return strings[rand.Intn(len(strings))]
}
