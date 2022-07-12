package tutorial

import (
	"log"
)

func Variables() {
	// one after one
	var firstName string // set as ""
	var lastName string  // set as ""
	var age int          // set as 0, 64 bits wide on 64 bits systems, 32-bit wide on 32 bits systems
	var weight float32   // set as 0.0
	var boy bool         // set as False
	var blood rune       // rune is the 'Char' type, equivalent to int32, set as '' by default
	var bloodBis byte    // byte is also the 'Char' type, equivalent to uint8, set as '' by default
	log.Printf("firstName: %s, lastName: %s, age: %d, weight: %f, boy: %t, blood group: %d, blood group: %d\n",
		firstName, lastName, age, weight, boy, blood, bloodBis)

	// all at once
	var (
		FirstNameBis string
		LastNameBis  string
		ageBis       int
		boyBis       bool
	)
	log.Printf("firstName: %s, lastName: %s, age: %d, gender: %t\n", FirstNameBis, LastNameBis, ageBis, boyBis)

	// Type inference
	firstName2 := "Sam Sagace" // string type
	lastName2 := "Gamegie"     // string type
	age2 := 38                 // no .x, so it is an int
	weight2 := 70.0            // .x, so it is a float32
	boy2 := true               // bool type
	blood2 := 'O'              // rune type
	log.Printf("firstName: %s, lastName: %s, age: %d, weight: %f, gender: %t, blood group: %d\n",
		firstName2, lastName2, age2, weight2, boy2, blood2)

	// constants
	// in Go, const are untyped by default
	const untyped1 = 1
	const untyped2 = "not any type"
	const untyped3 = false
	// always type constants
	const favoriteLanguage string = "Golang"
	const currentMajorVersion int = 1
	const currentMinorVersion int = 18
	const typedInteger = currentMajorVersion // this var is typed since it is instantiated from an int
	// cannot be changed
	// favoriteLanguage = "Python" // will provoke compiler error
}
