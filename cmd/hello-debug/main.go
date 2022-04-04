package main

import (
	"fmt"
	"golang_starter/pkg/greetings"
	"log"
	"rsc.io/quote"
	"time"
)

func main() {
	// set logger
	//log.SetPrefix("greetings: ")
	//log.SetFlags(0)

	// built-in package
	fmt.Println(quote.Go())

	// Will say hello every seconds
	names := []string{"Kaio", "Maty", "Sokka"}
	for {
		sayHello(names)
		time.Sleep(1 * time.Second)
	}
}

func sayHello(names []string) {
	// custom package
	message, err := greetings.Hellos(names)
	// raising an error (check the function definition)
	//message, err := greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
