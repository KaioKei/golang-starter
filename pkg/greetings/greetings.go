package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// The 'init' function allows us to set up database connections, or register with various service registries, or perform
// any number of other tasks that you typically only want to do once.
// Otherwise, check 'New()' functions to initiate such things once per structure and not per callback.
// It’s incredibly important to note that you cannot rely upon the order of execution of your init() functions.
// It’s instead better to focus on writing your systems in such a way that the order does not matter.
func init() {
	// initiate math/rand once after 'greetings' package call
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	// A slice of message formats
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you %v!",
		"Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}

// randomFormat returns one of a set of greeting messages. The returned message is selected at random.

// Hello returns string (a greeting for the named person) or an error if an empty string is provided as input
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("no name was provided")
	}
	// name: name to print when greeting
	// Return a greeting that embeds the name in a message.
	// declare message in one line
	message := fmt.Sprintf(randomFormat(), name)

	// return the message
	// and 'nil'. That way, the caller can see that the function succeeded.
	return message, nil
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	// Initialize a map with the following syntax: make(map[key-type]value-type)
	messages := make(map[string]string)

	// Loop through the received slice of names, calling
	// the Hello function to get a message for each name.
	// Don't need the index, so you use the Go blank identifier (an underscore) to ignore it
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		// In the map, associate the retrieved message with
		// the name.
		messages[name] = message
	}

	return messages, nil
}
