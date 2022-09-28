package tutorial

import (
	"errors"
	"log"
)

func Functions() {
	// no return
	Say("Functions tutorial")

	// return
	average := average(11, 16)
	log.Printf("Average of 11 & 16 is: %f\n", average)

	// shorter function definitions
	person("Frodo", "Baggins", 33) // see the shorter function definition

	// multiple return
	var oldPrice float32 = 12.50
	var newPrice float32 = 22.75
	difference, percentage := getPriceChange(oldPrice, newPrice)
	log.Printf("From %f to %f, the price changed by %f which represents %d%",
		oldPrice, newPrice, difference, percentage)

	// raise error
	_, shouldBeError := divide(10, 0)
	log.Println("Result is an error: ", shouldBeError)

	// named return
	defaultResult, changedResult := namedReturn()
	log.Println("Default result should be 0: ", defaultResult)
	log.Println("Changed result should be 10: ", changedResult)

	// blank return
	isZero, _ := namedReturn()
	_, isTen := namedReturn()
	log.Println("Result should be 0: ", isZero)
	log.Println("Result should be 10: ", isTen)

	// say 'world' when the function has returned its result (here, return nil)
	deferHelloWorld()
}

// Say accessible outside this package due to the upper case 'S'
func Say(s string) {
	log.Printf("%s", s)
}

// average not accessible outside this package due to the lower case 'a'
func average(x float32, y float32) float32 {
	return (x + y) / 2
}

// same but shorter
func person(firstName, lastName string, age int) {
	log.Printf("Hi, I am %s %s and I am %d years old.", firstName, lastName, age)
}

func getPriceChange(oldPrice, newPrice float32) (float32, int) {
	diff := newPrice - oldPrice
	ratio := int((diff / oldPrice) * 100)
	return diff, ratio
}

func divide(x, y float32) (float32, error) {
	if y == 0 {
		err := errors.New("cannot divide by 0")
		return 0, err
	}
	return x / y, nil
}

func namedReturn() (resultByDefault, resultChanged int) {
	// Returned value are already instantiated from the function definition with default's type value
	// So we cannot change it with ':=' but only with '='
	resultChanged = 10
	// It allows as to return both the variables without naming it
	// so the returned values are :
	//   - 'resultByDefault=0' since it as an int and we have not modified it
	//   - 'resultChanged=10' since we modified it
	return
}

func deferHelloWorld() {
	//  A defer statement defers the execution of a function until the surrounding function returns.
	// The deferred call's arguments are evaluated immediately, but the function call is not executed
	// until the surrounding function returns.
	defer log.Println("world !")

	log.Println("hello")
}
