package tutorial

import (
	"errors"
	"log"
)

// Number
// First, type the enumeration Values
// Numbers for examples (first is 1, second is 2, etc ...)
// Using 'uint' type is safer, so that you cant use negative enum values
type Number uint

// Greek
// Here, we enumerate seasons with numbers
type Greek uint

// Season is a string type, will be used for string enums
type Season string

// Fellow is a slug structure
type Fellow struct {
	slug string
}

// Define the accepted values
// Here the values are explicit
const (
	Zero Number = 0
	One  Number = 1
	Two  Number = 2
)

// Here, we wil use 'iota'
// 'iota' do the same as enums with numbers from 0 to the amount of defined variables
// Define 'Unknown' for the zero value is a best practice because 0 is also the 'empty' value of the enum
// So NEVER define an enum expected value from 0
const (
	Unknown Greek = iota
	Alpha
	Beta
	Gamma
)

// string enum
const (
	UnknownSeason Season = ""
	FirstSeason   Season = "spring"
	SecondSeason  Season = "summer"
	ThirdSeason   Season = "Autumn"
	FourthSeason  Season = "winter"
)

// a slug uses vars, not const
var (
	UnknownFellow = Fellow{""}
	Frodo         = Fellow{"Frodo"}
	Sam           = Fellow{"Sam"}
	Pipin         = Fellow{"Pipin"}
	Mery          = Fellow{"Mery"}
)

// getNumFromGreek function will help us to make safe enum calls
// like dealing with empty '0' values
func getNumFromGreek(g Greek) {
	if g == Unknown {
		log.Print("Return Panic Error ! This is an unknown value")
	}
	log.Printf("This is the matching greek number : '%d'", g)
}

func getSeason(s Season) {
	if s == UnknownSeason {
		log.Print("Return Panic Error ! This is an unknown value")
	}
	log.Printf("This is the given season : '%s'", s)
}

// getFellow return the value of a fellow
func getFellow(f Fellow) string {
	return f.slug
}

// FellowFromString returns a Fellow var from a string value
func FellowFromString(s string) (Fellow, error) {
	switch s {
	case Frodo.slug:
		return Frodo, nil
	case Sam.slug:
		return Sam, nil
	case Pipin.slug:
		return Pipin, nil
	case Mery.slug:
		return Mery, nil
	}

	return UnknownFellow, errors.New("Unknown fellow : " + s)
}

func Enums() {
	// basic enum
	log.Printf("This is the second number value : '%d'", Two)

	// safer enum
	// good values
	getNumFromGreek(Alpha)
	getNumFromGreek(Beta)

	// string enum
	getSeason(FirstSeason)

	// slug : structure enumeration
	// this is a best practice to apply
	getFellow(Sam)
	f, _ := FellowFromString("Sam")
	log.Printf("Build a fellow from string : %v", f)
}
