package main

import (
	"golang_starter/internal/tutorial"
	"log"
)

func main() {
	tutorial.Variables()
	tutorial.Statements()

	// functions
	tutorial.Say("'Say' is a public function from 'functions' go file")
	tutorial.Functions()

	tutorial.Packages()
	tutorial.Arrays()
	tutorial.Slices()
	tutorial.Maps()
	tutorial.Pointers()

	// structures
	tutorial.Structures()
	//tutorial.confidential cannot be accessed
	//tutorial.Public{packageSecret: "secret"} raises an error
	s1 := tutorial.Public{PackagePublic: "public"} // is ok
	log.Print(s1)

	tutorial.Methods()
	tutorial.Interfaces()
}
