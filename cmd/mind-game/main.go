package main

import (
	"fmt"
	"golang_starter/internal/randomizer"
	"golang_starter/internal/utils"
	"log"
	"strconv"
)

type magicNumber struct {
	Value        uint8
	UserAttempts uint8
}

func main() {
	fmt.Println("----------")
	fmt.Println("Welcome to the mind game !")
	fmt.Println("The goal is to guess the magic number between 1 and 100. Ready ?")
	fmt.Println("----------")
	var random = randomizer.GetRandomNumber(1, 100)
	mn := magicNumber{Value: random}
	intNumber := uint8(101) // impossible value for purpose

	startGameLoop(intNumber, &mn)
}

func startGameLoop(number uint8, pMagicNumber *magicNumber) {
	for number != pMagicNumber.Value {
		pMagicNumber.UserAttempts++
		// we must instantiate err because 'number' already has a value
		// thus we have to affect new values using '=' and not instantiate them with ':=' because of 'number'.
		var err error
		number, err = processInput()
		if err != nil {
			// process error at the top level code as a failure
			log.Fatal(err)
		}

		analyseNumber(number, pMagicNumber)
	}
}

// processInput
// propagate err through the code to the upper code level by returning it alongside the proper return value type
func processInput() (uint8, error) {
	strNumber, err := utils.GetUserInput("Try a number : ")
	if err != nil {
		log.Fatal(err)
	}

	intNumber, err := strconv.Atoi(strNumber)
	if err != nil {
		log.Fatal(err)
	}
	// propagate err through the code to the upper code level
	return uint8(intNumber), err
}

func analyseNumber(userNumber uint8, pMagicNumber *magicNumber) {
	if userNumber > pMagicNumber.Value {
		fmt.Println("Try lower ...")
	} else if userNumber < pMagicNumber.Value {
		fmt.Println("Try higher ...")
	} else {
		fmt.Println("")
		fmt.Printf("You win ! The magic number was %d !\n", pMagicNumber.Value)
		fmt.Printf("Number of attempts: %d\n", pMagicNumber.UserAttempts)
	}
}
