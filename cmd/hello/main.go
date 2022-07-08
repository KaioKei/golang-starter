package main

import (
	"golang_starter/pkg/greetings"
	"log"
	"rsc.io/quote"
)

func main() {
	// set logger
	//log.SetPrefix("greetings: ")
	//log.SetFlags(0)

	// built-in package
	log.Println(quote.Go())

	// input
	names := []string{"Kaio", "Maty", "Sokka"}

	// custom package
	message, err := greetings.Hellos(names)
	// raising an error (check the function definition)
	//message, err := greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(message)
}
