package utils

import (
	"errors"
	"fmt"
)

// GetUserInput
// return the user input or an error
func GetUserInput(inputInfo string) (string, error) {
	fmt.Print(inputInfo)
	var userInput string
	_, err := fmt.Scanln(&userInput)
	if err != nil {
		return "", errors.New("cannot process user input")
	}
	return userInput, nil
}
